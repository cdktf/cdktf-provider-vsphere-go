// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hostportgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v7/hostportgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/host_port_group vsphere_host_port_group}.
type HostPortGroup interface {
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
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CheckBeacon() interface{}
	SetCheckBeacon(val interface{})
	CheckBeaconInput() interface{}
	ComputedPolicy() cdktf.StringMap
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
	Key() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NotifySwitches() interface{}
	SetNotifySwitches(val interface{})
	NotifySwitchesInput() interface{}
	Ports() HostPortGroupPortsList
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
	VirtualSwitchName() *string
	SetVirtualSwitchName(val *string)
	VirtualSwitchNameInput() *string
	VlanId() *float64
	SetVlanId(val *float64)
	VlanIdInput() *float64
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
	ResetActiveNics()
	ResetAllowForgedTransmits()
	ResetAllowMacChanges()
	ResetAllowPromiscuous()
	ResetCheckBeacon()
	ResetFailback()
	ResetId()
	ResetNotifySwitches()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetShapingAverageBandwidth()
	ResetShapingBurstSize()
	ResetShapingEnabled()
	ResetShapingPeakBandwidth()
	ResetStandbyNics()
	ResetTeamingPolicy()
	ResetVlanId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for HostPortGroup
type jsiiProxy_HostPortGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_HostPortGroup) ActiveNics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"activeNics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) ActiveNicsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"activeNicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) AllowForgedTransmits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForgedTransmits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) AllowForgedTransmitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForgedTransmitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) AllowMacChanges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMacChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) AllowMacChangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMacChangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) AllowPromiscuous() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPromiscuous",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) AllowPromiscuousInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPromiscuousInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) CheckBeacon() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkBeacon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) CheckBeaconInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkBeaconInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) ComputedPolicy() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"computedPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) Failback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) FailbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) HostSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) HostSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) NotifySwitches() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifySwitches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) NotifySwitchesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifySwitchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) Ports() HostPortGroupPortsList {
	var returns HostPortGroupPortsList
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) ShapingAverageBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shapingAverageBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) ShapingAverageBandwidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shapingAverageBandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) ShapingBurstSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shapingBurstSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) ShapingBurstSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shapingBurstSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) ShapingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shapingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) ShapingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shapingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) ShapingPeakBandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shapingPeakBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) ShapingPeakBandwidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shapingPeakBandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) StandbyNics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"standbyNics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) StandbyNicsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"standbyNicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) TeamingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) TeamingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) VirtualSwitchName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualSwitchName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) VirtualSwitchNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualSwitchNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) VlanId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPortGroup) VlanIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlanIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/host_port_group vsphere_host_port_group} Resource.
func NewHostPortGroup(scope constructs.Construct, id *string, config *HostPortGroupConfig) HostPortGroup {
	_init_.Initialize()

	if err := validateNewHostPortGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_HostPortGroup{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.hostPortGroup.HostPortGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/host_port_group vsphere_host_port_group} Resource.
func NewHostPortGroup_Override(h HostPortGroup, scope constructs.Construct, id *string, config *HostPortGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.hostPortGroup.HostPortGroup",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HostPortGroup)SetActiveNics(val *[]*string) {
	if err := j.validateSetActiveNicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeNics",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetAllowForgedTransmits(val interface{}) {
	if err := j.validateSetAllowForgedTransmitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowForgedTransmits",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetAllowMacChanges(val interface{}) {
	if err := j.validateSetAllowMacChangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowMacChanges",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetAllowPromiscuous(val interface{}) {
	if err := j.validateSetAllowPromiscuousParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPromiscuous",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetCheckBeacon(val interface{}) {
	if err := j.validateSetCheckBeaconParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkBeacon",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetFailback(val interface{}) {
	if err := j.validateSetFailbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failback",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetHostSystemId(val *string) {
	if err := j.validateSetHostSystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostSystemId",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetNotifySwitches(val interface{}) {
	if err := j.validateSetNotifySwitchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifySwitches",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetShapingAverageBandwidth(val *float64) {
	if err := j.validateSetShapingAverageBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shapingAverageBandwidth",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetShapingBurstSize(val *float64) {
	if err := j.validateSetShapingBurstSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shapingBurstSize",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetShapingEnabled(val interface{}) {
	if err := j.validateSetShapingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shapingEnabled",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetShapingPeakBandwidth(val *float64) {
	if err := j.validateSetShapingPeakBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shapingPeakBandwidth",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetStandbyNics(val *[]*string) {
	if err := j.validateSetStandbyNicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"standbyNics",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetTeamingPolicy(val *string) {
	if err := j.validateSetTeamingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamingPolicy",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetVirtualSwitchName(val *string) {
	if err := j.validateSetVirtualSwitchNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualSwitchName",
		val,
	)
}

func (j *jsiiProxy_HostPortGroup)SetVlanId(val *float64) {
	if err := j.validateSetVlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vlanId",
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
func HostPortGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHostPortGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.hostPortGroup.HostPortGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HostPortGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHostPortGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.hostPortGroup.HostPortGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HostPortGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHostPortGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.hostPortGroup.HostPortGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HostPortGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vsphere.hostPortGroup.HostPortGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HostPortGroup) AddOverride(path *string, value interface{}) {
	if err := h.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HostPortGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HostPortGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HostPortGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HostPortGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HostPortGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HostPortGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HostPortGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HostPortGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HostPortGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HostPortGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HostPortGroup) OverrideLogicalId(newLogicalId *string) {
	if err := h.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HostPortGroup) ResetActiveNics() {
	_jsii_.InvokeVoid(
		h,
		"resetActiveNics",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostPortGroup) ResetAllowForgedTransmits() {
	_jsii_.InvokeVoid(
		h,
		"resetAllowForgedTransmits",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostPortGroup) ResetAllowMacChanges() {
	_jsii_.InvokeVoid(
		h,
		"resetAllowMacChanges",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostPortGroup) ResetAllowPromiscuous() {
	_jsii_.InvokeVoid(
		h,
		"resetAllowPromiscuous",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostPortGroup) ResetCheckBeacon() {
	_jsii_.InvokeVoid(
		h,
		"resetCheckBeacon",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostPortGroup) ResetFailback() {
	_jsii_.InvokeVoid(
		h,
		"resetFailback",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostPortGroup) ResetId() {
	_jsii_.InvokeVoid(
		h,
		"resetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostPortGroup) ResetNotifySwitches() {
	_jsii_.InvokeVoid(
		h,
		"resetNotifySwitches",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostPortGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostPortGroup) ResetShapingAverageBandwidth() {
	_jsii_.InvokeVoid(
		h,
		"resetShapingAverageBandwidth",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostPortGroup) ResetShapingBurstSize() {
	_jsii_.InvokeVoid(
		h,
		"resetShapingBurstSize",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostPortGroup) ResetShapingEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetShapingEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostPortGroup) ResetShapingPeakBandwidth() {
	_jsii_.InvokeVoid(
		h,
		"resetShapingPeakBandwidth",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostPortGroup) ResetStandbyNics() {
	_jsii_.InvokeVoid(
		h,
		"resetStandbyNics",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostPortGroup) ResetTeamingPolicy() {
	_jsii_.InvokeVoid(
		h,
		"resetTeamingPolicy",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostPortGroup) ResetVlanId() {
	_jsii_.InvokeVoid(
		h,
		"resetVlanId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostPortGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostPortGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostPortGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostPortGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

