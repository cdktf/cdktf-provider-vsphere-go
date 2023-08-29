// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hostportgroup

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.hostPortGroup.HostPortGroup",
		reflect.TypeOf((*HostPortGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "activeNics", GoGetter: "ActiveNics"},
			_jsii_.MemberProperty{JsiiProperty: "activeNicsInput", GoGetter: "ActiveNicsInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowForgedTransmits", GoGetter: "AllowForgedTransmits"},
			_jsii_.MemberProperty{JsiiProperty: "allowForgedTransmitsInput", GoGetter: "AllowForgedTransmitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowMacChanges", GoGetter: "AllowMacChanges"},
			_jsii_.MemberProperty{JsiiProperty: "allowMacChangesInput", GoGetter: "AllowMacChangesInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowPromiscuous", GoGetter: "AllowPromiscuous"},
			_jsii_.MemberProperty{JsiiProperty: "allowPromiscuousInput", GoGetter: "AllowPromiscuousInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "checkBeacon", GoGetter: "CheckBeacon"},
			_jsii_.MemberProperty{JsiiProperty: "checkBeaconInput", GoGetter: "CheckBeaconInput"},
			_jsii_.MemberProperty{JsiiProperty: "computedPolicy", GoGetter: "ComputedPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "failback", GoGetter: "Failback"},
			_jsii_.MemberProperty{JsiiProperty: "failbackInput", GoGetter: "FailbackInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "hostSystemId", GoGetter: "HostSystemId"},
			_jsii_.MemberProperty{JsiiProperty: "hostSystemIdInput", GoGetter: "HostSystemIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notifySwitches", GoGetter: "NotifySwitches"},
			_jsii_.MemberProperty{JsiiProperty: "notifySwitchesInput", GoGetter: "NotifySwitchesInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ports", GoGetter: "Ports"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetActiveNics", GoMethod: "ResetActiveNics"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowForgedTransmits", GoMethod: "ResetAllowForgedTransmits"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowMacChanges", GoMethod: "ResetAllowMacChanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowPromiscuous", GoMethod: "ResetAllowPromiscuous"},
			_jsii_.MemberMethod{JsiiMethod: "resetCheckBeacon", GoMethod: "ResetCheckBeacon"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailback", GoMethod: "ResetFailback"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotifySwitches", GoMethod: "ResetNotifySwitches"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetShapingAverageBandwidth", GoMethod: "ResetShapingAverageBandwidth"},
			_jsii_.MemberMethod{JsiiMethod: "resetShapingBurstSize", GoMethod: "ResetShapingBurstSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetShapingEnabled", GoMethod: "ResetShapingEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetShapingPeakBandwidth", GoMethod: "ResetShapingPeakBandwidth"},
			_jsii_.MemberMethod{JsiiMethod: "resetStandbyNics", GoMethod: "ResetStandbyNics"},
			_jsii_.MemberMethod{JsiiMethod: "resetTeamingPolicy", GoMethod: "ResetTeamingPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetVlanId", GoMethod: "ResetVlanId"},
			_jsii_.MemberProperty{JsiiProperty: "shapingAverageBandwidth", GoGetter: "ShapingAverageBandwidth"},
			_jsii_.MemberProperty{JsiiProperty: "shapingAverageBandwidthInput", GoGetter: "ShapingAverageBandwidthInput"},
			_jsii_.MemberProperty{JsiiProperty: "shapingBurstSize", GoGetter: "ShapingBurstSize"},
			_jsii_.MemberProperty{JsiiProperty: "shapingBurstSizeInput", GoGetter: "ShapingBurstSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "shapingEnabled", GoGetter: "ShapingEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "shapingEnabledInput", GoGetter: "ShapingEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "shapingPeakBandwidth", GoGetter: "ShapingPeakBandwidth"},
			_jsii_.MemberProperty{JsiiProperty: "shapingPeakBandwidthInput", GoGetter: "ShapingPeakBandwidthInput"},
			_jsii_.MemberProperty{JsiiProperty: "standbyNics", GoGetter: "StandbyNics"},
			_jsii_.MemberProperty{JsiiProperty: "standbyNicsInput", GoGetter: "StandbyNicsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "teamingPolicy", GoGetter: "TeamingPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "teamingPolicyInput", GoGetter: "TeamingPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "virtualSwitchName", GoGetter: "VirtualSwitchName"},
			_jsii_.MemberProperty{JsiiProperty: "virtualSwitchNameInput", GoGetter: "VirtualSwitchNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "vlanId", GoGetter: "VlanId"},
			_jsii_.MemberProperty{JsiiProperty: "vlanIdInput", GoGetter: "VlanIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_HostPortGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vsphere.hostPortGroup.HostPortGroupConfig",
		reflect.TypeOf((*HostPortGroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vsphere.hostPortGroup.HostPortGroupPorts",
		reflect.TypeOf((*HostPortGroupPorts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.hostPortGroup.HostPortGroupPortsList",
		reflect.TypeOf((*HostPortGroupPortsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_HostPortGroupPortsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.hostPortGroup.HostPortGroupPortsOutputReference",
		reflect.TypeOf((*HostPortGroupPortsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "macAddresses", GoGetter: "MacAddresses"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_HostPortGroupPortsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
