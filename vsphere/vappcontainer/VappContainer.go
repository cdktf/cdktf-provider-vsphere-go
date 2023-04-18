package vappcontainer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v4/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v4/vappcontainer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/vapp_container vsphere_vapp_container}.
type VappContainer interface {
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
	ParentFolderId() *string
	SetParentFolderId(val *string)
	ParentFolderIdInput() *string
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
	ResetParentFolderId()
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

// The jsii proxy struct for VappContainer
type jsiiProxy_VappContainer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VappContainer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) CpuExpandable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuExpandable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) CpuExpandableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuExpandableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) CpuLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) CpuLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) CpuReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) CpuReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) CpuShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) CpuShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) CpuShares() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuShares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) CpuSharesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuSharesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) CustomAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) CustomAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) MemoryExpandable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryExpandable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) MemoryExpandableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryExpandableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) MemoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) MemoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) MemoryReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) MemoryReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) MemoryShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) MemoryShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) MemoryShares() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryShares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) MemorySharesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySharesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) ParentFolderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentFolderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) ParentFolderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentFolderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) ParentResourcePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentResourcePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) ParentResourcePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentResourcePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VappContainer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/vapp_container vsphere_vapp_container} Resource.
func NewVappContainer(scope constructs.Construct, id *string, config *VappContainerConfig) VappContainer {
	_init_.Initialize()

	if err := validateNewVappContainerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VappContainer{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.vappContainer.VappContainer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/vapp_container vsphere_vapp_container} Resource.
func NewVappContainer_Override(v VappContainer, scope constructs.Construct, id *string, config *VappContainerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.vappContainer.VappContainer",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VappContainer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetCpuExpandable(val interface{}) {
	if err := j.validateSetCpuExpandableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuExpandable",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetCpuLimit(val *float64) {
	if err := j.validateSetCpuLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuLimit",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetCpuReservation(val *float64) {
	if err := j.validateSetCpuReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuReservation",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetCpuShareLevel(val *string) {
	if err := j.validateSetCpuShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuShareLevel",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetCpuShares(val *float64) {
	if err := j.validateSetCpuSharesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuShares",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetCustomAttributes(val *map[string]*string) {
	if err := j.validateSetCustomAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAttributes",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetMemoryExpandable(val interface{}) {
	if err := j.validateSetMemoryExpandableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryExpandable",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetMemoryLimit(val *float64) {
	if err := j.validateSetMemoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryLimit",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetMemoryReservation(val *float64) {
	if err := j.validateSetMemoryReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryReservation",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetMemoryShareLevel(val *string) {
	if err := j.validateSetMemoryShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryShareLevel",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetMemoryShares(val *float64) {
	if err := j.validateSetMemorySharesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryShares",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetParentFolderId(val *string) {
	if err := j.validateSetParentFolderIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentFolderId",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetParentResourcePoolId(val *string) {
	if err := j.validateSetParentResourcePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentResourcePoolId",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VappContainer)SetTags(val *[]*string) {
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
func VappContainer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVappContainer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.vappContainer.VappContainer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VappContainer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVappContainer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.vappContainer.VappContainer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VappContainer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVappContainer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.vappContainer.VappContainer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VappContainer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vsphere.vappContainer.VappContainer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VappContainer) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VappContainer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VappContainer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VappContainer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VappContainer) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VappContainer) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VappContainer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VappContainer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VappContainer) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VappContainer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VappContainer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VappContainer) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VappContainer) ResetCpuExpandable() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuExpandable",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappContainer) ResetCpuLimit() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuLimit",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappContainer) ResetCpuReservation() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuReservation",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappContainer) ResetCpuShareLevel() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuShareLevel",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappContainer) ResetCpuShares() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuShares",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappContainer) ResetCustomAttributes() {
	_jsii_.InvokeVoid(
		v,
		"resetCustomAttributes",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappContainer) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappContainer) ResetMemoryExpandable() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryExpandable",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappContainer) ResetMemoryLimit() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryLimit",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappContainer) ResetMemoryReservation() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryReservation",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappContainer) ResetMemoryShareLevel() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryShareLevel",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappContainer) ResetMemoryShares() {
	_jsii_.InvokeVoid(
		v,
		"resetMemoryShares",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappContainer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappContainer) ResetParentFolderId() {
	_jsii_.InvokeVoid(
		v,
		"resetParentFolderId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappContainer) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VappContainer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappContainer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappContainer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VappContainer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

