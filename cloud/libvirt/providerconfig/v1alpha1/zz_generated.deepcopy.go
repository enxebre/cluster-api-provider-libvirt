// +build !ignore_autogenerated

// Copyright © 2018 The Kubernetes Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibvirtClusterProviderConfig) DeepCopyInto(out *LibvirtClusterProviderConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibvirtClusterProviderConfig.
func (in *LibvirtClusterProviderConfig) DeepCopy() *LibvirtClusterProviderConfig {
	if in == nil {
		return nil
	}
	out := new(LibvirtClusterProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LibvirtClusterProviderConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibvirtClusterProviderStatus) DeepCopyInto(out *LibvirtClusterProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibvirtClusterProviderStatus.
func (in *LibvirtClusterProviderStatus) DeepCopy() *LibvirtClusterProviderStatus {
	if in == nil {
		return nil
	}
	out := new(LibvirtClusterProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LibvirtClusterProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibvirtMachineProviderCondition) DeepCopyInto(out *LibvirtMachineProviderCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibvirtMachineProviderCondition.
func (in *LibvirtMachineProviderCondition) DeepCopy() *LibvirtMachineProviderCondition {
	if in == nil {
		return nil
	}
	out := new(LibvirtMachineProviderCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibvirtMachineProviderConfig) DeepCopyInto(out *LibvirtMachineProviderConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(Volume)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibvirtMachineProviderConfig.
func (in *LibvirtMachineProviderConfig) DeepCopy() *LibvirtMachineProviderConfig {
	if in == nil {
		return nil
	}
	out := new(LibvirtMachineProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LibvirtMachineProviderConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibvirtMachineProviderStatus) DeepCopyInto(out *LibvirtMachineProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.InstanceState != nil {
		in, out := &in.InstanceState, &out.InstanceState
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]LibvirtMachineProviderCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibvirtMachineProviderStatus.
func (in *LibvirtMachineProviderStatus) DeepCopy() *LibvirtMachineProviderStatus {
	if in == nil {
		return nil
	}
	out := new(LibvirtMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LibvirtMachineProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Volume.
func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}
