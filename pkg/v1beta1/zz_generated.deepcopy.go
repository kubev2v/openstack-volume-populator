//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Red Hat Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackVolumePopulator) DeepCopyInto(out *OpenstackVolumePopulator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackVolumePopulator.
func (in *OpenstackVolumePopulator) DeepCopy() *OpenstackVolumePopulator {
	if in == nil {
		return nil
	}
	out := new(OpenstackVolumePopulator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenstackVolumePopulator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackVolumePopulatorList) DeepCopyInto(out *OpenstackVolumePopulatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenstackVolumePopulator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackVolumePopulatorList.
func (in *OpenstackVolumePopulatorList) DeepCopy() *OpenstackVolumePopulatorList {
	if in == nil {
		return nil
	}
	out := new(OpenstackVolumePopulatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenstackVolumePopulatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackVolumePopulatorSpec) DeepCopyInto(out *OpenstackVolumePopulatorSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackVolumePopulatorSpec.
func (in *OpenstackVolumePopulatorSpec) DeepCopy() *OpenstackVolumePopulatorSpec {
	if in == nil {
		return nil
	}
	out := new(OpenstackVolumePopulatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackVolumePopulatorStatus) DeepCopyInto(out *OpenstackVolumePopulatorStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackVolumePopulatorStatus.
func (in *OpenstackVolumePopulatorStatus) DeepCopy() *OpenstackVolumePopulatorStatus {
	if in == nil {
		return nil
	}
	out := new(OpenstackVolumePopulatorStatus)
	in.DeepCopyInto(out)
	return out
}
