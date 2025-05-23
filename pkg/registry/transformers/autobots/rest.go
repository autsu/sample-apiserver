package autobots

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/sample-apiserver/pkg/apis/transformers"
	transformersv1alpha1 "k8s.io/sample-apiserver/pkg/apis/transformers/v1alpha1"
	//"k8s.io/sample-apiserver/pkg/registry"
)

var (
	_ rest.Storage = &REST{}
	_ rest.Creater = &REST{}
	_ rest.Updater = &REST{}
	_ rest.Getter  = &REST{}
	_ rest.Lister  = &REST{}
	_ rest.Watcher = &REST{}
)

type REST struct{}

// Watch implements rest.Watcher.
func (r *REST) Watch(ctx context.Context, options *internalversion.ListOptions) (watch.Interface, error) {
	panic("unimplemented")
}

// ConvertToTable implements rest.Lister.
func (r *REST) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*v1.Table, error) {
	panic("unimplemented")
}

// List implements rest.Lister.
func (r *REST) List(ctx context.Context, options *internalversion.ListOptions) (runtime.Object, error) {
	return &transformers.AutobotList{}, nil
}

// NewList implements rest.Lister.
func (r *REST) NewList() runtime.Object {
	return &transformers.AutobotList{}
}

// Get implements rest.Getter.
func (r *REST) Get(ctx context.Context, name string, options *v1.GetOptions) (runtime.Object, error) {
	// return r.Store.Get(ctx, name, options)
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

// Update implements rest.Updater.
func (r *REST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *v1.UpdateOptions) (runtime.Object, bool, error) {
	panic("unimplemented")
}

// Create implements rest.Creater.
func (r *REST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *v1.CreateOptions) (runtime.Object, error) {
	panic("unimplemented")
}

// Destroy implements rest.Storage.
func (r *REST) Destroy() {
}

// New implements rest.Storage.
func (r *REST) New() runtime.Object {
	return &transformers.Autobot{}
}

// NewREST 返回主资源的 REST 存储接口
func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*REST, error) {
	strategy := NewStrategy(scheme)

	store := &genericregistry.Store{
		NewFunc:                   func() runtime.Object { return &transformers.Autobot{} },
		NewListFunc:               func() runtime.Object { return &transformers.AutobotList{} },
		PredicateFunc:             MatchAutobot,
		DefaultQualifiedResource:  transformers.Resource("autobots"),
		SingularQualifiedResource: transformers.Resource("autobot"),

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,

		// TODO: define table converter that exposes more than name/creation timestamp
		TableConvertor: rest.NewDefaultTableConvertor(transformers.Resource("autobots")),
	}

	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}

	return &REST{}, nil
	// return &registry.REST{Store: store}, nil
}
