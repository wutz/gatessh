package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var _ = metav1.Condition{}

func (in *SSHGatewayClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *SSHGatewayClass) DeepCopy() *SSHGatewayClass {
	if in == nil {
		return nil
	}
	out := new(SSHGatewayClass)
	in.DeepCopyInto(out)
	return out
}

func (in *SSHGatewayClass) DeepCopyInto(out *SSHGatewayClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

func (in *SSHGatewayClassSpec) DeepCopyInto(out *SSHGatewayClassSpec) {
	*out = *in
	if in.Description != nil {
		out.Description = new(string)
		*out.Description = *in.Description
	}
	if in.ParametersRef != nil {
		out.ParametersRef = new(ParametersReference)
		*out.ParametersRef = *in.ParametersRef
	}
}

func (in *SSHGatewayClassStatus) DeepCopyInto(out *SSHGatewayClassStatus) {
	*out = *in
	if in.Conditions != nil {
		out.Conditions = make([]metav1.Condition, len(in.Conditions))
		for i := range in.Conditions {
			in.Conditions[i].DeepCopyInto(&out.Conditions[i])
		}
	}
}

func (in *SSHGatewayClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *SSHGatewayClassList) DeepCopy() *SSHGatewayClassList {
	if in == nil {
		return nil
	}
	out := new(SSHGatewayClassList)
	in.DeepCopyInto(out)
	return out
}

func (in *SSHGatewayClassList) DeepCopyInto(out *SSHGatewayClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		out.Items = make([]SSHGatewayClass, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}
}

func (in *SSHGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *SSHGateway) DeepCopy() *SSHGateway {
	if in == nil {
		return nil
	}
	out := new(SSHGateway)
	in.DeepCopyInto(out)
	return out
}

func (in *SSHGateway) DeepCopyInto(out *SSHGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

func (in *SSHGatewaySpec) DeepCopyInto(out *SSHGatewaySpec) {
	*out = *in
	if in.Listeners != nil {
		out.Listeners = make([]SSHListener, len(in.Listeners))
		for i := range in.Listeners {
			in.Listeners[i].DeepCopyInto(&out.Listeners[i])
		}
	}
}

func (in *SSHListener) DeepCopyInto(out *SSHListener) {
	*out = *in
	if in.HostKey != nil {
		out.HostKey = new(SecretRef)
		*out.HostKey = *in.HostKey
	}
	in.Auth.DeepCopyInto(&out.Auth)
}

func (in *SSHAuthConfig) DeepCopyInto(out *SSHAuthConfig) {
	*out = *in
	if in.Methods != nil {
		out.Methods = make([]SSHAuthMethod, len(in.Methods))
		copy(out.Methods, in.Methods)
	}
	if in.MaxRetries != nil {
		out.MaxRetries = new(int32)
		*out.MaxRetries = *in.MaxRetries
	}
}

func (in *SSHGatewayStatus) DeepCopyInto(out *SSHGatewayStatus) {
	*out = *in
	if in.Conditions != nil {
		out.Conditions = make([]metav1.Condition, len(in.Conditions))
		for i := range in.Conditions {
			in.Conditions[i].DeepCopyInto(&out.Conditions[i])
		}
	}
	if in.Listeners != nil {
		out.Listeners = make([]ListenerStatus, len(in.Listeners))
		for i := range in.Listeners {
			in.Listeners[i].DeepCopyInto(&out.Listeners[i])
		}
	}
	if in.Addresses != nil {
		out.Addresses = make([]GatewayAddress, len(in.Addresses))
		copy(out.Addresses, in.Addresses)
	}
}

func (in *ListenerStatus) DeepCopyInto(out *ListenerStatus) {
	*out = *in
	if in.Conditions != nil {
		out.Conditions = make([]metav1.Condition, len(in.Conditions))
		for i := range in.Conditions {
			in.Conditions[i].DeepCopyInto(&out.Conditions[i])
		}
	}
}

func (in *SSHGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *SSHGatewayList) DeepCopy() *SSHGatewayList {
	if in == nil {
		return nil
	}
	out := new(SSHGatewayList)
	in.DeepCopyInto(out)
	return out
}

func (in *SSHGatewayList) DeepCopyInto(out *SSHGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		out.Items = make([]SSHGateway, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}
}

func (in *SSHRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *SSHRoute) DeepCopy() *SSHRoute {
	if in == nil {
		return nil
	}
	out := new(SSHRoute)
	in.DeepCopyInto(out)
	return out
}

func (in *SSHRoute) DeepCopyInto(out *SSHRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

func (in *SSHRouteSpec) DeepCopyInto(out *SSHRouteSpec) {
	*out = *in
	if in.ParentRefs != nil {
		out.ParentRefs = make([]ParentReference, len(in.ParentRefs))
		for i := range in.ParentRefs {
			in.ParentRefs[i].DeepCopyInto(&out.ParentRefs[i])
		}
	}
	if in.Rules != nil {
		out.Rules = make([]SSHRouteRule, len(in.Rules))
		for i := range in.Rules {
			in.Rules[i].DeepCopyInto(&out.Rules[i])
		}
	}
}

func (in *ParentReference) DeepCopyInto(out *ParentReference) {
	*out = *in
	if in.Port != nil {
		out.Port = new(int32)
		*out.Port = *in.Port
	}
}

func (in *SSHRouteRule) DeepCopyInto(out *SSHRouteRule) {
	*out = *in
	if in.Matches != nil {
		out.Matches = make([]SSHRouteMatch, len(in.Matches))
		for i := range in.Matches {
			in.Matches[i].DeepCopyInto(&out.Matches[i])
		}
	}
	if in.BackendRefs != nil {
		out.BackendRefs = make([]SSHBackendRef, len(in.BackendRefs))
		for i := range in.BackendRefs {
			in.BackendRefs[i].DeepCopyInto(&out.BackendRefs[i])
		}
	}
	if in.Filters != nil {
		out.Filters = make([]SSHRouteFilter, len(in.Filters))
		for i := range in.Filters {
			in.Filters[i].DeepCopyInto(&out.Filters[i])
		}
	}
}

func (in *SSHRouteMatch) DeepCopyInto(out *SSHRouteMatch) {
	*out = *in
	if in.Users != nil {
		out.Users = make([]string, len(in.Users))
		copy(out.Users, in.Users)
	}
	if in.Commands != nil {
		out.Commands = make([]CommandMatch, len(in.Commands))
		copy(out.Commands, in.Commands)
	}
}

func (in *SSHBackendRef) DeepCopyInto(out *SSHBackendRef) {
	*out = *in
	if in.Port != nil {
		out.Port = new(int32)
		*out.Port = *in.Port
	}
	if in.Weight != nil {
		out.Weight = new(int32)
		*out.Weight = *in.Weight
	}
}

func (in *SSHRouteFilter) DeepCopyInto(out *SSHRouteFilter) {
	*out = *in
	if in.CommandTemplate != nil {
		out.CommandTemplate = new(CommandTemplateFilter)
		in.CommandTemplate.DeepCopyInto(out.CommandTemplate)
	}
	if in.RecordSession != nil {
		out.RecordSession = new(RecordSessionFilter)
		*out.RecordSession = *in.RecordSession
	}
}

func (in *CommandTemplateFilter) DeepCopyInto(out *CommandTemplateFilter) {
	*out = *in
	if in.AllowedCommands != nil {
		out.AllowedCommands = make([]string, len(in.AllowedCommands))
		copy(out.AllowedCommands, in.AllowedCommands)
	}
	if in.DeniedCommands != nil {
		out.DeniedCommands = make([]string, len(in.DeniedCommands))
		copy(out.DeniedCommands, in.DeniedCommands)
	}
}

func (in *SSHRouteStatus) DeepCopyInto(out *SSHRouteStatus) {
	*out = *in
	if in.Parents != nil {
		out.Parents = make([]RouteParentStatus, len(in.Parents))
		for i := range in.Parents {
			in.Parents[i].DeepCopyInto(&out.Parents[i])
		}
	}
}

func (in *RouteParentStatus) DeepCopyInto(out *RouteParentStatus) {
	*out = *in
	in.ParentRef.DeepCopyInto(&out.ParentRef)
	if in.Conditions != nil {
		out.Conditions = make([]metav1.Condition, len(in.Conditions))
		for i := range in.Conditions {
			in.Conditions[i].DeepCopyInto(&out.Conditions[i])
		}
	}
}

func (in *SSHRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *SSHRouteList) DeepCopy() *SSHRouteList {
	if in == nil {
		return nil
	}
	out := new(SSHRouteList)
	in.DeepCopyInto(out)
	return out
}

func (in *SSHRouteList) DeepCopyInto(out *SSHRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		out.Items = make([]SSHRoute, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}
}
