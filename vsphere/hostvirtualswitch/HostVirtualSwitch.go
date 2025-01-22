// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hostvirtualswitch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v9/hostvirtualswitch/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/resources/host_virtual_switch vsphere_host_virtual_switch}.
type HostVirtualSwitch interface {
	cdktf.TerraformResource
	ActiveNics() *[]*string
	SetActiveNics(val *[]*string)
	ActiveNicsInput() *[]*string
	AllowForgedTransmits() interface{}
	SetAllowForgedTransmits(val interface{})
	AllowForgedTransmitsInput() interface{}
	AllowMacChanges() interface{}
	SetAllowMacChanges(val interface{})
	AllowMacChangesInput() interface{}
	AllowPromiscuous() interface{}
	SetAllowPromiscuous(val interface{})
	AllowPromiscuousInput() interface{}
	BeaconInterval() *float64
	SetBeaconInterval(val *float64)
	BeaconIntervalInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CheckBeacon() interface{}
	SetCheckBeacon(val interface{})
	CheckBeaconInput() interface{}
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	HostSystemId() *string
	SetHostSystemId(val *string)
	HostSystemIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
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
	Mtu() *float64
	SetMtu(val *float64)
	MtuInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkAdapters() *[]*string
	SetNetworkAdapters(val *[]*string)
	NetworkAdaptersInput() *[]*string
	// The tree node.
	Node() constructs.Node
	NotifySwitches() interface{}
	SetNotifySwitches(val interface{})
	NotifySwitchesInput() interface{}
	NumberOfPorts() *float64
	SetNumberOfPorts(val *float64)
	NumberOfPortsInput() *float64
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
	ShapingAverageBandwidth() *float64
	SetShapingAverageBandwidth(val *float64)
	ShapingAverageBandwidthInput() *float64
	ShapingBurstSize() *float64
	SetShapingBurstSize(val *float64)
	ShapingBurstSizeInput() *float64
	ShapingEnabled() interface{}
	SetShapingEnabled(val interface{})
	ShapingEnabledInput() interface{}
	ShapingPeakBandwidth() *float64
	SetShapingPeakBandwidth(val *float64)
	ShapingPeakBandwidthInput() *float64
	StandbyNics() *[]*string
	SetStandbyNics(val *[]*string)
	StandbyNicsInput() *[]*string
	TeamingPolicy() *string
	SetTeamingPolicy(val *string)
	TeamingPolicyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetAllowForgedTransmits()
	ResetAllowMacChanges()
	ResetAllowPromiscuous()
	ResetBeaconInterval()
	ResetCheckBeacon()
	ResetFailback()
	ResetId()
	ResetLinkDiscoveryOperation()
	ResetLinkDiscoveryProtocol()
	ResetMtu()
	ResetNotifySwitches()
	ResetNumberOfPorts()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetShapingAverageBandwidth()
	ResetShapingBurstSize()
	ResetShapingEnabled()
	ResetShapingPeakBandwidth()
	ResetStandbyNics()
	ResetTeamingPolicy()
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

// The jsii proxy struct for HostVirtualSwitch
type jsiiProxy_HostVirtualSwitch struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_HostVirtualSwitch) ActiveNics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"activeNics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) ActiveNicsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"activeNicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) AllowForgedTransmits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForgedTransmits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) AllowForgedTransmitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForgedTransmitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) AllowMacChanges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMacChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) AllowMacChangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMacChangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) AllowPromiscuous() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPromiscuous",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) AllowPromiscuousInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPromiscuousInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) BeaconInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"beaconInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) BeaconIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"beaconIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) CheckBeacon() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkBeacon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) CheckBeaconInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkBeaconInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) Failback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) FailbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) HostSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) HostSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) LinkDiscoveryOperation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkDiscoveryOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) LinkDiscoveryOperationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkDiscoveryOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) LinkDiscoveryProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkDiscoveryProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) LinkDiscoveryProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkDiscoveryProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) Mtu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mtu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) MtuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mtuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) NetworkAdapters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkAdapters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) NetworkAdaptersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkAdaptersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) NotifySwitches() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifySwitches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) NotifySwitchesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifySwitchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) NumberOfPorts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) NumberOfPortsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) ShapingAverageBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shapingAverageBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) ShapingAverageBandwidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shapingAverageBandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) ShapingBurstSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shapingBurstSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) ShapingBurstSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shapingBurstSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) ShapingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shapingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) ShapingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shapingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) ShapingPeakBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shapingPeakBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) ShapingPeakBandwidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shapingPeakBandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) StandbyNics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"standbyNics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) StandbyNicsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"standbyNicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) TeamingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) TeamingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostVirtualSwitch) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/resources/host_virtual_switch vsphere_host_virtual_switch} Resource.
func NewHostVirtualSwitch(scope constructs.Construct, id *string, config *HostVirtualSwitchConfig) HostVirtualSwitch {
	_init_.Initialize()

	if err := validateNewHostVirtualSwitchParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_HostVirtualSwitch{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.hostVirtualSwitch.HostVirtualSwitch",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/resources/host_virtual_switch vsphere_host_virtual_switch} Resource.
func NewHostVirtualSwitch_Override(h HostVirtualSwitch, scope constructs.Construct, id *string, config *HostVirtualSwitchConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.hostVirtualSwitch.HostVirtualSwitch",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetActiveNics(val *[]*string) {
	if err := j.validateSetActiveNicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeNics",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetAllowForgedTransmits(val interface{}) {
	if err := j.validateSetAllowForgedTransmitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowForgedTransmits",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetAllowMacChanges(val interface{}) {
	if err := j.validateSetAllowMacChangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowMacChanges",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetAllowPromiscuous(val interface{}) {
	if err := j.validateSetAllowPromiscuousParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPromiscuous",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetBeaconInterval(val *float64) {
	if err := j.validateSetBeaconIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"beaconInterval",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetCheckBeacon(val interface{}) {
	if err := j.validateSetCheckBeaconParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkBeacon",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetFailback(val interface{}) {
	if err := j.validateSetFailbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failback",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetHostSystemId(val *string) {
	if err := j.validateSetHostSystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostSystemId",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetLinkDiscoveryOperation(val *string) {
	if err := j.validateSetLinkDiscoveryOperationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkDiscoveryOperation",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetLinkDiscoveryProtocol(val *string) {
	if err := j.validateSetLinkDiscoveryProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkDiscoveryProtocol",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetMtu(val *float64) {
	if err := j.validateSetMtuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mtu",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetNetworkAdapters(val *[]*string) {
	if err := j.validateSetNetworkAdaptersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkAdapters",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetNotifySwitches(val interface{}) {
	if err := j.validateSetNotifySwitchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifySwitches",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetNumberOfPorts(val *float64) {
	if err := j.validateSetNumberOfPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfPorts",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetShapingAverageBandwidth(val *float64) {
	if err := j.validateSetShapingAverageBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shapingAverageBandwidth",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetShapingBurstSize(val *float64) {
	if err := j.validateSetShapingBurstSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shapingBurstSize",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetShapingEnabled(val interface{}) {
	if err := j.validateSetShapingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shapingEnabled",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetShapingPeakBandwidth(val *float64) {
	if err := j.validateSetShapingPeakBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shapingPeakBandwidth",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetStandbyNics(val *[]*string) {
	if err := j.validateSetStandbyNicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"standbyNics",
		val,
	)
}

func (j *jsiiProxy_HostVirtualSwitch)SetTeamingPolicy(val *string) {
	if err := j.validateSetTeamingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamingPolicy",
		val,
	)
}

// Generates CDKTF code for importing a HostVirtualSwitch resource upon running "cdktf plan <stack-name>".
func HostVirtualSwitch_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateHostVirtualSwitch_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.hostVirtualSwitch.HostVirtualSwitch",
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
func HostVirtualSwitch_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHostVirtualSwitch_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.hostVirtualSwitch.HostVirtualSwitch",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HostVirtualSwitch_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHostVirtualSwitch_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.hostVirtualSwitch.HostVirtualSwitch",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HostVirtualSwitch_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHostVirtualSwitch_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.hostVirtualSwitch.HostVirtualSwitch",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HostVirtualSwitch_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vsphere.hostVirtualSwitch.HostVirtualSwitch",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HostVirtualSwitch) AddMoveTarget(moveTarget *string) {
	if err := h.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (h *jsiiProxy_HostVirtualSwitch) AddOverride(path *string, value interface{}) {
	if err := h.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HostVirtualSwitch) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostVirtualSwitch) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostVirtualSwitch) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostVirtualSwitch) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostVirtualSwitch) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostVirtualSwitch) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostVirtualSwitch) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostVirtualSwitch) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostVirtualSwitch) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostVirtualSwitch) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostVirtualSwitch) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := h.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (h *jsiiProxy_HostVirtualSwitch) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostVirtualSwitch) MoveFromId(id *string) {
	if err := h.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveFromId",
		[]interface{}{id},
	)
}

