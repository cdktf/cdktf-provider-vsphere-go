// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datavsphereovfvmtemplate

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.dataVsphereOvfVmTemplate.DataVsphereOvfVmTemplate",
		reflect.TypeOf((*DataVsphereOvfVmTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowUnverifiedSslCert", GoGetter: "AllowUnverifiedSslCert"},
			_jsii_.MemberProperty{JsiiProperty: "allowUnverifiedSslCertInput", GoGetter: "AllowUnverifiedSslCertInput"},
			_jsii_.MemberProperty{JsiiProperty: "alternateGuestName", GoGetter: "AlternateGuestName"},
			_jsii_.MemberProperty{JsiiProperty: "annotation", GoGetter: "Annotation"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "cpuHotAddEnabled", GoGetter: "CpuHotAddEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "cpuHotRemoveEnabled", GoGetter: "CpuHotRemoveEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "cpuPerformanceCountersEnabled", GoGetter: "CpuPerformanceCountersEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "datastoreId", GoGetter: "DatastoreId"},
			_jsii_.MemberProperty{JsiiProperty: "datastoreIdInput", GoGetter: "DatastoreIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentOption", GoGetter: "DeploymentOption"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentOptionInput", GoGetter: "DeploymentOptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "diskProvisioning", GoGetter: "DiskProvisioning"},
			_jsii_.MemberProperty{JsiiProperty: "diskProvisioningInput", GoGetter: "DiskProvisioningInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableHiddenProperties", GoGetter: "EnableHiddenProperties"},
			_jsii_.MemberProperty{JsiiProperty: "enableHiddenPropertiesInput", GoGetter: "EnableHiddenPropertiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "firmware", GoGetter: "Firmware"},
			_jsii_.MemberProperty{JsiiProperty: "folder", GoGetter: "Folder"},
			_jsii_.MemberProperty{JsiiProperty: "folderInput", GoGetter: "FolderInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "guestId", GoGetter: "GuestId"},
			_jsii_.MemberProperty{JsiiProperty: "hostSystemId", GoGetter: "HostSystemId"},
			_jsii_.MemberProperty{JsiiProperty: "hostSystemIdInput", GoGetter: "HostSystemIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "ideControllerCount", GoGetter: "IdeControllerCount"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipAllocationPolicy", GoGetter: "IpAllocationPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "ipAllocationPolicyInput", GoGetter: "IpAllocationPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipProtocol", GoGetter: "IpProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "ipProtocolInput", GoGetter: "IpProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "localOvfPath", GoGetter: "LocalOvfPath"},
			_jsii_.MemberProperty{JsiiProperty: "localOvfPathInput", GoGetter: "LocalOvfPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "memory", GoGetter: "Memory"},
			_jsii_.MemberProperty{JsiiProperty: "memoryHotAddEnabled", GoGetter: "MemoryHotAddEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "nestedHvEnabled", GoGetter: "NestedHvEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "numCoresPerSocket", GoGetter: "NumCoresPerSocket"},
			_jsii_.MemberProperty{JsiiProperty: "numCpus", GoGetter: "NumCpus"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ovfNetworkMap", GoGetter: "OvfNetworkMap"},
			_jsii_.MemberProperty{JsiiProperty: "ovfNetworkMapInput", GoGetter: "OvfNetworkMapInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "remoteOvfUrl", GoGetter: "RemoteOvfUrl"},
			_jsii_.MemberProperty{JsiiProperty: "remoteOvfUrlInput", GoGetter: "RemoteOvfUrlInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowUnverifiedSslCert", GoMethod: "ResetAllowUnverifiedSslCert"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatastoreId", GoMethod: "ResetDatastoreId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeploymentOption", GoMethod: "ResetDeploymentOption"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiskProvisioning", GoMethod: "ResetDiskProvisioning"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableHiddenProperties", GoMethod: "ResetEnableHiddenProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resetFolder", GoMethod: "ResetFolder"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpAllocationPolicy", GoMethod: "ResetIpAllocationPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpProtocol", GoMethod: "ResetIpProtocol"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocalOvfPath", GoMethod: "ResetLocalOvfPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOvfNetworkMap", GoMethod: "ResetOvfNetworkMap"},
			_jsii_.MemberMethod{JsiiMethod: "resetRemoteOvfUrl", GoMethod: "ResetRemoteOvfUrl"},
			_jsii_.MemberProperty{JsiiProperty: "resourcePoolId", GoGetter: "ResourcePoolId"},
			_jsii_.MemberProperty{JsiiProperty: "resourcePoolIdInput", GoGetter: "ResourcePoolIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "sataControllerCount", GoGetter: "SataControllerCount"},
			_jsii_.MemberProperty{JsiiProperty: "scsiControllerCount", GoGetter: "ScsiControllerCount"},
			_jsii_.MemberProperty{JsiiProperty: "scsiType", GoGetter: "ScsiType"},
			_jsii_.MemberProperty{JsiiProperty: "swapPlacementPolicy", GoGetter: "SwapPlacementPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataVsphereOvfVmTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vsphere.dataVsphereOvfVmTemplate.DataVsphereOvfVmTemplateConfig",
		reflect.TypeOf((*DataVsphereOvfVmTemplateConfig)(nil)).Elem(),
	)
}
