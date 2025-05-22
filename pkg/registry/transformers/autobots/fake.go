package autobots

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/registry/rest"
	transformersv1alpha1 "k8s.io/sample-apiserver/pkg/apis/transformers/v1alpha1"
)

var (
	_ rest.Storage              = &FakeREST{}
	_ rest.Creater              = &FakeREST{}
	_ rest.Updater              = &FakeREST{}
	_ rest.Getter               = &FakeREST{}
	_ rest.Lister               = &FakeREST{}
	_ rest.Watcher              = &FakeREST{}
	_ rest.Scoper               = &FakeREST{} // 必须实现这个接口
	_ rest.SingularNameProvider = &FakeREST{} // 必须实现这个接口
)

type FakeREST struct{}

// GetSingularName implements rest.SingularNameProvider.
func (f *FakeREST) GetSingularName() string {
	return "autobot"
}

// NamespaceScoped implements rest.Scoper.
func (f *FakeREST) NamespaceScoped() bool {
	return true
}

// Watch implements rest.Watcher.
func (f *FakeREST) Watch(ctx context.Context, options *internalversion.ListOptions) (watch.Interface, error) {
	panic("unimplemented")
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
	panic("unimplemented")
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
	autobot.Name = autobot.Name + "-fake"
	return autobot, nil
}

func (f *FakeREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return &transformersv1alpha1.Autobot{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
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
