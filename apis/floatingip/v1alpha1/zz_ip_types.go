// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IPInitParameters struct {

	// Enable or disable delete protection. See "Delete Protection" in the Provider Docs for details.
	DeleteProtection *bool `json:"deleteProtection,omitempty" tf:"delete_protection,omitempty"`

	// Description of the Floating IP.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of home location (routing is optimized for that location). Optional if server_id argument is passed.
	HomeLocation *string `json:"homeLocation,omitempty" tf:"home_location,omitempty"`

	// User-defined labels (key-value pairs) should be created with.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the Floating IP.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Server to assign the Floating IP to.
	ServerID *float64 `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Type of the Floating IP. ipv4 ipv6
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IPObservation struct {

	// Enable or disable delete protection. See "Delete Protection" in the Provider Docs for details.
	DeleteProtection *bool `json:"deleteProtection,omitempty" tf:"delete_protection,omitempty"`

	// Description of the Floating IP.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of home location (routing is optimized for that location). Optional if server_id argument is passed.
	HomeLocation *string `json:"homeLocation,omitempty" tf:"home_location,omitempty"`

	// (int) Unique ID of the Floating IP.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (string) IP Address of the Floating IP.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// (string) IPv6 subnet. (Only set if type is ipv6)
	IPNetwork *string `json:"ipNetwork,omitempty" tf:"ip_network,omitempty"`

	// User-defined labels (key-value pairs) should be created with.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the Floating IP.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Server to assign the Floating IP to.
	ServerID *float64 `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Type of the Floating IP. ipv4 ipv6
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IPParameters struct {

	// Enable or disable delete protection. See "Delete Protection" in the Provider Docs for details.
	// +kubebuilder:validation:Optional
	DeleteProtection *bool `json:"deleteProtection,omitempty" tf:"delete_protection,omitempty"`

	// Description of the Floating IP.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of home location (routing is optimized for that location). Optional if server_id argument is passed.
	// +kubebuilder:validation:Optional
	HomeLocation *string `json:"homeLocation,omitempty" tf:"home_location,omitempty"`

	// User-defined labels (key-value pairs) should be created with.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the Floating IP.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Server to assign the Floating IP to.
	// +kubebuilder:validation:Optional
	ServerID *float64 `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Type of the Floating IP. ipv4 ipv6
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// IPSpec defines the desired state of IP
type IPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPParameters `json:"forProvider"`
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
	InitProvider IPInitParameters `json:"initProvider,omitempty"`
}

// IPStatus defines the observed state of IP.
type IPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IP is the Schema for the IPs API. Provides a Hetzner Cloud Floating IP to represent a publicly-accessible static IP address that can be mapped to one of your servers.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hcloud}
type IP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   IPSpec   `json:"spec"`
	Status IPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPList contains a list of IPs
type IPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IP `json:"items"`
}

// Repository type metadata.
var (
	IP_Kind             = "IP"
	IP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IP_Kind}.String()
	IP_KindAPIVersion   = IP_Kind + "." + CRDGroupVersion.String()
	IP_GroupVersionKind = CRDGroupVersion.WithKind(IP_Kind)
)

func init() {
	SchemeBuilder.Register(&IP{}, &IPList{})
}