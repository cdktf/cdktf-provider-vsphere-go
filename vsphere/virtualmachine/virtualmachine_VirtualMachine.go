package virtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v3/virtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/vsphere/r/virtual_machine vsphere_virtual_machine}.
type VirtualMachine interface {
	cdktf.TerraformResource
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
	Cdrom() VirtualMachineCdromOutputReference
	CdromInput() *VirtualMachineCdrom
	ChangeVersion() *string
	Clone() VirtualMachineCloneOutputReference
	CloneInput() *VirtualMachineClone
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	CustomAttributes() *map[string]*string
	SetCustomAttributes(val *map[string]*string)
	CustomAttributesInput() *map[string]*string
	DatacenterId() *string
	SetDatacenterId(val *string)
	DatacenterIdInput() *string
	DatastoreClusterId() *string
	SetDatastoreClusterId(val *string)
	DatastoreClusterIdInput() *string
	DatastoreId() *string
	SetDatastoreId(val *string)
	DatastoreIdInput() *string
	DefaultIpAddress() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Disk() VirtualMachineDiskList
	DiskInput() interface{}
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
	Firmware() *string
	SetFirmware(val *string)
	FirmwareInput() *string
	Folder() *string
	SetFolder(val *string)
	FolderInput() *string
	ForcePowerOff() interface{}
	SetForcePowerOff(val interface{})
	ForcePowerOffInput() interface{}
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
	HostSystemId() *string
	SetHostSystemId(val *string)
	HostSystemIdInput() *string
	HvMode() *string
	SetHvMode(val *string)
	HvModeInput() *string
	Id() *string
	SetId(val *string)
	IdeControllerCount() *float64
	SetIdeControllerCount(val *float64)
	IdeControllerCountInput() *float64
	IdInput() *string
	IgnoredGuestIps() *[]*string
	SetIgnoredGuestIps(val *[]*string)
	IgnoredGuestIpsInput() *[]*string
	Imported() cdktf.IResolvable
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
	MemoryShareCount() *float64
	SetMemoryShareCount(val *float64)
	MemoryShareCountInput() *float64
	MemoryShareLevel() *string
	SetMemoryShareLevel(val *string)
	MemoryShareLevelInput() *string
	MigrateWaitTimeout() *float64
	SetMigrateWaitTimeout(val *float64)
	MigrateWaitTimeoutInput() *float64
	Moid() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NestedHvEnabled() interface{}
	SetNestedHvEnabled(val interface{})
	NestedHvEnabledInput() interface{}
	NetworkInterface() VirtualMachineNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	// The tree node.
	Node() constructs.Node
	NumCoresPerSocket() *float64
	SetNumCoresPerSocket(val *float64)
	NumCoresPerSocketInput() *float64
	NumCpus() *float64
	SetNumCpus(val *float64)
	NumCpusInput() *float64
	OvfDeploy() VirtualMachineOvfDeployOutputReference
	OvfDeployInput() *VirtualMachineOvfDeploy
	PciDeviceId() *[]*string
	SetPciDeviceId(val *[]*string)
	PciDeviceIdInput() *[]*string
	PoweronTimeout() *float64
	SetPoweronTimeout(val *float64)
	PoweronTimeoutInput() *float64
	PowerState() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RebootRequired() cdktf.IResolvable
	ReplaceTrigger() *string
	SetReplaceTrigger(val *string)
	ReplaceTriggerInput() *string
	ResourcePoolId() *string
	SetResourcePoolId(val *string)
	ResourcePoolIdInput() *string
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
	SataControllerCount() *float64
	SetSataControllerCount(val *float64)
	SataControllerCountInput() *float64
	ScsiBusSharing() *string
	SetScsiBusSharing(val *string)
	ScsiBusSharingInput() *string
	ScsiControllerCount() *float64
	SetScsiControllerCount(val *float64)
	ScsiControllerCountInput() *float64
	ScsiType() *string
	SetScsiType(val *string)
	ScsiTypeInput() *string
	ShutdownWaitTimeout() *float64
	SetShutdownWaitTimeout(val *float64)
	ShutdownWaitTimeoutInput() *float64
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
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
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
	Vapp() VirtualMachineVappOutputReference
	VappInput() *VirtualMachineVapp
	VappTransport() *[]*string
	VbsEnabled() interface{}
	SetVbsEnabled(val interface{})
	VbsEnabledInput() interface{}
	VmwareToolsStatus() *string
	VmxPath() *string
	VvtdEnabled() interface{}
	SetVvtdEnabled(val interface{})
	VvtdEnabledInput() interface{}
	WaitForGuestIpTimeout() *float64
	SetWaitForGuestIpTimeout(val *float64)
	WaitForGuestIpTimeoutInput() *float64
	WaitForGuestNetRoutable() interface{}
	SetWaitForGuestNetRoutable(val interface{})
	WaitForGuestNetRoutableInput() interface{}
	WaitForGuestNetTimeout() *float64
	SetWaitForGuestNetTimeout(val *float64)
	WaitForGuestNetTimeoutInput() *float64
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
	PutCdrom(value *VirtualMachineCdrom)
	PutClone(value *VirtualMachineClone)
	PutDisk(value interface{})
	PutNetworkInterface(value interface{})
	PutOvfDeploy(value *VirtualMachineOvfDeploy)
	PutVapp(value *VirtualMachineVapp)
	ResetAlternateGuestName()
	ResetAnnotation()
	ResetBootDelay()
	ResetBootRetryDelay()
	ResetBootRetryEnabled()
	ResetCdrom()
	ResetClone()
	ResetCpuHotAddEnabled()
	ResetCpuHotRemoveEnabled()
	ResetCpuLimit()
	ResetCpuPerformanceCountersEnabled()
	ResetCpuReservation()
	ResetCpuShareCount()
	ResetCpuShareLevel()
	ResetCustomAttributes()
	ResetDatacenterId()
	ResetDatastoreClusterId()
	ResetDatastoreId()
	ResetDisk()
	ResetEfiSecureBootEnabled()
	ResetEnableDiskUuid()
	ResetEnableLogging()
	ResetEptRviMode()
	ResetExtraConfig()
	ResetFirmware()
	ResetFolder()
	ResetForcePowerOff()
	ResetGuestId()
	ResetHardwareVersion()
	ResetHostSystemId()
	ResetHvMode()
	ResetId()
	ResetIdeControllerCount()
	ResetIgnoredGuestIps()
	ResetLatencySensitivity()
	ResetMemory()
	ResetMemoryHotAddEnabled()
	ResetMemoryLimit()
	ResetMemoryReservation()
	ResetMemoryShareCount()
	ResetMemoryShareLevel()
	ResetMigrateWaitTimeout()
	ResetNestedHvEnabled()
	ResetNetworkInterface()
	ResetNumCoresPerSocket()
	ResetNumCpus()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOvfDeploy()
	ResetPciDeviceId()
	ResetPoweronTimeout()
	ResetReplaceTrigger()
	ResetRunToolsScriptsAfterPowerOn()
	ResetRunToolsScriptsAfterResume()
	ResetRunToolsScriptsBeforeGuestReboot()
	ResetRunToolsScriptsBeforeGuestShutdown()
	ResetRunToolsScriptsBeforeGuestStandby()
	ResetSataControllerCount()
	ResetScsiBusSharing()
	ResetScsiControllerCount()
	ResetScsiType()
	ResetShutdownWaitTimeout()
	ResetStoragePolicyId()
	ResetSwapPlacementPolicy()
	ResetSyncTimeWithHost()
	ResetSyncTimeWithHostPeriodically()
	ResetTags()
	ResetToolsUpgradePolicy()
	ResetVapp()
	ResetVbsEnabled()
	ResetVvtdEnabled()
	ResetWaitForGuestIpTimeout()
	ResetWaitForGuestNetRoutable()
	ResetWaitForGuestNetTimeout()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VirtualMachine
type jsiiProxy_VirtualMachine struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VirtualMachine) AlternateGuestName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alternateGuestName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) AlternateGuestNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alternateGuestNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Annotation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) AnnotationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) BootDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) BootDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) BootRetryDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootRetryDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) BootRetryDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootRetryDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) BootRetryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootRetryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) BootRetryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootRetryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Cdrom() VirtualMachineCdromOutputReference {
	var returns VirtualMachineCdromOutputReference
	_jsii_.Get(
		j,
		"cdrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CdromInput() *VirtualMachineCdrom {
	var returns *VirtualMachineCdrom
	_jsii_.Get(
		j,
		"cdromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ChangeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Clone() VirtualMachineCloneOutputReference {
	var returns VirtualMachineCloneOutputReference
	_jsii_.Get(
		j,
		"clone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CloneInput() *VirtualMachineClone {
	var returns *VirtualMachineClone
	_jsii_.Get(
		j,
		"cloneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CpuHotAddEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuHotAddEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CpuHotAddEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuHotAddEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CpuHotRemoveEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuHotRemoveEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CpuHotRemoveEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuHotRemoveEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CpuLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CpuLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CpuPerformanceCountersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuPerformanceCountersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CpuPerformanceCountersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuPerformanceCountersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CpuReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CpuReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CpuShareCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuShareCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CpuShareCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuShareCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CpuShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CpuShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CustomAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) CustomAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) DatacenterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) DatacenterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) DatastoreClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) DatastoreClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) DatastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) DatastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) DefaultIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Disk() VirtualMachineDiskList {
	var returns VirtualMachineDiskList
	_jsii_.Get(
		j,
		"disk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) DiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) EfiSecureBootEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efiSecureBootEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) EfiSecureBootEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efiSecureBootEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) EnableDiskUuid() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDiskUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) EnableDiskUuidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDiskUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) EnableLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) EnableLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) EptRviMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eptRviMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) EptRviModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eptRviModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ExtraConfig() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ExtraConfigInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Firmware() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firmware",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) FirmwareInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firmwareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Folder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) FolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ForcePowerOff() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forcePowerOff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ForcePowerOffInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forcePowerOffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) GuestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) GuestIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guestIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) GuestIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guestIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) HardwareVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hardwareVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) HardwareVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hardwareVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) HostSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) HostSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) HvMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hvMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) HvModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hvModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) IdeControllerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ideControllerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) IdeControllerCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ideControllerCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) IgnoredGuestIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoredGuestIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) IgnoredGuestIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoredGuestIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Imported() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"imported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) LatencySensitivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latencySensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) LatencySensitivityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latencySensitivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) MemoryHotAddEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryHotAddEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) MemoryHotAddEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryHotAddEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) MemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) MemoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) MemoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) MemoryReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) MemoryReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) MemoryShareCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryShareCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) MemoryShareCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryShareCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) MemoryShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) MemoryShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) MigrateWaitTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"migrateWaitTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) MigrateWaitTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"migrateWaitTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Moid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"moid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) NestedHvEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nestedHvEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) NestedHvEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nestedHvEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) NetworkInterface() VirtualMachineNetworkInterfaceList {
	var returns VirtualMachineNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) NumCoresPerSocket() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCoresPerSocket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) NumCoresPerSocketInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCoresPerSocketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) NumCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) NumCpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) OvfDeploy() VirtualMachineOvfDeployOutputReference {
	var returns VirtualMachineOvfDeployOutputReference
	_jsii_.Get(
		j,
		"ovfDeploy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) OvfDeployInput() *VirtualMachineOvfDeploy {
	var returns *VirtualMachineOvfDeploy
	_jsii_.Get(
		j,
		"ovfDeployInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) PciDeviceId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pciDeviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) PciDeviceIdInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pciDeviceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) PoweronTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"poweronTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) PoweronTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"poweronTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) PowerState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powerState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) RebootRequired() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"rebootRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ReplaceTrigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ReplaceTriggerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ResourcePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ResourcePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) RunToolsScriptsAfterPowerOn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsAfterPowerOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) RunToolsScriptsAfterPowerOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsAfterPowerOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) RunToolsScriptsAfterResume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsAfterResume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) RunToolsScriptsAfterResumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsAfterResumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) RunToolsScriptsBeforeGuestReboot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsBeforeGuestReboot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) RunToolsScriptsBeforeGuestRebootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsBeforeGuestRebootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) RunToolsScriptsBeforeGuestShutdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsBeforeGuestShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) RunToolsScriptsBeforeGuestShutdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsBeforeGuestShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) RunToolsScriptsBeforeGuestStandby() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsBeforeGuestStandby",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) RunToolsScriptsBeforeGuestStandbyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runToolsScriptsBeforeGuestStandbyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) SataControllerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sataControllerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) SataControllerCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sataControllerCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ScsiBusSharing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scsiBusSharing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ScsiBusSharingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scsiBusSharingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ScsiControllerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scsiControllerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ScsiControllerCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scsiControllerCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ScsiType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scsiType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ScsiTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scsiTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ShutdownWaitTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shutdownWaitTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ShutdownWaitTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shutdownWaitTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) StoragePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) StoragePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) SwapPlacementPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"swapPlacementPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) SwapPlacementPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"swapPlacementPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) SyncTimeWithHost() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncTimeWithHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) SyncTimeWithHostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncTimeWithHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) SyncTimeWithHostPeriodically() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncTimeWithHostPeriodically",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) SyncTimeWithHostPeriodicallyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncTimeWithHostPeriodicallyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ToolsUpgradePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolsUpgradePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) ToolsUpgradePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolsUpgradePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) Vapp() VirtualMachineVappOutputReference {
	var returns VirtualMachineVappOutputReference
	_jsii_.Get(
		j,
		"vapp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) VappInput() *VirtualMachineVapp {
	var returns *VirtualMachineVapp
	_jsii_.Get(
		j,
		"vappInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) VappTransport() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vappTransport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) VbsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vbsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) VbsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vbsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) VmwareToolsStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmwareToolsStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) VmxPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmxPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) VvtdEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vvtdEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) VvtdEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vvtdEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) WaitForGuestIpTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitForGuestIpTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) WaitForGuestIpTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitForGuestIpTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) WaitForGuestNetRoutable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForGuestNetRoutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) WaitForGuestNetRoutableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForGuestNetRoutableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) WaitForGuestNetTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitForGuestNetTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachine) WaitForGuestNetTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitForGuestNetTimeoutInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/vsphere/r/virtual_machine vsphere_virtual_machine} Resource.
func NewVirtualMachine(scope constructs.Construct, id *string, config *VirtualMachineConfig) VirtualMachine {
	_init_.Initialize()

	if err := validateNewVirtualMachineParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachine{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachine",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/vsphere/r/virtual_machine vsphere_virtual_machine} Resource.
func NewVirtualMachine_Override(v VirtualMachine, scope constructs.Construct, id *string, config *VirtualMachineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachine",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VirtualMachine)SetAlternateGuestName(val *string) {
	if err := j.validateSetAlternateGuestNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alternateGuestName",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetAnnotation(val *string) {
	if err := j.validateSetAnnotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotation",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetBootDelay(val *float64) {
	if err := j.validateSetBootDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDelay",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetBootRetryDelay(val *float64) {
	if err := j.validateSetBootRetryDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootRetryDelay",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetBootRetryEnabled(val interface{}) {
	if err := j.validateSetBootRetryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootRetryEnabled",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetCpuHotAddEnabled(val interface{}) {
	if err := j.validateSetCpuHotAddEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuHotAddEnabled",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetCpuHotRemoveEnabled(val interface{}) {
	if err := j.validateSetCpuHotRemoveEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuHotRemoveEnabled",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetCpuLimit(val *float64) {
	if err := j.validateSetCpuLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuLimit",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetCpuPerformanceCountersEnabled(val interface{}) {
	if err := j.validateSetCpuPerformanceCountersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuPerformanceCountersEnabled",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetCpuReservation(val *float64) {
	if err := j.validateSetCpuReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuReservation",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetCpuShareCount(val *float64) {
	if err := j.validateSetCpuShareCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuShareCount",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetCpuShareLevel(val *string) {
	if err := j.validateSetCpuShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuShareLevel",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetCustomAttributes(val *map[string]*string) {
	if err := j.validateSetCustomAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAttributes",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetDatacenterId(val *string) {
	if err := j.validateSetDatacenterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datacenterId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetDatastoreClusterId(val *string) {
	if err := j.validateSetDatastoreClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datastoreClusterId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetDatastoreId(val *string) {
	if err := j.validateSetDatastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datastoreId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetEfiSecureBootEnabled(val interface{}) {
	if err := j.validateSetEfiSecureBootEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"efiSecureBootEnabled",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetEnableDiskUuid(val interface{}) {
	if err := j.validateSetEnableDiskUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDiskUuid",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetEnableLogging(val interface{}) {
	if err := j.validateSetEnableLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLogging",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetEptRviMode(val *string) {
	if err := j.validateSetEptRviModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eptRviMode",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetExtraConfig(val *map[string]*string) {
	if err := j.validateSetExtraConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraConfig",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetFirmware(val *string) {
	if err := j.validateSetFirmwareParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firmware",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetFolder(val *string) {
	if err := j.validateSetFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"folder",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetForcePowerOff(val interface{}) {
	if err := j.validateSetForcePowerOffParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forcePowerOff",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetGuestId(val *string) {
	if err := j.validateSetGuestIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guestId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetHardwareVersion(val *float64) {
	if err := j.validateSetHardwareVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hardwareVersion",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetHostSystemId(val *string) {
	if err := j.validateSetHostSystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostSystemId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetHvMode(val *string) {
	if err := j.validateSetHvModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hvMode",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetIdeControllerCount(val *float64) {
	if err := j.validateSetIdeControllerCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ideControllerCount",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetIgnoredGuestIps(val *[]*string) {
	if err := j.validateSetIgnoredGuestIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoredGuestIps",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetLatencySensitivity(val *string) {
	if err := j.validateSetLatencySensitivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latencySensitivity",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetMemory(val *float64) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetMemoryHotAddEnabled(val interface{}) {
	if err := j.validateSetMemoryHotAddEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryHotAddEnabled",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetMemoryLimit(val *float64) {
	if err := j.validateSetMemoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryLimit",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetMemoryReservation(val *float64) {
	if err := j.validateSetMemoryReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryReservation",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetMemoryShareCount(val *float64) {
	if err := j.validateSetMemoryShareCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryShareCount",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetMemoryShareLevel(val *string) {
	if err := j.validateSetMemoryShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryShareLevel",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetMigrateWaitTimeout(val *float64) {
	if err := j.validateSetMigrateWaitTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"migrateWaitTimeout",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetNestedHvEnabled(val interface{}) {
	if err := j.validateSetNestedHvEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nestedHvEnabled",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetNumCoresPerSocket(val *float64) {
	if err := j.validateSetNumCoresPerSocketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numCoresPerSocket",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetNumCpus(val *float64) {
	if err := j.validateSetNumCpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numCpus",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetPciDeviceId(val *[]*string) {
	if err := j.validateSetPciDeviceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pciDeviceId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetPoweronTimeout(val *float64) {
	if err := j.validateSetPoweronTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"poweronTimeout",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetReplaceTrigger(val *string) {
	if err := j.validateSetReplaceTriggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceTrigger",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetResourcePoolId(val *string) {
	if err := j.validateSetResourcePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePoolId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetRunToolsScriptsAfterPowerOn(val interface{}) {
	if err := j.validateSetRunToolsScriptsAfterPowerOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runToolsScriptsAfterPowerOn",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetRunToolsScriptsAfterResume(val interface{}) {
	if err := j.validateSetRunToolsScriptsAfterResumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runToolsScriptsAfterResume",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetRunToolsScriptsBeforeGuestReboot(val interface{}) {
	if err := j.validateSetRunToolsScriptsBeforeGuestRebootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runToolsScriptsBeforeGuestReboot",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetRunToolsScriptsBeforeGuestShutdown(val interface{}) {
	if err := j.validateSetRunToolsScriptsBeforeGuestShutdownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runToolsScriptsBeforeGuestShutdown",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetRunToolsScriptsBeforeGuestStandby(val interface{}) {
	if err := j.validateSetRunToolsScriptsBeforeGuestStandbyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runToolsScriptsBeforeGuestStandby",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetSataControllerCount(val *float64) {
	if err := j.validateSetSataControllerCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sataControllerCount",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetScsiBusSharing(val *string) {
	if err := j.validateSetScsiBusSharingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scsiBusSharing",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetScsiControllerCount(val *float64) {
	if err := j.validateSetScsiControllerCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scsiControllerCount",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetScsiType(val *string) {
	if err := j.validateSetScsiTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scsiType",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetShutdownWaitTimeout(val *float64) {
	if err := j.validateSetShutdownWaitTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shutdownWaitTimeout",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetStoragePolicyId(val *string) {
	if err := j.validateSetStoragePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storagePolicyId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetSwapPlacementPolicy(val *string) {
	if err := j.validateSetSwapPlacementPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"swapPlacementPolicy",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetSyncTimeWithHost(val interface{}) {
	if err := j.validateSetSyncTimeWithHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncTimeWithHost",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetSyncTimeWithHostPeriodically(val interface{}) {
	if err := j.validateSetSyncTimeWithHostPeriodicallyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncTimeWithHostPeriodically",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetToolsUpgradePolicy(val *string) {
	if err := j.validateSetToolsUpgradePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toolsUpgradePolicy",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetVbsEnabled(val interface{}) {
	if err := j.validateSetVbsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vbsEnabled",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetVvtdEnabled(val interface{}) {
	if err := j.validateSetVvtdEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vvtdEnabled",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetWaitForGuestIpTimeout(val *float64) {
	if err := j.validateSetWaitForGuestIpTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForGuestIpTimeout",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetWaitForGuestNetRoutable(val interface{}) {
	if err := j.validateSetWaitForGuestNetRoutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForGuestNetRoutable",
		val,
	)
}

func (j *jsiiProxy_VirtualMachine)SetWaitForGuestNetTimeout(val *float64) {
	if err := j.validateSetWaitForGuestNetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForGuestNetTimeout",
		val,
	)
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
func VirtualMachine_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachine_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachine",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualMachine_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachine_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachine",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualMachine_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachine_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachine",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VirtualMachine_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachine",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VirtualMachine) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VirtualMachine) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachine) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachine) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachine) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachine) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachine) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachine) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachine) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachine) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachine) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachine) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VirtualMachine) PutCdrom(value *VirtualMachineCdrom) {
	if err := v.validatePutCdromParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putCdrom",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachine) PutClone(value *VirtualMachineClone) {
	if err := v.validatePutCloneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putClone",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachine) PutDisk(value interface{}) {
	if err := v.validatePutDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDisk",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachine) PutNetworkInterface(value interface{}) {
	if err := v.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachine) PutOvfDeploy(value *VirtualMachineOvfDeploy) {
	if err := v.validatePutOvfDeployParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putOvfDeploy",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachine) PutVapp(value *VirtualMachineVapp) {
	if err := v.validatePutVappParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putVapp",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachine) ResetAlternateGuestName() {
	_jsii_.InvokeVoid(
		v,
		"resetAlternateGuestName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetAnnotation() {
	_jsii_.InvokeVoid(
		v,
		"resetAnnotation",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetBootDelay() {
	_jsii_.InvokeVoid(
		v,
		"resetBootDelay",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetBootRetryDelay() {
	_jsii_.InvokeVoid(
		v,
		"resetBootRetryDelay",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetBootRetryEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetBootRetryEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetCdrom() {
	_jsii_.InvokeVoid(
		v,
		"resetCdrom",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetClone() {
	_jsii_.InvokeVoid(
		v,
		"resetClone",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetCpuHotAddEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuHotAddEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetCpuHotRemoveEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuHotRemoveEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetCpuLimit() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuLimit",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetCpuPerformanceCountersEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuPerformanceCountersEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetCpuReservation() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuReservation",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetCpuShareCount() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuShareCount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetCpuShareLevel() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuShareLevel",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetCustomAttributes() {
	_jsii_.InvokeVoid(
		v,
		"resetCustomAttributes",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetDatacenterId() {
	_jsii_.InvokeVoid(
		v,
		"resetDatacenterId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetDatastoreClusterId() {
	_jsii_.InvokeVoid(
		v,
		"resetDatastoreClusterId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetDatastoreId() {
	_jsii_.InvokeVoid(
		v,
		"resetDatastoreId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetDisk() {
	_jsii_.InvokeVoid(
		v,
		"resetDisk",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetEfiSecureBootEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetEfiSecureBootEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetEnableDiskUuid() {
	_jsii_.InvokeVoid(
		v,
		"resetEnableDiskUuid",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetEnableLogging() {
	_jsii_.InvokeVoid(
		v,
		"resetEnableLogging",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetEptRviMode() {
	_jsii_.InvokeVoid(
		v,
		"resetEptRviMode",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetExtraConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetExtraConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetFirmware() {
	_jsii_.InvokeVoid(
		v,
		"resetFirmware",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetFolder() {
	_jsii_.InvokeVoid(
		v,
		"resetFolder",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetForcePowerOff() {
	_jsii_.InvokeVoid(
		v,
		"resetForcePowerOff",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetGuestId() {
	_jsii_.InvokeVoid(
		v,
		"resetGuestId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetHardwareVersion() {
	_jsii_.InvokeVoid(
		v,
		"resetHardwareVersion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetHostSystemId() {
	_jsii_.InvokeVoid(
		v,
		"resetHostSystemId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetHvMode() {
	_jsii_.InvokeVoid(
		v,
		"resetHvMode",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetIdeControllerCount() {
	_jsii_.InvokeVoid(
		v,
		"resetIdeControllerCount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetIgnoredGuestIps() {
	_jsii_.InvokeVoid(
		v,
		"resetIgnoredGuestIps",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetLatencySensitivity() {
	_jsii_.InvokeVoid(
		v,
		"resetLatencySensitivity",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetMemory() {
	_jsii_.InvokeVoid(
		v,
		"resetMemory",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetMemoryHotAddEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryHotAddEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetMemoryLimit() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryLimit",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetMemoryReservation() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryReservation",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetMemoryShareCount() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryShareCount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetMemoryShareLevel() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryShareLevel",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetMigrateWaitTimeout() {
	_jsii_.InvokeVoid(
		v,
		"resetMigrateWaitTimeout",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetNestedHvEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetNestedHvEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		v,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetNumCoresPerSocket() {
	_jsii_.InvokeVoid(
		v,
		"resetNumCoresPerSocket",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetNumCpus() {
	_jsii_.InvokeVoid(
		v,
		"resetNumCpus",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetOvfDeploy() {
	_jsii_.InvokeVoid(
		v,
		"resetOvfDeploy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetPciDeviceId() {
	_jsii_.InvokeVoid(
		v,
		"resetPciDeviceId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetPoweronTimeout() {
	_jsii_.InvokeVoid(
		v,
		"resetPoweronTimeout",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetReplaceTrigger() {
	_jsii_.InvokeVoid(
		v,
		"resetReplaceTrigger",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetRunToolsScriptsAfterPowerOn() {
	_jsii_.InvokeVoid(
		v,
		"resetRunToolsScriptsAfterPowerOn",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetRunToolsScriptsAfterResume() {
	_jsii_.InvokeVoid(
		v,
		"resetRunToolsScriptsAfterResume",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetRunToolsScriptsBeforeGuestReboot() {
	_jsii_.InvokeVoid(
		v,
		"resetRunToolsScriptsBeforeGuestReboot",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetRunToolsScriptsBeforeGuestShutdown() {
	_jsii_.InvokeVoid(
		v,
		"resetRunToolsScriptsBeforeGuestShutdown",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetRunToolsScriptsBeforeGuestStandby() {
	_jsii_.InvokeVoid(
		v,
		"resetRunToolsScriptsBeforeGuestStandby",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetSataControllerCount() {
	_jsii_.InvokeVoid(
		v,
		"resetSataControllerCount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetScsiBusSharing() {
	_jsii_.InvokeVoid(
		v,
		"resetScsiBusSharing",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetScsiControllerCount() {
	_jsii_.InvokeVoid(
		v,
		"resetScsiControllerCount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetScsiType() {
	_jsii_.InvokeVoid(
		v,
		"resetScsiType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetShutdownWaitTimeout() {
	_jsii_.InvokeVoid(
		v,
		"resetShutdownWaitTimeout",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetStoragePolicyId() {
	_jsii_.InvokeVoid(
		v,
		"resetStoragePolicyId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetSwapPlacementPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetSwapPlacementPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetSyncTimeWithHost() {
	_jsii_.InvokeVoid(
		v,
		"resetSyncTimeWithHost",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetSyncTimeWithHostPeriodically() {
	_jsii_.InvokeVoid(
		v,
		"resetSyncTimeWithHostPeriodically",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetToolsUpgradePolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetToolsUpgradePolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetVapp() {
	_jsii_.InvokeVoid(
		v,
		"resetVapp",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetVbsEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetVbsEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetVvtdEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetVvtdEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetWaitForGuestIpTimeout() {
	_jsii_.InvokeVoid(
		v,
		"resetWaitForGuestIpTimeout",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetWaitForGuestNetRoutable() {
	_jsii_.InvokeVoid(
		v,
		"resetWaitForGuestNetRoutable",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) ResetWaitForGuestNetTimeout() {
	_jsii_.InvokeVoid(
		v,
		"resetWaitForGuestNetTimeout",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachine) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachine) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachine) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

