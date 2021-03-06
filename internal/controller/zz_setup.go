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

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	applicationprofile "github.com/crossplane-contrib/provider-jet-aci/internal/controller/applications/applicationprofile"
	contract "github.com/crossplane-contrib/provider-jet-aci/internal/controller/applications/contract"
	contractsubject "github.com/crossplane-contrib/provider-jet-aci/internal/controller/applications/contractsubject"
	epg "github.com/crossplane-contrib/provider-jet-aci/internal/controller/applications/epg"
	epgtocontract "github.com/crossplane-contrib/provider-jet-aci/internal/controller/applications/epgtocontract"
	esg "github.com/crossplane-contrib/provider-jet-aci/internal/controller/applications/esg"
	esgepgselector "github.com/crossplane-contrib/provider-jet-aci/internal/controller/applications/esgepgselector"
	esgselector "github.com/crossplane-contrib/provider-jet-aci/internal/controller/applications/esgselector"
	esgtagselector "github.com/crossplane-contrib/provider-jet-aci/internal/controller/applications/esgtagselector"
	filter "github.com/crossplane-contrib/provider-jet-aci/internal/controller/applications/filter"
	filterentry "github.com/crossplane-contrib/provider-jet-aci/internal/controller/applications/filterentry"
	importedcontract "github.com/crossplane-contrib/provider-jet-aci/internal/controller/applications/importedcontract"
	taboocontract "github.com/crossplane-contrib/provider-jet-aci/internal/controller/applications/taboocontract"
	bridgedomain "github.com/crossplane-contrib/provider-jet-aci/internal/controller/networks/bridgedomain"
	l3domainprofile "github.com/crossplane-contrib/provider-jet-aci/internal/controller/networks/l3domainprofile"
	l3externalnetworkinstanceprofile "github.com/crossplane-contrib/provider-jet-aci/internal/controller/networks/l3externalnetworkinstanceprofile"
	l3externalsubnet "github.com/crossplane-contrib/provider-jet-aci/internal/controller/networks/l3externalsubnet"
	l3logicalinterfaceprofile "github.com/crossplane-contrib/provider-jet-aci/internal/controller/networks/l3logicalinterfaceprofile"
	l3logicalnodeprofile "github.com/crossplane-contrib/provider-jet-aci/internal/controller/networks/l3logicalnodeprofile"
	l3logicalnodetofabricnode "github.com/crossplane-contrib/provider-jet-aci/internal/controller/networks/l3logicalnodetofabricnode"
	l3ospfexternalpolicy "github.com/crossplane-contrib/provider-jet-aci/internal/controller/networks/l3ospfexternalpolicy"
	l3ospfinterfaceprofile "github.com/crossplane-contrib/provider-jet-aci/internal/controller/networks/l3ospfinterfaceprofile"
	l3outside "github.com/crossplane-contrib/provider-jet-aci/internal/controller/networks/l3outside"
	l3pathattachment "github.com/crossplane-contrib/provider-jet-aci/internal/controller/networks/l3pathattachment"
	ospfinterfacepolicy "github.com/crossplane-contrib/provider-jet-aci/internal/controller/networks/ospfinterfacepolicy"
	ranges "github.com/crossplane-contrib/provider-jet-aci/internal/controller/networks/ranges"
	subnet "github.com/crossplane-contrib/provider-jet-aci/internal/controller/networks/subnet"
	vlanpool "github.com/crossplane-contrib/provider-jet-aci/internal/controller/networks/vlanpool"
	vmmcredential "github.com/crossplane-contrib/provider-jet-aci/internal/controller/networks/vmmcredential"
	vmmdomain "github.com/crossplane-contrib/provider-jet-aci/internal/controller/networks/vmmdomain"
	vrf "github.com/crossplane-contrib/provider-jet-aci/internal/controller/networks/vrf"
	providerconfig "github.com/crossplane-contrib/provider-jet-aci/internal/controller/providerconfig"
	tenant "github.com/crossplane-contrib/provider-jet-aci/internal/controller/root/tenant"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		applicationprofile.Setup,
		contract.Setup,
		contractsubject.Setup,
		epg.Setup,
		epgtocontract.Setup,
		esg.Setup,
		esgepgselector.Setup,
		esgselector.Setup,
		esgtagselector.Setup,
		filter.Setup,
		filterentry.Setup,
		importedcontract.Setup,
		taboocontract.Setup,
		bridgedomain.Setup,
		l3domainprofile.Setup,
		l3externalnetworkinstanceprofile.Setup,
		l3externalsubnet.Setup,
		l3logicalinterfaceprofile.Setup,
		l3logicalnodeprofile.Setup,
		l3logicalnodetofabricnode.Setup,
		l3ospfexternalpolicy.Setup,
		l3ospfinterfaceprofile.Setup,
		l3outside.Setup,
		l3pathattachment.Setup,
		ospfinterfacepolicy.Setup,
		ranges.Setup,
		subnet.Setup,
		vlanpool.Setup,
		vmmcredential.Setup,
		vmmdomain.Setup,
		vrf.Setup,
		providerconfig.Setup,
		tenant.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
