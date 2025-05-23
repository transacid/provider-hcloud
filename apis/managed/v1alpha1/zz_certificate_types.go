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

type CertificateInitParameters struct {

	// Domain names for which a certificate
	// should be obtained.
	// +listType=set
	DomainNames []*string `json:"domainNames,omitempty" tf:"domain_names,omitempty"`

	// User-defined labels (key-value pairs) the
	// certificate should be created with.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the Certificate.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type CertificateObservation struct {

	// (string) PEM encoded TLS certificate.
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// (string) Point in time when the Certificate was created at Hetzner Cloud (in ISO-8601 format).
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	// Domain names for which a certificate
	// should be obtained.
	// +listType=set
	DomainNames []*string `json:"domainNames,omitempty" tf:"domain_names,omitempty"`

	// (string) Fingerprint of the certificate.
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	// (int) Unique ID of the certificate.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// User-defined labels (key-value pairs) the
	// certificate should be created with.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the Certificate.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (string) Point in time when the Certificate stops being valid (in ISO-8601 format).
	NotValidAfter *string `json:"notValidAfter,omitempty" tf:"not_valid_after,omitempty"`

	// (string) Point in time when the Certificate becomes valid (in ISO-8601 format).
	NotValidBefore *string `json:"notValidBefore,omitempty" tf:"not_valid_before,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CertificateParameters struct {

	// Domain names for which a certificate
	// should be obtained.
	// +kubebuilder:validation:Optional
	// +listType=set
	DomainNames []*string `json:"domainNames,omitempty" tf:"domain_names,omitempty"`

	// User-defined labels (key-value pairs) the
	// certificate should be created with.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the Certificate.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateParameters `json:"forProvider"`
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
	InitProvider CertificateInitParameters `json:"initProvider,omitempty"`
}

// CertificateStatus defines the observed state of Certificate.
type CertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Certificate is the Schema for the Certificates API. Obtain a TLS Certificate managed by Hetzner Cloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hc}
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domainNames) || (has(self.initProvider) && has(self.initProvider.domainNames))",message="spec.forProvider.domainNames is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   CertificateSpec   `json:"spec"`
	Status CertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateList contains a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

// Repository type metadata.
var (
	Certificate_Kind             = "Certificate"
	Certificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Certificate_Kind}.String()
	Certificate_KindAPIVersion   = Certificate_Kind + "." + CRDGroupVersion.String()
	Certificate_GroupVersionKind = CRDGroupVersion.WithKind(Certificate_Kind)
)

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}
