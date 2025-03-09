// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"

	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	v1alpha1 "github.com/transacid/provider-hcloud/apis/network/v1alpha1"
	v1alpha11 "github.com/transacid/provider-hcloud/apis/placementgroup/v1alpha1"
)

// ResolveReferences of this Network.
func (mg *Network) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1alpha1.NetworkList{},
			Managed: &v1alpha1.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.ServerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServerIDRef,
		Selector:     mg.Spec.ForProvider.ServerIDSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerID")
	}
	mg.Spec.ForProvider.ServerID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.InitProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.NetworkIDRef,
		Selector:     mg.Spec.InitProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1alpha1.NetworkList{},
			Managed: &v1alpha1.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NetworkID")
	}
	mg.Spec.InitProvider.NetworkID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.InitProvider.ServerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ServerIDRef,
		Selector:     mg.Spec.InitProvider.ServerIDSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServerID")
	}
	mg.Spec.InitProvider.ServerID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Server.
func (mg *Server) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.PlacementGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PlacementGroupIDRef,
		Selector:     mg.Spec.ForProvider.PlacementGroupIDSelector,
		To: reference.To{
			List:    &v1alpha11.GroupList{},
			Managed: &v1alpha11.Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PlacementGroupID")
	}
	mg.Spec.ForProvider.PlacementGroupID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PlacementGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.InitProvider.PlacementGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.PlacementGroupIDRef,
		Selector:     mg.Spec.InitProvider.PlacementGroupIDSelector,
		To: reference.To{
			List:    &v1alpha11.GroupList{},
			Managed: &v1alpha11.Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PlacementGroupID")
	}
	mg.Spec.InitProvider.PlacementGroupID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PlacementGroupIDRef = rsp.ResolvedReference

	return nil
}
