// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package distributedvirtualswitchpvlanmapping

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v11/distributedvirtualswitchpvlanmapping/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.0/docs/resources/distributed_virtual_switch_pvlan_mapping vsphere_distributed_virtual_switch_pvlan_mapping}.
type DistributedVirtualSwitchPvlanMappingA interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	DistributedVirtualSwitchId() *string
	SetDistributedVirtualSwitchId(val *string)
	DistributedVirtualSwitchIdInput() *string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PrimaryVlanId() *float64
	SetPrimaryVlanId(val *float64)
	PrimaryVlanIdInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PvlanType() *string
	SetPvlanType(val *string)
	PvlanTypeInput() *string
	// Experimental.
	RawOverrides() interface{}
	SecondaryVlanId() *float64
	SetSecondaryVlanId(val *float64)
	SecondaryVlanIdInput() *float64
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
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for DistributedVirtualSwitchPvlanMappingA
type jsiiProxy_DistributedVirtualSwitchPvlanMappingA struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) DistributedVirtualSwitchId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributedVirtualSwitchId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) DistributedVirtualSwitchIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributedVirtualSwitchIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) PrimaryVlanId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"primaryVlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) PrimaryVlanIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"primaryVlanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) PvlanType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pvlanType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) PvlanTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pvlanTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) SecondaryVlanId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondaryVlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) SecondaryVlanIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondaryVlanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.0/docs/resources/distributed_virtual_switch_pvlan_mapping vsphere_distributed_virtual_switch_pvlan_mapping} Resource.
func NewDistributedVirtualSwitchPvlanMappingA(scope constructs.Construct, id *string, config *DistributedVirtualSwitchPvlanMappingAConfig) DistributedVirtualSwitchPvlanMappingA {
	_init_.Initialize()

	if err := validateNewDistributedVirtualSwitchPvlanMappingAParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DistributedVirtualSwitchPvlanMappingA{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.distributedVirtualSwitchPvlanMapping.DistributedVirtualSwitchPvlanMappingA",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.0/docs/resources/distributed_virtual_switch_pvlan_mapping vsphere_distributed_virtual_switch_pvlan_mapping} Resource.
func NewDistributedVirtualSwitchPvlanMappingA_Override(d DistributedVirtualSwitchPvlanMappingA, scope constructs.Construct, id *string, config *DistributedVirtualSwitchPvlanMappingAConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.distributedVirtualSwitchPvlanMapping.DistributedVirtualSwitchPvlanMappingA",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA)SetDistributedVirtualSwitchId(val *string) {
	if err := j.validateSetDistributedVirtualSwitchIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distributedVirtualSwitchId",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA)SetPrimaryVlanId(val *float64) {
	if err := j.validateSetPrimaryVlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryVlanId",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA)SetPvlanType(val *string) {
	if err := j.validateSetPvlanTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pvlanType",
		val,
	)
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingA)SetSecondaryVlanId(val *float64) {
	if err := j.validateSetSecondaryVlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryVlanId",
		val,
	)
}

// Generates CDKTF code for importing a DistributedVirtualSwitchPvlanMappingA resource upon running "cdktf plan <stack-name>".
func DistributedVirtualSwitchPvlanMappingA_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDistributedVirtualSwitchPvlanMappingA_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.distributedVirtualSwitchPvlanMapping.DistributedVirtualSwitchPvlanMappingA",
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
func DistributedVirtualSwitchPvlanMappingA_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDistributedVirtualSwitchPvlanMappingA_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.distributedVirtualSwitchPvlanMapping.DistributedVirtualSwitchPvlanMappingA",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DistributedVirtualSwitchPvlanMappingA_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDistributedVirtualSwitchPvlanMappingA_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.distributedVirtualSwitchPvlanMapping.DistributedVirtualSwitchPvlanMappingA",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DistributedVirtualSwitchPvlanMappingA_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDistributedVirtualSwitchPvlanMappingA_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.distributedVirtualSwitchPvlanMapping.DistributedVirtualSwitchPvlanMappingA",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DistributedVirtualSwitchPvlanMappingA_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vsphere.distributedVirtualSwitchPvlanMapping.DistributedVirtualSwitchPvlanMappingA",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingA) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

