// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datavspherevirtualmachine

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.dataVsphereVirtualMachine.DataVsphereVirtualMachine",
		reflect.TypeOf((*DataVsphereVirtualMachine)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alternateGuestName", GoGetter: "AlternateGuestName"},
			_jsii_.MemberProperty{JsiiProperty: "alternateGuestNameInput", GoGetter: "AlternateGuestNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "annotation", GoGetter: "Annotation"},
			_jsii_.MemberProperty{JsiiProperty: "annotationInput", GoGetter: "AnnotationInput"},
			_jsii_.MemberProperty{JsiiProperty: "bootDelay", GoGetter: "BootDelay"},
			_jsii_.MemberProperty{JsiiProperty: "bootDelayInput", GoGetter: "BootDelayInput"},
			_jsii_.MemberProperty{JsiiProperty: "bootRetryDelay", GoGetter: "BootRetryDelay"},
			_jsii_.MemberProperty{JsiiProperty: "bootRetryDelayInput", GoGetter: "BootRetryDelayInput"},
			_jsii_.MemberProperty{JsiiProperty: "bootRetryEnabled", GoGetter: "BootRetryEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "bootRetryEnabledInput", GoGetter: "BootRetryEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "changeVersion", GoGetter: "ChangeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "cpuHotAddEnabled", GoGetter: "CpuHotAddEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "cpuHotAddEnabledInput", GoGetter: "CpuHotAddEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "cpuHotRemoveEnabled", GoGetter: "CpuHotRemoveEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "cpuHotRemoveEnabledInput", GoGetter: "CpuHotRemoveEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "cpuLimit", GoGetter: "CpuLimit"},
			_jsii_.MemberProperty{JsiiProperty: "cpuLimitInput", GoGetter: "CpuLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "cpuPerformanceCountersEnabled", GoGetter: "CpuPerformanceCountersEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "cpuPerformanceCountersEnabledInput", GoGetter: "CpuPerformanceCountersEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "cpuReservation", GoGetter: "CpuReservation"},
			_jsii_.MemberProperty{JsiiProperty: "cpuReservationInput", GoGetter: "CpuReservationInput"},
			_jsii_.MemberProperty{JsiiProperty: "cpuShareCount", GoGetter: "CpuShareCount"},
			_jsii_.MemberProperty{JsiiProperty: "cpuShareCountInput", GoGetter: "CpuShareCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "cpuShareLevel", GoGetter: "CpuShareLevel"},
			_jsii_.MemberProperty{JsiiProperty: "cpuShareLevelInput", GoGetter: "CpuShareLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "datacenterId", GoGetter: "DatacenterId"},
			_jsii_.MemberProperty{JsiiProperty: "datacenterIdInput", GoGetter: "DatacenterIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultIpAddress", GoGetter: "DefaultIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "disks", GoGetter: "Disks"},
			_jsii_.MemberProperty{JsiiProperty: "efiSecureBootEnabled", GoGetter: "EfiSecureBootEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "efiSecureBootEnabledInput", GoGetter: "EfiSecureBootEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableDiskUuid", GoGetter: "EnableDiskUuid"},
			_jsii_.MemberProperty{JsiiProperty: "enableDiskUuidInput", GoGetter: "EnableDiskUuidInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableLogging", GoGetter: "EnableLogging"},
			_jsii_.MemberProperty{JsiiProperty: "enableLoggingInput", GoGetter: "EnableLoggingInput"},
			_jsii_.MemberProperty{JsiiProperty: "eptRviMode", GoGetter: "EptRviMode"},
			_jsii_.MemberProperty{JsiiProperty: "eptRviModeInput", GoGetter: "EptRviModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "extraConfig", GoGetter: "ExtraConfig"},
			_jsii_.MemberProperty{JsiiProperty: "extraConfigInput", GoGetter: "ExtraConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "extraConfigRebootRequired", GoGetter: "ExtraConfigRebootRequired"},
			_jsii_.MemberProperty{JsiiProperty: "extraConfigRebootRequiredInput", GoGetter: "ExtraConfigRebootRequiredInput"},
			_jsii_.MemberProperty{JsiiProperty: "firmware", GoGetter: "Firmware"},
			_jsii_.MemberProperty{JsiiProperty: "firmwareInput", GoGetter: "FirmwareInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "guestIdInput", GoGetter: "GuestIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "guestIpAddresses", GoGetter: "GuestIpAddresses"},
			_jsii_.MemberProperty{JsiiProperty: "hardwareVersion", GoGetter: "HardwareVersion"},
			_jsii_.MemberProperty{JsiiProperty: "hardwareVersionInput", GoGetter: "HardwareVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "hvMode", GoGetter: "HvMode"},
			_jsii_.MemberProperty{JsiiProperty: "hvModeInput", GoGetter: "HvModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "ideControllerScanCount", GoGetter: "IdeControllerScanCount"},
			_jsii_.MemberProperty{JsiiProperty: "ideControllerScanCountInput", GoGetter: "IdeControllerScanCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "instanceUuid", GoGetter: "InstanceUuid"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "latencySensitivity", GoGetter: "LatencySensitivity"},
			_jsii_.MemberProperty{JsiiProperty: "latencySensitivityInput", GoGetter: "LatencySensitivityInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "memory", GoGetter: "Memory"},
			_jsii_.MemberProperty{JsiiProperty: "memoryHotAddEnabled", GoGetter: "MemoryHotAddEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "memoryHotAddEnabledInput", GoGetter: "MemoryHotAddEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "memoryInput", GoGetter: "MemoryInput"},
			_jsii_.MemberProperty{JsiiProperty: "memoryLimit", GoGetter: "MemoryLimit"},
			_jsii_.MemberProperty{JsiiProperty: "memoryLimitInput", GoGetter: "MemoryLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "memoryReservation", GoGetter: "MemoryReservation"},
			_jsii_.MemberProperty{JsiiProperty: "memoryReservationInput", GoGetter: "MemoryReservationInput"},
			_jsii_.MemberProperty{JsiiProperty: "memoryReservationLockedToMax", GoGetter: "MemoryReservationLockedToMax"},
			_jsii_.MemberProperty{JsiiProperty: "memoryReservationLockedToMaxInput", GoGetter: "MemoryReservationLockedToMaxInput"},
			_jsii_.MemberProperty{JsiiProperty: "memoryShareCount", GoGetter: "MemoryShareCount"},
			_jsii_.MemberProperty{JsiiProperty: "memoryShareCountInput", GoGetter: "MemoryShareCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "memoryShareLevel", GoGetter: "MemoryShareLevel"},
			_jsii_.MemberProperty{JsiiProperty: "memoryShareLevelInput", GoGetter: "MemoryShareLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "moid", GoGetter: "Moid"},
			_jsii_.MemberProperty{JsiiProperty: "moidInput", GoGetter: "MoidInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "nestedHvEnabled", GoGetter: "NestedHvEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "nestedHvEnabledInput", GoGetter: "NestedHvEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaces", GoGetter: "NetworkInterfaces"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceTypes", GoGetter: "NetworkInterfaceTypes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "numCoresPerSocket", GoGetter: "NumCoresPerSocket"},
			_jsii_.MemberProperty{JsiiProperty: "numCoresPerSocketInput", GoGetter: "NumCoresPerSocketInput"},
			_jsii_.MemberProperty{JsiiProperty: "numCpus", GoGetter: "NumCpus"},
			_jsii_.MemberProperty{JsiiProperty: "numCpusInput", GoGetter: "NumCpusInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putVapp", GoMethod: "PutVapp"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "replaceTrigger", GoGetter: "ReplaceTrigger"},
			_jsii_.MemberProperty{JsiiProperty: "replaceTriggerInput", GoGetter: "ReplaceTriggerInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlternateGuestName", GoMethod: "ResetAlternateGuestName"},
			_jsii_.MemberMethod{JsiiMethod: "resetAnnotation", GoMethod: "ResetAnnotation"},
			_jsii_.MemberMethod{JsiiMethod: "resetBootDelay", GoMethod: "ResetBootDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resetBootRetryDelay", GoMethod: "ResetBootRetryDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resetBootRetryEnabled", GoMethod: "ResetBootRetryEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpuHotAddEnabled", GoMethod: "ResetCpuHotAddEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpuHotRemoveEnabled", GoMethod: "ResetCpuHotRemoveEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpuLimit", GoMethod: "ResetCpuLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpuPerformanceCountersEnabled", GoMethod: "ResetCpuPerformanceCountersEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpuReservation", GoMethod: "ResetCpuReservation"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpuShareCount", GoMethod: "ResetCpuShareCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpuShareLevel", GoMethod: "ResetCpuShareLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatacenterId", GoMethod: "ResetDatacenterId"},
			_jsii_.MemberMethod{JsiiMethod: "resetEfiSecureBootEnabled", GoMethod: "ResetEfiSecureBootEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableDiskUuid", GoMethod: "ResetEnableDiskUuid"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableLogging", GoMethod: "ResetEnableLogging"},
			_jsii_.MemberMethod{JsiiMethod: "resetEptRviMode", GoMethod: "ResetEptRviMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetExtraConfig", GoMethod: "ResetExtraConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetExtraConfigRebootRequired", GoMethod: "ResetExtraConfigRebootRequired"},
			_jsii_.MemberMethod{JsiiMethod: "resetFirmware", GoMethod: "ResetFirmware"},
			_jsii_.MemberMethod{JsiiMethod: "resetFolder", GoMethod: "ResetFolder"},
			_jsii_.MemberMethod{JsiiMethod: "resetGuestId", GoMethod: "ResetGuestId"},
			_jsii_.MemberMethod{JsiiMethod: "resetHardwareVersion", GoMethod: "ResetHardwareVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetHvMode", GoMethod: "ResetHvMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdeControllerScanCount", GoMethod: "ResetIdeControllerScanCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetLatencySensitivity", GoMethod: "ResetLatencySensitivity"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemory", GoMethod: "ResetMemory"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemoryHotAddEnabled", GoMethod: "ResetMemoryHotAddEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemoryLimit", GoMethod: "ResetMemoryLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemoryReservation", GoMethod: "ResetMemoryReservation"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemoryReservationLockedToMax", GoMethod: "ResetMemoryReservationLockedToMax"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemoryShareCount", GoMethod: "ResetMemoryShareCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemoryShareLevel", GoMethod: "ResetMemoryShareLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetMoid", GoMethod: "ResetMoid"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetNestedHvEnabled", GoMethod: "ResetNestedHvEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetNumCoresPerSocket", GoMethod: "ResetNumCoresPerSocket"},
			_jsii_.MemberMethod{JsiiMethod: "resetNumCpus", GoMethod: "ResetNumCpus"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplaceTrigger", GoMethod: "ResetReplaceTrigger"},
			_jsii_.MemberMethod{JsiiMethod: "resetRunToolsScriptsAfterPowerOn", GoMethod: "ResetRunToolsScriptsAfterPowerOn"},
			_jsii_.MemberMethod{JsiiMethod: "resetRunToolsScriptsAfterResume", GoMethod: "ResetRunToolsScriptsAfterResume"},
			_jsii_.MemberMethod{JsiiMethod: "resetRunToolsScriptsBeforeGuestReboot", GoMethod: "ResetRunToolsScriptsBeforeGuestReboot"},
			_jsii_.MemberMethod{JsiiMethod: "resetRunToolsScriptsBeforeGuestShutdown", GoMethod: "ResetRunToolsScriptsBeforeGuestShutdown"},
			_jsii_.MemberMethod{JsiiMethod: "resetRunToolsScriptsBeforeGuestStandby", GoMethod: "ResetRunToolsScriptsBeforeGuestStandby"},
			_jsii_.MemberMethod{JsiiMethod: "resetSataControllerScanCount", GoMethod: "ResetSataControllerScanCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetScsiControllerScanCount", GoMethod: "ResetScsiControllerScanCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetStoragePolicyId", GoMethod: "ResetStoragePolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSwapPlacementPolicy", GoMethod: "ResetSwapPlacementPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetSyncTimeWithHost", GoMethod: "ResetSyncTimeWithHost"},
			_jsii_.MemberMethod{JsiiMethod: "resetSyncTimeWithHostPeriodically", GoMethod: "ResetSyncTimeWithHostPeriodically"},
			_jsii_.MemberMethod{JsiiMethod: "resetToolsUpgradePolicy", GoMethod: "ResetToolsUpgradePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetUuid", GoMethod: "ResetUuid"},
			_jsii_.MemberMethod{JsiiMethod: "resetVapp", GoMethod: "ResetVapp"},
			_jsii_.MemberMethod{JsiiMethod: "resetVbsEnabled", GoMethod: "ResetVbsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetVvtdEnabled", GoMethod: "ResetVvtdEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "runToolsScriptsAfterPowerOn", GoGetter: "RunToolsScriptsAfterPowerOn"},
			_jsii_.MemberProperty{JsiiProperty: "runToolsScriptsAfterPowerOnInput", GoGetter: "RunToolsScriptsAfterPowerOnInput"},
			_jsii_.MemberProperty{JsiiProperty: "runToolsScriptsAfterResume", GoGetter: "RunToolsScriptsAfterResume"},
			_jsii_.MemberProperty{JsiiProperty: "runToolsScriptsAfterResumeInput", GoGetter: "RunToolsScriptsAfterResumeInput"},
			_jsii_.MemberProperty{JsiiProperty: "runToolsScriptsBeforeGuestReboot", GoGetter: "RunToolsScriptsBeforeGuestReboot"},
			_jsii_.MemberProperty{JsiiProperty: "runToolsScriptsBeforeGuestRebootInput", GoGetter: "RunToolsScriptsBeforeGuestRebootInput"},
			_jsii_.MemberProperty{JsiiProperty: "runToolsScriptsBeforeGuestShutdown", GoGetter: "RunToolsScriptsBeforeGuestShutdown"},
			_jsii_.MemberProperty{JsiiProperty: "runToolsScriptsBeforeGuestShutdownInput", GoGetter: "RunToolsScriptsBeforeGuestShutdownInput"},
			_jsii_.MemberProperty{JsiiProperty: "runToolsScriptsBeforeGuestStandby", GoGetter: "RunToolsScriptsBeforeGuestStandby"},
			_jsii_.MemberProperty{JsiiProperty: "runToolsScriptsBeforeGuestStandbyInput", GoGetter: "RunToolsScriptsBeforeGuestStandbyInput"},
			_jsii_.MemberProperty{JsiiProperty: "sataControllerScanCount", GoGetter: "SataControllerScanCount"},
			_jsii_.MemberProperty{JsiiProperty: "sataControllerScanCountInput", GoGetter: "SataControllerScanCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "scsiBusSharing", GoGetter: "ScsiBusSharing"},
			_jsii_.MemberProperty{JsiiProperty: "scsiControllerScanCount", GoGetter: "ScsiControllerScanCount"},
			_jsii_.MemberProperty{JsiiProperty: "scsiControllerScanCountInput", GoGetter: "ScsiControllerScanCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "scsiType", GoGetter: "ScsiType"},
			_jsii_.MemberProperty{JsiiProperty: "storagePolicyId", GoGetter: "StoragePolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "storagePolicyIdInput", GoGetter: "StoragePolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "swapPlacementPolicy", GoGetter: "SwapPlacementPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "swapPlacementPolicyInput", GoGetter: "SwapPlacementPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "syncTimeWithHost", GoGetter: "SyncTimeWithHost"},
			_jsii_.MemberProperty{JsiiProperty: "syncTimeWithHostInput", GoGetter: "SyncTimeWithHostInput"},
			_jsii_.MemberProperty{JsiiProperty: "syncTimeWithHostPeriodically", GoGetter: "SyncTimeWithHostPeriodically"},
			_jsii_.MemberProperty{JsiiProperty: "syncTimeWithHostPeriodicallyInput", GoGetter: "SyncTimeWithHostPeriodicallyInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "toolsUpgradePolicy", GoGetter: "ToolsUpgradePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "toolsUpgradePolicyInput", GoGetter: "ToolsUpgradePolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "uuid", GoGetter: "Uuid"},
			_jsii_.MemberProperty{JsiiProperty: "uuidInput", GoGetter: "UuidInput"},
			_jsii_.MemberProperty{JsiiProperty: "vapp", GoGetter: "Vapp"},
			_jsii_.MemberProperty{JsiiProperty: "vappInput", GoGetter: "VappInput"},
			_jsii_.MemberProperty{JsiiProperty: "vappTransport", GoGetter: "VappTransport"},
			_jsii_.MemberProperty{JsiiProperty: "vbsEnabled", GoGetter: "VbsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "vbsEnabledInput", GoGetter: "VbsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "vvtdEnabled", GoGetter: "VvtdEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "vvtdEnabledInput", GoGetter: "VvtdEnabledInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataVsphereVirtualMachine{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vsphere.dataVsphereVirtualMachine.DataVsphereVirtualMachineConfig",
		reflect.TypeOf((*DataVsphereVirtualMachineConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vsphere.dataVsphereVirtualMachine.DataVsphereVirtualMachineDisks",
		reflect.TypeOf((*DataVsphereVirtualMachineDisks)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.dataVsphereVirtualMachine.DataVsphereVirtualMachineDisksList",
		reflect.TypeOf((*DataVsphereVirtualMachineDisksList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_DataVsphereVirtualMachineDisksList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.dataVsphereVirtualMachine.DataVsphereVirtualMachineDisksOutputReference",
		reflect.TypeOf((*DataVsphereVirtualMachineDisksOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "eagerlyScrub", GoGetter: "EagerlyScrub"},
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
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "size", GoGetter: "Size"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "thinProvisioned", GoGetter: "ThinProvisioned"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "unitNumber", GoGetter: "UnitNumber"},
		},
		func() interface{} {
			j := jsiiProxy_DataVsphereVirtualMachineDisksOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vsphere.dataVsphereVirtualMachine.DataVsphereVirtualMachineNetworkInterfaces",
		reflect.TypeOf((*DataVsphereVirtualMachineNetworkInterfaces)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.dataVsphereVirtualMachine.DataVsphereVirtualMachineNetworkInterfacesList",
		reflect.TypeOf((*DataVsphereVirtualMachineNetworkInterfacesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_DataVsphereVirtualMachineNetworkInterfacesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.dataVsphereVirtualMachine.DataVsphereVirtualMachineNetworkInterfacesOutputReference",
		reflect.TypeOf((*DataVsphereVirtualMachineNetworkInterfacesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "adapterType", GoGetter: "AdapterType"},
			_jsii_.MemberProperty{JsiiProperty: "bandwidthLimit", GoGetter: "BandwidthLimit"},
			_jsii_.MemberProperty{JsiiProperty: "bandwidthReservation", GoGetter: "BandwidthReservation"},
			_jsii_.MemberProperty{JsiiProperty: "bandwidthShareCount", GoGetter: "BandwidthShareCount"},
			_jsii_.MemberProperty{JsiiProperty: "bandwidthShareLevel", GoGetter: "BandwidthShareLevel"},
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
			_jsii_.MemberProperty{JsiiProperty: "macAddress", GoGetter: "MacAddress"},
			_jsii_.MemberProperty{JsiiProperty: "networkId", GoGetter: "NetworkId"},
			_jsii_.MemberProperty{JsiiProperty: "physicalFunction", GoGetter: "PhysicalFunction"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataVsphereVirtualMachineNetworkInterfacesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vsphere.dataVsphereVirtualMachine.DataVsphereVirtualMachineVapp",
		reflect.TypeOf((*DataVsphereVirtualMachineVapp)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.dataVsphereVirtualMachine.DataVsphereVirtualMachineVappOutputReference",
		reflect.TypeOf((*DataVsphereVirtualMachineVappOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "properties", GoGetter: "Properties"},
			_jsii_.MemberProperty{JsiiProperty: "propertiesInput", GoGetter: "PropertiesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetProperties", GoMethod: "ResetProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataVsphereVirtualMachineVappOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
