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

type FilterEntryObservation_2 struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FilterEntryParameters_2 struct {

	// +kubebuilder:validation:Optional
	Annotation *string `json:"annotation,omitempty" tf:"annotation,omitempty"`

	// +kubebuilder:validation:Optional
	ApplyToFrag *string `json:"applyToFrag,omitempty" tf:"apply_to_frag,omitempty"`

	// +kubebuilder:validation:Optional
	ArpOpc *string `json:"arpOpc,omitempty" tf:"arp_opc,omitempty"`

	// +kubebuilder:validation:Optional
	DFromPort *string `json:"dFromPort,omitempty" tf:"d_from_port,omitempty"`

	// +kubebuilder:validation:Optional
	DToPort *string `json:"dToPort,omitempty" tf:"d_to_port,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	EtherT *string `json:"etherT,omitempty" tf:"ether_t,omitempty"`

	// +crossplane:generate:reference:type=Filter
	// +kubebuilder:validation:Optional
	FilterDn *string `json:"filterDn,omitempty" tf:"filter_dn,omitempty"`

	// +kubebuilder:validation:Optional
	FilterDnRef *v1.Reference `json:"filterDnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	FilterDnSelector *v1.Selector `json:"filterDnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Icmpv4T *string `json:"icmpv4T,omitempty" tf:"icmpv4_t,omitempty"`

	// +kubebuilder:validation:Optional
	Icmpv6T *string `json:"icmpv6T,omitempty" tf:"icmpv6_t,omitempty"`

	// +kubebuilder:validation:Optional
	MatchDscp *string `json:"matchDscp,omitempty" tf:"match_dscp,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NameAlias *string `json:"nameAlias,omitempty" tf:"name_alias,omitempty"`

	// +kubebuilder:validation:Optional
	Prot *string `json:"prot,omitempty" tf:"prot,omitempty"`

	// +kubebuilder:validation:Optional
	SFromPort *string `json:"sFromPort,omitempty" tf:"s_from_port,omitempty"`

	// +kubebuilder:validation:Optional
	SToPort *string `json:"sToPort,omitempty" tf:"s_to_port,omitempty"`

	// +kubebuilder:validation:Optional
	Stateful *string `json:"stateful,omitempty" tf:"stateful,omitempty"`

	// +kubebuilder:validation:Optional
	TCPRules []*string `json:"tcpRules,omitempty" tf:"tcp_rules,omitempty"`
}

// FilterEntrySpec defines the desired state of FilterEntry
type FilterEntrySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FilterEntryParameters_2 `json:"forProvider"`
}

// FilterEntryStatus defines the observed state of FilterEntry.
type FilterEntryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FilterEntryObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FilterEntry is the Schema for the FilterEntrys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,acijet}
type FilterEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FilterEntrySpec   `json:"spec"`
	Status            FilterEntryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FilterEntryList contains a list of FilterEntrys
type FilterEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FilterEntry `json:"items"`
}

// Repository type metadata.
var (
	FilterEntry_Kind             = "FilterEntry"
	FilterEntry_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FilterEntry_Kind}.String()
	FilterEntry_KindAPIVersion   = FilterEntry_Kind + "." + CRDGroupVersion.String()
	FilterEntry_GroupVersionKind = CRDGroupVersion.WithKind(FilterEntry_Kind)
)

func init() {
	SchemeBuilder.Register(&FilterEntry{}, &FilterEntryList{})
}
