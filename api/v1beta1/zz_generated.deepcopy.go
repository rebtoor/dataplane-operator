//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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
	"github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	"github.com/openstack-k8s-operators/lib-common/modules/storage"
	"github.com/openstack-k8s-operators/openstack-ansibleee-operator/api/v1alpha1"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleEESpec) DeepCopyInto(out *AnsibleEESpec) {
	*out = *in
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraMounts != nil {
		in, out := &in.ExtraMounts, &out.ExtraMounts
		*out = make([]storage.VolMounts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleEESpec.
func (in *AnsibleEESpec) DeepCopy() *AnsibleEESpec {
	if in == nil {
		return nil
	}
	out := new(AnsibleEESpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployStrategySection) DeepCopyInto(out *DeployStrategySection) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployStrategySection.
func (in *DeployStrategySection) DeepCopy() *DeployStrategySection {
	if in == nil {
		return nil
	}
	out := new(DeployStrategySection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfigSection) DeepCopyInto(out *NetworkConfigSection) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfigSection.
func (in *NetworkConfigSection) DeepCopy() *NetworkConfigSection {
	if in == nil {
		return nil
	}
	out := new(NetworkConfigSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworksSection) DeepCopyInto(out *NetworksSection) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworksSection.
func (in *NetworksSection) DeepCopy() *NetworksSection {
	if in == nil {
		return nil
	}
	out := new(NetworksSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSection) DeepCopyInto(out *NodeSection) {
	*out = *in
	out.NetworkConfig = in.NetworkConfig
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]NetworksSection, len(*in))
		copy(*out, *in)
	}
	if in.ExtraMounts != nil {
		in, out := &in.ExtraMounts, &out.ExtraMounts
		*out = make([]storage.VolMounts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.NetworkData != nil {
		in, out := &in.NetworkData, &out.NetworkData
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.Nova != nil {
		in, out := &in.Nova, &out.Nova
		*out = new(NovaTemplate)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSection.
func (in *NodeSection) DeepCopy() *NodeSection {
	if in == nil {
		return nil
	}
	out := new(NodeSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NovaTemplate) DeepCopyInto(out *NovaTemplate) {
	*out = *in
	if in.Deploy != nil {
		in, out := &in.Deploy, &out.Deploy
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NovaTemplate.
func (in *NovaTemplate) DeepCopy() *NovaTemplate {
	if in == nil {
		return nil
	}
	out := new(NovaTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlane) DeepCopyInto(out *OpenStackDataPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlane.
func (in *OpenStackDataPlane) DeepCopy() *OpenStackDataPlane {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackDataPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneList) DeepCopyInto(out *OpenStackDataPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackDataPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneList.
func (in *OpenStackDataPlaneList) DeepCopy() *OpenStackDataPlaneList {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackDataPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneNode) DeepCopyInto(out *OpenStackDataPlaneNode) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneNode.
func (in *OpenStackDataPlaneNode) DeepCopy() *OpenStackDataPlaneNode {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackDataPlaneNode) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneNodeList) DeepCopyInto(out *OpenStackDataPlaneNodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackDataPlaneNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneNodeList.
func (in *OpenStackDataPlaneNodeList) DeepCopy() *OpenStackDataPlaneNodeList {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneNodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackDataPlaneNodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneNodeSpec) DeepCopyInto(out *OpenStackDataPlaneNodeSpec) {
	*out = *in
	in.Node.DeepCopyInto(&out.Node)
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.DeployStrategy = in.DeployStrategy
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneNodeSpec.
func (in *OpenStackDataPlaneNodeSpec) DeepCopy() *OpenStackDataPlaneNodeSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneRole) DeepCopyInto(out *OpenStackDataPlaneRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneRole.
func (in *OpenStackDataPlaneRole) DeepCopy() *OpenStackDataPlaneRole {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackDataPlaneRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneRoleList) DeepCopyInto(out *OpenStackDataPlaneRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackDataPlaneRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneRoleList.
func (in *OpenStackDataPlaneRoleList) DeepCopy() *OpenStackDataPlaneRoleList {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackDataPlaneRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneRoleSpec) DeepCopyInto(out *OpenStackDataPlaneRoleSpec) {
	*out = *in
	in.BaremetalSetTemplate.DeepCopyInto(&out.BaremetalSetTemplate)
	in.NodeTemplate.DeepCopyInto(&out.NodeTemplate)
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.DeployStrategy = in.DeployStrategy
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneRoleSpec.
func (in *OpenStackDataPlaneRoleSpec) DeepCopy() *OpenStackDataPlaneRoleSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneService) DeepCopyInto(out *OpenStackDataPlaneService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneService.
func (in *OpenStackDataPlaneService) DeepCopy() *OpenStackDataPlaneService {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackDataPlaneService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneServiceList) DeepCopyInto(out *OpenStackDataPlaneServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackDataPlaneService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneServiceList.
func (in *OpenStackDataPlaneServiceList) DeepCopy() *OpenStackDataPlaneServiceList {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackDataPlaneServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneServiceSpec) DeepCopyInto(out *OpenStackDataPlaneServiceSpec) {
	*out = *in
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(v1alpha1.Role)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneServiceSpec.
func (in *OpenStackDataPlaneServiceSpec) DeepCopy() *OpenStackDataPlaneServiceSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneServiceStatus) DeepCopyInto(out *OpenStackDataPlaneServiceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneServiceStatus.
func (in *OpenStackDataPlaneServiceStatus) DeepCopy() *OpenStackDataPlaneServiceStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneSpec) DeepCopyInto(out *OpenStackDataPlaneSpec) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make(map[string]OpenStackDataPlaneNodeSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make(map[string]OpenStackDataPlaneRoleSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	out.DeployStrategy = in.DeployStrategy
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneSpec.
func (in *OpenStackDataPlaneSpec) DeepCopy() *OpenStackDataPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneStatus) DeepCopyInto(out *OpenStackDataPlaneStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneStatus.
func (in *OpenStackDataPlaneStatus) DeepCopy() *OpenStackDataPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneStatus)
	in.DeepCopyInto(out)
	return out
}
