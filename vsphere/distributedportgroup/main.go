// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package distributedportgroup

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.distributedPortGroup.DistributedPortGroup",
		reflect.TypeOf((*DistributedPortGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "activeUplinks", GoGetter: "ActiveUplinks"},
			_jsii_.MemberProperty{JsiiProperty: "activeUplinksInput", GoGetter: "ActiveUplinksInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowForgedTransmits", GoGetter: "AllowForgedTransmits"},
			_jsii_.MemberProperty{JsiiProperty: "allowForgedTransmitsInput", GoGetter: "AllowForgedTransmitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowMacChanges", GoGetter: "AllowMacChanges"},
			_jsii_.MemberProperty{JsiiProperty: "allowMacChangesInput", GoGetter: "AllowMacChangesInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowPromiscuous", GoGetter: "AllowPromiscuous"},
			_jsii_.MemberProperty{JsiiProperty: "allowPromiscuousInput", GoGetter: "AllowPromiscuousInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoExpand", GoGetter: "AutoExpand"},
			_jsii_.MemberProperty{JsiiProperty: "autoExpandInput", GoGetter: "AutoExpandInput"},
			_jsii_.MemberProperty{JsiiProperty: "blockAllPorts", GoGetter: "BlockAllPorts"},
			_jsii_.MemberProperty{JsiiProperty: "blockAllPortsInput", GoGetter: "BlockAllPortsInput"},
			_jsii_.MemberProperty{JsiiProperty: "blockOverrideAllowed", GoGetter: "BlockOverrideAllowed"},
			_jsii_.MemberProperty{JsiiProperty: "blockOverrideAllowedInput", GoGetter: "BlockOverrideAllowedInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "checkBeacon", GoGetter: "CheckBeacon"},
			_jsii_.MemberProperty{JsiiProperty: "checkBeaconInput", GoGetter: "CheckBeaconInput"},
			_jsii_.MemberProperty{JsiiProperty: "configVersion", GoGetter: "ConfigVersion"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customAttributes", GoGetter: "CustomAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "customAttributesInput", GoGetter: "CustomAttributesInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "directpathGen2Allowed", GoGetter: "DirectpathGen2Allowed"},
			_jsii_.MemberProperty{JsiiProperty: "directpathGen2AllowedInput", GoGetter: "DirectpathGen2AllowedInput"},
			_jsii_.MemberProperty{JsiiProperty: "distributedVirtualSwitchUuid", GoGetter: "DistributedVirtualSwitchUuid"},
			_jsii_.MemberProperty{JsiiProperty: "distributedVirtualSwitchUuidInput", GoGetter: "DistributedVirtualSwitchUuidInput"},
			_jsii_.MemberProperty{JsiiProperty: "egressShapingAverageBandwidth", GoGetter: "EgressShapingAverageBandwidth"},
			_jsii_.MemberProperty{JsiiProperty: "egressShapingAverageBandwidthInput", GoGetter: "EgressShapingAverageBandwidthInput"},
			_jsii_.MemberProperty{JsiiProperty: "egressShapingBurstSize", GoGetter: "EgressShapingBurstSize"},
			_jsii_.MemberProperty{JsiiProperty: "egressShapingBurstSizeInput", GoGetter: "EgressShapingBurstSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "egressShapingEnabled", GoGetter: "EgressShapingEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "egressShapingEnabledInput", GoGetter: "EgressShapingEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "egressShapingPeakBandwidth", GoGetter: "EgressShapingPeakBandwidth"},
			_jsii_.MemberProperty{JsiiProperty: "egressShapingPeakBandwidthInput", GoGetter: "EgressShapingPeakBandwidthInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "ingressShapingAverageBandwidth", GoGetter: "IngressShapingAverageBandwidth"},
			_jsii_.MemberProperty{JsiiProperty: "ingressShapingAverageBandwidthInput", GoGetter: "IngressShapingAverageBandwidthInput"},
			_jsii_.MemberProperty{JsiiProperty: "ingressShapingBurstSize", GoGetter: "IngressShapingBurstSize"},
			_jsii_.MemberProperty{JsiiProperty: "ingressShapingBurstSizeInput", GoGetter: "IngressShapingBurstSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "ingressShapingEnabled", GoGetter: "IngressShapingEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "ingressShapingEnabledInput", GoGetter: "IngressShapingEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "ingressShapingPeakBandwidth", GoGetter: "IngressShapingPeakBandwidth"},
			_jsii_.MemberProperty{JsiiProperty: "ingressShapingPeakBandwidthInput", GoGetter: "IngressShapingPeakBandwidthInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "lacpEnabled", GoGetter: "LacpEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "lacpEnabledInput", GoGetter: "LacpEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "lacpMode", GoGetter: "LacpMode"},
			_jsii_.MemberProperty{JsiiProperty: "lacpModeInput", GoGetter: "LacpModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "livePortMovingAllowed", GoGetter: "LivePortMovingAllowed"},
			_jsii_.MemberProperty{JsiiProperty: "livePortMovingAllowedInput", GoGetter: "LivePortMovingAllowedInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "netflowEnabled", GoGetter: "NetflowEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "netflowEnabledInput", GoGetter: "NetflowEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "netflowOverrideAllowed", GoGetter: "NetflowOverrideAllowed"},
			_jsii_.MemberProperty{JsiiProperty: "netflowOverrideAllowedInput", GoGetter: "NetflowOverrideAllowedInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkResourcePoolKey", GoGetter: "NetworkResourcePoolKey"},
			_jsii_.MemberProperty{JsiiProperty: "networkResourcePoolKeyInput", GoGetter: "NetworkResourcePoolKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkResourcePoolOverrideAllowed", GoGetter: "NetworkResourcePoolOverrideAllowed"},
			_jsii_.MemberProperty{JsiiProperty: "networkResourcePoolOverrideAllowedInput", GoGetter: "NetworkResourcePoolOverrideAllowedInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notifySwitches", GoGetter: "NotifySwitches"},
			_jsii_.MemberProperty{JsiiProperty: "notifySwitchesInput", GoGetter: "NotifySwitchesInput"},
			_jsii_.MemberProperty{JsiiProperty: "numberOfPorts", GoGetter: "NumberOfPorts"},
			_jsii_.MemberProperty{JsiiProperty: "numberOfPortsInput", GoGetter: "NumberOfPortsInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "portConfigResetAtDisconnect", GoGetter: "PortConfigResetAtDisconnect"},
			_jsii_.MemberProperty{JsiiProperty: "portConfigResetAtDisconnectInput", GoGetter: "PortConfigResetAtDisconnectInput"},
			_jsii_.MemberProperty{JsiiProperty: "portNameFormat", GoGetter: "PortNameFormat"},
			_jsii_.MemberProperty{JsiiProperty: "portNameFormatInput", GoGetter: "PortNameFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "portPrivateSecondaryVlanId", GoGetter: "PortPrivateSecondaryVlanId"},
			_jsii_.MemberProperty{JsiiProperty: "portPrivateSecondaryVlanIdInput", GoGetter: "PortPrivateSecondaryVlanIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putVlanRange", GoMethod: "PutVlanRange"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetActiveUplinks", GoMethod: "ResetActiveUplinks"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowForgedTransmits", GoMethod: "ResetAllowForgedTransmits"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowMacChanges", GoMethod: "ResetAllowMacChanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowPromiscuous", GoMethod: "ResetAllowPromiscuous"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoExpand", GoMethod: "ResetAutoExpand"},
			_jsii_.MemberMethod{JsiiMethod: "resetBlockAllPorts", GoMethod: "ResetBlockAllPorts"},
			_jsii_.MemberMethod{JsiiMethod: "resetBlockOverrideAllowed", GoMethod: "ResetBlockOverrideAllowed"},
			_jsii_.MemberMethod{JsiiMethod: "resetCheckBeacon", GoMethod: "ResetCheckBeacon"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomAttributes", GoMethod: "ResetCustomAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetDirectpathGen2Allowed", GoMethod: "ResetDirectpathGen2Allowed"},
			_jsii_.MemberMethod{JsiiMethod: "resetEgressShapingAverageBandwidth", GoMethod: "ResetEgressShapingAverageBandwidth"},
			_jsii_.MemberMethod{JsiiMethod: "resetEgressShapingBurstSize", GoMethod: "ResetEgressShapingBurstSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetEgressShapingEnabled", GoMethod: "ResetEgressShapingEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetEgressShapingPeakBandwidth", GoMethod: "ResetEgressShapingPeakBandwidth"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailback", GoMethod: "ResetFailback"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIngressShapingAverageBandwidth", GoMethod: "ResetIngressShapingAverageBandwidth"},
			_jsii_.MemberMethod{JsiiMethod: "resetIngressShapingBurstSize", GoMethod: "ResetIngressShapingBurstSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetIngressShapingEnabled", GoMethod: "ResetIngressShapingEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetIngressShapingPeakBandwidth", GoMethod: "ResetIngressShapingPeakBandwidth"},
			_jsii_.MemberMethod{JsiiMethod: "resetLacpEnabled", GoMethod: "ResetLacpEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLacpMode", GoMethod: "ResetLacpMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetLivePortMovingAllowed", GoMethod: "ResetLivePortMovingAllowed"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetflowEnabled", GoMethod: "ResetNetflowEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetflowOverrideAllowed", GoMethod: "ResetNetflowOverrideAllowed"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkResourcePoolKey", GoMethod: "ResetNetworkResourcePoolKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkResourcePoolOverrideAllowed", GoMethod: "ResetNetworkResourcePoolOverrideAllowed"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotifySwitches", GoMethod: "ResetNotifySwitches"},
			_jsii_.MemberMethod{JsiiMethod: "resetNumberOfPorts", GoMethod: "ResetNumberOfPorts"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPortConfigResetAtDisconnect", GoMethod: "ResetPortConfigResetAtDisconnect"},
			_jsii_.MemberMethod{JsiiMethod: "resetPortNameFormat", GoMethod: "ResetPortNameFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetPortPrivateSecondaryVlanId", GoMethod: "ResetPortPrivateSecondaryVlanId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityPolicyOverrideAllowed", GoMethod: "ResetSecurityPolicyOverrideAllowed"},
			_jsii_.MemberMethod{JsiiMethod: "resetShapingOverrideAllowed", GoMethod: "ResetShapingOverrideAllowed"},
			_jsii_.MemberMethod{JsiiMethod: "resetStandbyUplinks", GoMethod: "ResetStandbyUplinks"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTeamingPolicy", GoMethod: "ResetTeamingPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrafficFilterOverrideAllowed", GoMethod: "ResetTrafficFilterOverrideAllowed"},
			_jsii_.MemberMethod{JsiiMethod: "resetTxUplink", GoMethod: "ResetTxUplink"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resetUplinkTeamingOverrideAllowed", GoMethod: "ResetUplinkTeamingOverrideAllowed"},
			_jsii_.MemberMethod{JsiiMethod: "resetVlanId", GoMethod: "ResetVlanId"},
			_jsii_.MemberMethod{JsiiMethod: "resetVlanOverrideAllowed", GoMethod: "ResetVlanOverrideAllowed"},
			_jsii_.MemberMethod{JsiiMethod: "resetVlanRange", GoMethod: "ResetVlanRange"},
			_jsii_.MemberProperty{JsiiProperty: "securityPolicyOverrideAllowed", GoGetter: "SecurityPolicyOverrideAllowed"},
			_jsii_.MemberProperty{JsiiProperty: "securityPolicyOverrideAllowedInput", GoGetter: "SecurityPolicyOverrideAllowedInput"},
			_jsii_.MemberProperty{JsiiProperty: "shapingOverrideAllowed", GoGetter: "ShapingOverrideAllowed"},
			_jsii_.MemberProperty{JsiiProperty: "shapingOverrideAllowedInput", GoGetter: "ShapingOverrideAllowedInput"},
			_jsii_.MemberProperty{JsiiProperty: "standbyUplinks", GoGetter: "StandbyUplinks"},
			_jsii_.MemberProperty{JsiiProperty: "standbyUplinksInput", GoGetter: "StandbyUplinksInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "teamingPolicy", GoGetter: "TeamingPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "teamingPolicyInput", GoGetter: "TeamingPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "trafficFilterOverrideAllowed", GoGetter: "TrafficFilterOverrideAllowed"},
			_jsii_.MemberProperty{JsiiProperty: "trafficFilterOverrideAllowedInput", GoGetter: "TrafficFilterOverrideAllowedInput"},
			_jsii_.MemberProperty{JsiiProperty: "txUplink", GoGetter: "TxUplink"},
			_jsii_.MemberProperty{JsiiProperty: "txUplinkInput", GoGetter: "TxUplinkInput"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "uplinkTeamingOverrideAllowed", GoGetter: "UplinkTeamingOverrideAllowed"},
			_jsii_.MemberProperty{JsiiProperty: "uplinkTeamingOverrideAllowedInput", GoGetter: "UplinkTeamingOverrideAllowedInput"},
			_jsii_.MemberProperty{JsiiProperty: "vlanId", GoGetter: "VlanId"},
			_jsii_.MemberProperty{JsiiProperty: "vlanIdInput", GoGetter: "VlanIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "vlanOverrideAllowed", GoGetter: "VlanOverrideAllowed"},
			_jsii_.MemberProperty{JsiiProperty: "vlanOverrideAllowedInput", GoGetter: "VlanOverrideAllowedInput"},
			_jsii_.MemberProperty{JsiiProperty: "vlanRange", GoGetter: "VlanRange"},
			_jsii_.MemberProperty{JsiiProperty: "vlanRangeInput", GoGetter: "VlanRangeInput"},
		},
		func() interface{} {
			j := jsiiProxy_DistributedPortGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vsphere.distributedPortGroup.DistributedPortGroupConfig",
		reflect.TypeOf((*DistributedPortGroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vsphere.distributedPortGroup.DistributedPortGroupVlanRange",
		reflect.TypeOf((*DistributedPortGroupVlanRange)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.distributedPortGroup.DistributedPortGroupVlanRangeList",
		reflect.TypeOf((*DistributedPortGroupVlanRangeList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DistributedPortGroupVlanRangeList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.distributedPortGroup.DistributedPortGroupVlanRangeOutputReference",
		reflect.TypeOf((*DistributedPortGroupVlanRangeOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "maxVlan", GoGetter: "MaxVlan"},
			_jsii_.MemberProperty{JsiiProperty: "maxVlanInput", GoGetter: "MaxVlanInput"},
			_jsii_.MemberProperty{JsiiProperty: "minVlan", GoGetter: "MinVlan"},
			_jsii_.MemberProperty{JsiiProperty: "minVlanInput", GoGetter: "MinVlanInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DistributedPortGroupVlanRangeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
