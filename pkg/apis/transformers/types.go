package transformers

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Autobot struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   AutobotSpec
	Status AutobotStatus
}

type AutobotSpec struct {
	Mode     string `json:"mode,omitempty" protobuf:"bytes,2,opt,name=mode"`         // 模式
	Name     string `json:"name,omitempty" protobuf:"bytes,3,opt,name=name"`         // 名称
	Strength int    `json:"strength,omitempty" protobuf:"bytes,4,opt,name=strength"` // 力量
}

type AutobotStatus struct {
	HP     int    `json:"hp,omitempty" protobuf:"bytes,1,opt,name=hp"`
	Status string `json:"status,omitempty" protobuf:"bytes,2,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AutobotList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Autobot `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Decepticon struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   DecepticonSpec
	Status DecepticonStatus
}

type DecepticonSpec struct {
	Mode     string `json:"mode,omitempty" protobuf:"bytes,2,opt,name=mode"`
	Name     string `json:"name,omitempty" protobuf:"bytes,3,opt,name=name"`         // 名称
	Strength int    `json:"strength,omitempty" protobuf:"bytes,4,opt,name=strength"` // 力量
}

type DecepticonStatus struct {
	HP     int    `json:"hp,omitempty" protobuf:"bytes,1,opt,name=hp"`
	Status string `json:"status,omitempty" protobuf:"bytes,2,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DecepticonList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Decepticon `json:"items" protobuf:"bytes,2,rep,name=items"`
}
