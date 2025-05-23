package autobots

import (
	"context"
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/klog/v2"
	transformersv1alpha1 "k8s.io/sample-apiserver/pkg/apis/transformers/v1alpha1"
)

var (
	_ rest.Storage              = &FakeREST{}
	_ rest.Creater              = &FakeREST{}
	_ rest.Updater              = &FakeREST{}
	_ rest.Getter               = &FakeREST{}
	_ rest.Lister               = &FakeREST{}
	_ rest.Watcher              = &FakeREST{}
	_ rest.GracefulDeleter      = &FakeREST{}
	_ rest.CollectionDeleter    = &FakeREST{}
	_ rest.Scoper               = &FakeREST{} // 必须实现这个接口
	_ rest.SingularNameProvider = &FakeREST{} // 必须实现这个接口
	// _ rest.Connecter            = &FakeREST{} // 用于支持像 kubectl exec, kubectl attach, kubectl port-forward 这样的 "连接" 或 "流式" API 子资源
)

// 看起来判断请求 URL 是否是 connect 的依据是是否是子资源，如果路径有 8 层，则认为是子资源，比如
// /apis/{{group}}/{{version}}/namespaces/{{namespace}}/{{resource}}/{{name}}/{{subresource}}
// 如果路径有 7 层，则认为是资源，比如
// /apis/{{group}}/{{version}}/namespaces/{{namespace}}/{{resource}}/{{name}}

// 如果是父资源无需实现 Connecter 接口，子资源才需要

type FakeREST struct{}

// DeleteCollection implements rest.CollectionDeleter.
// 批量删除
func (f *FakeREST) DeleteCollection(ctx context.Context, deleteValidation rest.ValidateObjectFunc, options *metav1.DeleteOptions, listOptions *internalversion.ListOptions) (runtime.Object, error) {
	ls := listOptions.LabelSelector
	namespace, _ := request.NamespaceFrom(ctx)
	return &transformersv1alpha1.AutobotList{
		Items: []transformersv1alpha1.Autobot{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "autobot-" + ls.String(),
					Namespace: namespace,
				},
			},
		},
	}, nil
}

// Connect implements rest.Connecter.
func (f *FakeREST) Connect(ctx context.Context, id string, options runtime.Object, r rest.Responder) (http.Handler, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()

		ri, ok := request.RequestInfoFrom(req.Context())
		if !ok {
			klog.Errorf("Connect handler: no RequestInfo found in context for Autobot %s", id)
			http.Error(w, "Internal server error: no RequestInfo found in context", http.StatusInternalServerError)
			return
		}

		klog.Infof("Connect handler for Autobot '%s', Subresource: '%s', Path: '%s'", id, ri.Subresource, req.URL.Path)

		switch ri.Subresource {
		case "transform":
			if req.Method == http.MethodPost {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(fmt.Sprintf("Autobot '%s' POST transformation complete.", id)))
			} else if req.Method == http.MethodGet {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(fmt.Sprintf("GET status for transforming Autobot '%s'.", id)))
			} else {
				http.Error(w, fmt.Sprintf("Method %s not allowed for subresource 'transform'", req.Method), http.StatusMethodNotAllowed)
			}

		case "connect":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("Connect subresource for Autobot '%s' says: Hello, World!", id)))

		default:
			klog.Warningf("Subresource '%s' not supported for Autobot '%s'", ri.Subresource, id)
			http.Error(w, fmt.Sprintf("Subresource '%s' not supported", ri.Subresource), http.StatusNotFound)
		}
	}), nil
}

// ConnectMethods implements rest.Connecter.
func (f *FakeREST) ConnectMethods() []string {
	return []string{"GET", "POST"}
}

// NewConnectOptions implements rest.Connecter.
func (f *FakeREST) NewConnectOptions() (runtime.Object, bool, string) {
	return nil, false, ""
}

// Delete implements rest.GracefulDeleter.
func (f *FakeREST) Delete(ctx context.Context, name string, deleteValidation rest.ValidateObjectFunc, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	namespace, _ := request.NamespaceFrom(ctx)
	return &transformersv1alpha1.Autobot{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name + "-deleted",
			Namespace: namespace,
		},
	}, true, nil
}

