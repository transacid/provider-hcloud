// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha1 "github.com/transacid/provider-hcloud/apis/network/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this BalancerNetwork.
func (mg *BalancerNetwork) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.LoadBalancerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LoadBalancerIDRef,
		Selector:     mg.Spec.ForProvider.LoadBalancerIDSelector,
		To: reference.To{
			List:    &BalancerList{},
			Managed: &Balancer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoadBalancerID")
	}
	mg.Spec.ForProvider.LoadBalancerID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoadBalancerIDRef = rsp.ResolvedReference

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
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.InitProvider.LoadBalancerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LoadBalancerIDRef,
		Selector:     mg.Spec.InitProvider.LoadBalancerIDSelector,
		To: reference.To{
			List:    &BalancerList{},
			Managed: &Balancer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LoadBalancerID")
	}
	mg.Spec.InitProvider.LoadBalancerID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LoadBalancerIDRef = rsp.ResolvedReference

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

	return nil
}

// ResolveReferences of this BalancerService.
func (mg *BalancerService) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadBalancerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LoadBalancerIDRef,
		Selector:     mg.Spec.ForProvider.LoadBalancerIDSelector,
		To: reference.To{
			List:    &BalancerList{},
			Managed: &Balancer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoadBalancerID")
	}
	mg.Spec.ForProvider.LoadBalancerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoadBalancerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoadBalancerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LoadBalancerIDRef,
		Selector:     mg.Spec.InitProvider.LoadBalancerIDSelector,
		To: reference.To{
			List:    &BalancerList{},
			Managed: &Balancer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LoadBalancerID")
	}
	mg.Spec.InitProvider.LoadBalancerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LoadBalancerIDRef = rsp.ResolvedReference

	return nil
}
