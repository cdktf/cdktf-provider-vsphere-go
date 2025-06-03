// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datavspherevirtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v10/datavspherevirtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/data-sources/virtual_machine vsphere_virtual_machine}.
type DataVsphereVirtualMachine interface {
	cdktf.TerraformDataSource
	AlternateGuestName() *string
	SetAlternateGuestName(val *string)
	AlternateGuestNameInput() *string
	Annotation() *string
	SetAnnotation(val *string)
	AnnotationInput() *string
	BootDelay() *float64
	SetBootDelay(val *float64)
	BootDelayInput() *float64
	BootRetryDelay() *float64
	SetBootRetryDelay(val *float64)
	BootRetryDelayInput() *float64
	BootRetryEnabled() interface{}
	SetBootRetryEnabled(val interface{})
	BootRetryEnabledInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChangeVersion() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CpuHotAddEnabled() interface{}
	SetCpuHotAddEnabled(val interface{})
	CpuHotAddEnabledInput() interface{}
	CpuHotRemoveEnabled() interface{}
	SetCpuHotRemoveEnabled(val interface{})
	CpuHotRemoveEnabledInput() interface{}
	CpuLimit() *float64
	SetCpuLimit(val *float64)
	CpuLimitInput() *float64
	CpuPerformanceCountersEnabled() interface{}
	SetCpuPerformanceCountersEnabled(val interface{})
	CpuPerformanceCountersEnabledInput() interface{}
	CpuReservation() *float64
	SetCpuReservation(val *float64)
	CpuReservationInput() *float64
	CpuShareCount() *float64
	SetCpuShareCount(val *float64)
	CpuShareCountInput() *float64
	CpuShareLevel() *string
	SetCpuShareLevel(val *string)
	CpuShareLevelInput() *string
	DatacenterId() *string
	SetDatacenterId(val *string)
	DatacenterIdInput() *string
	DefaultIpAddress() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Disks() DataVsphereVirtualMachineDisksList
	EfiSecureBootEnabled() interface{}
	SetEfiSecureBootEnabled(val interface{})
	EfiSecureBootEnabledInput() interface{}
	EnableDiskUuid() interface{}
	SetEnableDiskUuid(val interface{})
	EnableDiskUuidInput() interface{}
	EnableLogging() interface{}
	SetEnableLogging(val interface{})
	EnableLoggingInput() interface{}
	EptRviMode() *string
	SetEptRviMode(val *string)
	EptRviModeInput() *string
	ExtraConfig() *map[string]*string
	SetExtraConfig(val *map[string]*string)
	ExtraConfigInput() *map[string]*string
	ExtraConfigRebootRequired() interface{}
	SetExtraConfigRebootRequired(val interface{})
	ExtraConfigRebootRequiredInput() interface{}
	Firmware() *string
	SetFirmware(val *string)
	FirmwareInput() *string
	Folder() *string
	SetFolder(val *string)
	FolderInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GuestId() *string
	SetGuestId(val *string)
	GuestIdInput() *string
	GuestIpAddresses() *[]*string
	HardwareVersion() *float64
	SetHardwareVersion(val *float64)
	HardwareVersionInput() *float64
	HvMode() *string
	SetHvMode(val *string)
	HvModeInput() *string
	Id() *string
	SetId(val *string)
	IdeControllerScanCount() *float64
	SetIdeControllerScanCount(val *float64)
	IdeControllerScanCountInput() *float64
	IdInput() *string
	InstanceUuid() *string
	LatencySensitivity() *string
	SetLatencySensitivity(val *string)
	LatencySensitivityInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Memory() *float64
	SetMemory(val *float64)
	MemoryHotAddEnabled() interface{}
	SetMemoryHotAddEnabled(val interface{})
	MemoryHotAddEnabledInput() interface{}
	MemoryInput() *float64
	MemoryLimit() *float64
	SetMemoryLimit(val *float64)
	MemoryLimitInput() *float64
	MemoryReservation() *float64
	SetMemoryReservation(val *float64)
	MemoryReservationInput() *float64
	MemoryReservationLockedToMax() interface{}
	SetMemoryReservationLockedToMax(val interface{})
	MemoryReservationLockedToMaxInput() interface{}
	MemoryShareCount() *float64
	SetMemoryShareCount(val *float64)
	MemoryShareCountInput() *float64
	MemoryShareLevel() *string
	SetMemoryShareLevel(val *string)
	MemoryShareLevelInput() *string
	Moid() *string
	SetMoid(val *string)
	MoidInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NestedHvEnabled() interface{}
	SetNestedHvEnabled(val interface{})
	NestedHvEnabledInput() interface{}
	NetworkInterfaces() DataVsphereVirtualMachineNetworkInterfacesList
	NetworkInterfaceTypes() *[]*string
	// The tree node.
	Node() constructs.Node
	NumCoresPerSocket() *float64
	SetNumCoresPerSocket(val *float64)
	NumCoresPerSocketInput() *float64
	NumCpus() *float64
	SetNumCpus(val *float64)
	NumCpusInput() *float64
	NvmeControllerScanCount() *float64
	SetNvmeControllerScanCount(val *float64)
	NvmeControllerScanCountInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ReplaceTrigger() *string
	SetReplaceTrigger(val *string)
	ReplaceTriggerInput() *string
	RunToolsScriptsAfterPowerOn() interface{}
	SetRunToolsScriptsAfterPowerOn(val interface{})
	RunToolsScriptsAfterPowerOnInput() interface{}
	RunToolsScriptsAfterResume() interface{}
	SetRunToolsScriptsAfterResume(val interface{})
	RunToolsScriptsAfterResumeInput() interface{}
	RunToolsScriptsBeforeGuestReboot() interface{}
	SetRunToolsScriptsBeforeGuestReboot(val interface{})
	RunToolsScriptsBeforeGuestRebootInput() interface{}
	RunToolsScriptsBeforeGuestShutdown() interface{}
	SetRunToolsScriptsBeforeGuestShutdown(val interface{})
	RunToolsScriptsBeforeGuestShutdownInput() interface{}
	RunToolsScriptsBeforeGuestStandby() interface{}
	SetRunToolsScriptsBeforeGuestStandby(val interface{})
	RunToolsScriptsBeforeGuestStandbyInput() interface{}
	SataControllerScanCount() *float64
	SetSataControllerScanCount(val *float64)
	SataControllerScanCountInput() *float64
	ScsiBusSharing() *string
	ScsiControllerScanCount() *float64
	SetScsiControllerScanCount(val *float64)
	ScsiControllerScanCountInput() *float64
	ScsiType() *string
	StoragePolicyId() *string
	SetStoragePolicyId(val *string)
	StoragePolicyIdInput() *string
	SwapPlacementPolicy() *string
	SetSwapPlacementPolicy(val *string)
	SwapPlacementPolicyInput() *string
	SyncTimeWithHost() interface{}
	SetSyncTimeWithHost(val interface{})
	SyncTimeWithHostInput() interface{}
	SyncTimeWithHostPeriodically() interface{}
	SetSyncTimeWithHostPeriodically(val interface{})
	SyncTimeWithHostPeriodicallyInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ToolsUpgradePolicy() *string
	SetToolsUpgradePolicy(val *string)
	ToolsUpgradePolicyInput() *string
	Uuid() *string
	SetUuid(val *string)
	UuidInput() *string
	Vapp() DataVsphereVirtualMachineVappOutputReference
	VappInput() *DataVsphereVirtualMachineVapp
	VappTransport() *[]*string
	VbsEnabled() interface{}
	SetVbsEnabled(val interface{})
	VbsEnabledInput() interface{}
	Vtpm() cdktf.IResolvable
	VvtdEnabled() interface{}
	SetVvtdEnabled(val interface{})
	VvtdEnabledInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutVapp(value *DataVsphereVirtualMachineVapp)
	ResetAlternateGuestName()
	ResetAnnotation()
	ResetBootDelay()
	ResetBootRetryDelay()
	ResetBootRetryEnabled()
	ResetCpuHotAddEnabled()
	ResetCpuHotRemoveEnabled()
	ResetCpuLimit()
	ResetCpuPerformanceCountersEnabled()
	ResetCpuReservation()
	ResetCpuShareCount()
	ResetCpuShareLevel()
	ResetDatacenterId()
	ResetEfiSecureBootEnabled()
	ResetEnableDiskUuid()
	ResetEnableLogging()
	ResetEptRviMode()
	ResetExtraConfig()
	ResetExtraConfigRebootRequired()
	ResetFirmware()
	ResetFolder()
	ResetGuestId()
	ResetHardwareVersion()
	ResetHvMode()
	ResetId()
	ResetIdeControllerScanCount()
	ResetLatencySensitivity()
	ResetMemory()
	ResetMemoryHotAddEnabled()
	ResetMemoryLimit()
	ResetMemoryReservation()
	ResetMemoryReservationLockedToMax()
	ResetMemoryShareCount()
	ResetMemoryShareLevel()
	ResetMoid()
	ResetName()
	ResetNestedHvEnabled()
	ResetNumCoresPerSocket()
	ResetNumCpus()
	ResetNvmeControllerScanCount()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReplaceTrigger()
	ResetRunToolsScriptsAfterPowerOn()
	ResetRunToolsScriptsAfterResume()
	ResetRunToolsScriptsBeforeGuestReboot()
	ResetRunToolsScriptsBeforeGuestShutdown()
	ResetRunToolsScriptsBeforeGuestStandby()
	ResetSataControllerScanCount()
	ResetScsiControllerScanCount()
	ResetStoragePolicyId()
	ResetSwapPlacementPolicy()
	ResetSyncTimeWithHost()
	ResetSyncTimeWithHostPeriodically()
	ResetToolsUpgradePolicy()
	ResetUuid()
	ResetVapp()
	ResetVbsEnabled()
	ResetVvtdEnabled()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataVsphereVirtualMachine
type jsiiProxy_DataVsphereVirtualMachine struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataVsphereVirtualMachine) AlternateGuestName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alternateGuestName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) AlternateGuestNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alternateGuestNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) Annotation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) AnnotationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) BootDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) BootDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) BootRetryDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootRetryDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) BootRetryDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootRetryDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) BootRetryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootRetryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) BootRetryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootRetryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) ChangeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) CpuHotAddEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuHotAddEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) CpuHotAddEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuHotAddEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) CpuHotRemoveEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuHotRemoveEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) CpuHotRemoveEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuHotRemoveEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) CpuLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) CpuLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) CpuPerformanceCountersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuPerformanceCountersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) CpuPerformanceCountersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuPerformanceCountersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) CpuReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) CpuReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) CpuShareCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuShareCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) CpuShareCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuShareCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) CpuShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) CpuShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) DatacenterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) DatacenterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) DefaultIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) Disks() DataVsphereVirtualMachineDisksList {
	var returns DataVsphereVirtualMachineDisksList
	_jsii_.Get(
		j,
		"disks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) EfiSecureBootEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efiSecureBootEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) EfiSecureBootEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efiSecureBootEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) EnableDiskUuid() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDiskUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) EnableDiskUuidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDiskUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) EnableLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) EnableLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) EptRviMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eptRviMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) EptRviModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eptRviModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) ExtraConfig() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) ExtraConfigInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) ExtraConfigRebootRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extraConfigRebootRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) ExtraConfigRebootRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extraConfigRebootRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) Firmware() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firmware",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) FirmwareInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firmwareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) Folder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) FolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) GuestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) GuestIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guestIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) GuestIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guestIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) HardwareVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hardwareVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) HardwareVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hardwareVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) HvMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hvMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) HvModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hvModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) IdeControllerScanCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ideControllerScanCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) IdeControllerScanCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ideControllerScanCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) InstanceUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) LatencySensitivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latencySensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) LatencySensitivityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latencySensitivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) MemoryHotAddEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryHotAddEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) MemoryHotAddEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryHotAddEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) MemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) MemoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) MemoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) MemoryReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) MemoryReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) MemoryReservationLockedToMax() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryReservationLockedToMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) MemoryReservationLockedToMaxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryReservationLockedToMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) MemoryShareCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryShareCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) MemoryShareCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryShareCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) MemoryShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) MemoryShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) Moid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"moid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) MoidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"moidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) NestedHvEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nestedHvEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) NestedHvEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nestedHvEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) NetworkInterfaces() DataVsphereVirtualMachineNetworkInterfacesList {
	var returns DataVsphereVirtualMachineNetworkInterfacesList
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) NetworkInterfaceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) NumCoresPerSocket() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCoresPerSocket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) NumCoresPerSocketInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCoresPerSocketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) NumCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) NumCpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) NvmeControllerScanCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nvmeControllerScanCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) NvmeControllerScanCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nvmeControllerScanCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) ReplaceTrigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) ReplaceTriggerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) RunToolsScriptsAfterPowerOn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsAfterPowerOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) RunToolsScriptsAfterPowerOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsAfterPowerOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) RunToolsScriptsAfterResume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsAfterResume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) RunToolsScriptsAfterResumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsAfterResumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) RunToolsScriptsBeforeGuestReboot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsBeforeGuestReboot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) RunToolsScriptsBeforeGuestRebootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsBeforeGuestRebootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) RunToolsScriptsBeforeGuestShutdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsBeforeGuestShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) RunToolsScriptsBeforeGuestShutdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsBeforeGuestShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) RunToolsScriptsBeforeGuestStandby() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsBeforeGuestStandby",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) RunToolsScriptsBeforeGuestStandbyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsBeforeGuestStandbyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) SataControllerScanCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sataControllerScanCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) SataControllerScanCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sataControllerScanCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) ScsiBusSharing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scsiBusSharing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) ScsiControllerScanCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scsiControllerScanCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) ScsiControllerScanCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scsiControllerScanCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) ScsiType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scsiType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) StoragePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) StoragePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) SwapPlacementPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"swapPlacementPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) SwapPlacementPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"swapPlacementPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) SyncTimeWithHost() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncTimeWithHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) SyncTimeWithHostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncTimeWithHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) SyncTimeWithHostPeriodically() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncTimeWithHostPeriodically",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) SyncTimeWithHostPeriodicallyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncTimeWithHostPeriodicallyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) ToolsUpgradePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolsUpgradePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) ToolsUpgradePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolsUpgradePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) UuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) Vapp() DataVsphereVirtualMachineVappOutputReference {
	var returns DataVsphereVirtualMachineVappOutputReference
	_jsii_.Get(
		j,
		"vapp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) VappInput() *DataVsphereVirtualMachineVapp {
	var returns *DataVsphereVirtualMachineVapp
	_jsii_.Get(
		j,
		"vappInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) VappTransport() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vappTransport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) VbsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vbsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) VbsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vbsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) Vtpm() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"vtpm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) VvtdEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vvtdEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereVirtualMachine) VvtdEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vvtdEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/data-sources/virtual_machine vsphere_virtual_machine} Data Source.
