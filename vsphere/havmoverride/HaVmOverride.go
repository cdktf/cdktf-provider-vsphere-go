package havmoverride

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v4/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v4/havmoverride/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/ha_vm_override vsphere_ha_vm_override}.
type HaVmOverride interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComputeClusterId() *string
	SetComputeClusterId(val *string)
	ComputeClusterIdInput() *string
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HaDatastoreApdRecoveryAction() *string
	SetHaDatastoreApdRecoveryAction(val *string)
	HaDatastoreApdRecoveryActionInput() *string
	HaDatastoreApdResponse() *string
	SetHaDatastoreApdResponse(val *string)
	HaDatastoreApdResponseDelay() *float64
	SetHaDatastoreApdResponseDelay(val *float64)
	HaDatastoreApdResponseDelayInput() *float64
	HaDatastoreApdResponseInput() *string
	HaDatastorePdlResponse() *string
	SetHaDatastorePdlResponse(val *string)
	HaDatastorePdlResponseInput() *string
	HaHostIsolationResponse() *string
	SetHaHostIsolationResponse(val *string)
	HaHostIsolationResponseInput() *string
	HaVmFailureInterval() *float64
	SetHaVmFailureInterval(val *float64)
	HaVmFailureIntervalInput() *float64
	HaVmMaximumFailureWindow() *float64
	SetHaVmMaximumFailureWindow(val *float64)
	HaVmMaximumFailureWindowInput() *float64
	HaVmMaximumResets() *float64
	SetHaVmMaximumResets(val *float64)
	HaVmMaximumResetsInput() *float64
	HaVmMinimumUptime() *float64
	SetHaVmMinimumUptime(val *float64)
	HaVmMinimumUptimeInput() *float64
	HaVmMonitoring() *string
	SetHaVmMonitoring(val *string)
	HaVmMonitoringInput() *string
	HaVmMonitoringUseClusterDefaults() interface{}
	SetHaVmMonitoringUseClusterDefaults(val interface{})
	HaVmMonitoringUseClusterDefaultsInput() interface{}
	HaVmRestartPriority() *string
	SetHaVmRestartPriority(val *string)
	HaVmRestartPriorityInput() *string
	HaVmRestartTimeout() *float64
	SetHaVmRestartTimeout(val *float64)
	HaVmRestartTimeoutInput() *float64
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
	ResetHaDatastoreApdRecoveryAction()
	ResetHaDatastoreApdResponse()
	ResetHaDatastoreApdResponseDelay()
	ResetHaDatastorePdlResponse()
	ResetHaHostIsolationResponse()
	ResetHaVmFailureInterval()
	ResetHaVmMaximumFailureWindow()
	ResetHaVmMaximumResets()
	ResetHaVmMinimumUptime()
	ResetHaVmMonitoring()
	ResetHaVmMonitoringUseClusterDefaults()
	ResetHaVmRestartPriority()
	ResetHaVmRestartTimeout()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for HaVmOverride
