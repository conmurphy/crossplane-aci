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

type BridgeDomainObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BridgeDomainParameters struct {

	// +kubebuilder:validation:Optional
	Annotation *string `json:"annotation,omitempty" tf:"annotation,omitempty"`

	// +kubebuilder:validation:Optional
	ArpFlood *string `json:"arpFlood,omitempty" tf:"arp_flood,omitempty"`

	// +kubebuilder:validation:Optional
	BridgeDomainType *string `json:"bridgeDomainType,omitempty" tf:"bridge_domain_type,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	EpClear *string `json:"epClear,omitempty" tf:"ep_clear,omitempty"`

	// +kubebuilder:validation:Optional
	EpMoveDetectMode *string `json:"epMoveDetectMode,omitempty" tf:"ep_move_detect_mode,omitempty"`

	// +kubebuilder:validation:Optional
	HostBasedRouting *string `json:"hostBasedRouting,omitempty" tf:"host_based_routing,omitempty"`

	// +kubebuilder:validation:Optional
	IPLearning *string `json:"ipLearning,omitempty" tf:"ip_learning,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6McastAllow *string `json:"ipv6McastAllow,omitempty" tf:"ipv6_mcast_allow,omitempty"`

	// +kubebuilder:validation:Optional
	IntersiteBumTrafficAllow *string `json:"intersiteBumTrafficAllow,omitempty" tf:"intersite_bum_traffic_allow,omitempty"`

	// +kubebuilder:validation:Optional
	IntersiteL2Stretch *string `json:"intersiteL2Stretch,omitempty" tf:"intersite_l2_stretch,omitempty"`

	// +kubebuilder:validation:Optional
	LimitIPLearnToSubnets *string `json:"limitIpLearnToSubnets,omitempty" tf:"limit_ip_learn_to_subnets,omitempty"`

	// +kubebuilder:validation:Optional
	LlAddr *string `json:"llAddr,omitempty" tf:"ll_addr,omitempty"`

	// +kubebuilder:validation:Optional
	Mac *string `json:"mac,omitempty" tf:"mac,omitempty"`

	// +kubebuilder:validation:Optional
	McastAllow *string `json:"mcastAllow,omitempty" tf:"mcast_allow,omitempty"`

	// +kubebuilder:validation:Optional
	MultiDstPktAct *string `json:"multiDstPktAct,omitempty" tf:"multi_dst_pkt_act,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NameAlias *string `json:"nameAlias,omitempty" tf:"name_alias,omitempty"`

	// +kubebuilder:validation:Optional
	OptimizeWanBandwidth *string `json:"optimizeWanBandwidth,omitempty" tf:"optimize_wan_bandwidth,omitempty"`

	// +kubebuilder:validation:Optional
	RelationFvRsAbdPolMonPol *string `json:"relationFvRsAbdPolMonPol,omitempty" tf:"relation_fv_rs_abd_pol_mon_pol,omitempty"`

	// +kubebuilder:validation:Optional
	RelationFvRsBdFloodTo []*string `json:"relationFvRsBdFloodTo,omitempty" tf:"relation_fv_rs_bd_flood_to,omitempty"`

	// +kubebuilder:validation:Optional
	RelationFvRsBdToEpRet *string `json:"relationFvRsBdToEpRet,omitempty" tf:"relation_fv_rs_bd_to_ep_ret,omitempty"`

	// +kubebuilder:validation:Optional
	RelationFvRsBdToFhs *string `json:"relationFvRsBdToFhs,omitempty" tf:"relation_fv_rs_bd_to_fhs,omitempty"`

	// +kubebuilder:validation:Optional
	RelationFvRsBdToNdP *string `json:"relationFvRsBdToNdP,omitempty" tf:"relation_fv_rs_bd_to_nd_p,omitempty"`

	// +kubebuilder:validation:Optional
	RelationFvRsBdToNetflowMonitorPol []RelationFvRsBdToNetflowMonitorPolParameters `json:"relationFvRsBdToNetflowMonitorPol,omitempty" tf:"relation_fv_rs_bd_to_netflow_monitor_pol,omitempty"`

	// +kubebuilder:validation:Optional
	RelationFvRsBdToOut []*string `json:"relationFvRsBdToOut,omitempty" tf:"relation_fv_rs_bd_to_out,omitempty"`

	// +kubebuilder:validation:Optional
	RelationFvRsBdToProfile *string `json:"relationFvRsBdToProfile,omitempty" tf:"relation_fv_rs_bd_to_profile,omitempty"`

	// +kubebuilder:validation:Optional
	RelationFvRsBdToRelayP *string `json:"relationFvRsBdToRelayP,omitempty" tf:"relation_fv_rs_bd_to_relay_p,omitempty"`

	// +crossplane:generate:reference:type=Vrf
	// +kubebuilder:validation:Optional
	RelationFvRsCtx *string `json:"relationFvRsCtx,omitempty" tf:"relation_fv_rs_ctx,omitempty"`

	// +kubebuilder:validation:Optional
	RelationFvRsCtxRef *v1.Reference `json:"relationFvRsCtxRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RelationFvRsCtxSelector *v1.Selector `json:"relationFvRsCtxSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RelationFvRsIgmpsn *string `json:"relationFvRsIgmpsn,omitempty" tf:"relation_fv_rs_igmpsn,omitempty"`

	// +kubebuilder:validation:Optional
	RelationFvRsMldsn *string `json:"relationFvRsMldsn,omitempty" tf:"relation_fv_rs_mldsn,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aci/apis/root/v1alpha1.Tenant
	// +kubebuilder:validation:Optional
	TenantDn *string `json:"tenantDn,omitempty" tf:"tenant_dn,omitempty"`

	// +kubebuilder:validation:Optional
	TenantDnRef *v1.Reference `json:"tenantDnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TenantDnSelector *v1.Selector `json:"tenantDnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	UnicastRoute *string `json:"unicastRoute,omitempty" tf:"unicast_route,omitempty"`

	// +kubebuilder:validation:Optional
	UnkMacUcastAct *string `json:"unkMacUcastAct,omitempty" tf:"unk_mac_ucast_act,omitempty"`

	// +kubebuilder:validation:Optional
	UnkMcastAct *string `json:"unkMcastAct,omitempty" tf:"unk_mcast_act,omitempty"`

	// +kubebuilder:validation:Optional
	V6UnkMcastAct *string `json:"v6unkMcastAct,omitempty" tf:"v6unk_mcast_act,omitempty"`

	// +kubebuilder:validation:Optional
	Vmac *string `json:"vmac,omitempty" tf:"vmac,omitempty"`
}

