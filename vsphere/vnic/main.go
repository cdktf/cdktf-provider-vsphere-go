package vnic

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.vnic.Vnic",
		reflect.TypeOf((*Vnic)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "distributedPortGroup", GoGetter: "DistributedPortGroup"},
			_jsii_.MemberProperty{JsiiProperty: "distributedPortGroupInput", GoGetter: "DistributedPortGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "distributedSwitchPort", GoGetter: "DistributedSwitchPort"},
			_jsii_.MemberProperty{JsiiProperty: "distributedSwitchPortInput", GoGetter: "DistributedSwitchPortInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "host", GoGetter: "Host"},
			_jsii_.MemberProperty{JsiiProperty: "hostInput", GoGetter: "HostInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4", GoGetter: "Ipv4"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4Input", GoGetter: "Ipv4Input"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6", GoGetter: "Ipv6"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Input", GoGetter: "Ipv6Input"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mac", GoGetter: "Mac"},
			_jsii_.MemberProperty{JsiiProperty: "macInput", GoGetter: "MacInput"},
			_jsii_.MemberProperty{JsiiProperty: "mtu", GoGetter: "Mtu"},
			_jsii_.MemberProperty{JsiiProperty: "mtuInput", GoGetter: "MtuInput"},
			_jsii_.MemberProperty{JsiiProperty: "netstack", GoGetter: "Netstack"},
			_jsii_.MemberProperty{JsiiProperty: "netstackInput", GoGetter: "NetstackInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "portgroup", GoGetter: "Portgroup"},
			_jsii_.MemberProperty{JsiiProperty: "portgroupInput", GoGetter: "PortgroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putIpv4", GoMethod: "PutIpv4"},
			_jsii_.MemberMethod{JsiiMethod: "putIpv6", GoMethod: "PutIpv6"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDistributedPortGroup", GoMethod: "ResetDistributedPortGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetDistributedSwitchPort", GoMethod: "ResetDistributedSwitchPort"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv4", GoMethod: "ResetIpv4"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6", GoMethod: "ResetIpv6"},
			_jsii_.MemberMethod{JsiiMethod: "resetMac", GoMethod: "ResetMac"},
			_jsii_.MemberMethod{JsiiMethod: "resetMtu", GoMethod: "ResetMtu"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetstack", GoMethod: "ResetNetstack"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPortgroup", GoMethod: "ResetPortgroup"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_Vnic{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vsphere.vnic.VnicConfig",
		reflect.TypeOf((*VnicConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vsphere.vnic.VnicIpv4",
		reflect.TypeOf((*VnicIpv4)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.vnic.VnicIpv4OutputReference",
		reflect.TypeOf((*VnicIpv4OutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dhcp", GoGetter: "Dhcp"},
			_jsii_.MemberProperty{JsiiProperty: "dhcpInput", GoGetter: "DhcpInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "gw", GoGetter: "Gw"},
			_jsii_.MemberProperty{JsiiProperty: "gwInput", GoGetter: "GwInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ip", GoGetter: "Ip"},
			_jsii_.MemberProperty{JsiiProperty: "ipInput", GoGetter: "IpInput"},
			_jsii_.MemberProperty{JsiiProperty: "netmask", GoGetter: "Netmask"},
			_jsii_.MemberProperty{JsiiProperty: "netmaskInput", GoGetter: "NetmaskInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDhcp", GoMethod: "ResetDhcp"},
			_jsii_.MemberMethod{JsiiMethod: "resetGw", GoMethod: "ResetGw"},
			_jsii_.MemberMethod{JsiiMethod: "resetIp", GoMethod: "ResetIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetmask", GoMethod: "ResetNetmask"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VnicIpv4OutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vsphere.vnic.VnicIpv6",
		reflect.TypeOf((*VnicIpv6)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.vnic.VnicIpv6OutputReference",
		reflect.TypeOf((*VnicIpv6OutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "addresses", GoGetter: "Addresses"},
			_jsii_.MemberProperty{JsiiProperty: "addressesInput", GoGetter: "AddressesInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoconfig", GoGetter: "Autoconfig"},
			_jsii_.MemberProperty{JsiiProperty: "autoconfigInput", GoGetter: "AutoconfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dhcp", GoGetter: "Dhcp"},
			_jsii_.MemberProperty{JsiiProperty: "dhcpInput", GoGetter: "DhcpInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "gw", GoGetter: "Gw"},
			_jsii_.MemberProperty{JsiiProperty: "gwInput", GoGetter: "GwInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAddresses", GoMethod: "ResetAddresses"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoconfig", GoMethod: "ResetAutoconfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetDhcp", GoMethod: "ResetDhcp"},
			_jsii_.MemberMethod{JsiiMethod: "resetGw", GoMethod: "ResetGw"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VnicIpv6OutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
