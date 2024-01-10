// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hostvirtualswitch

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.hostVirtualSwitch.HostVirtualSwitch",
		reflect.TypeOf((*HostVirtualSwitch)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "activeNics", GoGetter: "ActiveNics"},
			_jsii_.MemberProperty{JsiiProperty: "activeNicsInput", GoGetter: "ActiveNicsInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowForgedTransmits", GoGetter: "AllowForgedTransmits"},
			_jsii_.MemberProperty{JsiiProperty: "allowForgedTransmitsInput", GoGetter: "AllowForgedTransmitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowMacChanges", GoGetter: "AllowMacChanges"},
			_jsii_.MemberProperty{JsiiProperty: "allowMacChangesInput", GoGetter: "AllowMacChangesInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowPromiscuous", GoGetter: "AllowPromiscuous"},
			_jsii_.MemberProperty{JsiiProperty: "allowPromiscuousInput", GoGetter: "AllowPromiscuousInput"},
			_jsii_.MemberProperty{JsiiProperty: "beaconInterval", GoGetter: "BeaconInterval"},
			_jsii_.MemberProperty{JsiiProperty: "beaconIntervalInput", GoGetter: "BeaconIntervalInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "checkBeacon", GoGetter: "CheckBeacon"},
			_jsii_.MemberProperty{JsiiProperty: "checkBeaconInput", GoGetter: "CheckBeaconInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "hostSystemId", GoGetter: "HostSystemId"},
			_jsii_.MemberProperty{JsiiProperty: "hostSystemIdInput", GoGetter: "HostSystemIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "linkDiscoveryOperation", GoGetter: "LinkDiscoveryOperation"},
			_jsii_.MemberProperty{JsiiProperty: "linkDiscoveryOperationInput", GoGetter: "LinkDiscoveryOperationInput"},
			_jsii_.MemberProperty{JsiiProperty: "linkDiscoveryProtocol", GoGetter: "LinkDiscoveryProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "linkDiscoveryProtocolInput", GoGetter: "LinkDiscoveryProtocolInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "mtu", GoGetter: "Mtu"},
			_jsii_.MemberProperty{JsiiProperty: "mtuInput", GoGetter: "MtuInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkAdapters", GoGetter: "NetworkAdapters"},
			_jsii_.MemberProperty{JsiiProperty: "networkAdaptersInput", GoGetter: "NetworkAdaptersInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notifySwitches", GoGetter: "NotifySwitches"},
			_jsii_.MemberProperty{JsiiProperty: "notifySwitchesInput", GoGetter: "NotifySwitchesInput"},
			_jsii_.MemberProperty{JsiiProperty: "numberOfPorts", GoGetter: "NumberOfPorts"},
			_jsii_.MemberProperty{JsiiProperty: "numberOfPortsInput", GoGetter: "NumberOfPortsInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowForgedTransmits", GoMethod: "ResetAllowForgedTransmits"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowMacChanges", GoMethod: "ResetAllowMacChanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowPromiscuous", GoMethod: "ResetAllowPromiscuous"},
			_jsii_.MemberMethod{JsiiMethod: "resetBeaconInterval", GoMethod: "ResetBeaconInterval"},
			_jsii_.MemberMethod{JsiiMethod: "resetCheckBeacon", GoMethod: "ResetCheckBeacon"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailback", GoMethod: "ResetFailback"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLinkDiscoveryOperation", GoMethod: "ResetLinkDiscoveryOperation"},
			_jsii_.MemberMethod{JsiiMethod: "resetLinkDiscoveryProtocol", GoMethod: "ResetLinkDiscoveryProtocol"},
			_jsii_.MemberMethod{JsiiMethod: "resetMtu", GoMethod: "ResetMtu"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotifySwitches", GoMethod: "ResetNotifySwitches"},
			_jsii_.MemberMethod{JsiiMethod: "resetNumberOfPorts", GoMethod: "ResetNumberOfPorts"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetShapingAverageBandwidth", GoMethod: "ResetShapingAverageBandwidth"},
			_jsii_.MemberMethod{JsiiMethod: "resetShapingBurstSize", GoMethod: "ResetShapingBurstSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetShapingEnabled", GoMethod: "ResetShapingEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetShapingPeakBandwidth", GoMethod: "ResetShapingPeakBandwidth"},
			_jsii_.MemberMethod{JsiiMethod: "resetStandbyNics", GoMethod: "ResetStandbyNics"},
			_jsii_.MemberMethod{JsiiMethod: "resetTeamingPolicy", GoMethod: "ResetTeamingPolicy"},
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
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "teamingPolicy", GoGetter: "TeamingPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "teamingPolicyInput", GoGetter: "TeamingPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_HostVirtualSwitch{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vsphere.hostVirtualSwitch.HostVirtualSwitchConfig",
		reflect.TypeOf((*HostVirtualSwitchConfig)(nil)).Elem(),
	)
}
