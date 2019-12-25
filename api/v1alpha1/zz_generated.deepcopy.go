// +build !ignore_autogenerated

/*

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

// autogenerated by controller-gen object, do not modify manually

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinimalAzure) DeepCopyInto(out *MinimalAzure) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinimalAzure.
func (in *MinimalAzure) DeepCopy() *MinimalAzure {
	if in == nil {
		return nil
	}
	out := new(MinimalAzure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MinimalAzure) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinimalAzureList) DeepCopyInto(out *MinimalAzureList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MinimalAzure, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinimalAzureList.
func (in *MinimalAzureList) DeepCopy() *MinimalAzureList {
	if in == nil {
		return nil
	}
	out := new(MinimalAzureList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MinimalAzureList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinimalAzureSpec) DeepCopyInto(out *MinimalAzureSpec) {
	*out = *in
	out.CredentialsSecretRef = in.CredentialsSecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinimalAzureSpec.
func (in *MinimalAzureSpec) DeepCopy() *MinimalAzureSpec {
	if in == nil {
		return nil
	}
	out := new(MinimalAzureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinimalAzureStatus) DeepCopyInto(out *MinimalAzureStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinimalAzureStatus.
func (in *MinimalAzureStatus) DeepCopy() *MinimalAzureStatus {
	if in == nil {
		return nil
	}
	out := new(MinimalAzureStatus)
	in.DeepCopyInto(out)
	return out
}
