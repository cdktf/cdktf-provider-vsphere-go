// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package supervisor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v9/supervisor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs/resources/supervisor vsphere_supervisor}.
type Supervisor interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Cluster() *string
	SetCluster(val *string)
	ClusterInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContentLibrary() *string
	SetContentLibrary(val *string)
	ContentLibraryInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DvsUuid() *string
	SetDvsUuid(val *string)
	DvsUuidInput() *string
	EdgeCluster() *string
	SetEdgeCluster(val *string)
	EdgeClusterInput() *string
	EgressCidr() SupervisorEgressCidrList
	EgressCidrInput() interface{}
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
	IngressCidr() SupervisorIngressCidrList
	IngressCidrInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MainDns() *[]*string
	SetMainDns(val *[]*string)
	MainDnsInput() *[]*string
	ManagementNetwork() SupervisorManagementNetworkOutputReference
	ManagementNetworkInput() *SupervisorManagementNetwork
	Namespace() SupervisorNamespaceList
	NamespaceInput() interface{}
	// The tree node.
	Node() constructs.Node
	PodCidr() SupervisorPodCidrList
	PodCidrInput() interface{}
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
	SearchDomains() *[]*string
	SetSearchDomains(val *[]*string)
	SearchDomainsInput() *[]*string
	ServiceCidr() SupervisorServiceCidrOutputReference
	ServiceCidrInput() *SupervisorServiceCidr
	SizingHint() *string
	SetSizingHint(val *string)
	SizingHintInput() *string
	StoragePolicy() *string
	SetStoragePolicy(val *string)
	StoragePolicyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WorkerDns() *[]*string
	SetWorkerDns(val *[]*string)
	WorkerDnsInput() *[]*string
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
	PutEgressCidr(value interface{})
	PutIngressCidr(value interface{})
	PutManagementNetwork(value *SupervisorManagementNetwork)
	PutNamespace(value interface{})
	PutPodCidr(value interface{})
	PutServiceCidr(value *SupervisorServiceCidr)
	ResetId()
	ResetNamespace()
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

// The jsii proxy struct for Supervisor
type jsiiProxy_Supervisor struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Supervisor) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) ClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) ContentLibrary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentLibrary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) ContentLibraryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentLibraryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) DvsUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dvsUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) DvsUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dvsUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) EdgeCluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) EdgeClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) EgressCidr() SupervisorEgressCidrList {
	var returns SupervisorEgressCidrList
	_jsii_.Get(
		j,
		"egressCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) EgressCidrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"egressCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) IngressCidr() SupervisorIngressCidrList {
	var returns SupervisorIngressCidrList
	_jsii_.Get(
		j,
		"ingressCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) IngressCidrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingressCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) MainDns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mainDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) MainDnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mainDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) ManagementNetwork() SupervisorManagementNetworkOutputReference {
	var returns SupervisorManagementNetworkOutputReference
	_jsii_.Get(
		j,
		"managementNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) ManagementNetworkInput() *SupervisorManagementNetwork {
	var returns *SupervisorManagementNetwork
	_jsii_.Get(
		j,
		"managementNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) Namespace() SupervisorNamespaceList {
	var returns SupervisorNamespaceList
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) NamespaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) PodCidr() SupervisorPodCidrList {
	var returns SupervisorPodCidrList
	_jsii_.Get(
		j,
		"podCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) PodCidrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"podCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) SearchDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searchDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) SearchDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searchDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) ServiceCidr() SupervisorServiceCidrOutputReference {
	var returns SupervisorServiceCidrOutputReference
	_jsii_.Get(
		j,
		"serviceCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) ServiceCidrInput() *SupervisorServiceCidr {
	var returns *SupervisorServiceCidr
	_jsii_.Get(
		j,
		"serviceCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) SizingHint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingHint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) SizingHintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingHintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) StoragePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) StoragePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) WorkerDns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"workerDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Supervisor) WorkerDnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"workerDnsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs/resources/supervisor vsphere_supervisor} Resource.
func NewSupervisor(scope constructs.Construct, id *string, config *SupervisorConfig) Supervisor {
	_init_.Initialize()

	if err := validateNewSupervisorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Supervisor{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.supervisor.Supervisor",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs/resources/supervisor vsphere_supervisor} Resource.
func NewSupervisor_Override(s Supervisor, scope constructs.Construct, id *string, config *SupervisorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.supervisor.Supervisor",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_Supervisor)SetCluster(val *string) {
	if err := j.validateSetClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_Supervisor)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Supervisor)SetContentLibrary(val *string) {
	if err := j.validateSetContentLibraryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentLibrary",
		val,
	)
}

func (j *jsiiProxy_Supervisor)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Supervisor)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Supervisor)SetDvsUuid(val *string) {
	if err := j.validateSetDvsUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dvsUuid",
		val,
	)
}

func (j *jsiiProxy_Supervisor)SetEdgeCluster(val *string) {
	if err := j.validateSetEdgeClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeCluster",
		val,
	)
}

func (j *jsiiProxy_Supervisor)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Supervisor)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Supervisor)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Supervisor)SetMainDns(val *[]*string) {
	if err := j.validateSetMainDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mainDns",
		val,
	)
}

func (j *jsiiProxy_Supervisor)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Supervisor)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Supervisor)SetSearchDomains(val *[]*string) {
	if err := j.validateSetSearchDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchDomains",
		val,
	)
}

func (j *jsiiProxy_Supervisor)SetSizingHint(val *string) {
	if err := j.validateSetSizingHintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizingHint",
		val,
	)
}

func (j *jsiiProxy_Supervisor)SetStoragePolicy(val *string) {
	if err := j.validateSetStoragePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storagePolicy",
		val,
	)
}

func (j *jsiiProxy_Supervisor)SetWorkerDns(val *[]*string) {
	if err := j.validateSetWorkerDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workerDns",
		val,
	)
}

// Generates CDKTF code for importing a Supervisor resource upon running "cdktf plan <stack-name>".
func Supervisor_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSupervisor_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.supervisor.Supervisor",
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
func Supervisor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSupervisor_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.supervisor.Supervisor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Supervisor_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSupervisor_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.supervisor.Supervisor",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Supervisor_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSupervisor_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.supervisor.Supervisor",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Supervisor_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vsphere.supervisor.Supervisor",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_Supervisor) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_Supervisor) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_Supervisor) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_Supervisor) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_Supervisor) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_Supervisor) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_Supervisor) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_Supervisor) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_Supervisor) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_Supervisor) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_Supervisor) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_Supervisor) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Supervisor) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_Supervisor) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_Supervisor) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_Supervisor) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_Supervisor) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_Supervisor) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_Supervisor) PutEgressCidr(value interface{}) {
	if err := s.validatePutEgressCidrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEgressCidr",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Supervisor) PutIngressCidr(value interface{}) {
	if err := s.validatePutIngressCidrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIngressCidr",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Supervisor) PutManagementNetwork(value *SupervisorManagementNetwork) {
	if err := s.validatePutManagementNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putManagementNetwork",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Supervisor) PutNamespace(value interface{}) {
	if err := s.validatePutNamespaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNamespace",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Supervisor) PutPodCidr(value interface{}) {
	if err := s.validatePutPodCidrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPodCidr",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Supervisor) PutServiceCidr(value *SupervisorServiceCidr) {
	if err := s.validatePutServiceCidrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putServiceCidr",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Supervisor) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Supervisor) ResetNamespace() {
	_jsii_.InvokeVoid(
		s,
		"resetNamespace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Supervisor) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Supervisor) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Supervisor) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Supervisor) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Supervisor) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Supervisor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Supervisor) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

