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

type VmmDomainObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VmmDomainParameters struct {

	// +kubebuilder:validation:Optional
	AccessMode *string `json:"accessMode,omitempty" tf:"access_mode,omitempty"`

	// +kubebuilder:validation:Optional
	Annotation *string `json:"annotation,omitempty" tf:"annotation,omitempty"`

	// +kubebuilder:validation:Optional
	ArpLearning *string `json:"arpLearning,omitempty" tf:"arp_learning,omitempty"`

	// +kubebuilder:validation:Optional
	AveTimeOut *string `json:"aveTimeOut,omitempty" tf:"ave_time_out,omitempty"`

	// +kubebuilder:validation:Optional
	ConfigInfraPg *string `json:"configInfraPg,omitempty" tf:"config_infra_pg,omitempty"`

	// +kubebuilder:validation:Optional
	CtrlKnob *string `json:"ctrlKnob,omitempty" tf:"ctrl_knob,omitempty"`

	// +kubebuilder:validation:Optional
	Delimiter *string `json:"delimiter,omitempty" tf:"delimiter,omitempty"`

	// +kubebuilder:validation:Optional
	EnableAve *string `json:"enableAve,omitempty" tf:"enable_ave,omitempty"`

	// +kubebuilder:validation:Optional
	EnableTag *string `json:"enableTag,omitempty" tf:"enable_tag,omitempty"`

	// +kubebuilder:validation:Optional
	EncapMode *string `json:"encapMode,omitempty" tf:"encap_mode,omitempty"`

	// +kubebuilder:validation:Optional
	EnfPref *string `json:"enfPref,omitempty" tf:"enf_pref,omitempty"`

	// +kubebuilder:validation:Optional
	EpInventoryType *string `json:"epInventoryType,omitempty" tf:"ep_inventory_type,omitempty"`

	// +kubebuilder:validation:Optional
	EpRetTime *string `json:"epRetTime,omitempty" tf:"ep_ret_time,omitempty"`

	// +kubebuilder:validation:Optional
	HvAvailMonitor *string `json:"hvAvailMonitor,omitempty" tf:"hv_avail_monitor,omitempty"`

	// +kubebuilder:validation:Optional
	McastAddr *string `json:"mcastAddr,omitempty" tf:"mcast_addr,omitempty"`

	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NameAlias *string `json:"nameAlias,omitempty" tf:"name_alias,omitempty"`

	// +kubebuilder:validation:Optional
	PrefEncapMode *string `json:"prefEncapMode,omitempty" tf:"pref_encap_mode,omitempty"`

	// +kubebuilder:validation:Required
	ProviderProfileDn *string `json:"providerProfileDn" tf:"provider_profile_dn,omitempty"`

	// +kubebuilder:validation:Optional
	RelationInfraRsDomVxlanNsDef *string `json:"relationInfraRsDomVxlanNsDef,omitempty" tf:"relation_infra_rs_dom_vxlan_ns_def,omitempty"`

	// +kubebuilder:validation:Optional
	RelationInfraRsVipAddrNs *string `json:"relationInfraRsVipAddrNs,omitempty" tf:"relation_infra_rs_vip_addr_ns,omitempty"`

	// +kubebuilder:validation:Optional
	RelationInfraRsVlanNs *string `json:"relationInfraRsVlanNs,omitempty" tf:"relation_infra_rs_vlan_ns,omitempty"`

	// +kubebuilder:validation:Optional
	RelationInfraRsVlanNsDef *string `json:"relationInfraRsVlanNsDef,omitempty" tf:"relation_infra_rs_vlan_ns_def,omitempty"`

	// +kubebuilder:validation:Optional
	RelationVmmRsDefaultCdpIfPol *string `json:"relationVmmRsDefaultCdpIfPol,omitempty" tf:"relation_vmm_rs_default_cdp_if_pol,omitempty"`

	// +kubebuilder:validation:Optional
	RelationVmmRsDefaultFwPol *string `json:"relationVmmRsDefaultFwPol,omitempty" tf:"relation_vmm_rs_default_fw_pol,omitempty"`

	// +kubebuilder:validation:Optional
	RelationVmmRsDefaultL2InstPol *string `json:"relationVmmRsDefaultL2InstPol,omitempty" tf:"relation_vmm_rs_default_l2_inst_pol,omitempty"`

	// +kubebuilder:validation:Optional
	RelationVmmRsDefaultLacpLagPol *string `json:"relationVmmRsDefaultLacpLagPol,omitempty" tf:"relation_vmm_rs_default_lacp_lag_pol,omitempty"`

	// +kubebuilder:validation:Optional
	RelationVmmRsDefaultLldpIfPol *string `json:"relationVmmRsDefaultLldpIfPol,omitempty" tf:"relation_vmm_rs_default_lldp_if_pol,omitempty"`

	// +kubebuilder:validation:Optional
	RelationVmmRsDefaultStpIfPol *string `json:"relationVmmRsDefaultStpIfPol,omitempty" tf:"relation_vmm_rs_default_stp_if_pol,omitempty"`

	// +kubebuilder:validation:Optional
	RelationVmmRsDomMcastAddrNs *string `json:"relationVmmRsDomMcastAddrNs,omitempty" tf:"relation_vmm_rs_dom_mcast_addr_ns,omitempty"`

	// +kubebuilder:validation:Optional
	RelationVmmRsPrefEnhancedLagPol *string `json:"relationVmmRsPrefEnhancedLagPol,omitempty" tf:"relation_vmm_rs_pref_enhanced_lag_pol,omitempty"`
}

// VmmDomainSpec defines the desired state of VmmDomain
type VmmDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VmmDomainParameters `json:"forProvider"`
}

// VmmDomainStatus defines the observed state of VmmDomain.
type VmmDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VmmDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VmmDomain is the Schema for the VmmDomains API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,acijet}
type VmmDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VmmDomainSpec   `json:"spec"`
	Status            VmmDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VmmDomainList contains a list of VmmDomains
type VmmDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VmmDomain `json:"items"`
}

// Repository type metadata.
var (
	VmmDomain_Kind             = "VmmDomain"
	VmmDomain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VmmDomain_Kind}.String()
	VmmDomain_KindAPIVersion   = VmmDomain_Kind + "." + CRDGroupVersion.String()
	VmmDomain_GroupVersionKind = CRDGroupVersion.WithKind(VmmDomain_Kind)
)

func init() {
	SchemeBuilder.Register(&VmmDomain{}, &VmmDomainList{})
}