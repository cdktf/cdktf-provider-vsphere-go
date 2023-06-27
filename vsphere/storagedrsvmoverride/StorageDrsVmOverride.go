package storagedrsvmoverride

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v6/storagedrsvmoverride/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/storage_drs_vm_override vsphere_storage_drs_vm_override}.
type StorageDrsVmOverride interface {
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
	DatastoreClusterId() *string
	SetDatastoreClusterId(val *string)
	DatastoreClusterIdInput() *string
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
	// The tree node.
	Node() constructs.Node
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
	SdrsAutomationLevel() *string
	SetSdrsAutomationLevel(val *string)
	SdrsAutomationLevelInput() *string
	SdrsEnabled() *string
	SetSdrsEnabled(val *string)
	SdrsEnabledInput() *string
	SdrsIntraVmAffinity() *string
	SetSdrsIntraVmAffinity(val *string)
	SdrsIntraVmAffinityInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VirtualMachineId() *string
	SetVirtualMachineId(val *string)
	VirtualMachineIdInput() *string
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
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSdrsAutomationLevel()
	ResetSdrsEnabled()
	ResetSdrsIntraVmAffinity()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for StorageDrsVmOverride
type jsiiProxy_StorageDrsVmOverride struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StorageDrsVmOverride) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) DatastoreClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) DatastoreClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) SdrsAutomationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsAutomationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) SdrsAutomationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsAutomationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) SdrsEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) SdrsEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) SdrsIntraVmAffinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsIntraVmAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) SdrsIntraVmAffinityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsIntraVmAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) VirtualMachineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageDrsVmOverride) VirtualMachineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/storage_drs_vm_override vsphere_storage_drs_vm_override} Resource.
func NewStorageDrsVmOverride(scope constructs.Construct, id *string, config *StorageDrsVmOverrideConfig) StorageDrsVmOverride {
	_init_.Initialize()

	if err := validateNewStorageDrsVmOverrideParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageDrsVmOverride{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.storageDrsVmOverride.StorageDrsVmOverride",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/storage_drs_vm_override vsphere_storage_drs_vm_override} Resource.
func NewStorageDrsVmOverride_Override(s StorageDrsVmOverride, scope constructs.Construct, id *string, config *StorageDrsVmOverrideConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.storageDrsVmOverride.StorageDrsVmOverride",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StorageDrsVmOverride)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StorageDrsVmOverride)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StorageDrsVmOverride)SetDatastoreClusterId(val *string) {
	if err := j.validateSetDatastoreClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datastoreClusterId",
		val,
	)
}

func (j *jsiiProxy_StorageDrsVmOverride)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StorageDrsVmOverride)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StorageDrsVmOverride)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StorageDrsVmOverride)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StorageDrsVmOverride)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StorageDrsVmOverride)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StorageDrsVmOverride)SetSdrsAutomationLevel(val *string) {
	if err := j.validateSetSdrsAutomationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsAutomationLevel",
		val,
	)
}

func (j *jsiiProxy_StorageDrsVmOverride)SetSdrsEnabled(val *string) {
	if err := j.validateSetSdrsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageDrsVmOverride)SetSdrsIntraVmAffinity(val *string) {
	if err := j.validateSetSdrsIntraVmAffinityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsIntraVmAffinity",
		val,
	)
}

func (j *jsiiProxy_StorageDrsVmOverride)SetVirtualMachineId(val *string) {
	if err := j.validateSetVirtualMachineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualMachineId",
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
func StorageDrsVmOverride_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageDrsVmOverride_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.storageDrsVmOverride.StorageDrsVmOverride",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StorageDrsVmOverride_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageDrsVmOverride_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.storageDrsVmOverride.StorageDrsVmOverride",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StorageDrsVmOverride_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageDrsVmOverride_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.storageDrsVmOverride.StorageDrsVmOverride",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StorageDrsVmOverride_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vsphere.storageDrsVmOverride.StorageDrsVmOverride",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StorageDrsVmOverride) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StorageDrsVmOverride) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageDrsVmOverride) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageDrsVmOverride) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageDrsVmOverride) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageDrsVmOverride) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageDrsVmOverride) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageDrsVmOverride) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageDrsVmOverride) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageDrsVmOverride) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageDrsVmOverride) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageDrsVmOverride) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StorageDrsVmOverride) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageDrsVmOverride) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageDrsVmOverride) ResetSdrsAutomationLevel() {
	_jsii_.InvokeVoid(
		s,
		"resetSdrsAutomationLevel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageDrsVmOverride) ResetSdrsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSdrsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageDrsVmOverride) ResetSdrsIntraVmAffinity() {
	_jsii_.InvokeVoid(
		s,
		"resetSdrsIntraVmAffinity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageDrsVmOverride) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageDrsVmOverride) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageDrsVmOverride) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageDrsVmOverride) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

