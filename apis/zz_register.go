// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/transacid/provider-hcloud/apis/firewall/v1alpha1"
	v1alpha1floatingip "github.com/transacid/provider-hcloud/apis/floatingip/v1alpha1"
	v1alpha1loadbalancer "github.com/transacid/provider-hcloud/apis/loadbalancer/v1alpha1"
	v1alpha1managed "github.com/transacid/provider-hcloud/apis/managed/v1alpha1"
	v1alpha1network "github.com/transacid/provider-hcloud/apis/network/v1alpha1"
	v1alpha1placementgroup "github.com/transacid/provider-hcloud/apis/placementgroup/v1alpha1"
	v1alpha1primaryip "github.com/transacid/provider-hcloud/apis/primaryip/v1alpha1"
	v1alpha1rdns "github.com/transacid/provider-hcloud/apis/rdns/v1alpha1"
	v1alpha1server "github.com/transacid/provider-hcloud/apis/server/v1alpha1"
	v1alpha1snapshot "github.com/transacid/provider-hcloud/apis/snapshot/v1alpha1"
	v1alpha1sshkey "github.com/transacid/provider-hcloud/apis/sshkey/v1alpha1"
	v1alpha1uploaded "github.com/transacid/provider-hcloud/apis/uploaded/v1alpha1"
	v1alpha1apis "github.com/transacid/provider-hcloud/apis/v1alpha1"
	v1beta1 "github.com/transacid/provider-hcloud/apis/v1beta1"
	v1alpha1volume "github.com/transacid/provider-hcloud/apis/volume/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1floatingip.SchemeBuilder.AddToScheme,
		v1alpha1loadbalancer.SchemeBuilder.AddToScheme,
		v1alpha1managed.SchemeBuilder.AddToScheme,
		v1alpha1network.SchemeBuilder.AddToScheme,
		v1alpha1placementgroup.SchemeBuilder.AddToScheme,
		v1alpha1primaryip.SchemeBuilder.AddToScheme,
		v1alpha1rdns.SchemeBuilder.AddToScheme,
		v1alpha1server.SchemeBuilder.AddToScheme,
		v1alpha1snapshot.SchemeBuilder.AddToScheme,
		v1alpha1sshkey.SchemeBuilder.AddToScheme,
		v1alpha1uploaded.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1volume.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
