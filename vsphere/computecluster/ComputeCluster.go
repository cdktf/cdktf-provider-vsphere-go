// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v9/computecluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs/resources/compute_cluster vsphere_compute_cluster}.
type ComputeCluster interface {
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
	DpmAutomationLevel() *string
	SetDpmAutomationLevel(val *string)
	DpmAutomationLevelInput() *string
	DpmEnabled() interface{}
	SetDpmEnabled(val interface{})
	DpmEnabledInput() interface{}
	DpmThreshold() *float64
	SetDpmThreshold(val *float64)
	DpmThresholdInput() *float64
	DrsAdvancedOptions() *map[string]*string
	SetDrsAdvancedOptions(val *map[string]*string)
	DrsAdvancedOptionsInput() *map[string]*string
	DrsAutomationLevel() *string
	SetDrsAutomationLevel(val *string)
	DrsAutomationLevelInput() *string
	DrsEnabled() interface{}
	SetDrsEnabled(val interface{})
	DrsEnabledInput() interface{}
	DrsEnablePredictiveDrs() interface{}
	SetDrsEnablePredictiveDrs(val interface{})
	DrsEnablePredictiveDrsInput() interface{}
	DrsEnableVmOverrides() interface{}
	SetDrsEnableVmOverrides(val interface{})
	DrsEnableVmOverridesInput() interface{}
	DrsMigrationThreshold() *float64
	SetDrsMigrationThreshold(val *float64)
	DrsMigrationThresholdInput() *float64
	DrsScaleDescendantsShares() *string
	SetDrsScaleDescendantsShares(val *string)
	DrsScaleDescendantsSharesInput() *string
	Folder() *string
	SetFolder(val *string)
	FolderInput() *string
	ForceEvacuateOnDestroy() interface{}
	SetForceEvacuateOnDestroy(val interface{})
	ForceEvacuateOnDestroyInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HaAdmissionControlFailoverHostSystemIds() *[]*string
	SetHaAdmissionControlFailoverHostSystemIds(val *[]*string)
	HaAdmissionControlFailoverHostSystemIdsInput() *[]*string
	HaAdmissionControlHostFailureTolerance() *float64
	SetHaAdmissionControlHostFailureTolerance(val *float64)
	HaAdmissionControlHostFailureToleranceInput() *float64
	HaAdmissionControlPerformanceTolerance() *float64
	SetHaAdmissionControlPerformanceTolerance(val *float64)
	HaAdmissionControlPerformanceToleranceInput() *float64
	HaAdmissionControlPolicy() *string
	SetHaAdmissionControlPolicy(val *string)
	HaAdmissionControlPolicyInput() *string
	HaAdmissionControlResourcePercentageAutoCompute() interface{}
	SetHaAdmissionControlResourcePercentageAutoCompute(val interface{})
	HaAdmissionControlResourcePercentageAutoComputeInput() interface{}
	HaAdmissionControlResourcePercentageCpu() *float64
	SetHaAdmissionControlResourcePercentageCpu(val *float64)
	HaAdmissionControlResourcePercentageCpuInput() *float64
	HaAdmissionControlResourcePercentageMemory() *float64
	SetHaAdmissionControlResourcePercentageMemory(val *float64)
	HaAdmissionControlResourcePercentageMemoryInput() *float64
	HaAdmissionControlSlotPolicyExplicitCpu() *float64
	SetHaAdmissionControlSlotPolicyExplicitCpu(val *float64)
	HaAdmissionControlSlotPolicyExplicitCpuInput() *float64
	HaAdmissionControlSlotPolicyExplicitMemory() *float64
	SetHaAdmissionControlSlotPolicyExplicitMemory(val *float64)
	HaAdmissionControlSlotPolicyExplicitMemoryInput() *float64
	HaAdmissionControlSlotPolicyUseExplicitSize() interface{}
	SetHaAdmissionControlSlotPolicyUseExplicitSize(val interface{})
	HaAdmissionControlSlotPolicyUseExplicitSizeInput() interface{}
	HaAdvancedOptions() *map[string]*string
	SetHaAdvancedOptions(val *map[string]*string)
	HaAdvancedOptionsInput() *map[string]*string
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
	HaEnabled() interface{}
	SetHaEnabled(val interface{})
	HaEnabledInput() interface{}
	HaHeartbeatDatastoreIds() *[]*string
	SetHaHeartbeatDatastoreIds(val *[]*string)
	HaHeartbeatDatastoreIdsInput() *[]*string
	HaHeartbeatDatastorePolicy() *string
	SetHaHeartbeatDatastorePolicy(val *string)
	HaHeartbeatDatastorePolicyInput() *string
	HaHostIsolationResponse() *string
	SetHaHostIsolationResponse(val *string)
	HaHostIsolationResponseInput() *string
	HaHostMonitoring() *string
	SetHaHostMonitoring(val *string)
	HaHostMonitoringInput() *string
	HaVmComponentProtection() *string
	SetHaVmComponentProtection(val *string)
	HaVmComponentProtectionInput() *string
	HaVmDependencyRestartCondition() *string
	SetHaVmDependencyRestartCondition(val *string)
	HaVmDependencyRestartConditionInput() *string
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
	HaVmRestartAdditionalDelay() *float64
	SetHaVmRestartAdditionalDelay(val *float64)
	HaVmRestartAdditionalDelayInput() *float64
	HaVmRestartPriority() *string
	SetHaVmRestartPriority(val *string)
	HaVmRestartPriorityInput() *string
	HaVmRestartTimeout() *float64
	SetHaVmRestartTimeout(val *float64)
	HaVmRestartTimeoutInput() *float64
	HostClusterExitTimeout() *float64
	SetHostClusterExitTimeout(val *float64)
	HostClusterExitTimeoutInput() *float64
	HostImage() ComputeClusterHostImageOutputReference
	HostImageInput() *ComputeClusterHostImage
	HostManaged() interface{}
	SetHostManaged(val interface{})
	HostManagedInput() interface{}
	HostSystemIds() *[]*string
	SetHostSystemIds(val *[]*string)
	HostSystemIdsInput() *[]*string
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
	ProactiveHaAutomationLevel() *string
	SetProactiveHaAutomationLevel(val *string)
	ProactiveHaAutomationLevelInput() *string
	ProactiveHaEnabled() interface{}
	SetProactiveHaEnabled(val interface{})
	ProactiveHaEnabledInput() interface{}
	ProactiveHaModerateRemediation() *string
	SetProactiveHaModerateRemediation(val *string)
	ProactiveHaModerateRemediationInput() *string
	ProactiveHaProviderIds() *[]*string
	SetProactiveHaProviderIds(val *[]*string)
	ProactiveHaProviderIdsInput() *[]*string
	ProactiveHaSevereRemediation() *string
	SetProactiveHaSevereRemediation(val *string)
	ProactiveHaSevereRemediationInput() *string
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
	ResourcePoolId() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VsanCompressionEnabled() interface{}
	SetVsanCompressionEnabled(val interface{})
	VsanCompressionEnabledInput() interface{}
	VsanDedupEnabled() interface{}
	SetVsanDedupEnabled(val interface{})
	VsanDedupEnabledInput() interface{}
	VsanDiskGroup() ComputeClusterVsanDiskGroupList
	VsanDiskGroupInput() interface{}
	VsanDitEncryptionEnabled() interface{}
	SetVsanDitEncryptionEnabled(val interface{})
	VsanDitEncryptionEnabledInput() interface{}
	VsanDitRekeyInterval() *float64
	SetVsanDitRekeyInterval(val *float64)
	VsanDitRekeyIntervalInput() *float64
	VsanEnabled() interface{}
	SetVsanEnabled(val interface{})
	VsanEnabledInput() interface{}
	VsanEsaEnabled() interface{}
	SetVsanEsaEnabled(val interface{})
	VsanEsaEnabledInput() interface{}
	VsanFaultDomains() ComputeClusterVsanFaultDomainsList
	VsanFaultDomainsInput() interface{}
	VsanNetworkDiagnosticModeEnabled() interface{}
	SetVsanNetworkDiagnosticModeEnabled(val interface{})
	VsanNetworkDiagnosticModeEnabledInput() interface{}
	VsanPerformanceEnabled() interface{}
	SetVsanPerformanceEnabled(val interface{})
	VsanPerformanceEnabledInput() interface{}
	VsanRemoteDatastoreIds() *[]*string
	SetVsanRemoteDatastoreIds(val *[]*string)
	VsanRemoteDatastoreIdsInput() *[]*string
	VsanStretchedCluster() ComputeClusterVsanStretchedClusterOutputReference
	VsanStretchedClusterInput() *ComputeClusterVsanStretchedCluster
	VsanUnmapEnabled() interface{}
	SetVsanUnmapEnabled(val interface{})
	VsanUnmapEnabledInput() interface{}
	VsanVerboseModeEnabled() interface{}
	SetVsanVerboseModeEnabled(val interface{})
	VsanVerboseModeEnabledInput() interface{}
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
	PutHostImage(value *ComputeClusterHostImage)
	PutVsanDiskGroup(value interface{})
	PutVsanFaultDomains(value interface{})
	PutVsanStretchedCluster(value *ComputeClusterVsanStretchedCluster)
	ResetCustomAttributes()
	ResetDpmAutomationLevel()
	ResetDpmEnabled()
	ResetDpmThreshold()
	ResetDrsAdvancedOptions()
	ResetDrsAutomationLevel()
	ResetDrsEnabled()
	ResetDrsEnablePredictiveDrs()
	ResetDrsEnableVmOverrides()
	ResetDrsMigrationThreshold()
	ResetDrsScaleDescendantsShares()
	ResetFolder()
	ResetForceEvacuateOnDestroy()
	ResetHaAdmissionControlFailoverHostSystemIds()
	ResetHaAdmissionControlHostFailureTolerance()
	ResetHaAdmissionControlPerformanceTolerance()
	ResetHaAdmissionControlPolicy()
	ResetHaAdmissionControlResourcePercentageAutoCompute()
	ResetHaAdmissionControlResourcePercentageCpu()
	ResetHaAdmissionControlResourcePercentageMemory()
	ResetHaAdmissionControlSlotPolicyExplicitCpu()
	ResetHaAdmissionControlSlotPolicyExplicitMemory()
	ResetHaAdmissionControlSlotPolicyUseExplicitSize()
	ResetHaAdvancedOptions()
	ResetHaDatastoreApdRecoveryAction()
	ResetHaDatastoreApdResponse()
	ResetHaDatastoreApdResponseDelay()
	ResetHaDatastorePdlResponse()
	ResetHaEnabled()
	ResetHaHeartbeatDatastoreIds()
	ResetHaHeartbeatDatastorePolicy()
	ResetHaHostIsolationResponse()
	ResetHaHostMonitoring()
	ResetHaVmComponentProtection()
	ResetHaVmDependencyRestartCondition()
	ResetHaVmFailureInterval()
	ResetHaVmMaximumFailureWindow()
	ResetHaVmMaximumResets()
	ResetHaVmMinimumUptime()
	ResetHaVmMonitoring()
	ResetHaVmRestartAdditionalDelay()
	ResetHaVmRestartPriority()
	ResetHaVmRestartTimeout()
	ResetHostClusterExitTimeout()
	ResetHostImage()
	ResetHostManaged()
	ResetHostSystemIds()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProactiveHaAutomationLevel()
	ResetProactiveHaEnabled()
	ResetProactiveHaModerateRemediation()
	ResetProactiveHaProviderIds()
	ResetProactiveHaSevereRemediation()
	ResetTags()
	ResetVsanCompressionEnabled()
	ResetVsanDedupEnabled()
	ResetVsanDiskGroup()
	ResetVsanDitEncryptionEnabled()
	ResetVsanDitRekeyInterval()
	ResetVsanEnabled()
	ResetVsanEsaEnabled()
	ResetVsanFaultDomains()
	ResetVsanNetworkDiagnosticModeEnabled()
	ResetVsanPerformanceEnabled()
	ResetVsanRemoteDatastoreIds()
	ResetVsanStretchedCluster()
	ResetVsanUnmapEnabled()
	ResetVsanVerboseModeEnabled()
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

// The jsii proxy struct for ComputeCluster
type jsiiProxy_ComputeCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) CustomAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) CustomAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DatacenterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DatacenterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DpmAutomationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dpmAutomationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DpmAutomationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dpmAutomationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DpmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dpmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DpmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dpmEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DpmThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dpmThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DpmThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dpmThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DrsAdvancedOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"drsAdvancedOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DrsAdvancedOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"drsAdvancedOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DrsAutomationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drsAutomationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DrsAutomationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drsAutomationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DrsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DrsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DrsEnablePredictiveDrs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drsEnablePredictiveDrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DrsEnablePredictiveDrsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drsEnablePredictiveDrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DrsEnableVmOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drsEnableVmOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DrsEnableVmOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drsEnableVmOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DrsMigrationThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drsMigrationThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DrsMigrationThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drsMigrationThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DrsScaleDescendantsShares() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drsScaleDescendantsShares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) DrsScaleDescendantsSharesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drsScaleDescendantsSharesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) Folder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) FolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) ForceEvacuateOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceEvacuateOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) ForceEvacuateOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceEvacuateOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlFailoverHostSystemIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"haAdmissionControlFailoverHostSystemIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlFailoverHostSystemIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"haAdmissionControlFailoverHostSystemIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlHostFailureTolerance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haAdmissionControlHostFailureTolerance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlHostFailureToleranceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haAdmissionControlHostFailureToleranceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlPerformanceTolerance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haAdmissionControlPerformanceTolerance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlPerformanceToleranceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haAdmissionControlPerformanceToleranceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haAdmissionControlPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haAdmissionControlPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlResourcePercentageAutoCompute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"haAdmissionControlResourcePercentageAutoCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlResourcePercentageAutoComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"haAdmissionControlResourcePercentageAutoComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlResourcePercentageCpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haAdmissionControlResourcePercentageCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlResourcePercentageCpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haAdmissionControlResourcePercentageCpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlResourcePercentageMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haAdmissionControlResourcePercentageMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlResourcePercentageMemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haAdmissionControlResourcePercentageMemoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlSlotPolicyExplicitCpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haAdmissionControlSlotPolicyExplicitCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlSlotPolicyExplicitCpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haAdmissionControlSlotPolicyExplicitCpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlSlotPolicyExplicitMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haAdmissionControlSlotPolicyExplicitMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlSlotPolicyExplicitMemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haAdmissionControlSlotPolicyExplicitMemoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlSlotPolicyUseExplicitSize() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"haAdmissionControlSlotPolicyUseExplicitSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdmissionControlSlotPolicyUseExplicitSizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"haAdmissionControlSlotPolicyUseExplicitSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdvancedOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"haAdvancedOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaAdvancedOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"haAdvancedOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaDatastoreApdRecoveryAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haDatastoreApdRecoveryAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaDatastoreApdRecoveryActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haDatastoreApdRecoveryActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaDatastoreApdResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haDatastoreApdResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaDatastoreApdResponseDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haDatastoreApdResponseDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaDatastoreApdResponseDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haDatastoreApdResponseDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaDatastoreApdResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haDatastoreApdResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaDatastorePdlResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haDatastorePdlResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaDatastorePdlResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haDatastorePdlResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"haEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"haEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaHeartbeatDatastoreIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"haHeartbeatDatastoreIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaHeartbeatDatastoreIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"haHeartbeatDatastoreIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaHeartbeatDatastorePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haHeartbeatDatastorePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaHeartbeatDatastorePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haHeartbeatDatastorePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaHostIsolationResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haHostIsolationResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaHostIsolationResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haHostIsolationResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaHostMonitoring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haHostMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaHostMonitoringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haHostMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmComponentProtection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haVmComponentProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmComponentProtectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haVmComponentProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmDependencyRestartCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haVmDependencyRestartCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmDependencyRestartConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haVmDependencyRestartConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmFailureInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmFailureInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmFailureIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmFailureIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmMaximumFailureWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmMaximumFailureWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmMaximumFailureWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmMaximumFailureWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmMaximumResets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmMaximumResets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmMaximumResetsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmMaximumResetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmMinimumUptime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmMinimumUptime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmMinimumUptimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmMinimumUptimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmMonitoring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haVmMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmMonitoringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haVmMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmRestartAdditionalDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmRestartAdditionalDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmRestartAdditionalDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmRestartAdditionalDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmRestartPriority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haVmRestartPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmRestartPriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"haVmRestartPriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmRestartTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmRestartTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HaVmRestartTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haVmRestartTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HostClusterExitTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hostClusterExitTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HostClusterExitTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hostClusterExitTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HostImage() ComputeClusterHostImageOutputReference {
	var returns ComputeClusterHostImageOutputReference
	_jsii_.Get(
		j,
		"hostImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HostImageInput() *ComputeClusterHostImage {
	var returns *ComputeClusterHostImage
	_jsii_.Get(
		j,
		"hostImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HostManaged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HostManagedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostManagedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HostSystemIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostSystemIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) HostSystemIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostSystemIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) ProactiveHaAutomationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proactiveHaAutomationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) ProactiveHaAutomationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proactiveHaAutomationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) ProactiveHaEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proactiveHaEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) ProactiveHaEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proactiveHaEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) ProactiveHaModerateRemediation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proactiveHaModerateRemediation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) ProactiveHaModerateRemediationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proactiveHaModerateRemediationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) ProactiveHaProviderIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"proactiveHaProviderIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) ProactiveHaProviderIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"proactiveHaProviderIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) ProactiveHaSevereRemediation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proactiveHaSevereRemediation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) ProactiveHaSevereRemediationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proactiveHaSevereRemediationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) ResourcePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanCompressionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanCompressionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanCompressionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanCompressionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanDedupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanDedupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanDedupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanDedupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanDiskGroup() ComputeClusterVsanDiskGroupList {
	var returns ComputeClusterVsanDiskGroupList
	_jsii_.Get(
		j,
		"vsanDiskGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanDiskGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanDiskGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanDitEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanDitEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanDitEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanDitEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanDitRekeyInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vsanDitRekeyInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanDitRekeyIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vsanDitRekeyIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanEsaEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanEsaEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanEsaEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanEsaEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanFaultDomains() ComputeClusterVsanFaultDomainsList {
	var returns ComputeClusterVsanFaultDomainsList
	_jsii_.Get(
		j,
		"vsanFaultDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanFaultDomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanFaultDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanNetworkDiagnosticModeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanNetworkDiagnosticModeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanNetworkDiagnosticModeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanNetworkDiagnosticModeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanPerformanceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanPerformanceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanPerformanceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanPerformanceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanRemoteDatastoreIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vsanRemoteDatastoreIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanRemoteDatastoreIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vsanRemoteDatastoreIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanStretchedCluster() ComputeClusterVsanStretchedClusterOutputReference {
	var returns ComputeClusterVsanStretchedClusterOutputReference
	_jsii_.Get(
		j,
		"vsanStretchedCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanStretchedClusterInput() *ComputeClusterVsanStretchedCluster {
	var returns *ComputeClusterVsanStretchedCluster
	_jsii_.Get(
		j,
		"vsanStretchedClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanUnmapEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanUnmapEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanUnmapEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanUnmapEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanVerboseModeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanVerboseModeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeCluster) VsanVerboseModeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vsanVerboseModeEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs/resources/compute_cluster vsphere_compute_cluster} Resource.
func NewComputeCluster(scope constructs.Construct, id *string, config *ComputeClusterConfig) ComputeCluster {
	_init_.Initialize()

	if err := validateNewComputeClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeCluster{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.computeCluster.ComputeCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs/resources/compute_cluster vsphere_compute_cluster} Resource.
func NewComputeCluster_Override(c ComputeCluster, scope constructs.Construct, id *string, config *ComputeClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.computeCluster.ComputeCluster",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetCustomAttributes(val *map[string]*string) {
	if err := j.validateSetCustomAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAttributes",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetDatacenterId(val *string) {
	if err := j.validateSetDatacenterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datacenterId",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetDpmAutomationLevel(val *string) {
	if err := j.validateSetDpmAutomationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dpmAutomationLevel",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetDpmEnabled(val interface{}) {
	if err := j.validateSetDpmEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dpmEnabled",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetDpmThreshold(val *float64) {
	if err := j.validateSetDpmThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dpmThreshold",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetDrsAdvancedOptions(val *map[string]*string) {
	if err := j.validateSetDrsAdvancedOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drsAdvancedOptions",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetDrsAutomationLevel(val *string) {
	if err := j.validateSetDrsAutomationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drsAutomationLevel",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetDrsEnabled(val interface{}) {
	if err := j.validateSetDrsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drsEnabled",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetDrsEnablePredictiveDrs(val interface{}) {
	if err := j.validateSetDrsEnablePredictiveDrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drsEnablePredictiveDrs",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetDrsEnableVmOverrides(val interface{}) {
	if err := j.validateSetDrsEnableVmOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drsEnableVmOverrides",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetDrsMigrationThreshold(val *float64) {
	if err := j.validateSetDrsMigrationThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drsMigrationThreshold",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetDrsScaleDescendantsShares(val *string) {
	if err := j.validateSetDrsScaleDescendantsSharesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drsScaleDescendantsShares",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetFolder(val *string) {
	if err := j.validateSetFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"folder",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetForceEvacuateOnDestroy(val interface{}) {
	if err := j.validateSetForceEvacuateOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceEvacuateOnDestroy",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaAdmissionControlFailoverHostSystemIds(val *[]*string) {
	if err := j.validateSetHaAdmissionControlFailoverHostSystemIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haAdmissionControlFailoverHostSystemIds",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaAdmissionControlHostFailureTolerance(val *float64) {
	if err := j.validateSetHaAdmissionControlHostFailureToleranceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haAdmissionControlHostFailureTolerance",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaAdmissionControlPerformanceTolerance(val *float64) {
	if err := j.validateSetHaAdmissionControlPerformanceToleranceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haAdmissionControlPerformanceTolerance",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaAdmissionControlPolicy(val *string) {
	if err := j.validateSetHaAdmissionControlPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haAdmissionControlPolicy",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaAdmissionControlResourcePercentageAutoCompute(val interface{}) {
	if err := j.validateSetHaAdmissionControlResourcePercentageAutoComputeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haAdmissionControlResourcePercentageAutoCompute",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaAdmissionControlResourcePercentageCpu(val *float64) {
	if err := j.validateSetHaAdmissionControlResourcePercentageCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haAdmissionControlResourcePercentageCpu",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaAdmissionControlResourcePercentageMemory(val *float64) {
	if err := j.validateSetHaAdmissionControlResourcePercentageMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haAdmissionControlResourcePercentageMemory",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaAdmissionControlSlotPolicyExplicitCpu(val *float64) {
	if err := j.validateSetHaAdmissionControlSlotPolicyExplicitCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haAdmissionControlSlotPolicyExplicitCpu",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaAdmissionControlSlotPolicyExplicitMemory(val *float64) {
	if err := j.validateSetHaAdmissionControlSlotPolicyExplicitMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haAdmissionControlSlotPolicyExplicitMemory",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaAdmissionControlSlotPolicyUseExplicitSize(val interface{}) {
	if err := j.validateSetHaAdmissionControlSlotPolicyUseExplicitSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haAdmissionControlSlotPolicyUseExplicitSize",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaAdvancedOptions(val *map[string]*string) {
	if err := j.validateSetHaAdvancedOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haAdvancedOptions",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaDatastoreApdRecoveryAction(val *string) {
	if err := j.validateSetHaDatastoreApdRecoveryActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haDatastoreApdRecoveryAction",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaDatastoreApdResponse(val *string) {
	if err := j.validateSetHaDatastoreApdResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haDatastoreApdResponse",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaDatastoreApdResponseDelay(val *float64) {
	if err := j.validateSetHaDatastoreApdResponseDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haDatastoreApdResponseDelay",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaDatastorePdlResponse(val *string) {
	if err := j.validateSetHaDatastorePdlResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haDatastorePdlResponse",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaEnabled(val interface{}) {
	if err := j.validateSetHaEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haEnabled",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaHeartbeatDatastoreIds(val *[]*string) {
	if err := j.validateSetHaHeartbeatDatastoreIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haHeartbeatDatastoreIds",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaHeartbeatDatastorePolicy(val *string) {
	if err := j.validateSetHaHeartbeatDatastorePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haHeartbeatDatastorePolicy",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaHostIsolationResponse(val *string) {
	if err := j.validateSetHaHostIsolationResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haHostIsolationResponse",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaHostMonitoring(val *string) {
	if err := j.validateSetHaHostMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haHostMonitoring",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaVmComponentProtection(val *string) {
	if err := j.validateSetHaVmComponentProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haVmComponentProtection",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaVmDependencyRestartCondition(val *string) {
	if err := j.validateSetHaVmDependencyRestartConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haVmDependencyRestartCondition",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaVmFailureInterval(val *float64) {
	if err := j.validateSetHaVmFailureIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haVmFailureInterval",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaVmMaximumFailureWindow(val *float64) {
	if err := j.validateSetHaVmMaximumFailureWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haVmMaximumFailureWindow",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaVmMaximumResets(val *float64) {
	if err := j.validateSetHaVmMaximumResetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haVmMaximumResets",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaVmMinimumUptime(val *float64) {
	if err := j.validateSetHaVmMinimumUptimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haVmMinimumUptime",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaVmMonitoring(val *string) {
	if err := j.validateSetHaVmMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haVmMonitoring",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaVmRestartAdditionalDelay(val *float64) {
	if err := j.validateSetHaVmRestartAdditionalDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haVmRestartAdditionalDelay",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaVmRestartPriority(val *string) {
	if err := j.validateSetHaVmRestartPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haVmRestartPriority",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHaVmRestartTimeout(val *float64) {
	if err := j.validateSetHaVmRestartTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haVmRestartTimeout",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHostClusterExitTimeout(val *float64) {
	if err := j.validateSetHostClusterExitTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostClusterExitTimeout",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHostManaged(val interface{}) {
	if err := j.validateSetHostManagedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostManaged",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetHostSystemIds(val *[]*string) {
	if err := j.validateSetHostSystemIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostSystemIds",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetProactiveHaAutomationLevel(val *string) {
	if err := j.validateSetProactiveHaAutomationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proactiveHaAutomationLevel",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetProactiveHaEnabled(val interface{}) {
	if err := j.validateSetProactiveHaEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proactiveHaEnabled",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetProactiveHaModerateRemediation(val *string) {
	if err := j.validateSetProactiveHaModerateRemediationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proactiveHaModerateRemediation",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetProactiveHaProviderIds(val *[]*string) {
	if err := j.validateSetProactiveHaProviderIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proactiveHaProviderIds",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetProactiveHaSevereRemediation(val *string) {
	if err := j.validateSetProactiveHaSevereRemediationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proactiveHaSevereRemediation",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetVsanCompressionEnabled(val interface{}) {
	if err := j.validateSetVsanCompressionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vsanCompressionEnabled",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetVsanDedupEnabled(val interface{}) {
	if err := j.validateSetVsanDedupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vsanDedupEnabled",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetVsanDitEncryptionEnabled(val interface{}) {
	if err := j.validateSetVsanDitEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vsanDitEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetVsanDitRekeyInterval(val *float64) {
	if err := j.validateSetVsanDitRekeyIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vsanDitRekeyInterval",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetVsanEnabled(val interface{}) {
	if err := j.validateSetVsanEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vsanEnabled",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetVsanEsaEnabled(val interface{}) {
	if err := j.validateSetVsanEsaEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vsanEsaEnabled",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetVsanNetworkDiagnosticModeEnabled(val interface{}) {
	if err := j.validateSetVsanNetworkDiagnosticModeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vsanNetworkDiagnosticModeEnabled",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetVsanPerformanceEnabled(val interface{}) {
	if err := j.validateSetVsanPerformanceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vsanPerformanceEnabled",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetVsanRemoteDatastoreIds(val *[]*string) {
	if err := j.validateSetVsanRemoteDatastoreIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vsanRemoteDatastoreIds",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetVsanUnmapEnabled(val interface{}) {
	if err := j.validateSetVsanUnmapEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vsanUnmapEnabled",
		val,
	)
}

func (j *jsiiProxy_ComputeCluster)SetVsanVerboseModeEnabled(val interface{}) {
	if err := j.validateSetVsanVerboseModeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vsanVerboseModeEnabled",
		val,
	)
}

// Generates CDKTF code for importing a ComputeCluster resource upon running "cdktf plan <stack-name>".
func ComputeCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateComputeCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.computeCluster.ComputeCluster",
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
func ComputeCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.computeCluster.ComputeCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.computeCluster.ComputeCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.computeCluster.ComputeCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vsphere.computeCluster.ComputeCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeCluster) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ComputeCluster) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ComputeCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeCluster) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ComputeCluster) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeCluster) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeCluster) PutHostImage(value *ComputeClusterHostImage) {
	if err := c.validatePutHostImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHostImage",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeCluster) PutVsanDiskGroup(value interface{}) {
	if err := c.validatePutVsanDiskGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVsanDiskGroup",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeCluster) PutVsanFaultDomains(value interface{}) {
	if err := c.validatePutVsanFaultDomainsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVsanFaultDomains",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeCluster) PutVsanStretchedCluster(value *ComputeClusterVsanStretchedCluster) {
	if err := c.validatePutVsanStretchedClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVsanStretchedCluster",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeCluster) ResetCustomAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetDpmAutomationLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetDpmAutomationLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetDpmEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetDpmEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetDpmThreshold() {
	_jsii_.InvokeVoid(
		c,
		"resetDpmThreshold",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetDrsAdvancedOptions() {
	_jsii_.InvokeVoid(
		c,
		"resetDrsAdvancedOptions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetDrsAutomationLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetDrsAutomationLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetDrsEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetDrsEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetDrsEnablePredictiveDrs() {
	_jsii_.InvokeVoid(
		c,
		"resetDrsEnablePredictiveDrs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetDrsEnableVmOverrides() {
	_jsii_.InvokeVoid(
		c,
		"resetDrsEnableVmOverrides",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetDrsMigrationThreshold() {
	_jsii_.InvokeVoid(
		c,
		"resetDrsMigrationThreshold",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetDrsScaleDescendantsShares() {
	_jsii_.InvokeVoid(
		c,
		"resetDrsScaleDescendantsShares",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetFolder() {
	_jsii_.InvokeVoid(
		c,
		"resetFolder",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetForceEvacuateOnDestroy() {
	_jsii_.InvokeVoid(
		c,
		"resetForceEvacuateOnDestroy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaAdmissionControlFailoverHostSystemIds() {
	_jsii_.InvokeVoid(
		c,
		"resetHaAdmissionControlFailoverHostSystemIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaAdmissionControlHostFailureTolerance() {
	_jsii_.InvokeVoid(
		c,
		"resetHaAdmissionControlHostFailureTolerance",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaAdmissionControlPerformanceTolerance() {
	_jsii_.InvokeVoid(
		c,
		"resetHaAdmissionControlPerformanceTolerance",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaAdmissionControlPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetHaAdmissionControlPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaAdmissionControlResourcePercentageAutoCompute() {
	_jsii_.InvokeVoid(
		c,
		"resetHaAdmissionControlResourcePercentageAutoCompute",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaAdmissionControlResourcePercentageCpu() {
	_jsii_.InvokeVoid(
		c,
		"resetHaAdmissionControlResourcePercentageCpu",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaAdmissionControlResourcePercentageMemory() {
	_jsii_.InvokeVoid(
		c,
		"resetHaAdmissionControlResourcePercentageMemory",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaAdmissionControlSlotPolicyExplicitCpu() {
	_jsii_.InvokeVoid(
		c,
		"resetHaAdmissionControlSlotPolicyExplicitCpu",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaAdmissionControlSlotPolicyExplicitMemory() {
	_jsii_.InvokeVoid(
		c,
		"resetHaAdmissionControlSlotPolicyExplicitMemory",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaAdmissionControlSlotPolicyUseExplicitSize() {
	_jsii_.InvokeVoid(
		c,
		"resetHaAdmissionControlSlotPolicyUseExplicitSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaAdvancedOptions() {
	_jsii_.InvokeVoid(
		c,
		"resetHaAdvancedOptions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaDatastoreApdRecoveryAction() {
	_jsii_.InvokeVoid(
		c,
		"resetHaDatastoreApdRecoveryAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaDatastoreApdResponse() {
	_jsii_.InvokeVoid(
		c,
		"resetHaDatastoreApdResponse",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaDatastoreApdResponseDelay() {
	_jsii_.InvokeVoid(
		c,
		"resetHaDatastoreApdResponseDelay",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaDatastorePdlResponse() {
	_jsii_.InvokeVoid(
		c,
		"resetHaDatastorePdlResponse",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetHaEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaHeartbeatDatastoreIds() {
	_jsii_.InvokeVoid(
		c,
		"resetHaHeartbeatDatastoreIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaHeartbeatDatastorePolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetHaHeartbeatDatastorePolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaHostIsolationResponse() {
	_jsii_.InvokeVoid(
		c,
		"resetHaHostIsolationResponse",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaHostMonitoring() {
	_jsii_.InvokeVoid(
		c,
		"resetHaHostMonitoring",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaVmComponentProtection() {
	_jsii_.InvokeVoid(
		c,
		"resetHaVmComponentProtection",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaVmDependencyRestartCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetHaVmDependencyRestartCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaVmFailureInterval() {
	_jsii_.InvokeVoid(
		c,
		"resetHaVmFailureInterval",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaVmMaximumFailureWindow() {
	_jsii_.InvokeVoid(
		c,
		"resetHaVmMaximumFailureWindow",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaVmMaximumResets() {
	_jsii_.InvokeVoid(
		c,
		"resetHaVmMaximumResets",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaVmMinimumUptime() {
	_jsii_.InvokeVoid(
		c,
		"resetHaVmMinimumUptime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaVmMonitoring() {
	_jsii_.InvokeVoid(
		c,
		"resetHaVmMonitoring",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaVmRestartAdditionalDelay() {
	_jsii_.InvokeVoid(
		c,
		"resetHaVmRestartAdditionalDelay",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaVmRestartPriority() {
	_jsii_.InvokeVoid(
		c,
		"resetHaVmRestartPriority",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHaVmRestartTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetHaVmRestartTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHostClusterExitTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetHostClusterExitTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHostImage() {
	_jsii_.InvokeVoid(
		c,
		"resetHostImage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHostManaged() {
	_jsii_.InvokeVoid(
		c,
		"resetHostManaged",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetHostSystemIds() {
	_jsii_.InvokeVoid(
		c,
		"resetHostSystemIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetProactiveHaAutomationLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetProactiveHaAutomationLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetProactiveHaEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetProactiveHaEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetProactiveHaModerateRemediation() {
	_jsii_.InvokeVoid(
		c,
		"resetProactiveHaModerateRemediation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetProactiveHaProviderIds() {
	_jsii_.InvokeVoid(
		c,
		"resetProactiveHaProviderIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetProactiveHaSevereRemediation() {
	_jsii_.InvokeVoid(
		c,
		"resetProactiveHaSevereRemediation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetVsanCompressionEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetVsanCompressionEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetVsanDedupEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetVsanDedupEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetVsanDiskGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetVsanDiskGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetVsanDitEncryptionEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetVsanDitEncryptionEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetVsanDitRekeyInterval() {
	_jsii_.InvokeVoid(
		c,
		"resetVsanDitRekeyInterval",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetVsanEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetVsanEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetVsanEsaEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetVsanEsaEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetVsanFaultDomains() {
	_jsii_.InvokeVoid(
		c,
		"resetVsanFaultDomains",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetVsanNetworkDiagnosticModeEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetVsanNetworkDiagnosticModeEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetVsanPerformanceEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetVsanPerformanceEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetVsanRemoteDatastoreIds() {
	_jsii_.InvokeVoid(
		c,
		"resetVsanRemoteDatastoreIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetVsanStretchedCluster() {
	_jsii_.InvokeVoid(
		c,
		"resetVsanStretchedCluster",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetVsanUnmapEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetVsanUnmapEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) ResetVsanVerboseModeEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetVsanVerboseModeEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

