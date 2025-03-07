// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha1 "github.com/transacid/provider-hcloud/apis/server/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Attachment.
func (mg *Attachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromFloatPtrValues(mg.Spec.ForProvider.ServerIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.ServerIdsRefs,
		Selector:      mg.Spec.ForProvider.ServerIdsSelector,
		To: reference.To{
			List:    &v1alpha1.ServerList{},
			Managed: &v1alpha1.Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerIds")
	}
	mg.Spec.ForProvider.ServerIds = reference.ToFloatPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.ServerIdsRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromFloatPtrValues(mg.Spec.InitProvider.ServerIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.ServerIdsRefs,
		Selector:      mg.Spec.InitProvider.ServerIdsSelector,
		To: reference.To{
			List:    &v1alpha1.ServerList{},
			Managed: &v1alpha1.Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServerIds")
	}
	mg.Spec.InitProvider.ServerIds = reference.ToFloatPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.ServerIdsRefs = mrsp.ResolvedReferences

	return nil
}
