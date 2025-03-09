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

type RdnsInitParameters struct {

	// The DNS address the ip_address should resolve to.
	DNSPtr *string `json:"dnsPtr,omitempty" tf:"dns_ptr,omitempty"`

	// The Floating IP the ip_address belongs to.
	FloatingIPID *float64 `json:"floatingIpId,omitempty" tf:"floating_ip_id,omitempty"`

	// The IP address that should point to dns_ptr.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The Load Balancer the ip_address belongs to.
	LoadBalancerID *float64 `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// The Primary IP the ip_address belongs to.
	PrimaryIPID *float64 `json:"primaryIpId,omitempty" tf:"primary_ip_id,omitempty"`

	// The server the ip_address belongs to.
	ServerID *float64 `json:"serverId,omitempty" tf:"server_id,omitempty"`
}

type RdnsObservation struct {

	// The DNS address the ip_address should resolve to.
	DNSPtr *string `json:"dnsPtr,omitempty" tf:"dns_ptr,omitempty"`

	// The Floating IP the ip_address belongs to.
	FloatingIPID *float64 `json:"floatingIpId,omitempty" tf:"floating_ip_id,omitempty"`

	// (int) Unique ID of the Reverse DNS Entry.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP address that should point to dns_ptr.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The Load Balancer the ip_address belongs to.
	LoadBalancerID *float64 `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// The Primary IP the ip_address belongs to.
	PrimaryIPID *float64 `json:"primaryIpId,omitempty" tf:"primary_ip_id,omitempty"`

	// The server the ip_address belongs to.
	ServerID *float64 `json:"serverId,omitempty" tf:"server_id,omitempty"`
}

type RdnsParameters struct {

	// The DNS address the ip_address should resolve to.
	// +kubebuilder:validation:Optional
	DNSPtr *string `json:"dnsPtr,omitempty" tf:"dns_ptr,omitempty"`

	// The Floating IP the ip_address belongs to.
	// +kubebuilder:validation:Optional
	FloatingIPID *float64 `json:"floatingIpId,omitempty" tf:"floating_ip_id,omitempty"`

	// The IP address that should point to dns_ptr.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The Load Balancer the ip_address belongs to.
	// +kubebuilder:validation:Optional
	LoadBalancerID *float64 `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// The Primary IP the ip_address belongs to.
	// +kubebuilder:validation:Optional
	PrimaryIPID *float64 `json:"primaryIpId,omitempty" tf:"primary_ip_id,omitempty"`

	// The server the ip_address belongs to.
	// +kubebuilder:validation:Optional
	ServerID *float64 `json:"serverId,omitempty" tf:"server_id,omitempty"`
}

// RdnsSpec defines the desired state of Rdns
type RdnsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RdnsParameters `json:"forProvider"`
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
	InitProvider RdnsInitParameters `json:"initProvider,omitempty"`
}

// RdnsStatus defines the observed state of Rdns.
type RdnsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RdnsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Rdns is the Schema for the Rdnss API. Provides a Hetzner Cloud Reverse DNS Entry to create, modify and reset reverse dns entries for Hetzner Cloud Servers, Primary IPs, Floating IPs or Load Balancers.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hc}
type Rdns struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dnsPtr) || (has(self.initProvider) && has(self.initProvider.dnsPtr))",message="spec.forProvider.dnsPtr is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipAddress) || (has(self.initProvider) && has(self.initProvider.ipAddress))",message="spec.forProvider.ipAddress is a required parameter"
	Spec   RdnsSpec   `json:"spec"`
	Status RdnsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RdnsList contains a list of Rdnss
type RdnsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rdns `json:"items"`
}

// Repository type metadata.
var (
	Rdns_Kind             = "Rdns"
	Rdns_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Rdns_Kind}.String()
	Rdns_KindAPIVersion   = Rdns_Kind + "." + CRDGroupVersion.String()
	Rdns_GroupVersionKind = CRDGroupVersion.WithKind(Rdns_Kind)
)

func init() {
	SchemeBuilder.Register(&Rdns{}, &RdnsList{})
}
