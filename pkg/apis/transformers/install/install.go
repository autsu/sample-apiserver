package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/sample-apiserver/pkg/apis/transformers"
	"k8s.io/sample-apiserver/pkg/apis/transformers/v1alpha1"
)

func Install(scheme *runtime.Scheme) {
	utilruntime.Must(transformers.AddToScheme(scheme))
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must(scheme.SetVersionPriority(v1alpha1.SchemeGroupVersion))
}