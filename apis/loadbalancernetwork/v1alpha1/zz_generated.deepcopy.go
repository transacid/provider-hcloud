//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BalancerNetwork) DeepCopyInto(out *BalancerNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BalancerNetwork.
func (in *BalancerNetwork) DeepCopy() *BalancerNetwork {
	if in == nil {
		return nil
	}
	out := new(BalancerNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BalancerNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BalancerNetworkInitParameters) DeepCopyInto(out *BalancerNetworkInitParameters) {
	*out = *in
	if in.EnablePublicInterface != nil {
		in, out := &in.EnablePublicInterface, &out.EnablePublicInterface
		*out = new(bool)
		**out = **in
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.LoadBalancerID != nil {
		in, out := &in.LoadBalancerID, &out.LoadBalancerID
		*out = new(float64)
		**out = **in
	}
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(float64)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BalancerNetworkInitParameters.
func (in *BalancerNetworkInitParameters) DeepCopy() *BalancerNetworkInitParameters {
	if in == nil {
		return nil
	}
	out := new(BalancerNetworkInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BalancerNetworkList) DeepCopyInto(out *BalancerNetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BalancerNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BalancerNetworkList.
func (in *BalancerNetworkList) DeepCopy() *BalancerNetworkList {
	if in == nil {
		return nil
	}
	out := new(BalancerNetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BalancerNetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BalancerNetworkObservation) DeepCopyInto(out *BalancerNetworkObservation) {
	*out = *in
	if in.EnablePublicInterface != nil {
		in, out := &in.EnablePublicInterface, &out.EnablePublicInterface
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.LoadBalancerID != nil {
		in, out := &in.LoadBalancerID, &out.LoadBalancerID
		*out = new(float64)
		**out = **in
	}
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(float64)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BalancerNetworkObservation.
func (in *BalancerNetworkObservation) DeepCopy() *BalancerNetworkObservation {
	if in == nil {
		return nil
	}
	out := new(BalancerNetworkObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BalancerNetworkParameters) DeepCopyInto(out *BalancerNetworkParameters) {
	*out = *in
	if in.EnablePublicInterface != nil {
		in, out := &in.EnablePublicInterface, &out.EnablePublicInterface
		*out = new(bool)
		**out = **in
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.LoadBalancerID != nil {
		in, out := &in.LoadBalancerID, &out.LoadBalancerID
		*out = new(float64)
		**out = **in
	}
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(float64)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BalancerNetworkParameters.
func (in *BalancerNetworkParameters) DeepCopy() *BalancerNetworkParameters {
	if in == nil {
		return nil
	}
	out := new(BalancerNetworkParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BalancerNetworkSpec) DeepCopyInto(out *BalancerNetworkSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BalancerNetworkSpec.
func (in *BalancerNetworkSpec) DeepCopy() *BalancerNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(BalancerNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BalancerNetworkStatus) DeepCopyInto(out *BalancerNetworkStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BalancerNetworkStatus.
func (in *BalancerNetworkStatus) DeepCopy() *BalancerNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(BalancerNetworkStatus)
	in.DeepCopyInto(out)
	return out
}