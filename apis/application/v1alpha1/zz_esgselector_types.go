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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type EsgSelectorObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EsgSelectorParameters struct {

	// +kubebuilder:validation:Optional
	Annotation *string `json:"annotation,omitempty" tf:"annotation,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +crossplane:generate:reference:type=Esg
	// +kubebuilder:validation:Optional
	EndpointSecurityGroupDn *string `json:"endpointSecurityGroupDn,omitempty" tf:"endpoint_security_group_dn,omitempty"`

	// +kubebuilder:validation:Optional
	EndpointSecurityGroupDnRef *v1.Reference `json:"endpointSecurityGroupDnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	EndpointSecurityGroupDnSelector *v1.Selector `json:"endpointSecurityGroupDnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	MatchExpression *string `json:"matchExpression" tf:"match_expression,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NameAlias *string `json:"nameAlias,omitempty" tf:"name_alias,omitempty"`
}

// EsgSelectorSpec defines the desired state of EsgSelector
type EsgSelectorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EsgSelectorParameters `json:"forProvider"`
}

// EsgSelectorStatus defines the observed state of EsgSelector.
type EsgSelectorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EsgSelectorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EsgSelector is the Schema for the EsgSelectors API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,acijet}
type EsgSelector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EsgSelectorSpec   `json:"spec"`
	Status            EsgSelectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EsgSelectorList contains a list of EsgSelectors
type EsgSelectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EsgSelector `json:"items"`
}

// Repository type metadata.
var (
	EsgSelector_Kind             = "EsgSelector"
	EsgSelector_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EsgSelector_Kind}.String()
	EsgSelector_KindAPIVersion   = EsgSelector_Kind + "." + CRDGroupVersion.String()
	EsgSelector_GroupVersionKind = CRDGroupVersion.WithKind(EsgSelector_Kind)
)

func init() {
	SchemeBuilder.Register(&EsgSelector{}, &EsgSelectorList{})
}
