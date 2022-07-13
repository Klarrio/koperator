//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright © 2019 Banzai Cloud
//
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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterReference) DeepCopyInto(out *ClusterReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterReference.
func (in *ClusterReference) DeepCopy() *ClusterReference {
	if in == nil {
		return nil
	}
	out := new(ClusterReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaTopic) DeepCopyInto(out *KafkaTopic) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaTopic.
func (in *KafkaTopic) DeepCopy() *KafkaTopic {
	if in == nil {
		return nil
	}
	out := new(KafkaTopic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaTopic) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaTopicList) DeepCopyInto(out *KafkaTopicList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KafkaTopic, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaTopicList.
func (in *KafkaTopicList) DeepCopy() *KafkaTopicList {
	if in == nil {
		return nil
	}
	out := new(KafkaTopicList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaTopicList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaTopicSpec) DeepCopyInto(out *KafkaTopicSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.ClusterRef = in.ClusterRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaTopicSpec.
func (in *KafkaTopicSpec) DeepCopy() *KafkaTopicSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaTopicSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaTopicStatus) DeepCopyInto(out *KafkaTopicStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaTopicStatus.
func (in *KafkaTopicStatus) DeepCopy() *KafkaTopicStatus {
	if in == nil {
		return nil
	}
	out := new(KafkaTopicStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaUser) DeepCopyInto(out *KafkaUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaUser.
func (in *KafkaUser) DeepCopy() *KafkaUser {
	if in == nil {
		return nil
	}
	out := new(KafkaUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaUserList) DeepCopyInto(out *KafkaUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KafkaUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaUserList.
func (in *KafkaUserList) DeepCopy() *KafkaUserList {
	if in == nil {
		return nil
	}
	out := new(KafkaUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaUserSpec) DeepCopyInto(out *KafkaUserSpec) {
	*out = *in
	out.ClusterRef = in.ClusterRef
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DNSNames != nil {
		in, out := &in.DNSNames, &out.DNSNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TopicGrants != nil {
		in, out := &in.TopicGrants, &out.TopicGrants
		*out = make([]UserTopicGrant, len(*in))
		copy(*out, *in)
	}
	if in.CreateCert != nil {
		in, out := &in.CreateCert, &out.CreateCert
		*out = new(bool)
		**out = **in
	}
	if in.PKIBackendSpec != nil {
		in, out := &in.PKIBackendSpec, &out.PKIBackendSpec
		*out = new(PKIBackendSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaUserSpec.
func (in *KafkaUserSpec) DeepCopy() *KafkaUserSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaUserStatus) DeepCopyInto(out *KafkaUserStatus) {
	*out = *in
	if in.ACLs != nil {
		in, out := &in.ACLs, &out.ACLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaUserStatus.
func (in *KafkaUserStatus) DeepCopy() *KafkaUserStatus {
	if in == nil {
		return nil
	}
	out := new(KafkaUserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PKIBackendSpec) DeepCopyInto(out *PKIBackendSpec) {
	*out = *in
	if in.IssuerRef != nil {
		in, out := &in.IssuerRef, &out.IssuerRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PKIBackendSpec.
func (in *PKIBackendSpec) DeepCopy() *PKIBackendSpec {
	if in == nil {
		return nil
	}
	out := new(PKIBackendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserTopicGrant) DeepCopyInto(out *UserTopicGrant) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserTopicGrant.
func (in *UserTopicGrant) DeepCopy() *UserTopicGrant {
	if in == nil {
		return nil
	}
	out := new(UserTopicGrant)
	in.DeepCopyInto(out)
	return out
}