func NewDataVsphereVirtualMachine(scope constructs.Construct, id *string, config *DataVsphereVirtualMachineConfig) DataVsphereVirtualMachine {
	_init_.Initialize()

	if err := validateNewDataVsphereVirtualMachineParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVsphereVirtualMachine{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.dataVsphereVirtualMachine.DataVsphereVirtualMachine",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/data-sources/virtual_machine vsphere_virtual_machine} Data Source.
func NewDataVsphereVirtualMachine_Override(d DataVsphereVirtualMachine, scope constructs.Construct, id *string, config *DataVsphereVirtualMachineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.dataVsphereVirtualMachine.DataVsphereVirtualMachine",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetAlternateGuestName(val *string) {
	if err := j.validateSetAlternateGuestNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alternateGuestName",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetAnnotation(val *string) {
	if err := j.validateSetAnnotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotation",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetBootDelay(val *float64) {
	if err := j.validateSetBootDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDelay",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetBootRetryDelay(val *float64) {
	if err := j.validateSetBootRetryDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootRetryDelay",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetBootRetryEnabled(val interface{}) {
	if err := j.validateSetBootRetryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootRetryEnabled",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetCpuHotAddEnabled(val interface{}) {
	if err := j.validateSetCpuHotAddEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuHotAddEnabled",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetCpuHotRemoveEnabled(val interface{}) {
	if err := j.validateSetCpuHotRemoveEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuHotRemoveEnabled",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetCpuLimit(val *float64) {
	if err := j.validateSetCpuLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuLimit",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetCpuPerformanceCountersEnabled(val interface{}) {
	if err := j.validateSetCpuPerformanceCountersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuPerformanceCountersEnabled",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetCpuReservation(val *float64) {
	if err := j.validateSetCpuReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuReservation",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetCpuShareCount(val *float64) {
	if err := j.validateSetCpuShareCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuShareCount",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetCpuShareLevel(val *string) {
	if err := j.validateSetCpuShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuShareLevel",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetDatacenterId(val *string) {
	if err := j.validateSetDatacenterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datacenterId",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetEfiSecureBootEnabled(val interface{}) {
	if err := j.validateSetEfiSecureBootEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"efiSecureBootEnabled",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetEnableDiskUuid(val interface{}) {
	if err := j.validateSetEnableDiskUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDiskUuid",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetEnableLogging(val interface{}) {
	if err := j.validateSetEnableLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLogging",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetEptRviMode(val *string) {
	if err := j.validateSetEptRviModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eptRviMode",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetExtraConfig(val *map[string]*string) {
	if err := j.validateSetExtraConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraConfig",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetExtraConfigRebootRequired(val interface{}) {
	if err := j.validateSetExtraConfigRebootRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraConfigRebootRequired",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetFirmware(val *string) {
	if err := j.validateSetFirmwareParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firmware",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetFolder(val *string) {
	if err := j.validateSetFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"folder",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetGuestId(val *string) {
	if err := j.validateSetGuestIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guestId",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetHardwareVersion(val *float64) {
	if err := j.validateSetHardwareVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hardwareVersion",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetHvMode(val *string) {
	if err := j.validateSetHvModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hvMode",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetIdeControllerScanCount(val *float64) {
	if err := j.validateSetIdeControllerScanCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ideControllerScanCount",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetLatencySensitivity(val *string) {
	if err := j.validateSetLatencySensitivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latencySensitivity",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetMemory(val *float64) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetMemoryHotAddEnabled(val interface{}) {
	if err := j.validateSetMemoryHotAddEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryHotAddEnabled",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetMemoryLimit(val *float64) {
	if err := j.validateSetMemoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryLimit",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetMemoryReservation(val *float64) {
	if err := j.validateSetMemoryReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryReservation",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetMemoryReservationLockedToMax(val interface{}) {
	if err := j.validateSetMemoryReservationLockedToMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryReservationLockedToMax",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetMemoryShareCount(val *float64) {
	if err := j.validateSetMemoryShareCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryShareCount",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetMemoryShareLevel(val *string) {
	if err := j.validateSetMemoryShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryShareLevel",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetMoid(val *string) {
	if err := j.validateSetMoidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"moid",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetNestedHvEnabled(val interface{}) {
	if err := j.validateSetNestedHvEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nestedHvEnabled",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetNumCoresPerSocket(val *float64) {
	if err := j.validateSetNumCoresPerSocketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numCoresPerSocket",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetNumCpus(val *float64) {
	if err := j.validateSetNumCpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numCpus",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetNvmeControllerScanCount(val *float64) {
	if err := j.validateSetNvmeControllerScanCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nvmeControllerScanCount",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetReplaceTrigger(val *string) {
	if err := j.validateSetReplaceTriggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceTrigger",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetRunToolsScriptsAfterPowerOn(val interface{}) {
	if err := j.validateSetRunToolsScriptsAfterPowerOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runToolsScriptsAfterPowerOn",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetRunToolsScriptsAfterResume(val interface{}) {
	if err := j.validateSetRunToolsScriptsAfterResumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runToolsScriptsAfterResume",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetRunToolsScriptsBeforeGuestReboot(val interface{}) {
	if err := j.validateSetRunToolsScriptsBeforeGuestRebootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runToolsScriptsBeforeGuestReboot",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetRunToolsScriptsBeforeGuestShutdown(val interface{}) {
	if err := j.validateSetRunToolsScriptsBeforeGuestShutdownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runToolsScriptsBeforeGuestShutdown",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetRunToolsScriptsBeforeGuestStandby(val interface{}) {
	if err := j.validateSetRunToolsScriptsBeforeGuestStandbyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runToolsScriptsBeforeGuestStandby",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetSataControllerScanCount(val *float64) {
	if err := j.validateSetSataControllerScanCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sataControllerScanCount",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetScsiControllerScanCount(val *float64) {
	if err := j.validateSetScsiControllerScanCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scsiControllerScanCount",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetStoragePolicyId(val *string) {
	if err := j.validateSetStoragePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storagePolicyId",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetSwapPlacementPolicy(val *string) {
	if err := j.validateSetSwapPlacementPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"swapPlacementPolicy",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetSyncTimeWithHost(val interface{}) {
	if err := j.validateSetSyncTimeWithHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncTimeWithHost",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetSyncTimeWithHostPeriodically(val interface{}) {
	if err := j.validateSetSyncTimeWithHostPeriodicallyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncTimeWithHostPeriodically",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetToolsUpgradePolicy(val *string) {
	if err := j.validateSetToolsUpgradePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toolsUpgradePolicy",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetUuid(val *string) {
	if err := j.validateSetUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uuid",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetVbsEnabled(val interface{}) {
	if err := j.validateSetVbsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vbsEnabled",
		val,
	)
}

func (j *jsiiProxy_DataVsphereVirtualMachine)SetVvtdEnabled(val interface{}) {
	if err := j.validateSetVvtdEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vvtdEnabled",
		val,
	)
}

// Generates CDKTF code for importing a DataVsphereVirtualMachine resource upon running "cdktf plan <stack-name>".
func DataVsphereVirtualMachine_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataVsphereVirtualMachine_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.dataVsphereVirtualMachine.DataVsphereVirtualMachine",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DataVsphereVirtualMachine_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVsphereVirtualMachine_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.dataVsphereVirtualMachine.DataVsphereVirtualMachine",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVsphereVirtualMachine_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVsphereVirtualMachine_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.dataVsphereVirtualMachine.DataVsphereVirtualMachine",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVsphereVirtualMachine_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVsphereVirtualMachine_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.dataVsphereVirtualMachine.DataVsphereVirtualMachine",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVsphereVirtualMachine_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vsphere.dataVsphereVirtualMachine.DataVsphereVirtualMachine",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVsphereVirtualMachine) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereVirtualMachine) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereVirtualMachine) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereVirtualMachine) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereVirtualMachine) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereVirtualMachine) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereVirtualMachine) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereVirtualMachine) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereVirtualMachine) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereVirtualMachine) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereVirtualMachine) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) PutVapp(value *DataVsphereVirtualMachineVapp) {
	if err := d.validatePutVappParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVapp",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetAlternateGuestName() {
	_jsii_.InvokeVoid(
		d,
		"resetAlternateGuestName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetAnnotation() {
	_jsii_.InvokeVoid(
		d,
		"resetAnnotation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetBootDelay() {
	_jsii_.InvokeVoid(
		d,
		"resetBootDelay",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetBootRetryDelay() {
	_jsii_.InvokeVoid(
		d,
		"resetBootRetryDelay",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetBootRetryEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetBootRetryEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetCpuHotAddEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetCpuHotAddEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetCpuHotRemoveEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetCpuHotRemoveEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetCpuLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetCpuLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetCpuPerformanceCountersEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetCpuPerformanceCountersEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetCpuReservation() {
	_jsii_.InvokeVoid(
		d,
		"resetCpuReservation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetCpuShareCount() {
	_jsii_.InvokeVoid(
		d,
		"resetCpuShareCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetCpuShareLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetCpuShareLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetDatacenterId() {
	_jsii_.InvokeVoid(
		d,
		"resetDatacenterId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetEfiSecureBootEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetEfiSecureBootEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetEnableDiskUuid() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableDiskUuid",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetEnableLogging() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableLogging",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetEptRviMode() {
	_jsii_.InvokeVoid(
		d,
		"resetEptRviMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetExtraConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetExtraConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetExtraConfigRebootRequired() {
	_jsii_.InvokeVoid(
		d,
		"resetExtraConfigRebootRequired",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetFirmware() {
	_jsii_.InvokeVoid(
		d,
		"resetFirmware",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetFolder() {
	_jsii_.InvokeVoid(
		d,
		"resetFolder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetGuestId() {
	_jsii_.InvokeVoid(
		d,
		"resetGuestId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetHardwareVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetHardwareVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetHvMode() {
	_jsii_.InvokeVoid(
		d,
		"resetHvMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetIdeControllerScanCount() {
	_jsii_.InvokeVoid(
		d,
		"resetIdeControllerScanCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetLatencySensitivity() {
	_jsii_.InvokeVoid(
		d,
		"resetLatencySensitivity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetMemory() {
	_jsii_.InvokeVoid(
		d,
		"resetMemory",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetMemoryHotAddEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetMemoryHotAddEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetMemoryLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetMemoryLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetMemoryReservation() {
	_jsii_.InvokeVoid(
		d,
		"resetMemoryReservation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetMemoryReservationLockedToMax() {
	_jsii_.InvokeVoid(
		d,
		"resetMemoryReservationLockedToMax",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetMemoryShareCount() {
	_jsii_.InvokeVoid(
		d,
		"resetMemoryShareCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetMemoryShareLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetMemoryShareLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetMoid() {
	_jsii_.InvokeVoid(
		d,
		"resetMoid",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetNestedHvEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetNestedHvEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetNumCoresPerSocket() {
	_jsii_.InvokeVoid(
		d,
		"resetNumCoresPerSocket",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetNumCpus() {
	_jsii_.InvokeVoid(
		d,
		"resetNumCpus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetNvmeControllerScanCount() {
	_jsii_.InvokeVoid(
		d,
		"resetNvmeControllerScanCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetReplaceTrigger() {
	_jsii_.InvokeVoid(
		d,
		"resetReplaceTrigger",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetRunToolsScriptsAfterPowerOn() {
	_jsii_.InvokeVoid(
		d,
		"resetRunToolsScriptsAfterPowerOn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetRunToolsScriptsAfterResume() {
	_jsii_.InvokeVoid(
		d,
		"resetRunToolsScriptsAfterResume",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetRunToolsScriptsBeforeGuestReboot() {
	_jsii_.InvokeVoid(
		d,
		"resetRunToolsScriptsBeforeGuestReboot",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetRunToolsScriptsBeforeGuestShutdown() {
	_jsii_.InvokeVoid(
		d,
		"resetRunToolsScriptsBeforeGuestShutdown",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetRunToolsScriptsBeforeGuestStandby() {
	_jsii_.InvokeVoid(
		d,
		"resetRunToolsScriptsBeforeGuestStandby",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetSataControllerScanCount() {
	_jsii_.InvokeVoid(
		d,
		"resetSataControllerScanCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetScsiControllerScanCount() {
	_jsii_.InvokeVoid(
		d,
		"resetScsiControllerScanCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetStoragePolicyId() {
	_jsii_.InvokeVoid(
		d,
		"resetStoragePolicyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetSwapPlacementPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetSwapPlacementPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetSyncTimeWithHost() {
	_jsii_.InvokeVoid(
		d,
		"resetSyncTimeWithHost",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetSyncTimeWithHostPeriodically() {
	_jsii_.InvokeVoid(
		d,
		"resetSyncTimeWithHostPeriodically",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetToolsUpgradePolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetToolsUpgradePolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetUuid() {
	_jsii_.InvokeVoid(
		d,
		"resetUuid",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetVapp() {
	_jsii_.InvokeVoid(
		d,
		"resetVapp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetVbsEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetVbsEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ResetVvtdEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetVvtdEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereVirtualMachine) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereVirtualMachine) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereVirtualMachine) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

