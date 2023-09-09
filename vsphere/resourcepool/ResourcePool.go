// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v7/resourcepool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.3/docs/resources/resource_pool vsphere_resource_pool}.
type ResourcePool interface {
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
	CpuExpandable() interface{}
	SetCpuExpandable(val interface{})
	CpuExpandableInput() interface{}
	CpuLimit() *float64
	SetCpuLimit(val *float64)
	CpuLimitInput() *float64
	CpuReservation() *float64
	SetCpuReservation(val *float64)
	CpuReservationInput() *float64
	CpuShareLevel() *string
	SetCpuShareLevel(val *string)
	CpuShareLevelInput() *string
	CpuShares() *float64
	SetCpuShares(val *float64)
	CpuSharesInput() *float64
	CustomAttributes() *map[string]*string
	SetCustomAttributes(val *map[string]*string)
	CustomAttributesInput() *map[string]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	MemoryExpandable() interface{}
	SetMemoryExpandable(val interface{})
	MemoryExpandableInput() interface{}
	MemoryLimit() *float64
	SetMemoryLimit(val *float64)
	MemoryLimitInput() *float64
	MemoryReservation() *float64
	SetMemoryReservation(val *float64)
	MemoryReservationInput() *float64
	MemoryShareLevel() *string
	SetMemoryShareLevel(val *string)
	MemoryShareLevelInput() *string
	MemoryShares() *float64
	SetMemoryShares(val *float64)
	MemorySharesInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ParentResourcePoolId() *string
	SetParentResourcePoolId(val *string)
	ParentResourcePoolIdInput() *string
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
	ScaleDescendantsShares() *string
	SetScaleDescendantsShares(val *string)
	ScaleDescendantsSharesInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetCpuExpandable()
	ResetCpuLimit()
	ResetCpuReservation()
	ResetCpuShareLevel()
	ResetCpuShares()
	ResetCustomAttributes()
	ResetId()
	ResetMemoryExpandable()
	ResetMemoryLimit()
	ResetMemoryReservation()
	ResetMemoryShareLevel()
	ResetMemoryShares()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetScaleDescendantsShares()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ResourcePool
type jsiiProxy_ResourcePool struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ResourcePool) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) CpuExpandable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuExpandable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) CpuExpandableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuExpandableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) CpuLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) CpuLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) CpuReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) CpuReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) CpuShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) CpuShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) CpuShares() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuShares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) CpuSharesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuSharesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) CustomAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) CustomAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) MemoryExpandable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryExpandable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) MemoryExpandableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryExpandableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) MemoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) MemoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) MemoryReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) MemoryReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) MemoryShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) MemoryShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) MemoryShares() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryShares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) MemorySharesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySharesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) ParentResourcePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentResourcePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) ParentResourcePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentResourcePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) ScaleDescendantsShares() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDescendantsShares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) ScaleDescendantsSharesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDescendantsSharesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.3/docs/resources/resource_pool vsphere_resource_pool} Resource.
func NewResourcePool(scope constructs.Construct, id *string, config *ResourcePoolConfig) ResourcePool {
	_init_.Initialize()

	if err := validateNewResourcePoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResourcePool{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.resourcePool.ResourcePool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.3/docs/resources/resource_pool vsphere_resource_pool} Resource.
func NewResourcePool_Override(r ResourcePool, scope constructs.Construct, id *string, config *ResourcePoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.resourcePool.ResourcePool",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_ResourcePool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetCpuExpandable(val interface{}) {
	if err := j.validateSetCpuExpandableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuExpandable",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetCpuLimit(val *float64) {
	if err := j.validateSetCpuLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuLimit",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetCpuReservation(val *float64) {
	if err := j.validateSetCpuReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuReservation",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetCpuShareLevel(val *string) {
	if err := j.validateSetCpuShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuShareLevel",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetCpuShares(val *float64) {
	if err := j.validateSetCpuSharesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuShares",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetCustomAttributes(val *map[string]*string) {
	if err := j.validateSetCustomAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAttributes",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetMemoryExpandable(val interface{}) {
	if err := j.validateSetMemoryExpandableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryExpandable",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetMemoryLimit(val *float64) {
	if err := j.validateSetMemoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryLimit",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetMemoryReservation(val *float64) {
	if err := j.validateSetMemoryReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryReservation",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetMemoryShareLevel(val *string) {
	if err := j.validateSetMemoryShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryShareLevel",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetMemoryShares(val *float64) {
	if err := j.validateSetMemorySharesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryShares",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetParentResourcePoolId(val *string) {
	if err := j.validateSetParentResourcePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentResourcePoolId",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetScaleDescendantsShares(val *string) {
	if err := j.validateSetScaleDescendantsSharesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleDescendantsShares",
		val,
	)
}

func (j *jsiiProxy_ResourcePool)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
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
func ResourcePool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourcePool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.resourcePool.ResourcePool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ResourcePool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourcePool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.resourcePool.ResourcePool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ResourcePool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourcePool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.resourcePool.ResourcePool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ResourcePool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vsphere.resourcePool.ResourcePool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_ResourcePool) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_ResourcePool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePool) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePool) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePool) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePool) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePool) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_ResourcePool) ResetCpuExpandable() {
	_jsii_.InvokeVoid(
		r,
		"resetCpuExpandable",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePool) ResetCpuLimit() {
	_jsii_.InvokeVoid(
		r,
		"resetCpuLimit",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePool) ResetCpuReservation() {
	_jsii_.InvokeVoid(
		r,
		"resetCpuReservation",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePool) ResetCpuShareLevel() {
	_jsii_.InvokeVoid(
		r,
		"resetCpuShareLevel",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePool) ResetCpuShares() {
	_jsii_.InvokeVoid(
		r,
		"resetCpuShares",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePool) ResetCustomAttributes() {
	_jsii_.InvokeVoid(
		r,
		"resetCustomAttributes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePool) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePool) ResetMemoryExpandable() {
	_jsii_.InvokeVoid(
		r,
		"resetMemoryExpandable",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePool) ResetMemoryLimit() {
	_jsii_.InvokeVoid(
		r,
		"resetMemoryLimit",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePool) ResetMemoryReservation() {
	_jsii_.InvokeVoid(
		r,
		"resetMemoryReservation",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePool) ResetMemoryShareLevel() {
	_jsii_.InvokeVoid(
		r,
		"resetMemoryShareLevel",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePool) ResetMemoryShares() {
	_jsii_.InvokeVoid(
		r,
		"resetMemoryShares",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePool) ResetScaleDescendantsShares() {
	_jsii_.InvokeVoid(
		r,
		"resetScaleDescendantsShares",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePool) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