// GetSingularName implements rest.SingularNameProvider.
func (f *FakeREST) GetSingularName() string {
	return "autobot"
}

// NamespaceScoped implements rest.Scoper.
func (f *FakeREST) NamespaceScoped() bool {
	return true
}

var _ watch.Interface = &FakeWatch{}

type FakeWatch struct {
	ch  chan watch.Event
	ctx context.Context
}

func (f *FakeWatch) Stop() {
	klog.V(2).Infof("FakeWatch stopping")
	close(f.ch)
	f.ctx.Done()
}

func (f *FakeWatch) ResultChan() <-chan watch.Event {
	go func() {
		for {
			tk := time.NewTicker(time.Duration(rand.Int32N(4)+1) * time.Second)
			select {
			case <-tk.C:
				f.ch <- watch.Event{
					Type: watch.Added,
					Object: &transformersv1alpha1.Autobot{
						ObjectMeta: metav1.ObjectMeta{
							Name: "autobot-" + strconv.Itoa(int(rand.Int32N(1000))),
						},
					},
				}
			case <-f.ctx.Done():
				klog.V(2).Infof("FakeWatch context done")
				return
			}
		}
	}()
	return f.ch
}

// Watch implements rest.Watcher.
func (f *FakeREST) Watch(ctx context.Context, opts *internalversion.ListOptions) (watch.Interface, error) {
	fw := &FakeWatch{
		ch:  make(chan watch.Event),
		ctx: ctx,
	}
	return fw, nil
}

// ConvertToTable implements rest.Lister.
func (f *FakeREST) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	return &metav1.Table{}, nil
}

// List implements rest.Lister.
func (f *FakeREST) List(ctx context.Context, options *internalversion.ListOptions) (runtime.Object, error) {
	return &transformersv1alpha1.AutobotList{
		Items: []transformersv1alpha1.Autobot{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "autobot-1",
				},
				Spec: transformersv1alpha1.AutobotSpec{
					GenericSpec: transformersv1alpha1.GenericSpec{
						Mode:     transformersv1alpha1.ModeCar,
						Name:     "autobot-1",
						Strength: 100,
					},
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "autobot-2",
				},
				Spec: transformersv1alpha1.AutobotSpec{
					GenericSpec: transformersv1alpha1.GenericSpec{
						Mode:     transformersv1alpha1.ModeCar,
						Name:     "autobot-2",
						Strength: 100,
					},
				},
			},
		},
	}, nil
}

// NewList implements rest.Lister.
func (f *FakeREST) NewList() runtime.Object {
	return &transformersv1alpha1.AutobotList{}
}

// Update implements rest.Updater.
func (f *FakeREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	newObj, err := objInfo.UpdatedObject(ctx, &transformersv1alpha1.Autobot{})
	if err != nil {
		return nil, false, err
	}
	autobot, ok := newObj.(*transformersv1alpha1.Autobot)
	if !ok {
		return nil, false, fmt.Errorf("obj is not a Autobot")
	}
	autobot.Name = autobot.Name + "-updated"
	return autobot, true, nil
}

// New implements rest.Storage.
func (f *FakeREST) New() runtime.Object {
	return &transformersv1alpha1.Autobot{}
}

func (f *FakeREST) Destroy() {}

func (f *FakeREST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	autobot, ok := obj.(*transformersv1alpha1.Autobot)
	if !ok {
		return nil, fmt.Errorf("obj is not a Autobot")
	}
	if createValidation != nil {
		if err := createValidation(ctx, obj); err != nil {
			return nil, err
		}
	}
	autobot.Name = autobot.Name + "-created"
	return autobot, nil
}

func (f *FakeREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	namespace, _ := request.NamespaceFrom(ctx)
	return &transformersv1alpha1.Autobot{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: transformersv1alpha1.AutobotSpec{
			GenericSpec: transformersv1alpha1.GenericSpec{
				Mode:     transformersv1alpha1.ModeCar,
				Name:     name,
				Strength: 100,
			},
		},
	}, nil
}
