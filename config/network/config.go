package network

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure ...
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aci_vrf", func(r *config.Resource) {
		r.ShortGroup = "networks"
		r.Kind = "Vrf"
		r.References["tenant_dn"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-aci/apis/root/v1alpha1.Tenant",
		}
	})
	p.AddResourceConfigurator("aci_bridge_domain", func(r *config.Resource) {
		r.ShortGroup = "networks"
		r.Kind = "BridgeDomain"
		r.References["tenant_dn"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-aci/apis/root/v1alpha1.Tenant",
		}
		r.References["relation_fv_rs_ctx"] = config.Reference{
			Type: "Vrf",
		}

	})
	p.AddResourceConfigurator("aci_subnet", func(r *config.Resource) {
		r.ShortGroup = "networks"
		r.Kind = "Subnet"
		r.References["parent_dn"] = config.Reference{
			Type: "BridgeDomain",
		}
	})

	// L3 Out
	p.AddResourceConfigurator("aci_l3_outside", func(r *config.Resource) {
		r.ShortGroup = "networks"
		r.Kind = "L3Outside"
		r.References["tenant_dn"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-aci/apis/root/v1alpha1.Tenant",
		}
		r.References["relation_l3ext_rs_ectx"] = config.Reference{
			Type: "Vrf",
		}
		r.References["relation_l3ext_rs_l3_dom_att"] = config.Reference{
			Type: "L3DomainProfile",
		}
	})
	p.AddResourceConfigurator("aci_external_network_instance_profile", func(r *config.Resource) {
		r.ShortGroup = "networks"
		r.Kind = "L3ExternalNetworkInstanceProfile"
		r.References["l3_outside_dn"] = config.Reference{
			Type: "L3Outside",
		}
	})
	p.AddResourceConfigurator("aci_l3_ext_subnet", func(r *config.Resource) {
		r.ShortGroup = "networks"
		r.Kind = "L3ExternalSubnet"
		r.References["external_network_instance_profile_dn"] = config.Reference{
			Type: "L3ExternalNetworkInstanceProfile",
		}
	})
	p.AddResourceConfigurator("aci_logical_node_profile", func(r *config.Resource) {
		r.ShortGroup = "networks"
		r.Kind = "L3LogicalNodeProfile"
		r.References["l3_outside_dn"] = config.Reference{
			Type: "L3Outside",
		}
	})
	p.AddResourceConfigurator("aci_logical_interface_profile", func(r *config.Resource) {
		r.ShortGroup = "networks"
		r.Kind = "L3LogicalInterfaceProfile"
		r.References["logical_node_profile_dn"] = config.Reference{
			Type: "L3LogicalNodeProfile",
		}
	})
	p.AddResourceConfigurator("aci_l3out_path_attachment", func(r *config.Resource) {
		r.ShortGroup = "networks"
		r.Kind = "L3PathAttachment"
		r.References["logical_interface_profile_dn"] = config.Reference{
			Type: "L3LogicalInterfaceProfile",
		}
	})
	p.AddResourceConfigurator("aci_l3out_ospf_interface_profile", func(r *config.Resource) {
		r.ShortGroup = "networks"
		r.Kind = "L3OspfInterfaceProfile"
		r.References["logical_interface_profile_dn"] = config.Reference{
			Type: "L3LogicalInterfaceProfile",
		}
	})
	p.AddResourceConfigurator("aci_l3_domain_profile", func(r *config.Resource) {
		r.ShortGroup = "networks"
		r.Kind = "L3DomainProfile"
		r.References["relation_infra_rs_vlan_ns"] = config.Reference{
			Type: "VlanPool",
		}
	})
	p.AddResourceConfigurator("aci_logical_node_to_fabric_node", func(r *config.Resource) {
		r.ShortGroup = "networks"
		r.Kind = "L3LogicalNodeToFabricNode"
		r.References["logical_node_profile_dn"] = config.Reference{
			Type: "L3LogicalNodeProfile",
		}
	})
	p.AddResourceConfigurator("aci_ospf_interface_policy", func(r *config.Resource) {
		r.ShortGroup = "networks"
		r.Kind = "OspfInterfacePolicy"
		r.References["tenant_dn"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-aci/apis/root/v1alpha1.Tenant",
		}
	})
	p.AddResourceConfigurator("aci_l3out_ospf_external_policy", func(r *config.Resource) {
		r.ShortGroup = "networks"
		r.Kind = "L3OspfExternalPolicy"
		r.References["l3_outside_dn"] = config.Reference{
			Type: "L3Outside",
		}
	})

	// VMM
	p.AddResourceConfigurator("aci_vmm_domain", func(r *config.Resource) {
		r.ShortGroup = "networks"
		r.Kind = "VmmDomain"
		r.References["logical_node_profile_dn"] = config.Reference{
			Type: "L3LogicalNodeProfile",
		}
	})
	p.AddResourceConfigurator("aci_vmm_credential", func(r *config.Resource) {
		r.ShortGroup = "networks"
		r.Kind = "VmmCredential"
		r.References["vmm_domain_dn"] = config.Reference{
			Type: "VmmDomain",
		}
	})

	// Access Policies
	p.AddResourceConfigurator("aci_vlan_pool", func(r *config.Resource) {
		r.ShortGroup = "networks"
		r.Kind = "VlanPool"
	})

	p.AddResourceConfigurator("aci_ranges", func(r *config.Resource) {
		r.ShortGroup = "networks"
		r.Kind = "Ranges"
		r.References["vlan_pool_dn"] = config.Reference{
			Type: "VlanPool",
		}
	})

}
