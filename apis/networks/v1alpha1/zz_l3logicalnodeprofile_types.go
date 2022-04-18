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

type L3LogicalNodeProfileObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type L3LogicalNodeProfileParameters struct {

	// +kubebuilder:validation:Optional
	Annotation *string `json:"annotation,omitempty" tf:"annotation,omitempty"`

	// +kubebuilder:validation:Optional
	ConfigIssues *string `json:"configIssues,omitempty" tf:"config_issues,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +crossplane:generate:reference:type=L3Outside
	// +kubebuilder:validation:Optional
	L3OutsideDn *string `json:"l3OutsideDn,omitempty" tf:"l3_outside_dn,omitempty"`

	// +kubebuilder:validation:Optional
	L3OutsideDnRef *v1.Reference `json:"l3OutsideDnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	L3OutsideDnSelector *v1.Selector `json:"l3OutsideDnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NameAlias *string `json:"nameAlias,omitempty" tf:"name_alias,omitempty"`

	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// +kubebuilder:validation:Optional
	TargetDscp *string `json:"targetDscp,omitempty" tf:"target_dscp,omitempty"`
}

// L3LogicalNodeProfileSpec defines the desired state of L3LogicalNodeProfile
type L3LogicalNodeProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     L3LogicalNodeProfileParameters `json:"forProvider"`
}

// L3LogicalNodeProfileStatus defines the observed state of L3LogicalNodeProfile.
type L3LogicalNodeProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        L3LogicalNodeProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// L3LogicalNodeProfile is the Schema for the L3LogicalNodeProfiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,acijet}
type L3LogicalNodeProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              L3LogicalNodeProfileSpec   `json:"spec"`
	Status            L3LogicalNodeProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// L3LogicalNodeProfileList contains a list of L3LogicalNodeProfiles
type L3LogicalNodeProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []L3LogicalNodeProfile `json:"items"`
}

// Repository type metadata.
var (
	L3LogicalNodeProfile_Kind             = "L3LogicalNodeProfile"
	L3LogicalNodeProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: L3LogicalNodeProfile_Kind}.String()
	L3LogicalNodeProfile_KindAPIVersion   = L3LogicalNodeProfile_Kind + "." + CRDGroupVersion.String()
	L3LogicalNodeProfile_GroupVersionKind = CRDGroupVersion.WithKind(L3LogicalNodeProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&L3LogicalNodeProfile{}, &L3LogicalNodeProfileList{})
}