type jsiiProxy_HaVmOverride struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_HaVmOverride) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) ComputeClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) ComputeClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaDatastoreApdRecoveryAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haDatastoreApdRecoveryAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaDatastoreApdRecoveryActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haDatastoreApdRecoveryActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaDatastoreApdResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haDatastoreApdResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaDatastoreApdResponseDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haDatastoreApdResponseDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaDatastoreApdResponseDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haDatastoreApdResponseDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaDatastoreApdResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haDatastoreApdResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaDatastorePdlResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haDatastorePdlResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaDatastorePdlResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haDatastorePdlResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaHostIsolationResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haHostIsolationResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaHostIsolationResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haHostIsolationResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaVmFailureInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmFailureInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaVmFailureIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmFailureIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaVmMaximumFailureWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmMaximumFailureWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaVmMaximumFailureWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmMaximumFailureWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaVmMaximumResets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmMaximumResets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaVmMaximumResetsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmMaximumResetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaVmMinimumUptime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmMinimumUptime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaVmMinimumUptimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmMinimumUptimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaVmMonitoring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haVmMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaVmMonitoringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haVmMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaVmMonitoringUseClusterDefaults() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"haVmMonitoringUseClusterDefaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaVmMonitoringUseClusterDefaultsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"haVmMonitoringUseClusterDefaultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaVmRestartPriority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haVmRestartPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaVmRestartPriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haVmRestartPriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaVmRestartTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmRestartTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) HaVmRestartTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmRestartTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) VirtualMachineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HaVmOverride) VirtualMachineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/ha_vm_override vsphere_ha_vm_override} Resource.
func NewHaVmOverride(scope constructs.Construct, id *string, config *HaVmOverrideConfig) HaVmOverride {
	_init_.Initialize()

	if err := validateNewHaVmOverrideParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_HaVmOverride{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.haVmOverride.HaVmOverride",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/ha_vm_override vsphere_ha_vm_override} Resource.
func NewHaVmOverride_Override(h HaVmOverride, scope constructs.Construct, id *string, config *HaVmOverrideConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.haVmOverride.HaVmOverride",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HaVmOverride)SetComputeClusterId(val *string) {
	if err := j.validateSetComputeClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeClusterId",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetHaDatastoreApdRecoveryAction(val *string) {
	if err := j.validateSetHaDatastoreApdRecoveryActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haDatastoreApdRecoveryAction",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetHaDatastoreApdResponse(val *string) {
	if err := j.validateSetHaDatastoreApdResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haDatastoreApdResponse",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetHaDatastoreApdResponseDelay(val *float64) {
	if err := j.validateSetHaDatastoreApdResponseDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haDatastoreApdResponseDelay",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetHaDatastorePdlResponse(val *string) {
	if err := j.validateSetHaDatastorePdlResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haDatastorePdlResponse",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetHaHostIsolationResponse(val *string) {
	if err := j.validateSetHaHostIsolationResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haHostIsolationResponse",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetHaVmFailureInterval(val *float64) {
	if err := j.validateSetHaVmFailureIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haVmFailureInterval",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetHaVmMaximumFailureWindow(val *float64) {
	if err := j.validateSetHaVmMaximumFailureWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haVmMaximumFailureWindow",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetHaVmMaximumResets(val *float64) {
	if err := j.validateSetHaVmMaximumResetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haVmMaximumResets",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetHaVmMinimumUptime(val *float64) {
	if err := j.validateSetHaVmMinimumUptimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haVmMinimumUptime",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetHaVmMonitoring(val *string) {
	if err := j.validateSetHaVmMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haVmMonitoring",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetHaVmMonitoringUseClusterDefaults(val interface{}) {
	if err := j.validateSetHaVmMonitoringUseClusterDefaultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haVmMonitoringUseClusterDefaults",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetHaVmRestartPriority(val *string) {
	if err := j.validateSetHaVmRestartPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haVmRestartPriority",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetHaVmRestartTimeout(val *float64) {
	if err := j.validateSetHaVmRestartTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haVmRestartTimeout",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_HaVmOverride)SetVirtualMachineId(val *string) {
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
func HaVmOverride_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHaVmOverride_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.haVmOverride.HaVmOverride",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HaVmOverride_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHaVmOverride_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.haVmOverride.HaVmOverride",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HaVmOverride_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHaVmOverride_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.haVmOverride.HaVmOverride",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HaVmOverride_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vsphere.haVmOverride.HaVmOverride",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HaVmOverride) AddOverride(path *string, value interface{}) {
	if err := h.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HaVmOverride) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HaVmOverride) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HaVmOverride) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HaVmOverride) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HaVmOverride) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HaVmOverride) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HaVmOverride) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HaVmOverride) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HaVmOverride) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HaVmOverride) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HaVmOverride) OverrideLogicalId(newLogicalId *string) {
	if err := h.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HaVmOverride) ResetHaDatastoreApdRecoveryAction() {
	_jsii_.InvokeVoid(
		h,
		"resetHaDatastoreApdRecoveryAction",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HaVmOverride) ResetHaDatastoreApdResponse() {
	_jsii_.InvokeVoid(
		h,
		"resetHaDatastoreApdResponse",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HaVmOverride) ResetHaDatastoreApdResponseDelay() {
	_jsii_.InvokeVoid(
		h,
		"resetHaDatastoreApdResponseDelay",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HaVmOverride) ResetHaDatastorePdlResponse() {
	_jsii_.InvokeVoid(
		h,
		"resetHaDatastorePdlResponse",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HaVmOverride) ResetHaHostIsolationResponse() {
	_jsii_.InvokeVoid(
		h,
		"resetHaHostIsolationResponse",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HaVmOverride) ResetHaVmFailureInterval() {
	_jsii_.InvokeVoid(
		h,
		"resetHaVmFailureInterval",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HaVmOverride) ResetHaVmMaximumFailureWindow() {
	_jsii_.InvokeVoid(
		h,
		"resetHaVmMaximumFailureWindow",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HaVmOverride) ResetHaVmMaximumResets() {
	_jsii_.InvokeVoid(
		h,
		"resetHaVmMaximumResets",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HaVmOverride) ResetHaVmMinimumUptime() {
	_jsii_.InvokeVoid(
		h,
		"resetHaVmMinimumUptime",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HaVmOverride) ResetHaVmMonitoring() {
	_jsii_.InvokeVoid(
		h,
		"resetHaVmMonitoring",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HaVmOverride) ResetHaVmMonitoringUseClusterDefaults() {
	_jsii_.InvokeVoid(
		h,
		"resetHaVmMonitoringUseClusterDefaults",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HaVmOverride) ResetHaVmRestartPriority() {
	_jsii_.InvokeVoid(
		h,
		"resetHaVmRestartPriority",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HaVmOverride) ResetHaVmRestartTimeout() {
	_jsii_.InvokeVoid(
		h,
		"resetHaVmRestartTimeout",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HaVmOverride) ResetId() {
	_jsii_.InvokeVoid(
		h,
		"resetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HaVmOverride) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HaVmOverride) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HaVmOverride) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HaVmOverride) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HaVmOverride) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

