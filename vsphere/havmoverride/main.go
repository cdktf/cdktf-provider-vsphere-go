// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package havmoverride

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.haVmOverride.HaVmOverride",
		reflect.TypeOf((*HaVmOverride)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "computeClusterId", GoGetter: "ComputeClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "computeClusterIdInput", GoGetter: "ComputeClusterIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
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
			_jsii_.MemberProperty{JsiiProperty: "haDatastoreApdRecoveryAction", GoGetter: "HaDatastoreApdRecoveryAction"},
			_jsii_.MemberProperty{JsiiProperty: "haDatastoreApdRecoveryActionInput", GoGetter: "HaDatastoreApdRecoveryActionInput"},
			_jsii_.MemberProperty{JsiiProperty: "haDatastoreApdResponse", GoGetter: "HaDatastoreApdResponse"},
			_jsii_.MemberProperty{JsiiProperty: "haDatastoreApdResponseDelay", GoGetter: "HaDatastoreApdResponseDelay"},
			_jsii_.MemberProperty{JsiiProperty: "haDatastoreApdResponseDelayInput", GoGetter: "HaDatastoreApdResponseDelayInput"},
			_jsii_.MemberProperty{JsiiProperty: "haDatastoreApdResponseInput", GoGetter: "HaDatastoreApdResponseInput"},
			_jsii_.MemberProperty{JsiiProperty: "haDatastorePdlResponse", GoGetter: "HaDatastorePdlResponse"},
			_jsii_.MemberProperty{JsiiProperty: "haDatastorePdlResponseInput", GoGetter: "HaDatastorePdlResponseInput"},
			_jsii_.MemberProperty{JsiiProperty: "haHostIsolationResponse", GoGetter: "HaHostIsolationResponse"},
			_jsii_.MemberProperty{JsiiProperty: "haHostIsolationResponseInput", GoGetter: "HaHostIsolationResponseInput"},
			_jsii_.MemberProperty{JsiiProperty: "haVmFailureInterval", GoGetter: "HaVmFailureInterval"},
			_jsii_.MemberProperty{JsiiProperty: "haVmFailureIntervalInput", GoGetter: "HaVmFailureIntervalInput"},
			_jsii_.MemberProperty{JsiiProperty: "haVmMaximumFailureWindow", GoGetter: "HaVmMaximumFailureWindow"},
			_jsii_.MemberProperty{JsiiProperty: "haVmMaximumFailureWindowInput", GoGetter: "HaVmMaximumFailureWindowInput"},
			_jsii_.MemberProperty{JsiiProperty: "haVmMaximumResets", GoGetter: "HaVmMaximumResets"},
			_jsii_.MemberProperty{JsiiProperty: "haVmMaximumResetsInput", GoGetter: "HaVmMaximumResetsInput"},
			_jsii_.MemberProperty{JsiiProperty: "haVmMinimumUptime", GoGetter: "HaVmMinimumUptime"},
			_jsii_.MemberProperty{JsiiProperty: "haVmMinimumUptimeInput", GoGetter: "HaVmMinimumUptimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "haVmMonitoring", GoGetter: "HaVmMonitoring"},
			_jsii_.MemberProperty{JsiiProperty: "haVmMonitoringInput", GoGetter: "HaVmMonitoringInput"},
			_jsii_.MemberProperty{JsiiProperty: "haVmMonitoringUseClusterDefaults", GoGetter: "HaVmMonitoringUseClusterDefaults"},
			_jsii_.MemberProperty{JsiiProperty: "haVmMonitoringUseClusterDefaultsInput", GoGetter: "HaVmMonitoringUseClusterDefaultsInput"},
			_jsii_.MemberProperty{JsiiProperty: "haVmRestartPriority", GoGetter: "HaVmRestartPriority"},
			_jsii_.MemberProperty{JsiiProperty: "haVmRestartPriorityInput", GoGetter: "HaVmRestartPriorityInput"},
			_jsii_.MemberProperty{JsiiProperty: "haVmRestartTimeout", GoGetter: "HaVmRestartTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "haVmRestartTimeoutInput", GoGetter: "HaVmRestartTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetHaDatastoreApdRecoveryAction", GoMethod: "ResetHaDatastoreApdRecoveryAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetHaDatastoreApdResponse", GoMethod: "ResetHaDatastoreApdResponse"},
			_jsii_.MemberMethod{JsiiMethod: "resetHaDatastoreApdResponseDelay", GoMethod: "ResetHaDatastoreApdResponseDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resetHaDatastorePdlResponse", GoMethod: "ResetHaDatastorePdlResponse"},
			_jsii_.MemberMethod{JsiiMethod: "resetHaHostIsolationResponse", GoMethod: "ResetHaHostIsolationResponse"},
			_jsii_.MemberMethod{JsiiMethod: "resetHaVmFailureInterval", GoMethod: "ResetHaVmFailureInterval"},
			_jsii_.MemberMethod{JsiiMethod: "resetHaVmMaximumFailureWindow", GoMethod: "ResetHaVmMaximumFailureWindow"},
			_jsii_.MemberMethod{JsiiMethod: "resetHaVmMaximumResets", GoMethod: "ResetHaVmMaximumResets"},
			_jsii_.MemberMethod{JsiiMethod: "resetHaVmMinimumUptime", GoMethod: "ResetHaVmMinimumUptime"},
			_jsii_.MemberMethod{JsiiMethod: "resetHaVmMonitoring", GoMethod: "ResetHaVmMonitoring"},
			_jsii_.MemberMethod{JsiiMethod: "resetHaVmMonitoringUseClusterDefaults", GoMethod: "ResetHaVmMonitoringUseClusterDefaults"},
			_jsii_.MemberMethod{JsiiMethod: "resetHaVmRestartPriority", GoMethod: "ResetHaVmRestartPriority"},
			_jsii_.MemberMethod{JsiiMethod: "resetHaVmRestartTimeout", GoMethod: "ResetHaVmRestartTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "virtualMachineId", GoGetter: "VirtualMachineId"},
			_jsii_.MemberProperty{JsiiProperty: "virtualMachineIdInput", GoGetter: "VirtualMachineIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_HaVmOverride{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vsphere.haVmOverride.HaVmOverrideConfig",
		reflect.TypeOf((*HaVmOverrideConfig)(nil)).Elem(),
	)
}
