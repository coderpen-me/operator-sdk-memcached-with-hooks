//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataStruct.
func (in *DataStruct) DeepCopy() *DataStruct {
	if in == nil {
		return nil
	}
	out := new(DataStruct)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Memcached) DeepCopyInto(out *Memcached) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Memcached.
func (in *Memcached) DeepCopy() *Memcached {
	if in == nil {
		return nil
	}
	out := new(Memcached)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Memcached) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcachedList) DeepCopyInto(out *MemcachedList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Memcached, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcachedList.
func (in *MemcachedList) DeepCopy() *MemcachedList {
	if in == nil {
		return nil
	}
	out := new(MemcachedList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MemcachedList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcachedSpec) DeepCopyInto(out *MemcachedSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]v1.JSON, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcachedSpec.
func (in *MemcachedSpec) DeepCopy() *MemcachedSpec {
	if in == nil {
		return nil
	}
	out := new(MemcachedSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcachedStatus) DeepCopyInto(out *MemcachedStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcachedStatus.
func (in *MemcachedStatus) DeepCopy() *MemcachedStatus {
	if in == nil {
		return nil
	}
	out := new(MemcachedStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MismatchedField) DeepCopyInto(out *MismatchedField) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MismatchedField.
func (in *MismatchedField) DeepCopy() *MismatchedField {
	if in == nil {
		return nil
	}
	out := new(MismatchedField)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MissingField) DeepCopyInto(out *MissingField) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MissingField.
func (in *MissingField) DeepCopy() *MissingField {
	if in == nil {
		return nil
	}
	out := new(MissingField)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugin1) DeepCopyInto(out *Plugin1) {
	*out = *in
	out.Plugin1Name = in.Plugin1Name
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugin1.
func (in *Plugin1) DeepCopy() *Plugin1 {
	if in == nil {
		return nil
	}
	out := new(Plugin1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugin2) DeepCopyInto(out *Plugin2) {
	*out = *in
	out.Plugin2Name = in.Plugin2Name
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugin2.
func (in *Plugin2) DeepCopy() *Plugin2 {
	if in == nil {
		return nil
	}
	out := new(Plugin2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatorResponse) DeepCopyInto(out *ValidatorResponse) {
	*out = *in
	in.Data.DeepCopyInto(&out.Data)
	if in.PluginError != nil {
		in, out := &in.PluginError, &out.PluginError
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatorResponse.
func (in *ValidatorResponse) DeepCopy() *ValidatorResponse {
	if in == nil {
		return nil
	}
	out := new(ValidatorResponse)
	in.DeepCopyInto(out)
	return out
}
