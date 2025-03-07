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

type SnapshotInitParameters struct {

	// Description of the snapshot.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// User-defined labels (key-value pairs) should be created with.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Server to the snapshot should be created from.
	// +crossplane:generate:reference:type=github.com/transacid/provider-hcloud/apis/server/v1alpha1.Server
	ServerID *float64 `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a Server in server to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a Server in server to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`
}

type SnapshotObservation struct {

	// Description of the snapshot.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (int) Unique ID of the snapshot.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// User-defined labels (key-value pairs) should be created with.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Server to the snapshot should be created from.
	ServerID *float64 `json:"serverId,omitempty" tf:"server_id,omitempty"`
}

type SnapshotParameters struct {

	// Description of the snapshot.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// User-defined labels (key-value pairs) should be created with.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Server to the snapshot should be created from.
	// +crossplane:generate:reference:type=github.com/transacid/provider-hcloud/apis/server/v1alpha1.Server
	// +kubebuilder:validation:Optional
	ServerID *float64 `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a Server in server to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a Server in server to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`
}

// SnapshotSpec defines the desired state of Snapshot
type SnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotParameters `json:"forProvider"`
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
	InitProvider SnapshotInitParameters `json:"initProvider,omitempty"`
}

// SnapshotStatus defines the observed state of Snapshot.
type SnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Snapshot is the Schema for the Snapshots API. Provides a Hetzner Cloud snapshot to represent an image with type snapshot in the Hetzner Cloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hc}
type Snapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotSpec   `json:"spec"`
	Status            SnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotList contains a list of Snapshots
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Snapshot `json:"items"`
}

// Repository type metadata.
var (
	Snapshot_Kind             = "Snapshot"
	Snapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Snapshot_Kind}.String()
	Snapshot_KindAPIVersion   = Snapshot_Kind + "." + CRDGroupVersion.String()
	Snapshot_GroupVersionKind = CRDGroupVersion.WithKind(Snapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&Snapshot{}, &SnapshotList{})
}
