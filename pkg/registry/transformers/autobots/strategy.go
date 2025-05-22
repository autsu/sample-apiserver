package autobots

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/sample-apiserver/pkg/apis/transformers"
)

var (
	_ rest.RESTCreateStrategy = &autobotsStrategy{}
	_ rest.RESTUpdateStrategy = &autobotsStrategy{}
	_ rest.RESTDeleteStrategy = &autobotsStrategy{}
	_ rest.Creater            = &autobotsStrategy{}
	_ rest.Updater            = &autobotsStrategy{}
	_ rest.Getter             = &autobotsStrategy{}
)

type autobotsStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

// Get implements rest.Getter.
func (o *autobotsStrategy) Get(ctx context.Context, name string, options *v1.GetOptions) (runtime.Object, error) {
	panic("unimplemented")
}

// Update implements rest.Updater.
func (o *autobotsStrategy) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *v1.UpdateOptions) (runtime.Object, bool, error) {
	panic("unimplemented")
}

// Create implements rest.Creater.
func (o *autobotsStrategy) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *v1.CreateOptions) (runtime.Object, error) {
	panic("unimplemented")
}

// New implements rest.Creater.
func (o *autobotsStrategy) New() runtime.Object {
	panic("unimplemented")
}

// AllowCreateOnUpdate implements rest.RESTUpdateStrategy.
// return true 表示：如果要更新的对象不存在，则直接创建
func (o *autobotsStrategy) AllowCreateOnUpdate() bool {
	return false
}

// AllowUnconditionalUpdate implements rest.RESTUpdateStrategy.
// 是否允许无条件更新，也就是跳过版本号校验的乐观锁检查
func (o *autobotsStrategy) AllowUnconditionalUpdate() bool {
	return false
}

// PrepareForUpdate implements rest.RESTUpdateStrategy.
// 在更新之前，对对象进行一些预处理
func (o *autobotsStrategy) PrepareForUpdate(ctx context.Context, obj runtime.Object, old runtime.Object) {
}

// ValidateUpdate implements rest.RESTUpdateStrategy.
func (o *autobotsStrategy) ValidateUpdate(ctx context.Context, obj runtime.Object, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

// WarningsOnUpdate implements rest.RESTUpdateStrategy.
// 返回在对象更新期间产生的任何非致命性警告信息
func (o *autobotsStrategy) WarningsOnUpdate(ctx context.Context, obj runtime.Object, old runtime.Object) []string {
	return nil
}

// Canonicalize implements rest.RESTCreateStrategy.
// 在对象持久化存储之前，将其转换为“规范”形式。这通常用于清理或标准化对象字段，确保一致性。例如，某些字段可能需要特定的格式或默认值。
func (o *autobotsStrategy) Canonicalize(obj runtime.Object) {
}

// NamespaceScoped implements rest.RESTCreateStrategy.
// 该资源是否是 Namespace 级别的，是的话返回 true
func (o *autobotsStrategy) NamespaceScoped() bool {
	return true
}

// PrepareForCreate implements rest.RESTCreateStrategy.
func (o *autobotsStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

// Validate implements rest.RESTCreateStrategy.
func (o *autobotsStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

// WarningsOnCreate implements rest.RESTCreateStrategy.
func (o *autobotsStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return nil
}

func NewStrategy(typer runtime.ObjectTyper) *autobotsStrategy {
	return &autobotsStrategy{typer, names.SimpleNameGenerator}
}

func SelectableFields(obj *transformers.Autobot) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

// MatchAutobot implements filtering based on labels.
func MatchAutobot(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// GetAttrs returns labels and fields of a given object for filtering purposes.
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	apiserver, ok := obj.(*transformers.Autobot)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a Autobot")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), nil
}
