package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Status string

const (
	StatusAlive   Status = "alive"   // 存活
	StatusWounded Status = "wounded" // 重伤
	StatusDead    Status = "dead"    // 死亡
)

type Mode string

const (
	ModeCar   Mode = "car"   // 汽车
	ModeRobot Mode = "robot" // 机器人
	ModePlane Mode = "plane" // 飞机
)

type GenericSpec struct {
	Mode     Mode   `json:"mode,omitempty" protobuf:"bytes,2,opt,name=mode"`         // 变形模式
	Name     string `json:"name,omitempty" protobuf:"bytes,3,opt,name=name"`         // 名称
	Strength int    `json:"strength,omitempty" protobuf:"bytes,4,opt,name=strength"` // 力量
}

type GenericStatus struct {
	HP     int    `json:"hp,omitempty" protobuf:"bytes,1,opt,name=hp"`         // 生命值
	Status Status `json:"status,omitempty" protobuf:"bytes,2,opt,name=status"` // 状态
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Autobot 汽车人
type Autobot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   AutobotSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status AutobotStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type AutobotSpec struct {
	GenericSpec `json:",inline"`
}

type AutobotStatus struct {
	GenericStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AutobotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Autobot `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Decepticon 霸天虎
type Decepticon struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   DecepticonSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status DecepticonStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type DecepticonSpec struct {
	GenericSpec `json:",inline"`
}

type DecepticonStatus struct {
	GenericStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DecepticonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Decepticon `json:"items" protobuf:"bytes,2,rep,name=items"`
}
