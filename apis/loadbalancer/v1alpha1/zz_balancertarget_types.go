// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BalancerTargetInitParameters struct {

	// IP address for an IP Target. Required if
	// type is ip.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// Label Selector selecting targets
	// for this Load Balancer. Required if type is label_selector.
	LabelSelector *string `json:"labelSelector,omitempty" tf:"label_selector,omitempty"`

	// ID of the Load Balancer to which
	// the target gets attached.
	LoadBalancerID *float64 `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// ID of the server which should be a
	// target for this Load Balancer. Required if type is server
	ServerID *float64 `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Type of the target. Possible values
	// server, label_selector, ip.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// use the private IP to connect to
	// Load Balancer targets. Only allowed if type is server or
	// label_selector.
	UsePrivateIP *bool `json:"usePrivateIp,omitempty" tf:"use_private_ip,omitempty"`
}

type BalancerTargetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP address for an IP Target. Required if
	// type is ip.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// Label Selector selecting targets
	// for this Load Balancer. Required if type is label_selector.
	LabelSelector *string `json:"labelSelector,omitempty" tf:"label_selector,omitempty"`

	// ID of the Load Balancer to which
	// the target gets attached.
	LoadBalancerID *float64 `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// ID of the server which should be a
	// target for this Load Balancer. Required if type is server
	ServerID *float64 `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Type of the target. Possible values
	// server, label_selector, ip.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// use the private IP to connect to
	// Load Balancer targets. Only allowed if type is server or
	// label_selector.
	UsePrivateIP *bool `json:"usePrivateIp,omitempty" tf:"use_private_ip,omitempty"`
}

type BalancerTargetParameters struct {

	// IP address for an IP Target. Required if
	// type is ip.
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// Label Selector selecting targets
	// for this Load Balancer. Required if type is label_selector.
	// +kubebuilder:validation:Optional
	LabelSelector *string `json:"labelSelector,omitempty" tf:"label_selector,omitempty"`

	// ID of the Load Balancer to which
	// the target gets attached.
	// +kubebuilder:validation:Optional
	LoadBalancerID *float64 `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// ID of the server which should be a
	// target for this Load Balancer. Required if type is server
	// +kubebuilder:validation:Optional
	ServerID *float64 `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Type of the target. Possible values
	// server, label_selector, ip.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// use the private IP to connect to
	// Load Balancer targets. Only allowed if type is server or
	// label_selector.
	// +kubebuilder:validation:Optional
	UsePrivateIP *bool `json:"usePrivateIp,omitempty" tf:"use_private_ip,omitempty"`
}

// BalancerTargetSpec defines the desired state of BalancerTarget
type BalancerTargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BalancerTargetParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider BalancerTargetInitParameters `json:"initProvider,omitempty"`
}

// BalancerTargetStatus defines the observed state of BalancerTarget.
type BalancerTargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BalancerTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BalancerTarget is the Schema for the BalancerTargets API. Adds a target to a Hetzner Cloud Load Balancer.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hc}
type BalancerTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.loadBalancerId) || (has(self.initProvider) && has(self.initProvider.loadBalancerId))",message="spec.forProvider.loadBalancerId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   BalancerTargetSpec   `json:"spec"`
	Status BalancerTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BalancerTargetList contains a list of BalancerTargets
type BalancerTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BalancerTarget `json:"items"`
}

// Repository type metadata.
var (
	BalancerTarget_Kind             = "BalancerTarget"
	BalancerTarget_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BalancerTarget_Kind}.String()
	BalancerTarget_KindAPIVersion   = BalancerTarget_Kind + "." + CRDGroupVersion.String()
	BalancerTarget_GroupVersionKind = CRDGroupVersion.WithKind(BalancerTarget_Kind)
)

func init() {
	SchemeBuilder.Register(&BalancerTarget{}, &BalancerTargetList{})
}