func (h *jsiiProxy_HostVirtualSwitch) MoveTo(moveTarget *string, index interface{}) {
	if err := h.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (h *jsiiProxy_HostVirtualSwitch) MoveToId(id *string) {
	if err := h.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveToId",
		[]interface{}{id},
	)
}

func (h *jsiiProxy_HostVirtualSwitch) OverrideLogicalId(newLogicalId *string) {
	if err := h.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetAllowForgedTransmits() {
	_jsii_.InvokeVoid(
		h,
		"resetAllowForgedTransmits",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetAllowMacChanges() {
	_jsii_.InvokeVoid(
		h,
		"resetAllowMacChanges",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetAllowPromiscuous() {
	_jsii_.InvokeVoid(
		h,
		"resetAllowPromiscuous",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetBeaconInterval() {
	_jsii_.InvokeVoid(
		h,
		"resetBeaconInterval",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetCheckBeacon() {
	_jsii_.InvokeVoid(
		h,
		"resetCheckBeacon",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetFailback() {
	_jsii_.InvokeVoid(
		h,
		"resetFailback",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetId() {
	_jsii_.InvokeVoid(
		h,
		"resetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetLinkDiscoveryOperation() {
	_jsii_.InvokeVoid(
		h,
		"resetLinkDiscoveryOperation",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetLinkDiscoveryProtocol() {
	_jsii_.InvokeVoid(
		h,
		"resetLinkDiscoveryProtocol",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetMtu() {
	_jsii_.InvokeVoid(
		h,
		"resetMtu",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetNotifySwitches() {
	_jsii_.InvokeVoid(
		h,
		"resetNotifySwitches",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetNumberOfPorts() {
	_jsii_.InvokeVoid(
		h,
		"resetNumberOfPorts",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetShapingAverageBandwidth() {
	_jsii_.InvokeVoid(
		h,
		"resetShapingAverageBandwidth",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetShapingBurstSize() {
	_jsii_.InvokeVoid(
		h,
		"resetShapingBurstSize",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetShapingEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetShapingEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetShapingPeakBandwidth() {
	_jsii_.InvokeVoid(
		h,
		"resetShapingPeakBandwidth",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetStandbyNics() {
	_jsii_.InvokeVoid(
		h,
		"resetStandbyNics",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) ResetTeamingPolicy() {
	_jsii_.InvokeVoid(
		h,
		"resetTeamingPolicy",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostVirtualSwitch) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostVirtualSwitch) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostVirtualSwitch) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostVirtualSwitch) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostVirtualSwitch) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostVirtualSwitch) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

