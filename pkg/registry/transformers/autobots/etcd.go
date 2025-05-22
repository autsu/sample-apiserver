package autobots

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/sample-apiserver/pkg/apis/transformers"
	"k8s.io/sample-apiserver/pkg/registry"
)

// NewREST 返回主资源的 REST 存储接口
func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*registry.REST, error) {
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

	return &registry.REST{Store: store}, nil
}
