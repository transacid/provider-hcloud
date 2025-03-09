// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type AlgorithmInitParameters struct {

	// Type of the Load Balancer Algorithm. round_robin or least_connections
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AlgorithmObservation struct {

	// Type of the Load Balancer Algorithm. round_robin or least_connections
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AlgorithmParameters struct {

	// Type of the Load Balancer Algorithm. round_robin or least_connections
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type BalancerInitParameters struct {

	// Configuration of the algorithm the Load Balancer use.
	Algorithm []AlgorithmInitParameters `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Enable or disable delete protection. See "Delete Protection" in the Provider Docs for details.
	DeleteProtection *bool `json:"deleteProtection,omitempty" tf:"delete_protection,omitempty"`

	// User-defined labels (key-value pairs) should be created with.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Type of the Load Balancer.
	LoadBalancerType *string `json:"loadBalancerType,omitempty" tf:"load_balancer_type,omitempty"`

	// The location name of the Load Balancer. Require when no network_zone is set. See the Hetzner Docs for more details about locations.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Name of the Load Balancer.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Network Zone of the Load Balancer. Require when no location is set.
	NetworkZone *string `json:"networkZone,omitempty" tf:"network_zone,omitempty"`

	Target []TargetInitParameters `json:"target,omitempty" tf:"target,omitempty"`
}

type BalancerObservation struct {

	// Configuration of the algorithm the Load Balancer use.
	Algorithm []AlgorithmObservation `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Enable or disable delete protection. See "Delete Protection" in the Provider Docs for details.
	DeleteProtection *bool `json:"deleteProtection,omitempty" tf:"delete_protection,omitempty"`

	// (int) Unique ID of the Load Balancer.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (string) IPv4 Address of the Load Balancer.
	IPv4 *string `json:"ipv4,omitempty" tf:"ipv4,omitempty"`

	// (string) IPv6 Address of the Load Balancer.
	IPv6 *string `json:"ipv6,omitempty" tf:"ipv6,omitempty"`

	// User-defined labels (key-value pairs) should be created with.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Type of the Load Balancer.
	LoadBalancerType *string `json:"loadBalancerType,omitempty" tf:"load_balancer_type,omitempty"`

	// The location name of the Load Balancer. Require when no network_zone is set. See the Hetzner Docs for more details about locations.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Name of the Load Balancer.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (int) ID of the first private network that this Load Balancer is connected to.
	NetworkID *float64 `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// (string) IP of the Load Balancer in the first private network that it is connected to.
	NetworkIP *string `json:"networkIp,omitempty" tf:"network_ip,omitempty"`

	// The Network Zone of the Load Balancer. Require when no location is set.
	NetworkZone *string `json:"networkZone,omitempty" tf:"network_zone,omitempty"`

	Target []TargetObservation `json:"target,omitempty" tf:"target,omitempty"`
}

type BalancerParameters struct {

	// Configuration of the algorithm the Load Balancer use.
	// +kubebuilder:validation:Optional
	Algorithm []AlgorithmParameters `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Enable or disable delete protection. See "Delete Protection" in the Provider Docs for details.
	// +kubebuilder:validation:Optional
	DeleteProtection *bool `json:"deleteProtection,omitempty" tf:"delete_protection,omitempty"`

	// User-defined labels (key-value pairs) should be created with.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Type of the Load Balancer.
	// +kubebuilder:validation:Optional
	LoadBalancerType *string `json:"loadBalancerType,omitempty" tf:"load_balancer_type,omitempty"`

	// The location name of the Load Balancer. Require when no network_zone is set. See the Hetzner Docs for more details about locations.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Name of the Load Balancer.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Network Zone of the Load Balancer. Require when no location is set.
	// +kubebuilder:validation:Optional
	NetworkZone *string `json:"networkZone,omitempty" tf:"network_zone,omitempty"`

	// +kubebuilder:validation:Optional
	Target []TargetParameters `json:"target,omitempty" tf:"target,omitempty"`
}

type TargetInitParameters struct {

	// (int) Unique ID of the Load Balancer.
	ServerID *float64 `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Type of the Load Balancer Algorithm. round_robin or least_connections
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	UsePrivateIP *bool `json:"usePrivateIp,omitempty" tf:"use_private_ip,omitempty"`
}

type TargetObservation struct {

	// (int) Unique ID of the Load Balancer.
	ServerID *float64 `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Type of the Load Balancer Algorithm. round_robin or least_connections
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	UsePrivateIP *bool `json:"usePrivateIp,omitempty" tf:"use_private_ip,omitempty"`
}

type TargetParameters struct {

	// (int) Unique ID of the Load Balancer.
	// +kubebuilder:validation:Optional
	ServerID *float64 `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Type of the Load Balancer Algorithm. round_robin or least_connections
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	UsePrivateIP *bool `json:"usePrivateIp,omitempty" tf:"use_private_ip,omitempty"`
}

// BalancerSpec defines the desired state of Balancer
type BalancerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BalancerParameters `json:"forProvider"`
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
	InitProvider BalancerInitParameters `json:"initProvider,omitempty"`
}

// BalancerStatus defines the observed state of Balancer.
type BalancerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BalancerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Balancer is the Schema for the Balancers API. Provides a Hetzner Cloud Load Balancer to represent a Load Balancer in the Hetzner Cloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hc}
type Balancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.loadBalancerType) || (has(self.initProvider) && has(self.initProvider.loadBalancerType))",message="spec.forProvider.loadBalancerType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   BalancerSpec   `json:"spec"`
	Status BalancerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BalancerList contains a list of Balancers
type BalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Balancer `json:"items"`
}

// Repository type metadata.
var (
	Balancer_Kind             = "Balancer"
	Balancer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Balancer_Kind}.String()
	Balancer_KindAPIVersion   = Balancer_Kind + "." + CRDGroupVersion.String()
	Balancer_GroupVersionKind = CRDGroupVersion.WithKind(Balancer_Kind)
)

func init() {
	SchemeBuilder.Register(&Balancer{}, &BalancerList{})
}
