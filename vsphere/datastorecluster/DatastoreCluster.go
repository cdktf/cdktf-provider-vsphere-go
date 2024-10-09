// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastorecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v9/datastorecluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/datastore_cluster vsphere_datastore_cluster}.
type DatastoreCluster interface {
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	SdrsAdvancedOptions() *map[string]*string
	SetSdrsAdvancedOptions(val *map[string]*string)
	SdrsAdvancedOptionsInput() *map[string]*string
	SdrsAutomationLevel() *string
	SetSdrsAutomationLevel(val *string)
	SdrsAutomationLevelInput() *string
	SdrsDefaultIntraVmAffinity() interface{}
	SetSdrsDefaultIntraVmAffinity(val interface{})
	SdrsDefaultIntraVmAffinityInput() interface{}
	SdrsEnabled() interface{}
	SetSdrsEnabled(val interface{})
	SdrsEnabledInput() interface{}
	SdrsFreeSpaceThreshold() *float64
	SetSdrsFreeSpaceThreshold(val *float64)
	SdrsFreeSpaceThresholdInput() *float64
	SdrsFreeSpaceThresholdMode() *string
	SetSdrsFreeSpaceThresholdMode(val *string)
	SdrsFreeSpaceThresholdModeInput() *string
	SdrsFreeSpaceUtilizationDifference() *float64
	SetSdrsFreeSpaceUtilizationDifference(val *float64)
	SdrsFreeSpaceUtilizationDifferenceInput() *float64
	SdrsIoBalanceAutomationLevel() *string
	SetSdrsIoBalanceAutomationLevel(val *string)
	SdrsIoBalanceAutomationLevelInput() *string
	SdrsIoLatencyThreshold() *float64
	SetSdrsIoLatencyThreshold(val *float64)
	SdrsIoLatencyThresholdInput() *float64
	SdrsIoLoadBalanceEnabled() interface{}
	SetSdrsIoLoadBalanceEnabled(val interface{})
	SdrsIoLoadBalanceEnabledInput() interface{}
	SdrsIoLoadImbalanceThreshold() *float64
	SetSdrsIoLoadImbalanceThreshold(val *float64)
	SdrsIoLoadImbalanceThresholdInput() *float64
	SdrsIoReservableIopsThreshold() *float64
	SetSdrsIoReservableIopsThreshold(val *float64)
	SdrsIoReservableIopsThresholdInput() *float64
	SdrsIoReservablePercentThreshold() *float64
	SetSdrsIoReservablePercentThreshold(val *float64)
	SdrsIoReservablePercentThresholdInput() *float64
	SdrsIoReservableThresholdMode() *string
	SetSdrsIoReservableThresholdMode(val *string)
	SdrsIoReservableThresholdModeInput() *string
	SdrsLoadBalanceInterval() *float64
	SetSdrsLoadBalanceInterval(val *float64)
	SdrsLoadBalanceIntervalInput() *float64
	SdrsPolicyEnforcementAutomationLevel() *string
	SetSdrsPolicyEnforcementAutomationLevel(val *string)
	SdrsPolicyEnforcementAutomationLevelInput() *string
	SdrsRuleEnforcementAutomationLevel() *string
	SetSdrsRuleEnforcementAutomationLevel(val *string)
	SdrsRuleEnforcementAutomationLevelInput() *string
	SdrsSpaceBalanceAutomationLevel() *string
	SetSdrsSpaceBalanceAutomationLevel(val *string)
	SdrsSpaceBalanceAutomationLevelInput() *string
	SdrsSpaceUtilizationThreshold() *float64
	SetSdrsSpaceUtilizationThreshold(val *float64)
	SdrsSpaceUtilizationThresholdInput() *float64
	SdrsVmEvacuationAutomationLevel() *string
	SetSdrsVmEvacuationAutomationLevel(val *string)
	SdrsVmEvacuationAutomationLevelInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
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
	ResetCustomAttributes()
	ResetFolder()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSdrsAdvancedOptions()
	ResetSdrsAutomationLevel()
	ResetSdrsDefaultIntraVmAffinity()
	ResetSdrsEnabled()
	ResetSdrsFreeSpaceThreshold()
	ResetSdrsFreeSpaceThresholdMode()
	ResetSdrsFreeSpaceUtilizationDifference()
	ResetSdrsIoBalanceAutomationLevel()
	ResetSdrsIoLatencyThreshold()
	ResetSdrsIoLoadBalanceEnabled()
	ResetSdrsIoLoadImbalanceThreshold()
	ResetSdrsIoReservableIopsThreshold()
	ResetSdrsIoReservablePercentThreshold()
	ResetSdrsIoReservableThresholdMode()
	ResetSdrsLoadBalanceInterval()
	ResetSdrsPolicyEnforcementAutomationLevel()
	ResetSdrsRuleEnforcementAutomationLevel()
	ResetSdrsSpaceBalanceAutomationLevel()
	ResetSdrsSpaceUtilizationThreshold()
	ResetSdrsVmEvacuationAutomationLevel()
	ResetTags()
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

// The jsii proxy struct for DatastoreCluster
type jsiiProxy_DatastoreCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatastoreCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) CustomAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) CustomAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) DatacenterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) DatacenterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) Folder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) FolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsAdvancedOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sdrsAdvancedOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsAdvancedOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sdrsAdvancedOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsAutomationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsAutomationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsAutomationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsAutomationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsDefaultIntraVmAffinity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sdrsDefaultIntraVmAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsDefaultIntraVmAffinityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sdrsDefaultIntraVmAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sdrsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sdrsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsFreeSpaceThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sdrsFreeSpaceThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsFreeSpaceThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sdrsFreeSpaceThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsFreeSpaceThresholdMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsFreeSpaceThresholdMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsFreeSpaceThresholdModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsFreeSpaceThresholdModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsFreeSpaceUtilizationDifference() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sdrsFreeSpaceUtilizationDifference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsFreeSpaceUtilizationDifferenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sdrsFreeSpaceUtilizationDifferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsIoBalanceAutomationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsIoBalanceAutomationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsIoBalanceAutomationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsIoBalanceAutomationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsIoLatencyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sdrsIoLatencyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsIoLatencyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sdrsIoLatencyThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsIoLoadBalanceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sdrsIoLoadBalanceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsIoLoadBalanceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sdrsIoLoadBalanceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsIoLoadImbalanceThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sdrsIoLoadImbalanceThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsIoLoadImbalanceThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sdrsIoLoadImbalanceThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsIoReservableIopsThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sdrsIoReservableIopsThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsIoReservableIopsThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sdrsIoReservableIopsThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsIoReservablePercentThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sdrsIoReservablePercentThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsIoReservablePercentThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sdrsIoReservablePercentThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsIoReservableThresholdMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsIoReservableThresholdMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsIoReservableThresholdModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsIoReservableThresholdModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsLoadBalanceInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sdrsLoadBalanceInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsLoadBalanceIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sdrsLoadBalanceIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsPolicyEnforcementAutomationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsPolicyEnforcementAutomationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsPolicyEnforcementAutomationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsPolicyEnforcementAutomationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsRuleEnforcementAutomationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsRuleEnforcementAutomationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsRuleEnforcementAutomationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsRuleEnforcementAutomationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsSpaceBalanceAutomationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsSpaceBalanceAutomationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsSpaceBalanceAutomationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsSpaceBalanceAutomationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsSpaceUtilizationThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sdrsSpaceUtilizationThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsSpaceUtilizationThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sdrsSpaceUtilizationThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsVmEvacuationAutomationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsVmEvacuationAutomationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) SdrsVmEvacuationAutomationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdrsVmEvacuationAutomationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastoreCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/datastore_cluster vsphere_datastore_cluster} Resource.
func NewDatastoreCluster(scope constructs.Construct, id *string, config *DatastoreClusterConfig) DatastoreCluster {
	_init_.Initialize()

	if err := validateNewDatastoreClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatastoreCluster{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.datastoreCluster.DatastoreCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/datastore_cluster vsphere_datastore_cluster} Resource.
func NewDatastoreCluster_Override(d DatastoreCluster, scope constructs.Construct, id *string, config *DatastoreClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.datastoreCluster.DatastoreCluster",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetCustomAttributes(val *map[string]*string) {
	if err := j.validateSetCustomAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAttributes",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetDatacenterId(val *string) {
	if err := j.validateSetDatacenterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datacenterId",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetFolder(val *string) {
	if err := j.validateSetFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"folder",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsAdvancedOptions(val *map[string]*string) {
	if err := j.validateSetSdrsAdvancedOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsAdvancedOptions",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsAutomationLevel(val *string) {
	if err := j.validateSetSdrsAutomationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsAutomationLevel",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsDefaultIntraVmAffinity(val interface{}) {
	if err := j.validateSetSdrsDefaultIntraVmAffinityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsDefaultIntraVmAffinity",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsEnabled(val interface{}) {
	if err := j.validateSetSdrsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsEnabled",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsFreeSpaceThreshold(val *float64) {
	if err := j.validateSetSdrsFreeSpaceThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsFreeSpaceThreshold",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsFreeSpaceThresholdMode(val *string) {
	if err := j.validateSetSdrsFreeSpaceThresholdModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsFreeSpaceThresholdMode",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsFreeSpaceUtilizationDifference(val *float64) {
	if err := j.validateSetSdrsFreeSpaceUtilizationDifferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsFreeSpaceUtilizationDifference",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsIoBalanceAutomationLevel(val *string) {
	if err := j.validateSetSdrsIoBalanceAutomationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsIoBalanceAutomationLevel",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsIoLatencyThreshold(val *float64) {
	if err := j.validateSetSdrsIoLatencyThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsIoLatencyThreshold",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsIoLoadBalanceEnabled(val interface{}) {
	if err := j.validateSetSdrsIoLoadBalanceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsIoLoadBalanceEnabled",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsIoLoadImbalanceThreshold(val *float64) {
	if err := j.validateSetSdrsIoLoadImbalanceThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsIoLoadImbalanceThreshold",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsIoReservableIopsThreshold(val *float64) {
	if err := j.validateSetSdrsIoReservableIopsThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsIoReservableIopsThreshold",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsIoReservablePercentThreshold(val *float64) {
	if err := j.validateSetSdrsIoReservablePercentThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsIoReservablePercentThreshold",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsIoReservableThresholdMode(val *string) {
	if err := j.validateSetSdrsIoReservableThresholdModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsIoReservableThresholdMode",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsLoadBalanceInterval(val *float64) {
	if err := j.validateSetSdrsLoadBalanceIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsLoadBalanceInterval",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsPolicyEnforcementAutomationLevel(val *string) {
	if err := j.validateSetSdrsPolicyEnforcementAutomationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsPolicyEnforcementAutomationLevel",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsRuleEnforcementAutomationLevel(val *string) {
	if err := j.validateSetSdrsRuleEnforcementAutomationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsRuleEnforcementAutomationLevel",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsSpaceBalanceAutomationLevel(val *string) {
	if err := j.validateSetSdrsSpaceBalanceAutomationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsSpaceBalanceAutomationLevel",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsSpaceUtilizationThreshold(val *float64) {
	if err := j.validateSetSdrsSpaceUtilizationThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsSpaceUtilizationThreshold",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetSdrsVmEvacuationAutomationLevel(val *string) {
	if err := j.validateSetSdrsVmEvacuationAutomationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdrsVmEvacuationAutomationLevel",
		val,
	)
}

func (j *jsiiProxy_DatastoreCluster)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a DatastoreCluster resource upon running "cdktf plan <stack-name>".
func DatastoreCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDatastoreCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.datastoreCluster.DatastoreCluster",
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
func DatastoreCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatastoreCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.datastoreCluster.DatastoreCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatastoreCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatastoreCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.datastoreCluster.DatastoreCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatastoreCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatastoreCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.datastoreCluster.DatastoreCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatastoreCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vsphere.datastoreCluster.DatastoreCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatastoreCluster) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DatastoreCluster) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatastoreCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatastoreCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastoreCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatastoreCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatastoreCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatastoreCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatastoreCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatastoreCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatastoreCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatastoreCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DatastoreCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastoreCluster) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatastoreCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DatastoreCluster) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatastoreCluster) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetCustomAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetFolder() {
	_jsii_.InvokeVoid(
		d,
		"resetFolder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsAdvancedOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsAdvancedOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsAutomationLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsAutomationLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsDefaultIntraVmAffinity() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsDefaultIntraVmAffinity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsFreeSpaceThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsFreeSpaceThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsFreeSpaceThresholdMode() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsFreeSpaceThresholdMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsFreeSpaceUtilizationDifference() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsFreeSpaceUtilizationDifference",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsIoBalanceAutomationLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsIoBalanceAutomationLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsIoLatencyThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsIoLatencyThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsIoLoadBalanceEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsIoLoadBalanceEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsIoLoadImbalanceThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsIoLoadImbalanceThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsIoReservableIopsThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsIoReservableIopsThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsIoReservablePercentThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsIoReservablePercentThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsIoReservableThresholdMode() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsIoReservableThresholdMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsLoadBalanceInterval() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsLoadBalanceInterval",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsPolicyEnforcementAutomationLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsPolicyEnforcementAutomationLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsRuleEnforcementAutomationLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsRuleEnforcementAutomationLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsSpaceBalanceAutomationLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsSpaceBalanceAutomationLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsSpaceUtilizationThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsSpaceUtilizationThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetSdrsVmEvacuationAutomationLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetSdrsVmEvacuationAutomationLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastoreCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

