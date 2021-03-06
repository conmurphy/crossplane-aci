/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha11 "github.com/crossplane-contrib/provider-jet-aci/apis/networks/v1alpha1"
	v1alpha1 "github.com/crossplane-contrib/provider-jet-aci/apis/root/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ApplicationProfile.
func (mg *ApplicationProfile) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TenantDn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TenantDnRef,
		Selector:     mg.Spec.ForProvider.TenantDnSelector,
		To: reference.To{
			List:    &v1alpha1.TenantList{},
			Managed: &v1alpha1.Tenant{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantDn")
	}
	mg.Spec.ForProvider.TenantDn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantDnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Contract.
func (mg *Contract) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TenantDn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TenantDnRef,
		Selector:     mg.Spec.ForProvider.TenantDnSelector,
		To: reference.To{
			List:    &v1alpha1.TenantList{},
			Managed: &v1alpha1.Tenant{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantDn")
	}
	mg.Spec.ForProvider.TenantDn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantDnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ContractSubject.
func (mg *ContractSubject) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ContractDn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ContractDnRef,
		Selector:     mg.Spec.ForProvider.ContractDnSelector,
		To: reference.To{
			List:    &ContractList{},
			Managed: &Contract{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ContractDn")
	}
	mg.Spec.ForProvider.ContractDn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ContractDnRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.RelationVzRsSubjFiltAtt),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.RelationVzRsSubjFiltAttRefs,
		Selector:      mg.Spec.ForProvider.RelationVzRsSubjFiltAttSelector,
		To: reference.To{
			List:    &FilterList{},
			Managed: &Filter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RelationVzRsSubjFiltAtt")
	}
	mg.Spec.ForProvider.RelationVzRsSubjFiltAtt = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.RelationVzRsSubjFiltAttRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this Epg.
func (mg *Epg) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationProfileDn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ApplicationProfileDnRef,
		Selector:     mg.Spec.ForProvider.ApplicationProfileDnSelector,
		To: reference.To{
			List:    &ApplicationProfileList{},
			Managed: &ApplicationProfile{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationProfileDn")
	}
	mg.Spec.ForProvider.ApplicationProfileDn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationProfileDnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EpgToContract.
func (mg *EpgToContract) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationEpgDn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ApplicationEpgDnRef,
		Selector:     mg.Spec.ForProvider.ApplicationEpgDnSelector,
		To: reference.To{
			List:    &EpgList{},
			Managed: &Epg{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationEpgDn")
	}
	mg.Spec.ForProvider.ApplicationEpgDn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationEpgDnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ContractDn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ContractDnRef,
		Selector:     mg.Spec.ForProvider.ContractDnSelector,
		To: reference.To{
			List:    &ContractList{},
			Managed: &Contract{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ContractDn")
	}
	mg.Spec.ForProvider.ContractDn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ContractDnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Esg.
func (mg *Esg) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationProfileDn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ApplicationProfileDnRef,
		Selector:     mg.Spec.ForProvider.ApplicationProfileDnSelector,
		To: reference.To{
			List:    &ApplicationProfileList{},
			Managed: &ApplicationProfile{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationProfileDn")
	}
	mg.Spec.ForProvider.ApplicationProfileDn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationProfileDnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RelationFvRsScope),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RelationFvRsScopeRef,
		Selector:     mg.Spec.ForProvider.RelationFvRsScopeSelector,
		To: reference.To{
			List:    &v1alpha11.VrfList{},
			Managed: &v1alpha11.Vrf{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RelationFvRsScope")
	}
	mg.Spec.ForProvider.RelationFvRsScope = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RelationFvRsScopeRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EsgEpgSelector.
func (mg *EsgEpgSelector) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EndpointSecurityGroupDn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EndpointSecurityGroupDnRef,
		Selector:     mg.Spec.ForProvider.EndpointSecurityGroupDnSelector,
		To: reference.To{
			List:    &EsgList{},
			Managed: &Esg{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EndpointSecurityGroupDn")
	}
	mg.Spec.ForProvider.EndpointSecurityGroupDn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EndpointSecurityGroupDnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MatchEpgDn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.MatchEpgDnRef,
		Selector:     mg.Spec.ForProvider.MatchEpgDnSelector,
		To: reference.To{
			List:    &EpgList{},
			Managed: &Epg{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MatchEpgDn")
	}
	mg.Spec.ForProvider.MatchEpgDn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MatchEpgDnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EsgSelector.
func (mg *EsgSelector) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EndpointSecurityGroupDn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EndpointSecurityGroupDnRef,
		Selector:     mg.Spec.ForProvider.EndpointSecurityGroupDnSelector,
		To: reference.To{
			List:    &EsgList{},
			Managed: &Esg{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EndpointSecurityGroupDn")
	}
	mg.Spec.ForProvider.EndpointSecurityGroupDn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EndpointSecurityGroupDnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EsgTagSelector.
func (mg *EsgTagSelector) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EndpointSecurityGroupDn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EndpointSecurityGroupDnRef,
		Selector:     mg.Spec.ForProvider.EndpointSecurityGroupDnSelector,
		To: reference.To{
			List:    &EsgList{},
			Managed: &Esg{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EndpointSecurityGroupDn")
	}
	mg.Spec.ForProvider.EndpointSecurityGroupDn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EndpointSecurityGroupDnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Filter.
func (mg *Filter) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TenantDn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TenantDnRef,
		Selector:     mg.Spec.ForProvider.TenantDnSelector,
		To: reference.To{
			List:    &v1alpha1.TenantList{},
			Managed: &v1alpha1.Tenant{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantDn")
	}
	mg.Spec.ForProvider.TenantDn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantDnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FilterEntry.
func (mg *FilterEntry) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FilterDn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FilterDnRef,
		Selector:     mg.Spec.ForProvider.FilterDnSelector,
		To: reference.To{
			List:    &FilterList{},
			Managed: &Filter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FilterDn")
	}
	mg.Spec.ForProvider.FilterDn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FilterDnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ImportedContract.
func (mg *ImportedContract) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TenantDn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TenantDnRef,
		Selector:     mg.Spec.ForProvider.TenantDnSelector,
		To: reference.To{
			List:    &v1alpha1.TenantList{},
			Managed: &v1alpha1.Tenant{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantDn")
	}
	mg.Spec.ForProvider.TenantDn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantDnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TabooContract.
func (mg *TabooContract) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TenantDn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TenantDnRef,
		Selector:     mg.Spec.ForProvider.TenantDnSelector,
		To: reference.To{
			List:    &v1alpha1.TenantList{},
			Managed: &v1alpha1.Tenant{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantDn")
	}
	mg.Spec.ForProvider.TenantDn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantDnRef = rsp.ResolvedReference

	return nil
}
