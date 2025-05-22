package v1alpha1

import (
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/sample-apiserver/pkg/apis/transformers"
)

func Convert_v1alpha1_AutobotSpec_To_transformers_AutobotSpec(in *AutobotSpec, out *transformers.AutobotSpec, s conversion.Scope) error {
	out.Mode = string(in.GenericSpec.Mode)
	out.Name = in.GenericSpec.Name
	out.Strength = in.GenericSpec.Strength
	return nil
}

func Convert_transformers_AutobotSpec_To_v1alpha1_AutobotSpec(in *transformers.AutobotSpec, out *AutobotSpec, s conversion.Scope) error {
	out.GenericSpec.Mode = Mode(in.Mode)
	out.GenericSpec.Name = in.Name
	out.GenericSpec.Strength = in.Strength
	return nil
}

func Convert_v1alpha1_DecepticonSpec_To_transformers_DecepticonSpec(in *DecepticonSpec, out *transformers.DecepticonSpec, s conversion.Scope) error {
	out.Mode = string(in.GenericSpec.Mode)
	out.Name = in.GenericSpec.Name
	out.Strength = in.GenericSpec.Strength
	return nil
}

func Convert_transformers_DecepticonSpec_To_v1alpha1_DecepticonSpec(in *transformers.DecepticonSpec, out *DecepticonSpec, s conversion.Scope) error {
	out.GenericSpec.Mode = Mode(in.Mode)
	out.GenericSpec.Name = in.Name
	out.GenericSpec.Strength = in.Strength
	return nil
}

func Convert_v1alpha1_AutobotStatus_To_transformers_AutobotStatus(in *AutobotStatus, out *transformers.AutobotStatus, s conversion.Scope) error {
	out.HP = in.HP
	out.Status = string(in.Status)
	return nil
}

func Convert_transformers_AutobotStatus_To_v1alpha1_AutobotStatus(in *transformers.AutobotStatus, out *AutobotStatus, s conversion.Scope) error {
	out.HP = in.HP
	out.Status = Status(in.Status)
	return nil
}

func Convert_v1alpha1_DecepticonStatus_To_transformers_DecepticonStatus(in *DecepticonStatus, out *transformers.DecepticonStatus, s conversion.Scope) error {
	out.HP = in.HP
	out.Status = string(in.Status)
	return nil
}

func Convert_transformers_DecepticonStatus_To_v1alpha1_DecepticonStatus(in *transformers.DecepticonStatus, out *DecepticonStatus, s conversion.Scope) error {
	out.HP = in.HP
	out.Status = Status(in.Status)
	return nil
}

