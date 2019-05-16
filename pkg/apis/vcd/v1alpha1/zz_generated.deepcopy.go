// +build !ignore_autogenerated

/*
Copyright 2019 Tyler French.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	common "sigs.k8s.io/cluster-api/pkg/apis/cluster/common"
	clusterv1alpha1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcdClusterProviderSpec) DeepCopyInto(out *VcdClusterProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcdClusterProviderSpec.
func (in *VcdClusterProviderSpec) DeepCopy() *VcdClusterProviderSpec {
	if in == nil {
		return nil
	}
	out := new(VcdClusterProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VcdClusterProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcdClusterProviderSpecList) DeepCopyInto(out *VcdClusterProviderSpecList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VcdClusterProviderSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcdClusterProviderSpecList.
func (in *VcdClusterProviderSpecList) DeepCopy() *VcdClusterProviderSpecList {
	if in == nil {
		return nil
	}
	out := new(VcdClusterProviderSpecList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VcdClusterProviderSpecList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcdClusterProviderSpecSpec) DeepCopyInto(out *VcdClusterProviderSpecSpec) {
	*out = *in
	in.ClusterNetwork.DeepCopyInto(&out.ClusterNetwork)
	in.ProviderSpec.DeepCopyInto(&out.ProviderSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcdClusterProviderSpecSpec.
func (in *VcdClusterProviderSpecSpec) DeepCopy() *VcdClusterProviderSpecSpec {
	if in == nil {
		return nil
	}
	out := new(VcdClusterProviderSpecSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcdClusterProviderSpecStatus) DeepCopyInto(out *VcdClusterProviderSpecStatus) {
	*out = *in
	if in.APIEndpoints != nil {
		in, out := &in.APIEndpoints, &out.APIEndpoints
		*out = make([]clusterv1alpha1.APIEndpoint, len(*in))
		copy(*out, *in)
	}
	if in.ProviderStatus != nil {
		in, out := &in.ProviderStatus, &out.ProviderStatus
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcdClusterProviderSpecStatus.
func (in *VcdClusterProviderSpecStatus) DeepCopy() *VcdClusterProviderSpecStatus {
	if in == nil {
		return nil
	}
	out := new(VcdClusterProviderSpecStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcdClusterProviderStatus) DeepCopyInto(out *VcdClusterProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcdClusterProviderStatus.
func (in *VcdClusterProviderStatus) DeepCopy() *VcdClusterProviderStatus {
	if in == nil {
		return nil
	}
	out := new(VcdClusterProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VcdClusterProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcdClusterProviderStatusList) DeepCopyInto(out *VcdClusterProviderStatusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VcdClusterProviderStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcdClusterProviderStatusList.
func (in *VcdClusterProviderStatusList) DeepCopy() *VcdClusterProviderStatusList {
	if in == nil {
		return nil
	}
	out := new(VcdClusterProviderStatusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VcdClusterProviderStatusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcdClusterProviderStatusSpec) DeepCopyInto(out *VcdClusterProviderStatusSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcdClusterProviderStatusSpec.
func (in *VcdClusterProviderStatusSpec) DeepCopy() *VcdClusterProviderStatusSpec {
	if in == nil {
		return nil
	}
	out := new(VcdClusterProviderStatusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcdClusterProviderStatusStatus) DeepCopyInto(out *VcdClusterProviderStatusStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcdClusterProviderStatusStatus.
func (in *VcdClusterProviderStatusStatus) DeepCopy() *VcdClusterProviderStatusStatus {
	if in == nil {
		return nil
	}
	out := new(VcdClusterProviderStatusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcdMachineProviderSpec) DeepCopyInto(out *VcdMachineProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcdMachineProviderSpec.
func (in *VcdMachineProviderSpec) DeepCopy() *VcdMachineProviderSpec {
	if in == nil {
		return nil
	}
	out := new(VcdMachineProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VcdMachineProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcdMachineProviderSpecList) DeepCopyInto(out *VcdMachineProviderSpecList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VcdMachineProviderSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcdMachineProviderSpecList.
func (in *VcdMachineProviderSpecList) DeepCopy() *VcdMachineProviderSpecList {
	if in == nil {
		return nil
	}
	out := new(VcdMachineProviderSpecList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VcdMachineProviderSpecList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcdMachineProviderSpecSpec) DeepCopyInto(out *VcdMachineProviderSpecSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]v1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ProviderSpec.DeepCopyInto(&out.ProviderSpec)
	out.Versions = in.Versions
	if in.ConfigSource != nil {
		in, out := &in.ConfigSource, &out.ConfigSource
		*out = new(v1.NodeConfigSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcdMachineProviderSpecSpec.
func (in *VcdMachineProviderSpecSpec) DeepCopy() *VcdMachineProviderSpecSpec {
	if in == nil {
		return nil
	}
	out := new(VcdMachineProviderSpecSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcdMachineProviderSpecStatus) DeepCopyInto(out *VcdMachineProviderSpecStatus) {
	*out = *in
	if in.NodeRef != nil {
		in, out := &in.NodeRef, &out.NodeRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.LastUpdated != nil {
		in, out := &in.LastUpdated, &out.LastUpdated
		*out = (*in).DeepCopy()
	}
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = new(clusterv1alpha1.MachineVersionInfo)
		**out = **in
	}
	if in.ErrorReason != nil {
		in, out := &in.ErrorReason, &out.ErrorReason
		*out = new(common.MachineStatusError)
		**out = **in
	}
	if in.ErrorMessage != nil {
		in, out := &in.ErrorMessage, &out.ErrorMessage
		*out = new(string)
		**out = **in
	}
	if in.ProviderStatus != nil {
		in, out := &in.ProviderStatus, &out.ProviderStatus
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1.NodeAddress, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.NodeCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastOperation != nil {
		in, out := &in.LastOperation, &out.LastOperation
		*out = new(clusterv1alpha1.LastOperation)
		(*in).DeepCopyInto(*out)
	}
	if in.Phase != nil {
		in, out := &in.Phase, &out.Phase
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcdMachineProviderSpecStatus.
func (in *VcdMachineProviderSpecStatus) DeepCopy() *VcdMachineProviderSpecStatus {
	if in == nil {
		return nil
	}
	out := new(VcdMachineProviderSpecStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcdMachineProviderStatus) DeepCopyInto(out *VcdMachineProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcdMachineProviderStatus.
func (in *VcdMachineProviderStatus) DeepCopy() *VcdMachineProviderStatus {
	if in == nil {
		return nil
	}
	out := new(VcdMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VcdMachineProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcdMachineProviderStatusList) DeepCopyInto(out *VcdMachineProviderStatusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VcdMachineProviderStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcdMachineProviderStatusList.
func (in *VcdMachineProviderStatusList) DeepCopy() *VcdMachineProviderStatusList {
	if in == nil {
		return nil
	}
	out := new(VcdMachineProviderStatusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VcdMachineProviderStatusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcdMachineProviderStatusSpec) DeepCopyInto(out *VcdMachineProviderStatusSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcdMachineProviderStatusSpec.
func (in *VcdMachineProviderStatusSpec) DeepCopy() *VcdMachineProviderStatusSpec {
	if in == nil {
		return nil
	}
	out := new(VcdMachineProviderStatusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcdMachineProviderStatusStatus) DeepCopyInto(out *VcdMachineProviderStatusStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcdMachineProviderStatusStatus.
func (in *VcdMachineProviderStatusStatus) DeepCopy() *VcdMachineProviderStatusStatus {
	if in == nil {
		return nil
	}
	out := new(VcdMachineProviderStatusStatus)
	in.DeepCopyInto(out)
	return out
}
