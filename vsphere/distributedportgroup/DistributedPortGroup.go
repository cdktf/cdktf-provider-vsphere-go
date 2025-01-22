// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package distributedportgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v9/distributedportgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/resources/distributed_port_group vsphere_distributed_port_group}.
type DistributedPortGroup interface {
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
	AutoExpand() interface{}
	SetAutoExpand(val interface{})
	AutoExpandInput() interface{}
	BlockAllPorts() interface{}
	SetBlockAllPorts(val interface{})
	BlockAllPortsInput() interface{}
	BlockOverrideAllowed() interface{}
	SetBlockOverrideAllowed(val interface{})
	BlockOverrideAllowedInput() interface{}
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
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomAttributes() *map[string]*string
	SetCustomAttributes(val *map[string]*string)
	CustomAttributesInput() *map[string]*string
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
	DistributedVirtualSwitchUuid() *string
	SetDistributedVirtualSwitchUuid(val *string)
	DistributedVirtualSwitchUuidInput() *string
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
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
	Key() *string
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
	LivePortMovingAllowed() interface{}
	SetLivePortMovingAllowed(val interface{})
	LivePortMovingAllowedInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetflowEnabled() interface{}
	SetNetflowEnabled(val interface{})
	NetflowEnabledInput() interface{}
	NetflowOverrideAllowed() interface{}
	SetNetflowOverrideAllowed(val interface{})
	NetflowOverrideAllowedInput() interface{}
	NetworkResourcePoolKey() *string
	SetNetworkResourcePoolKey(val *string)
	NetworkResourcePoolKeyInput() *string
	NetworkResourcePoolOverrideAllowed() interface{}
	SetNetworkResourcePoolOverrideAllowed(val interface{})
	NetworkResourcePoolOverrideAllowedInput() interface{}
	// The tree node.
	Node() constructs.Node
	NotifySwitches() interface{}
	SetNotifySwitches(val interface{})
	NotifySwitchesInput() interface{}
	NumberOfPorts() *float64
	SetNumberOfPorts(val *float64)
	NumberOfPortsInput() *float64
	PortConfigResetAtDisconnect() interface{}
	SetPortConfigResetAtDisconnect(val interface{})
	PortConfigResetAtDisconnectInput() interface{}
	PortNameFormat() *string
	SetPortNameFormat(val *string)
	PortNameFormatInput() *string
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
	// Experimental.
	RawOverrides() interface{}
	SecurityPolicyOverrideAllowed() interface{}
	SetSecurityPolicyOverrideAllowed(val interface{})
	SecurityPolicyOverrideAllowedInput() interface{}
	ShapingOverrideAllowed() interface{}
	SetShapingOverrideAllowed(val interface{})
	ShapingOverrideAllowedInput() interface{}
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
	TrafficFilterOverrideAllowed() interface{}
	SetTrafficFilterOverrideAllowed(val interface{})
	TrafficFilterOverrideAllowedInput() interface{}
	TxUplink() interface{}
	SetTxUplink(val interface{})
	TxUplinkInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UplinkTeamingOverrideAllowed() interface{}
	SetUplinkTeamingOverrideAllowed(val interface{})
	UplinkTeamingOverrideAllowedInput() interface{}
	VlanId() *float64
	SetVlanId(val *float64)
	VlanIdInput() *float64
	VlanOverrideAllowed() interface{}
	SetVlanOverrideAllowed(val interface{})
	VlanOverrideAllowedInput() interface{}
	VlanRange() DistributedPortGroupVlanRangeList
	VlanRangeInput() interface{}
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
	PutVlanRange(value interface{})
	ResetActiveUplinks()
	ResetAllowForgedTransmits()
	ResetAllowMacChanges()
	ResetAllowPromiscuous()
	ResetAutoExpand()
	ResetBlockAllPorts()
	ResetBlockOverrideAllowed()
	ResetCheckBeacon()
	ResetCustomAttributes()
	ResetDescription()
	ResetDirectpathGen2Allowed()
	ResetEgressShapingAverageBandwidth()
	ResetEgressShapingBurstSize()
	ResetEgressShapingEnabled()
	ResetEgressShapingPeakBandwidth()
	ResetFailback()
	ResetId()
	ResetIngressShapingAverageBandwidth()
	ResetIngressShapingBurstSize()
	ResetIngressShapingEnabled()
	ResetIngressShapingPeakBandwidth()
	ResetLacpEnabled()
	ResetLacpMode()
	ResetLivePortMovingAllowed()
	ResetNetflowEnabled()
	ResetNetflowOverrideAllowed()
	ResetNetworkResourcePoolKey()
	ResetNetworkResourcePoolOverrideAllowed()
	ResetNotifySwitches()
	ResetNumberOfPorts()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPortConfigResetAtDisconnect()
	ResetPortNameFormat()
	ResetPortPrivateSecondaryVlanId()
	ResetSecurityPolicyOverrideAllowed()
	ResetShapingOverrideAllowed()
	ResetStandbyUplinks()
	ResetTags()
	ResetTeamingPolicy()
	ResetTrafficFilterOverrideAllowed()
	ResetTxUplink()
	ResetType()
	ResetUplinkTeamingOverrideAllowed()
	ResetVlanId()
	ResetVlanOverrideAllowed()
	ResetVlanRange()
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

// The jsii proxy struct for DistributedPortGroup
type jsiiProxy_DistributedPortGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DistributedPortGroup) ActiveUplinks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"activeUplinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) ActiveUplinksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"activeUplinksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) AllowForgedTransmits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForgedTransmits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) AllowForgedTransmitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForgedTransmitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) AllowMacChanges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMacChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) AllowMacChangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMacChangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) AllowPromiscuous() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPromiscuous",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) AllowPromiscuousInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPromiscuousInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) AutoExpand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoExpand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) AutoExpandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoExpandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) BlockAllPorts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockAllPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) BlockAllPortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockAllPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) BlockOverrideAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockOverrideAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) BlockOverrideAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockOverrideAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) CheckBeacon() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkBeacon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) CheckBeaconInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkBeaconInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) ConfigVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) CustomAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) CustomAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) DirectpathGen2Allowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directpathGen2Allowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) DirectpathGen2AllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directpathGen2AllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) DistributedVirtualSwitchUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributedVirtualSwitchUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) DistributedVirtualSwitchUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributedVirtualSwitchUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) EgressShapingAverageBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"egressShapingAverageBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) EgressShapingAverageBandwidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"egressShapingAverageBandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) EgressShapingBurstSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"egressShapingBurstSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) EgressShapingBurstSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"egressShapingBurstSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) EgressShapingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"egressShapingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) EgressShapingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"egressShapingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) EgressShapingPeakBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"egressShapingPeakBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) EgressShapingPeakBandwidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"egressShapingPeakBandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) Failback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) FailbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) IngressShapingAverageBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressShapingAverageBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) IngressShapingAverageBandwidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressShapingAverageBandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) IngressShapingBurstSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressShapingBurstSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) IngressShapingBurstSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressShapingBurstSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) IngressShapingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingressShapingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) IngressShapingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingressShapingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) IngressShapingPeakBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressShapingPeakBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) IngressShapingPeakBandwidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressShapingPeakBandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) LacpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lacpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) LacpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lacpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) LacpMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lacpMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) LacpModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lacpModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) LivePortMovingAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"livePortMovingAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) LivePortMovingAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"livePortMovingAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) NetflowEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"netflowEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) NetflowEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"netflowEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) NetflowOverrideAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"netflowOverrideAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) NetflowOverrideAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"netflowOverrideAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) NetworkResourcePoolKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkResourcePoolKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) NetworkResourcePoolKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkResourcePoolKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) NetworkResourcePoolOverrideAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkResourcePoolOverrideAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) NetworkResourcePoolOverrideAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkResourcePoolOverrideAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) NotifySwitches() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifySwitches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) NotifySwitchesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifySwitchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) NumberOfPorts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) NumberOfPortsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) PortConfigResetAtDisconnect() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portConfigResetAtDisconnect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) PortConfigResetAtDisconnectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portConfigResetAtDisconnectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) PortNameFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portNameFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) PortNameFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portNameFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) PortPrivateSecondaryVlanId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portPrivateSecondaryVlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) PortPrivateSecondaryVlanIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portPrivateSecondaryVlanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) SecurityPolicyOverrideAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityPolicyOverrideAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) SecurityPolicyOverrideAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityPolicyOverrideAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) ShapingOverrideAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shapingOverrideAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) ShapingOverrideAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shapingOverrideAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) StandbyUplinks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"standbyUplinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) StandbyUplinksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"standbyUplinksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) TeamingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) TeamingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) TrafficFilterOverrideAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficFilterOverrideAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) TrafficFilterOverrideAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficFilterOverrideAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) TxUplink() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"txUplink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) TxUplinkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"txUplinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) UplinkTeamingOverrideAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uplinkTeamingOverrideAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) UplinkTeamingOverrideAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uplinkTeamingOverrideAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) VlanId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) VlanIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) VlanOverrideAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vlanOverrideAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) VlanOverrideAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vlanOverrideAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) VlanRange() DistributedPortGroupVlanRangeList {
	var returns DistributedPortGroupVlanRangeList
	_jsii_.Get(
		j,
		"vlanRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedPortGroup) VlanRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vlanRangeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/resources/distributed_port_group vsphere_distributed_port_group} Resource.
func NewDistributedPortGroup(scope constructs.Construct, id *string, config *DistributedPortGroupConfig) DistributedPortGroup {
	_init_.Initialize()

	if err := validateNewDistributedPortGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DistributedPortGroup{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.distributedPortGroup.DistributedPortGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/resources/distributed_port_group vsphere_distributed_port_group} Resource.
func NewDistributedPortGroup_Override(d DistributedPortGroup, scope constructs.Construct, id *string, config *DistributedPortGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.distributedPortGroup.DistributedPortGroup",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetActiveUplinks(val *[]*string) {
	if err := j.validateSetActiveUplinksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeUplinks",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetAllowForgedTransmits(val interface{}) {
	if err := j.validateSetAllowForgedTransmitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowForgedTransmits",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetAllowMacChanges(val interface{}) {
	if err := j.validateSetAllowMacChangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowMacChanges",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetAllowPromiscuous(val interface{}) {
	if err := j.validateSetAllowPromiscuousParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPromiscuous",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetAutoExpand(val interface{}) {
	if err := j.validateSetAutoExpandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoExpand",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetBlockAllPorts(val interface{}) {
	if err := j.validateSetBlockAllPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockAllPorts",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetBlockOverrideAllowed(val interface{}) {
	if err := j.validateSetBlockOverrideAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockOverrideAllowed",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetCheckBeacon(val interface{}) {
	if err := j.validateSetCheckBeaconParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkBeacon",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetCustomAttributes(val *map[string]*string) {
	if err := j.validateSetCustomAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAttributes",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetDirectpathGen2Allowed(val interface{}) {
	if err := j.validateSetDirectpathGen2AllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directpathGen2Allowed",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetDistributedVirtualSwitchUuid(val *string) {
	if err := j.validateSetDistributedVirtualSwitchUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distributedVirtualSwitchUuid",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetEgressShapingAverageBandwidth(val *float64) {
	if err := j.validateSetEgressShapingAverageBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressShapingAverageBandwidth",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetEgressShapingBurstSize(val *float64) {
	if err := j.validateSetEgressShapingBurstSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressShapingBurstSize",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetEgressShapingEnabled(val interface{}) {
	if err := j.validateSetEgressShapingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressShapingEnabled",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetEgressShapingPeakBandwidth(val *float64) {
	if err := j.validateSetEgressShapingPeakBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressShapingPeakBandwidth",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetFailback(val interface{}) {
	if err := j.validateSetFailbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failback",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetIngressShapingAverageBandwidth(val *float64) {
	if err := j.validateSetIngressShapingAverageBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressShapingAverageBandwidth",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetIngressShapingBurstSize(val *float64) {
	if err := j.validateSetIngressShapingBurstSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressShapingBurstSize",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetIngressShapingEnabled(val interface{}) {
	if err := j.validateSetIngressShapingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressShapingEnabled",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetIngressShapingPeakBandwidth(val *float64) {
	if err := j.validateSetIngressShapingPeakBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressShapingPeakBandwidth",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetLacpEnabled(val interface{}) {
	if err := j.validateSetLacpEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lacpEnabled",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetLacpMode(val *string) {
	if err := j.validateSetLacpModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lacpMode",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetLivePortMovingAllowed(val interface{}) {
	if err := j.validateSetLivePortMovingAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"livePortMovingAllowed",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetNetflowEnabled(val interface{}) {
	if err := j.validateSetNetflowEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netflowEnabled",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetNetflowOverrideAllowed(val interface{}) {
	if err := j.validateSetNetflowOverrideAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netflowOverrideAllowed",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetNetworkResourcePoolKey(val *string) {
	if err := j.validateSetNetworkResourcePoolKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkResourcePoolKey",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetNetworkResourcePoolOverrideAllowed(val interface{}) {
	if err := j.validateSetNetworkResourcePoolOverrideAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkResourcePoolOverrideAllowed",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetNotifySwitches(val interface{}) {
	if err := j.validateSetNotifySwitchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifySwitches",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetNumberOfPorts(val *float64) {
	if err := j.validateSetNumberOfPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfPorts",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetPortConfigResetAtDisconnect(val interface{}) {
	if err := j.validateSetPortConfigResetAtDisconnectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portConfigResetAtDisconnect",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetPortNameFormat(val *string) {
	if err := j.validateSetPortNameFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portNameFormat",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetPortPrivateSecondaryVlanId(val *float64) {
	if err := j.validateSetPortPrivateSecondaryVlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portPrivateSecondaryVlanId",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetSecurityPolicyOverrideAllowed(val interface{}) {
	if err := j.validateSetSecurityPolicyOverrideAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityPolicyOverrideAllowed",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetShapingOverrideAllowed(val interface{}) {
	if err := j.validateSetShapingOverrideAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shapingOverrideAllowed",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetStandbyUplinks(val *[]*string) {
	if err := j.validateSetStandbyUplinksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"standbyUplinks",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetTeamingPolicy(val *string) {
	if err := j.validateSetTeamingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamingPolicy",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetTrafficFilterOverrideAllowed(val interface{}) {
	if err := j.validateSetTrafficFilterOverrideAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficFilterOverrideAllowed",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetTxUplink(val interface{}) {
	if err := j.validateSetTxUplinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"txUplink",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetUplinkTeamingOverrideAllowed(val interface{}) {
	if err := j.validateSetUplinkTeamingOverrideAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uplinkTeamingOverrideAllowed",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetVlanId(val *float64) {
	if err := j.validateSetVlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vlanId",
		val,
	)
}

func (j *jsiiProxy_DistributedPortGroup)SetVlanOverrideAllowed(val interface{}) {
	if err := j.validateSetVlanOverrideAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vlanOverrideAllowed",
		val,
	)
}

// Generates CDKTF code for importing a DistributedPortGroup resource upon running "cdktf plan <stack-name>".
func DistributedPortGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDistributedPortGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.distributedPortGroup.DistributedPortGroup",
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
func DistributedPortGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDistributedPortGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.distributedPortGroup.DistributedPortGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DistributedPortGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDistributedPortGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.distributedPortGroup.DistributedPortGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DistributedPortGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDistributedPortGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.distributedPortGroup.DistributedPortGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DistributedPortGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vsphere.distributedPortGroup.DistributedPortGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DistributedPortGroup) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DistributedPortGroup) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DistributedPortGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DistributedPortGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DistributedPortGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DistributedPortGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DistributedPortGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DistributedPortGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DistributedPortGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DistributedPortGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DistributedPortGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DistributedPortGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedPortGroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DistributedPortGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DistributedPortGroup) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DistributedPortGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DistributedPortGroup) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DistributedPortGroup) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DistributedPortGroup) PutVlanRange(value interface{}) {
	if err := d.validatePutVlanRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVlanRange",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetActiveUplinks() {
	_jsii_.InvokeVoid(
		d,
		"resetActiveUplinks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetAllowForgedTransmits() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowForgedTransmits",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetAllowMacChanges() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowMacChanges",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetAllowPromiscuous() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowPromiscuous",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetAutoExpand() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoExpand",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetBlockAllPorts() {
	_jsii_.InvokeVoid(
		d,
		"resetBlockAllPorts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetBlockOverrideAllowed() {
	_jsii_.InvokeVoid(
		d,
		"resetBlockOverrideAllowed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetCheckBeacon() {
	_jsii_.InvokeVoid(
		d,
		"resetCheckBeacon",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetCustomAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetDirectpathGen2Allowed() {
	_jsii_.InvokeVoid(
		d,
		"resetDirectpathGen2Allowed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetEgressShapingAverageBandwidth() {
	_jsii_.InvokeVoid(
		d,
		"resetEgressShapingAverageBandwidth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetEgressShapingBurstSize() {
	_jsii_.InvokeVoid(
		d,
		"resetEgressShapingBurstSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetEgressShapingEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetEgressShapingEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetEgressShapingPeakBandwidth() {
	_jsii_.InvokeVoid(
		d,
		"resetEgressShapingPeakBandwidth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetFailback() {
	_jsii_.InvokeVoid(
		d,
		"resetFailback",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetIngressShapingAverageBandwidth() {
	_jsii_.InvokeVoid(
		d,
		"resetIngressShapingAverageBandwidth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetIngressShapingBurstSize() {
	_jsii_.InvokeVoid(
		d,
		"resetIngressShapingBurstSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetIngressShapingEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetIngressShapingEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetIngressShapingPeakBandwidth() {
	_jsii_.InvokeVoid(
		d,
		"resetIngressShapingPeakBandwidth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetLacpEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetLacpEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetLacpMode() {
	_jsii_.InvokeVoid(
		d,
		"resetLacpMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetLivePortMovingAllowed() {
	_jsii_.InvokeVoid(
		d,
		"resetLivePortMovingAllowed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetNetflowEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetNetflowEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetNetflowOverrideAllowed() {
	_jsii_.InvokeVoid(
		d,
		"resetNetflowOverrideAllowed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetNetworkResourcePoolKey() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkResourcePoolKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetNetworkResourcePoolOverrideAllowed() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkResourcePoolOverrideAllowed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetNotifySwitches() {
	_jsii_.InvokeVoid(
		d,
		"resetNotifySwitches",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetNumberOfPorts() {
	_jsii_.InvokeVoid(
		d,
		"resetNumberOfPorts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetPortConfigResetAtDisconnect() {
	_jsii_.InvokeVoid(
		d,
		"resetPortConfigResetAtDisconnect",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetPortNameFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetPortNameFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetPortPrivateSecondaryVlanId() {
	_jsii_.InvokeVoid(
		d,
		"resetPortPrivateSecondaryVlanId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetSecurityPolicyOverrideAllowed() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityPolicyOverrideAllowed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetShapingOverrideAllowed() {
	_jsii_.InvokeVoid(
		d,
		"resetShapingOverrideAllowed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetStandbyUplinks() {
	_jsii_.InvokeVoid(
		d,
		"resetStandbyUplinks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetTeamingPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetTeamingPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetTrafficFilterOverrideAllowed() {
	_jsii_.InvokeVoid(
		d,
		"resetTrafficFilterOverrideAllowed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetTxUplink() {
	_jsii_.InvokeVoid(
		d,
		"resetTxUplink",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetType() {
	_jsii_.InvokeVoid(
		d,
		"resetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetUplinkTeamingOverrideAllowed() {
	_jsii_.InvokeVoid(
		d,
		"resetUplinkTeamingOverrideAllowed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetVlanId() {
	_jsii_.InvokeVoid(
		d,
		"resetVlanId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetVlanOverrideAllowed() {
	_jsii_.InvokeVoid(
		d,
		"resetVlanOverrideAllowed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) ResetVlanRange() {
	_jsii_.InvokeVoid(
		d,
		"resetVlanRange",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedPortGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedPortGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedPortGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedPortGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedPortGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedPortGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

