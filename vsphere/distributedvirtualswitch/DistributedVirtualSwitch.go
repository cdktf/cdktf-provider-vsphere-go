// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package distributedvirtualswitch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v9/distributedvirtualswitch/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/distributed_virtual_switch vsphere_distributed_virtual_switch}.
type DistributedVirtualSwitch interface {
	cdktf.TerraformResource
	ActiveUplinks() *[]*string
	SetActiveUplinks(val *[]*string)
	ActiveUplinksInput() *[]*string
	AllowForgedTransmits() interface{}
	SetAllowForgedTransmits(val interface{})
	AllowForgedTransmitsInput() interface{}
	AllowMacChanges() interface{}
	SetAllowMacChanges(val interface{})
	AllowMacChangesInput() interface{}
	AllowPromiscuous() interface{}
	SetAllowPromiscuous(val interface{})
	AllowPromiscuousInput() interface{}
	BackupnfcMaximumMbit() *float64
	SetBackupnfcMaximumMbit(val *float64)
	BackupnfcMaximumMbitInput() *float64
	BackupnfcReservationMbit() *float64
	SetBackupnfcReservationMbit(val *float64)
	BackupnfcReservationMbitInput() *float64
	BackupnfcShareCount() *float64
	SetBackupnfcShareCount(val *float64)
	BackupnfcShareCountInput() *float64
	BackupnfcShareLevel() *string
	SetBackupnfcShareLevel(val *string)
	BackupnfcShareLevelInput() *string
	BlockAllPorts() interface{}
	SetBlockAllPorts(val interface{})
	BlockAllPortsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CheckBeacon() interface{}
	SetCheckBeacon(val interface{})
	CheckBeaconInput() interface{}
	ConfigVersion() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContactDetail() *string
	SetContactDetail(val *string)
	ContactDetailInput() *string
	ContactName() *string
	SetContactName(val *string)
	ContactNameInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomAttributes() *map[string]*string
	SetCustomAttributes(val *map[string]*string)
	CustomAttributesInput() *map[string]*string
	DatacenterId() *string
	SetDatacenterId(val *string)
	DatacenterIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DirectpathGen2Allowed() interface{}
	SetDirectpathGen2Allowed(val interface{})
	DirectpathGen2AllowedInput() interface{}
	EgressShapingAverageBandwidth() *float64
	SetEgressShapingAverageBandwidth(val *float64)
	EgressShapingAverageBandwidthInput() *float64
	EgressShapingBurstSize() *float64
	SetEgressShapingBurstSize(val *float64)
	EgressShapingBurstSizeInput() *float64
	EgressShapingEnabled() interface{}
	SetEgressShapingEnabled(val interface{})
	EgressShapingEnabledInput() interface{}
	EgressShapingPeakBandwidth() *float64
	SetEgressShapingPeakBandwidth(val *float64)
	EgressShapingPeakBandwidthInput() *float64
	Failback() interface{}
	SetFailback(val interface{})
	FailbackInput() interface{}
	FaulttoleranceMaximumMbit() *float64
	SetFaulttoleranceMaximumMbit(val *float64)
	FaulttoleranceMaximumMbitInput() *float64
	FaulttoleranceReservationMbit() *float64
	SetFaulttoleranceReservationMbit(val *float64)
	FaulttoleranceReservationMbitInput() *float64
	FaulttoleranceShareCount() *float64
	SetFaulttoleranceShareCount(val *float64)
	FaulttoleranceShareCountInput() *float64
	FaulttoleranceShareLevel() *string
	SetFaulttoleranceShareLevel(val *string)
	FaulttoleranceShareLevelInput() *string
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
	HbrMaximumMbit() *float64
	SetHbrMaximumMbit(val *float64)
	HbrMaximumMbitInput() *float64
	HbrReservationMbit() *float64
	SetHbrReservationMbit(val *float64)
	HbrReservationMbitInput() *float64
	HbrShareCount() *float64
	SetHbrShareCount(val *float64)
	HbrShareCountInput() *float64
	HbrShareLevel() *string
	SetHbrShareLevel(val *string)
	HbrShareLevelInput() *string
	Host() DistributedVirtualSwitchHostList
	HostInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	IgnoreOtherPvlanMappings() interface{}
	SetIgnoreOtherPvlanMappings(val interface{})
	IgnoreOtherPvlanMappingsInput() interface{}
	IngressShapingAverageBandwidth() *float64
	SetIngressShapingAverageBandwidth(val *float64)
	IngressShapingAverageBandwidthInput() *float64
	IngressShapingBurstSize() *float64
	SetIngressShapingBurstSize(val *float64)
	IngressShapingBurstSizeInput() *float64
	IngressShapingEnabled() interface{}
	SetIngressShapingEnabled(val interface{})
	IngressShapingEnabledInput() interface{}
	IngressShapingPeakBandwidth() *float64
	SetIngressShapingPeakBandwidth(val *float64)
	IngressShapingPeakBandwidthInput() *float64
	Ipv4Address() *string
	SetIpv4Address(val *string)
	Ipv4AddressInput() *string
	IscsiMaximumMbit() *float64
	SetIscsiMaximumMbit(val *float64)
	IscsiMaximumMbitInput() *float64
	IscsiReservationMbit() *float64
	SetIscsiReservationMbit(val *float64)
	IscsiReservationMbitInput() *float64
	IscsiShareCount() *float64
	SetIscsiShareCount(val *float64)
	IscsiShareCountInput() *float64
	IscsiShareLevel() *string
	SetIscsiShareLevel(val *string)
	IscsiShareLevelInput() *string
	LacpApiVersion() *string
	SetLacpApiVersion(val *string)
	LacpApiVersionInput() *string
	LacpEnabled() interface{}
	SetLacpEnabled(val interface{})
	LacpEnabledInput() interface{}
	LacpMode() *string
	SetLacpMode(val *string)
	LacpModeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LinkDiscoveryOperation() *string
	SetLinkDiscoveryOperation(val *string)
	LinkDiscoveryOperationInput() *string
	LinkDiscoveryProtocol() *string
	SetLinkDiscoveryProtocol(val *string)
	LinkDiscoveryProtocolInput() *string
	ManagementMaximumMbit() *float64
	SetManagementMaximumMbit(val *float64)
	ManagementMaximumMbitInput() *float64
	ManagementReservationMbit() *float64
	SetManagementReservationMbit(val *float64)
	ManagementReservationMbitInput() *float64
	ManagementShareCount() *float64
	SetManagementShareCount(val *float64)
	ManagementShareCountInput() *float64
	ManagementShareLevel() *string
	SetManagementShareLevel(val *string)
	ManagementShareLevelInput() *string
	MaxMtu() *float64
	SetMaxMtu(val *float64)
	MaxMtuInput() *float64
	MulticastFilteringMode() *string
	SetMulticastFilteringMode(val *string)
	MulticastFilteringModeInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetflowActiveFlowTimeout() *float64
	SetNetflowActiveFlowTimeout(val *float64)
	NetflowActiveFlowTimeoutInput() *float64
	NetflowCollectorIpAddress() *string
	SetNetflowCollectorIpAddress(val *string)
	NetflowCollectorIpAddressInput() *string
	NetflowCollectorPort() *float64
	SetNetflowCollectorPort(val *float64)
	NetflowCollectorPortInput() *float64
	NetflowEnabled() interface{}
	SetNetflowEnabled(val interface{})
	NetflowEnabledInput() interface{}
	NetflowIdleFlowTimeout() *float64
	SetNetflowIdleFlowTimeout(val *float64)
	NetflowIdleFlowTimeoutInput() *float64
	NetflowInternalFlowsOnly() interface{}
	SetNetflowInternalFlowsOnly(val interface{})
	NetflowInternalFlowsOnlyInput() interface{}
	NetflowObservationDomainId() *float64
	SetNetflowObservationDomainId(val *float64)
	NetflowObservationDomainIdInput() *float64
	NetflowSamplingRate() *float64
	SetNetflowSamplingRate(val *float64)
	NetflowSamplingRateInput() *float64
	NetworkResourceControlEnabled() interface{}
	SetNetworkResourceControlEnabled(val interface{})
	NetworkResourceControlEnabledInput() interface{}
	NetworkResourceControlVersion() *string
	SetNetworkResourceControlVersion(val *string)
	NetworkResourceControlVersionInput() *string
	NfsMaximumMbit() *float64
	SetNfsMaximumMbit(val *float64)
	NfsMaximumMbitInput() *float64
	NfsReservationMbit() *float64
	SetNfsReservationMbit(val *float64)
	NfsReservationMbitInput() *float64
	NfsShareCount() *float64
	SetNfsShareCount(val *float64)
	NfsShareCountInput() *float64
	NfsShareLevel() *string
	SetNfsShareLevel(val *string)
	NfsShareLevelInput() *string
	// The tree node.
	Node() constructs.Node
	NotifySwitches() interface{}
	SetNotifySwitches(val interface{})
	NotifySwitchesInput() interface{}
	PortPrivateSecondaryVlanId() *float64
	SetPortPrivateSecondaryVlanId(val *float64)
	PortPrivateSecondaryVlanIdInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PvlanMapping() DistributedVirtualSwitchPvlanMappingList
	PvlanMappingInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	StandbyUplinks() *[]*string
	SetStandbyUplinks(val *[]*string)
	StandbyUplinksInput() *[]*string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	TeamingPolicy() *string
	SetTeamingPolicy(val *string)
	TeamingPolicyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TxUplink() interface{}
	SetTxUplink(val interface{})
	TxUplinkInput() interface{}
	Uplinks() *[]*string
	SetUplinks(val *[]*string)
	UplinksInput() *[]*string
	VdpMaximumMbit() *float64
	SetVdpMaximumMbit(val *float64)
	VdpMaximumMbitInput() *float64
	VdpReservationMbit() *float64
	SetVdpReservationMbit(val *float64)
	VdpReservationMbitInput() *float64
	VdpShareCount() *float64
	SetVdpShareCount(val *float64)
	VdpShareCountInput() *float64
	VdpShareLevel() *string
	SetVdpShareLevel(val *string)
	VdpShareLevelInput() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	VirtualmachineMaximumMbit() *float64
	SetVirtualmachineMaximumMbit(val *float64)
	VirtualmachineMaximumMbitInput() *float64
	VirtualmachineReservationMbit() *float64
	SetVirtualmachineReservationMbit(val *float64)
	VirtualmachineReservationMbitInput() *float64
	VirtualmachineShareCount() *float64
	SetVirtualmachineShareCount(val *float64)
	VirtualmachineShareCountInput() *float64
	VirtualmachineShareLevel() *string
	SetVirtualmachineShareLevel(val *string)
	VirtualmachineShareLevelInput() *string
	VlanId() *float64
	SetVlanId(val *float64)
	VlanIdInput() *float64
	VlanRange() DistributedVirtualSwitchVlanRangeList
	VlanRangeInput() interface{}
	VmotionMaximumMbit() *float64
	SetVmotionMaximumMbit(val *float64)
	VmotionMaximumMbitInput() *float64
	VmotionReservationMbit() *float64
	SetVmotionReservationMbit(val *float64)
	VmotionReservationMbitInput() *float64
	VmotionShareCount() *float64
	SetVmotionShareCount(val *float64)
	VmotionShareCountInput() *float64
	VmotionShareLevel() *string
	SetVmotionShareLevel(val *string)
	VmotionShareLevelInput() *string
	VsanMaximumMbit() *float64
	SetVsanMaximumMbit(val *float64)
	VsanMaximumMbitInput() *float64
	VsanReservationMbit() *float64
	SetVsanReservationMbit(val *float64)
	VsanReservationMbitInput() *float64
	VsanShareCount() *float64
	SetVsanShareCount(val *float64)
	VsanShareCountInput() *float64
	VsanShareLevel() *string
	SetVsanShareLevel(val *string)
	VsanShareLevelInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutHost(value interface{})
	PutPvlanMapping(value interface{})
	PutVlanRange(value interface{})
	ResetActiveUplinks()
	ResetAllowForgedTransmits()
	ResetAllowMacChanges()
	ResetAllowPromiscuous()
	ResetBackupnfcMaximumMbit()
	ResetBackupnfcReservationMbit()
	ResetBackupnfcShareCount()
	ResetBackupnfcShareLevel()
	ResetBlockAllPorts()
	ResetCheckBeacon()
	ResetContactDetail()
	ResetContactName()
	ResetCustomAttributes()
	ResetDescription()
	ResetDirectpathGen2Allowed()
	ResetEgressShapingAverageBandwidth()
	ResetEgressShapingBurstSize()
	ResetEgressShapingEnabled()
	ResetEgressShapingPeakBandwidth()
	ResetFailback()
	ResetFaulttoleranceMaximumMbit()
	ResetFaulttoleranceReservationMbit()
	ResetFaulttoleranceShareCount()
	ResetFaulttoleranceShareLevel()
	ResetFolder()
	ResetHbrMaximumMbit()
	ResetHbrReservationMbit()
	ResetHbrShareCount()
	ResetHbrShareLevel()
	ResetHost()
	ResetId()
	ResetIgnoreOtherPvlanMappings()
	ResetIngressShapingAverageBandwidth()
	ResetIngressShapingBurstSize()
	ResetIngressShapingEnabled()
	ResetIngressShapingPeakBandwidth()
	ResetIpv4Address()
	ResetIscsiMaximumMbit()
	ResetIscsiReservationMbit()
	ResetIscsiShareCount()
	ResetIscsiShareLevel()
	ResetLacpApiVersion()
	ResetLacpEnabled()
	ResetLacpMode()
	ResetLinkDiscoveryOperation()
	ResetLinkDiscoveryProtocol()
	ResetManagementMaximumMbit()
	ResetManagementReservationMbit()
	ResetManagementShareCount()
	ResetManagementShareLevel()
	ResetMaxMtu()
	ResetMulticastFilteringMode()
	ResetNetflowActiveFlowTimeout()
	ResetNetflowCollectorIpAddress()
	ResetNetflowCollectorPort()
	ResetNetflowEnabled()
	ResetNetflowIdleFlowTimeout()
	ResetNetflowInternalFlowsOnly()
	ResetNetflowObservationDomainId()
	ResetNetflowSamplingRate()
	ResetNetworkResourceControlEnabled()
	ResetNetworkResourceControlVersion()
	ResetNfsMaximumMbit()
	ResetNfsReservationMbit()
	ResetNfsShareCount()
	ResetNfsShareLevel()
	ResetNotifySwitches()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPortPrivateSecondaryVlanId()
	ResetPvlanMapping()
	ResetStandbyUplinks()
	ResetTags()
	ResetTeamingPolicy()
	ResetTxUplink()
	ResetUplinks()
	ResetVdpMaximumMbit()
	ResetVdpReservationMbit()
	ResetVdpShareCount()
	ResetVdpShareLevel()
	ResetVersion()
	ResetVirtualmachineMaximumMbit()
	ResetVirtualmachineReservationMbit()
	ResetVirtualmachineShareCount()
	ResetVirtualmachineShareLevel()
	ResetVlanId()
	ResetVlanRange()
	ResetVmotionMaximumMbit()
	ResetVmotionReservationMbit()
	ResetVmotionShareCount()
	ResetVmotionShareLevel()
	ResetVsanMaximumMbit()
	ResetVsanReservationMbit()
	ResetVsanShareCount()
	ResetVsanShareLevel()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
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

// The jsii proxy struct for DistributedVirtualSwitch
type jsiiProxy_DistributedVirtualSwitch struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DistributedVirtualSwitch) ActiveUplinks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"activeUplinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) ActiveUplinksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"activeUplinksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) AllowForgedTransmits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForgedTransmits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) AllowForgedTransmitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForgedTransmitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) AllowMacChanges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMacChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) AllowMacChangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMacChangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) AllowPromiscuous() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPromiscuous",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) AllowPromiscuousInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPromiscuousInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) BackupnfcMaximumMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupnfcMaximumMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) BackupnfcMaximumMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupnfcMaximumMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) BackupnfcReservationMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupnfcReservationMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) BackupnfcReservationMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupnfcReservationMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) BackupnfcShareCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupnfcShareCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) BackupnfcShareCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupnfcShareCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) BackupnfcShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupnfcShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) BackupnfcShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupnfcShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) BlockAllPorts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockAllPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) BlockAllPortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockAllPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) CheckBeacon() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkBeacon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) CheckBeaconInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkBeaconInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) ConfigVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) ContactDetail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactDetail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) ContactDetailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactDetailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) ContactName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) ContactNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) CustomAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) CustomAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) DatacenterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) DatacenterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) DirectpathGen2Allowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directpathGen2Allowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) DirectpathGen2AllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directpathGen2AllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) EgressShapingAverageBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"egressShapingAverageBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) EgressShapingAverageBandwidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"egressShapingAverageBandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) EgressShapingBurstSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"egressShapingBurstSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) EgressShapingBurstSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"egressShapingBurstSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) EgressShapingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"egressShapingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) EgressShapingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"egressShapingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) EgressShapingPeakBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"egressShapingPeakBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) EgressShapingPeakBandwidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"egressShapingPeakBandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) Failback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) FailbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) FaulttoleranceMaximumMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"faulttoleranceMaximumMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) FaulttoleranceMaximumMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"faulttoleranceMaximumMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) FaulttoleranceReservationMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"faulttoleranceReservationMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) FaulttoleranceReservationMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"faulttoleranceReservationMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) FaulttoleranceShareCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"faulttoleranceShareCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) FaulttoleranceShareCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"faulttoleranceShareCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) FaulttoleranceShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faulttoleranceShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) FaulttoleranceShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faulttoleranceShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) Folder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) FolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) HbrMaximumMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hbrMaximumMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) HbrMaximumMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hbrMaximumMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) HbrReservationMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hbrReservationMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) HbrReservationMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hbrReservationMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) HbrShareCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hbrShareCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) HbrShareCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hbrShareCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) HbrShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hbrShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) HbrShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hbrShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) Host() DistributedVirtualSwitchHostList {
	var returns DistributedVirtualSwitchHostList
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) HostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IgnoreOtherPvlanMappings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreOtherPvlanMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IgnoreOtherPvlanMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreOtherPvlanMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IngressShapingAverageBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressShapingAverageBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IngressShapingAverageBandwidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressShapingAverageBandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IngressShapingBurstSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressShapingBurstSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IngressShapingBurstSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressShapingBurstSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IngressShapingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingressShapingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IngressShapingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingressShapingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IngressShapingPeakBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressShapingPeakBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IngressShapingPeakBandwidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressShapingPeakBandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) Ipv4Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) Ipv4AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IscsiMaximumMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iscsiMaximumMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IscsiMaximumMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iscsiMaximumMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IscsiReservationMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iscsiReservationMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IscsiReservationMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iscsiReservationMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IscsiShareCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iscsiShareCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IscsiShareCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iscsiShareCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IscsiShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iscsiShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) IscsiShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iscsiShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) LacpApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lacpApiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) LacpApiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lacpApiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) LacpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lacpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) LacpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lacpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) LacpMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lacpMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) LacpModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lacpModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) LinkDiscoveryOperation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkDiscoveryOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) LinkDiscoveryOperationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkDiscoveryOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) LinkDiscoveryProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkDiscoveryProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) LinkDiscoveryProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkDiscoveryProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) ManagementMaximumMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managementMaximumMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) ManagementMaximumMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managementMaximumMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) ManagementReservationMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managementReservationMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) ManagementReservationMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managementReservationMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) ManagementShareCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managementShareCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) ManagementShareCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managementShareCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) ManagementShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) ManagementShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) MaxMtu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMtu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) MaxMtuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMtuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) MulticastFilteringMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multicastFilteringMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) MulticastFilteringModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multicastFilteringModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetflowActiveFlowTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netflowActiveFlowTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetflowActiveFlowTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netflowActiveFlowTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetflowCollectorIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netflowCollectorIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetflowCollectorIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netflowCollectorIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetflowCollectorPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netflowCollectorPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetflowCollectorPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netflowCollectorPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetflowEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"netflowEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetflowEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"netflowEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetflowIdleFlowTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netflowIdleFlowTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetflowIdleFlowTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netflowIdleFlowTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetflowInternalFlowsOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"netflowInternalFlowsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetflowInternalFlowsOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"netflowInternalFlowsOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetflowObservationDomainId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netflowObservationDomainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetflowObservationDomainIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netflowObservationDomainIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetflowSamplingRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netflowSamplingRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetflowSamplingRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netflowSamplingRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetworkResourceControlEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkResourceControlEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetworkResourceControlEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkResourceControlEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetworkResourceControlVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkResourceControlVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NetworkResourceControlVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkResourceControlVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NfsMaximumMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nfsMaximumMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NfsMaximumMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nfsMaximumMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NfsReservationMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nfsReservationMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NfsReservationMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nfsReservationMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NfsShareCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nfsShareCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NfsShareCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nfsShareCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NfsShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nfsShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NfsShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nfsShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NotifySwitches() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifySwitches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) NotifySwitchesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifySwitchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) PortPrivateSecondaryVlanId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portPrivateSecondaryVlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) PortPrivateSecondaryVlanIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portPrivateSecondaryVlanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) PvlanMapping() DistributedVirtualSwitchPvlanMappingList {
	var returns DistributedVirtualSwitchPvlanMappingList
	_jsii_.Get(
		j,
		"pvlanMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) PvlanMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pvlanMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) StandbyUplinks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"standbyUplinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) StandbyUplinksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"standbyUplinksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) TeamingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) TeamingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) TxUplink() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"txUplink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) TxUplinkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"txUplinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) Uplinks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"uplinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) UplinksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"uplinksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VdpMaximumMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vdpMaximumMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VdpMaximumMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vdpMaximumMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VdpReservationMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vdpReservationMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VdpReservationMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vdpReservationMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VdpShareCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vdpShareCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VdpShareCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vdpShareCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VdpShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdpShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VdpShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdpShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VirtualmachineMaximumMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"virtualmachineMaximumMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VirtualmachineMaximumMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"virtualmachineMaximumMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VirtualmachineReservationMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"virtualmachineReservationMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VirtualmachineReservationMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"virtualmachineReservationMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VirtualmachineShareCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"virtualmachineShareCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VirtualmachineShareCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"virtualmachineShareCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VirtualmachineShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualmachineShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VirtualmachineShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualmachineShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VlanId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VlanIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VlanRange() DistributedVirtualSwitchVlanRangeList {
	var returns DistributedVirtualSwitchVlanRangeList
	_jsii_.Get(
		j,
		"vlanRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VlanRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vlanRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VmotionMaximumMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmotionMaximumMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VmotionMaximumMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmotionMaximumMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VmotionReservationMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmotionReservationMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VmotionReservationMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmotionReservationMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VmotionShareCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmotionShareCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VmotionShareCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmotionShareCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VmotionShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmotionShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VmotionShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmotionShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VsanMaximumMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vsanMaximumMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VsanMaximumMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vsanMaximumMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VsanReservationMbit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vsanReservationMbit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VsanReservationMbitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vsanReservationMbitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VsanShareCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vsanShareCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VsanShareCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vsanShareCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VsanShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vsanShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitch) VsanShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vsanShareLevelInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/distributed_virtual_switch vsphere_distributed_virtual_switch} Resource.
func NewDistributedVirtualSwitch(scope constructs.Construct, id *string, config *DistributedVirtualSwitchConfig) DistributedVirtualSwitch {
	_init_.Initialize()

	if err := validateNewDistributedVirtualSwitchParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DistributedVirtualSwitch{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.distributedVirtualSwitch.DistributedVirtualSwitch",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/distributed_virtual_switch vsphere_distributed_virtual_switch} Resource.
func NewDistributedVirtualSwitch_Override(d DistributedVirtualSwitch, scope constructs.Construct, id *string, config *DistributedVirtualSwitchConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.distributedVirtualSwitch.DistributedVirtualSwitch",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetActiveUplinks(val *[]*string) {
	if err := j.validateSetActiveUplinksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeUplinks",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetAllowForgedTransmits(val interface{}) {
	if err := j.validateSetAllowForgedTransmitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowForgedTransmits",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetAllowMacChanges(val interface{}) {
	if err := j.validateSetAllowMacChangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowMacChanges",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetAllowPromiscuous(val interface{}) {
	if err := j.validateSetAllowPromiscuousParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPromiscuous",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetBackupnfcMaximumMbit(val *float64) {
	if err := j.validateSetBackupnfcMaximumMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupnfcMaximumMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetBackupnfcReservationMbit(val *float64) {
	if err := j.validateSetBackupnfcReservationMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupnfcReservationMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetBackupnfcShareCount(val *float64) {
	if err := j.validateSetBackupnfcShareCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupnfcShareCount",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetBackupnfcShareLevel(val *string) {
	if err := j.validateSetBackupnfcShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupnfcShareLevel",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetBlockAllPorts(val interface{}) {
	if err := j.validateSetBlockAllPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockAllPorts",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetCheckBeacon(val interface{}) {
	if err := j.validateSetCheckBeaconParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkBeacon",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetContactDetail(val *string) {
	if err := j.validateSetContactDetailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactDetail",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetContactName(val *string) {
	if err := j.validateSetContactNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactName",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetCustomAttributes(val *map[string]*string) {
	if err := j.validateSetCustomAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAttributes",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetDatacenterId(val *string) {
	if err := j.validateSetDatacenterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datacenterId",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetDirectpathGen2Allowed(val interface{}) {
	if err := j.validateSetDirectpathGen2AllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directpathGen2Allowed",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetEgressShapingAverageBandwidth(val *float64) {
	if err := j.validateSetEgressShapingAverageBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressShapingAverageBandwidth",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetEgressShapingBurstSize(val *float64) {
	if err := j.validateSetEgressShapingBurstSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressShapingBurstSize",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetEgressShapingEnabled(val interface{}) {
	if err := j.validateSetEgressShapingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressShapingEnabled",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetEgressShapingPeakBandwidth(val *float64) {
	if err := j.validateSetEgressShapingPeakBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressShapingPeakBandwidth",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetFailback(val interface{}) {
	if err := j.validateSetFailbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failback",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetFaulttoleranceMaximumMbit(val *float64) {
	if err := j.validateSetFaulttoleranceMaximumMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"faulttoleranceMaximumMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetFaulttoleranceReservationMbit(val *float64) {
	if err := j.validateSetFaulttoleranceReservationMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"faulttoleranceReservationMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetFaulttoleranceShareCount(val *float64) {
	if err := j.validateSetFaulttoleranceShareCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"faulttoleranceShareCount",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetFaulttoleranceShareLevel(val *string) {
	if err := j.validateSetFaulttoleranceShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"faulttoleranceShareLevel",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetFolder(val *string) {
	if err := j.validateSetFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"folder",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetHbrMaximumMbit(val *float64) {
	if err := j.validateSetHbrMaximumMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hbrMaximumMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetHbrReservationMbit(val *float64) {
	if err := j.validateSetHbrReservationMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hbrReservationMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetHbrShareCount(val *float64) {
	if err := j.validateSetHbrShareCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hbrShareCount",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetHbrShareLevel(val *string) {
	if err := j.validateSetHbrShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hbrShareLevel",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetIgnoreOtherPvlanMappings(val interface{}) {
	if err := j.validateSetIgnoreOtherPvlanMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreOtherPvlanMappings",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetIngressShapingAverageBandwidth(val *float64) {
	if err := j.validateSetIngressShapingAverageBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressShapingAverageBandwidth",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetIngressShapingBurstSize(val *float64) {
	if err := j.validateSetIngressShapingBurstSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressShapingBurstSize",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetIngressShapingEnabled(val interface{}) {
	if err := j.validateSetIngressShapingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressShapingEnabled",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetIngressShapingPeakBandwidth(val *float64) {
	if err := j.validateSetIngressShapingPeakBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressShapingPeakBandwidth",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetIpv4Address(val *string) {
	if err := j.validateSetIpv4AddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Address",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetIscsiMaximumMbit(val *float64) {
	if err := j.validateSetIscsiMaximumMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iscsiMaximumMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetIscsiReservationMbit(val *float64) {
	if err := j.validateSetIscsiReservationMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iscsiReservationMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetIscsiShareCount(val *float64) {
	if err := j.validateSetIscsiShareCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iscsiShareCount",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetIscsiShareLevel(val *string) {
	if err := j.validateSetIscsiShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iscsiShareLevel",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetLacpApiVersion(val *string) {
	if err := j.validateSetLacpApiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lacpApiVersion",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetLacpEnabled(val interface{}) {
	if err := j.validateSetLacpEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lacpEnabled",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetLacpMode(val *string) {
	if err := j.validateSetLacpModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lacpMode",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetLinkDiscoveryOperation(val *string) {
	if err := j.validateSetLinkDiscoveryOperationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkDiscoveryOperation",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetLinkDiscoveryProtocol(val *string) {
	if err := j.validateSetLinkDiscoveryProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkDiscoveryProtocol",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetManagementMaximumMbit(val *float64) {
	if err := j.validateSetManagementMaximumMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managementMaximumMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetManagementReservationMbit(val *float64) {
	if err := j.validateSetManagementReservationMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managementReservationMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetManagementShareCount(val *float64) {
	if err := j.validateSetManagementShareCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managementShareCount",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetManagementShareLevel(val *string) {
	if err := j.validateSetManagementShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managementShareLevel",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetMaxMtu(val *float64) {
	if err := j.validateSetMaxMtuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxMtu",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetMulticastFilteringMode(val *string) {
	if err := j.validateSetMulticastFilteringModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multicastFilteringMode",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetNetflowActiveFlowTimeout(val *float64) {
	if err := j.validateSetNetflowActiveFlowTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netflowActiveFlowTimeout",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetNetflowCollectorIpAddress(val *string) {
	if err := j.validateSetNetflowCollectorIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netflowCollectorIpAddress",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetNetflowCollectorPort(val *float64) {
	if err := j.validateSetNetflowCollectorPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netflowCollectorPort",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetNetflowEnabled(val interface{}) {
	if err := j.validateSetNetflowEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netflowEnabled",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetNetflowIdleFlowTimeout(val *float64) {
	if err := j.validateSetNetflowIdleFlowTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netflowIdleFlowTimeout",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetNetflowInternalFlowsOnly(val interface{}) {
	if err := j.validateSetNetflowInternalFlowsOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netflowInternalFlowsOnly",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetNetflowObservationDomainId(val *float64) {
	if err := j.validateSetNetflowObservationDomainIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netflowObservationDomainId",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetNetflowSamplingRate(val *float64) {
	if err := j.validateSetNetflowSamplingRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netflowSamplingRate",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetNetworkResourceControlEnabled(val interface{}) {
	if err := j.validateSetNetworkResourceControlEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkResourceControlEnabled",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetNetworkResourceControlVersion(val *string) {
	if err := j.validateSetNetworkResourceControlVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkResourceControlVersion",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetNfsMaximumMbit(val *float64) {
	if err := j.validateSetNfsMaximumMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nfsMaximumMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetNfsReservationMbit(val *float64) {
	if err := j.validateSetNfsReservationMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nfsReservationMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetNfsShareCount(val *float64) {
	if err := j.validateSetNfsShareCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nfsShareCount",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetNfsShareLevel(val *string) {
	if err := j.validateSetNfsShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nfsShareLevel",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetNotifySwitches(val interface{}) {
	if err := j.validateSetNotifySwitchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifySwitches",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetPortPrivateSecondaryVlanId(val *float64) {
	if err := j.validateSetPortPrivateSecondaryVlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portPrivateSecondaryVlanId",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetStandbyUplinks(val *[]*string) {
	if err := j.validateSetStandbyUplinksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"standbyUplinks",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetTeamingPolicy(val *string) {
	if err := j.validateSetTeamingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamingPolicy",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetTxUplink(val interface{}) {
	if err := j.validateSetTxUplinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"txUplink",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetUplinks(val *[]*string) {
	if err := j.validateSetUplinksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uplinks",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetVdpMaximumMbit(val *float64) {
	if err := j.validateSetVdpMaximumMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vdpMaximumMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetVdpReservationMbit(val *float64) {
	if err := j.validateSetVdpReservationMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vdpReservationMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetVdpShareCount(val *float64) {
	if err := j.validateSetVdpShareCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vdpShareCount",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetVdpShareLevel(val *string) {
	if err := j.validateSetVdpShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vdpShareLevel",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetVirtualmachineMaximumMbit(val *float64) {
	if err := j.validateSetVirtualmachineMaximumMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualmachineMaximumMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetVirtualmachineReservationMbit(val *float64) {
	if err := j.validateSetVirtualmachineReservationMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualmachineReservationMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetVirtualmachineShareCount(val *float64) {
	if err := j.validateSetVirtualmachineShareCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualmachineShareCount",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetVirtualmachineShareLevel(val *string) {
	if err := j.validateSetVirtualmachineShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualmachineShareLevel",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetVlanId(val *float64) {
	if err := j.validateSetVlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vlanId",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetVmotionMaximumMbit(val *float64) {
	if err := j.validateSetVmotionMaximumMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmotionMaximumMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetVmotionReservationMbit(val *float64) {
	if err := j.validateSetVmotionReservationMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmotionReservationMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetVmotionShareCount(val *float64) {
	if err := j.validateSetVmotionShareCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmotionShareCount",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetVmotionShareLevel(val *string) {
	if err := j.validateSetVmotionShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmotionShareLevel",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetVsanMaximumMbit(val *float64) {
	if err := j.validateSetVsanMaximumMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vsanMaximumMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetVsanReservationMbit(val *float64) {
	if err := j.validateSetVsanReservationMbitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vsanReservationMbit",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetVsanShareCount(val *float64) {
	if err := j.validateSetVsanShareCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vsanShareCount",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitch)SetVsanShareLevel(val *string) {
	if err := j.validateSetVsanShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vsanShareLevel",
		val,
	)
}

// Generates CDKTF code for importing a DistributedVirtualSwitch resource upon running "cdktf plan <stack-name>".
func DistributedVirtualSwitch_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDistributedVirtualSwitch_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.distributedVirtualSwitch.DistributedVirtualSwitch",
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
func DistributedVirtualSwitch_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDistributedVirtualSwitch_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.distributedVirtualSwitch.DistributedVirtualSwitch",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DistributedVirtualSwitch_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDistributedVirtualSwitch_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.distributedVirtualSwitch.DistributedVirtualSwitch",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DistributedVirtualSwitch_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDistributedVirtualSwitch_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.distributedVirtualSwitch.DistributedVirtualSwitch",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DistributedVirtualSwitch_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vsphere.distributedVirtualSwitch.DistributedVirtualSwitch",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DistributedVirtualSwitch) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DistributedVirtualSwitch) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DistributedVirtualSwitch) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DistributedVirtualSwitch) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DistributedVirtualSwitch) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DistributedVirtualSwitch) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DistributedVirtualSwitch) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DistributedVirtualSwitch) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DistributedVirtualSwitch) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DistributedVirtualSwitch) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedVirtualSwitch) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DistributedVirtualSwitch) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) PutHost(value interface{}) {
	if err := d.validatePutHostParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHost",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) PutPvlanMapping(value interface{}) {
	if err := d.validatePutPvlanMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPvlanMapping",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) PutVlanRange(value interface{}) {
	if err := d.validatePutVlanRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVlanRange",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetActiveUplinks() {
	_jsii_.InvokeVoid(
		d,
		"resetActiveUplinks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetAllowForgedTransmits() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowForgedTransmits",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetAllowMacChanges() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowMacChanges",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetAllowPromiscuous() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowPromiscuous",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetBackupnfcMaximumMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupnfcMaximumMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetBackupnfcReservationMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupnfcReservationMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetBackupnfcShareCount() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupnfcShareCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetBackupnfcShareLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupnfcShareLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetBlockAllPorts() {
	_jsii_.InvokeVoid(
		d,
		"resetBlockAllPorts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetCheckBeacon() {
	_jsii_.InvokeVoid(
		d,
		"resetCheckBeacon",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetContactDetail() {
	_jsii_.InvokeVoid(
		d,
		"resetContactDetail",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetContactName() {
	_jsii_.InvokeVoid(
		d,
		"resetContactName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetCustomAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetDirectpathGen2Allowed() {
	_jsii_.InvokeVoid(
		d,
		"resetDirectpathGen2Allowed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetEgressShapingAverageBandwidth() {
	_jsii_.InvokeVoid(
		d,
		"resetEgressShapingAverageBandwidth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetEgressShapingBurstSize() {
	_jsii_.InvokeVoid(
		d,
		"resetEgressShapingBurstSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetEgressShapingEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetEgressShapingEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetEgressShapingPeakBandwidth() {
	_jsii_.InvokeVoid(
		d,
		"resetEgressShapingPeakBandwidth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetFailback() {
	_jsii_.InvokeVoid(
		d,
		"resetFailback",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetFaulttoleranceMaximumMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetFaulttoleranceMaximumMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetFaulttoleranceReservationMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetFaulttoleranceReservationMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetFaulttoleranceShareCount() {
	_jsii_.InvokeVoid(
		d,
		"resetFaulttoleranceShareCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetFaulttoleranceShareLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetFaulttoleranceShareLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetFolder() {
	_jsii_.InvokeVoid(
		d,
		"resetFolder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetHbrMaximumMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetHbrMaximumMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetHbrReservationMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetHbrReservationMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetHbrShareCount() {
	_jsii_.InvokeVoid(
		d,
		"resetHbrShareCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetHbrShareLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetHbrShareLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetHost() {
	_jsii_.InvokeVoid(
		d,
		"resetHost",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetIgnoreOtherPvlanMappings() {
	_jsii_.InvokeVoid(
		d,
		"resetIgnoreOtherPvlanMappings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetIngressShapingAverageBandwidth() {
	_jsii_.InvokeVoid(
		d,
		"resetIngressShapingAverageBandwidth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetIngressShapingBurstSize() {
	_jsii_.InvokeVoid(
		d,
		"resetIngressShapingBurstSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetIngressShapingEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetIngressShapingEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetIngressShapingPeakBandwidth() {
	_jsii_.InvokeVoid(
		d,
		"resetIngressShapingPeakBandwidth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetIpv4Address() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv4Address",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetIscsiMaximumMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetIscsiMaximumMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetIscsiReservationMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetIscsiReservationMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetIscsiShareCount() {
	_jsii_.InvokeVoid(
		d,
		"resetIscsiShareCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetIscsiShareLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetIscsiShareLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetLacpApiVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetLacpApiVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetLacpEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetLacpEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetLacpMode() {
	_jsii_.InvokeVoid(
		d,
		"resetLacpMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetLinkDiscoveryOperation() {
	_jsii_.InvokeVoid(
		d,
		"resetLinkDiscoveryOperation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetLinkDiscoveryProtocol() {
	_jsii_.InvokeVoid(
		d,
		"resetLinkDiscoveryProtocol",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetManagementMaximumMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetManagementMaximumMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetManagementReservationMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetManagementReservationMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetManagementShareCount() {
	_jsii_.InvokeVoid(
		d,
		"resetManagementShareCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetManagementShareLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetManagementShareLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetMaxMtu() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxMtu",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetMulticastFilteringMode() {
	_jsii_.InvokeVoid(
		d,
		"resetMulticastFilteringMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetNetflowActiveFlowTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetNetflowActiveFlowTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetNetflowCollectorIpAddress() {
	_jsii_.InvokeVoid(
		d,
		"resetNetflowCollectorIpAddress",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetNetflowCollectorPort() {
	_jsii_.InvokeVoid(
		d,
		"resetNetflowCollectorPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetNetflowEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetNetflowEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetNetflowIdleFlowTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetNetflowIdleFlowTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetNetflowInternalFlowsOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetNetflowInternalFlowsOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetNetflowObservationDomainId() {
	_jsii_.InvokeVoid(
		d,
		"resetNetflowObservationDomainId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetNetflowSamplingRate() {
	_jsii_.InvokeVoid(
		d,
		"resetNetflowSamplingRate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetNetworkResourceControlEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkResourceControlEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetNetworkResourceControlVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkResourceControlVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetNfsMaximumMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetNfsMaximumMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetNfsReservationMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetNfsReservationMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetNfsShareCount() {
	_jsii_.InvokeVoid(
		d,
		"resetNfsShareCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetNfsShareLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetNfsShareLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetNotifySwitches() {
	_jsii_.InvokeVoid(
		d,
		"resetNotifySwitches",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetPortPrivateSecondaryVlanId() {
	_jsii_.InvokeVoid(
		d,
		"resetPortPrivateSecondaryVlanId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetPvlanMapping() {
	_jsii_.InvokeVoid(
		d,
		"resetPvlanMapping",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetStandbyUplinks() {
	_jsii_.InvokeVoid(
		d,
		"resetStandbyUplinks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetTeamingPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetTeamingPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetTxUplink() {
	_jsii_.InvokeVoid(
		d,
		"resetTxUplink",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetUplinks() {
	_jsii_.InvokeVoid(
		d,
		"resetUplinks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVdpMaximumMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetVdpMaximumMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVdpReservationMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetVdpReservationMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVdpShareCount() {
	_jsii_.InvokeVoid(
		d,
		"resetVdpShareCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVdpShareLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetVdpShareLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVirtualmachineMaximumMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetVirtualmachineMaximumMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVirtualmachineReservationMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetVirtualmachineReservationMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVirtualmachineShareCount() {
	_jsii_.InvokeVoid(
		d,
		"resetVirtualmachineShareCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVirtualmachineShareLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetVirtualmachineShareLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVlanId() {
	_jsii_.InvokeVoid(
		d,
		"resetVlanId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVlanRange() {
	_jsii_.InvokeVoid(
		d,
		"resetVlanRange",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVmotionMaximumMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetVmotionMaximumMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVmotionReservationMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetVmotionReservationMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVmotionShareCount() {
	_jsii_.InvokeVoid(
		d,
		"resetVmotionShareCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVmotionShareLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetVmotionShareLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVsanMaximumMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetVsanMaximumMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVsanReservationMbit() {
	_jsii_.InvokeVoid(
		d,
		"resetVsanReservationMbit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVsanShareCount() {
	_jsii_.InvokeVoid(
		d,
		"resetVsanShareCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) ResetVsanShareLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetVsanShareLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitch) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedVirtualSwitch) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedVirtualSwitch) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedVirtualSwitch) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedVirtualSwitch) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedVirtualSwitch) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

