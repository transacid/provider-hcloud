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

type ApplyToInitParameters struct {

	// Label Selector to select servers the firewall should be applied to (only one
	// of server and label_selectorcan be applied in one block)
	LabelSelector *string `json:"labelSelector,omitempty" tf:"label_selector,omitempty"`

	// ID of the server you want to apply the firewall to (only one of server
	// and label_selectorcan be applied in one block)
	Server *float64 `json:"server,omitempty" tf:"server,omitempty"`
}

type ApplyToObservation struct {

	// Label Selector to select servers the firewall should be applied to (only one
	// of server and label_selectorcan be applied in one block)
	LabelSelector *string `json:"labelSelector,omitempty" tf:"label_selector,omitempty"`

	// ID of the server you want to apply the firewall to (only one of server
	// and label_selectorcan be applied in one block)
	Server *float64 `json:"server,omitempty" tf:"server,omitempty"`
}

type ApplyToParameters struct {

	// Label Selector to select servers the firewall should be applied to (only one
	// of server and label_selectorcan be applied in one block)
	// +kubebuilder:validation:Optional
	LabelSelector *string `json:"labelSelector,omitempty" tf:"label_selector,omitempty"`

	// ID of the server you want to apply the firewall to (only one of server
	// and label_selectorcan be applied in one block)
	// +kubebuilder:validation:Optional
	Server *float64 `json:"server,omitempty" tf:"server,omitempty"`
}

type FirewallInitParameters struct {

	// Resources the firewall should be assigned to
	ApplyTo []ApplyToInitParameters `json:"applyTo,omitempty" tf:"apply_to,omitempty"`

	// User-defined labels (key-value pairs) should be created with.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the Firewall.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration of a Rule from this Firewall.
	Rule []RuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type FirewallObservation struct {

	// Resources the firewall should be assigned to
	ApplyTo []ApplyToObservation `json:"applyTo,omitempty" tf:"apply_to,omitempty"`

	// (int) Unique ID of the Firewall.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// User-defined labels (key-value pairs) should be created with.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the Firewall.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration of a Rule from this Firewall.
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type FirewallParameters struct {

	// Resources the firewall should be assigned to
	// +kubebuilder:validation:Optional
	ApplyTo []ApplyToParameters `json:"applyTo,omitempty" tf:"apply_to,omitempty"`

	// User-defined labels (key-value pairs) should be created with.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the Firewall.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration of a Rule from this Firewall.
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type RuleInitParameters struct {

	// Description of the firewall rule
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of IPs or CIDRs that are allowed within this Firewall Rule (when direction
	// is out)
	// +listType=set
	DestinationIps []*string `json:"destinationIps,omitempty" tf:"destination_ips,omitempty"`

	// Direction of the Firewall Rule. in
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Port of the Firewall Rule. Required when protocol is tcp or udp. You can use any
	// to allow all ports for the specific protocol. Port ranges are also possible: 80-85 allows all ports between 80 and 85.
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol of the Firewall Rule. tcp, icmp, udp, gre, esp
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// List of IPs or CIDRs that are allowed within this Firewall Rule (when direction
	// is in)
	// +listType=set
	SourceIps []*string `json:"sourceIps,omitempty" tf:"source_ips,omitempty"`
}

type RuleObservation struct {

	// Description of the firewall rule
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of IPs or CIDRs that are allowed within this Firewall Rule (when direction
	// is out)
	// +listType=set
	DestinationIps []*string `json:"destinationIps,omitempty" tf:"destination_ips,omitempty"`

	// Direction of the Firewall Rule. in
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Port of the Firewall Rule. Required when protocol is tcp or udp. You can use any
	// to allow all ports for the specific protocol. Port ranges are also possible: 80-85 allows all ports between 80 and 85.
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol of the Firewall Rule. tcp, icmp, udp, gre, esp
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// List of IPs or CIDRs that are allowed within this Firewall Rule (when direction
	// is in)
	// +listType=set
	SourceIps []*string `json:"sourceIps,omitempty" tf:"source_ips,omitempty"`
}

type RuleParameters struct {

	// Description of the firewall rule
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of IPs or CIDRs that are allowed within this Firewall Rule (when direction
	// is out)
	// +kubebuilder:validation:Optional
	// +listType=set
	DestinationIps []*string `json:"destinationIps,omitempty" tf:"destination_ips,omitempty"`

	// Direction of the Firewall Rule. in
	// +kubebuilder:validation:Optional
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// Port of the Firewall Rule. Required when protocol is tcp or udp. You can use any
	// to allow all ports for the specific protocol. Port ranges are also possible: 80-85 allows all ports between 80 and 85.
	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol of the Firewall Rule. tcp, icmp, udp, gre, esp
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// List of IPs or CIDRs that are allowed within this Firewall Rule (when direction
	// is in)
	// +kubebuilder:validation:Optional
	// +listType=set
	SourceIps []*string `json:"sourceIps,omitempty" tf:"source_ips,omitempty"`
}

// FirewallSpec defines the desired state of Firewall
type FirewallSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallParameters `json:"forProvider"`
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
	InitProvider FirewallInitParameters `json:"initProvider,omitempty"`
}

// FirewallStatus defines the observed state of Firewall.
type FirewallStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Firewall is the Schema for the Firewalls API. Provides a Hetzner Cloud Firewall to represent a Firewall in the Hetzner Cloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hc}
type Firewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   FirewallSpec   `json:"spec"`
	Status FirewallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallList contains a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Firewall `json:"items"`
}

// Repository type metadata.
var (
	Firewall_Kind             = "Firewall"
	Firewall_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Firewall_Kind}.String()
	Firewall_KindAPIVersion   = Firewall_Kind + "." + CRDGroupVersion.String()
	Firewall_GroupVersionKind = CRDGroupVersion.WithKind(Firewall_Kind)
)

func init() {
	SchemeBuilder.Register(&Firewall{}, &FirewallList{})
}