type RelationFvRsBdToNetflowMonitorPolObservation struct {
}

type RelationFvRsBdToNetflowMonitorPolParameters struct {

	// +kubebuilder:validation:Required
	FltType *string `json:"fltType" tf:"flt_type,omitempty"`

	// +kubebuilder:validation:Required
	TnNetflowMonitorPolName *string `json:"tnNetflowMonitorPolName" tf:"tn_netflow_monitor_pol_name,omitempty"`
}

// BridgeDomainSpec defines the desired state of BridgeDomain
type BridgeDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BridgeDomainParameters `json:"forProvider"`
}

// BridgeDomainStatus defines the observed state of BridgeDomain.
type BridgeDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BridgeDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BridgeDomain is the Schema for the BridgeDomains API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,acijet}
type BridgeDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BridgeDomainSpec   `json:"spec"`
	Status            BridgeDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BridgeDomainList contains a list of BridgeDomains
type BridgeDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BridgeDomain `json:"items"`
}

// Repository type metadata.
var (
	BridgeDomain_Kind             = "BridgeDomain"
	BridgeDomain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BridgeDomain_Kind}.String()
	BridgeDomain_KindAPIVersion   = BridgeDomain_Kind + "." + CRDGroupVersion.String()
	BridgeDomain_GroupVersionKind = CRDGroupVersion.WithKind(BridgeDomain_Kind)
)

func init() {
	SchemeBuilder.Register(&BridgeDomain{}, &BridgeDomainList{})
}
