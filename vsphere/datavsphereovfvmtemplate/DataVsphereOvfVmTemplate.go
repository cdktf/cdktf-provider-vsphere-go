// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datavsphereovfvmtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v9/datavsphereovfvmtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.12.0/docs/data-sources/ovf_vm_template vsphere_ovf_vm_template}.
type DataVsphereOvfVmTemplate interface {
	cdktf.TerraformDataSource
	AllowUnverifiedSslCert() interface{}
	SetAllowUnverifiedSslCert(val interface{})
	AllowUnverifiedSslCertInput() interface{}
	AlternateGuestName() *string
	Annotation() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CpuHotAddEnabled() cdktf.IResolvable
	CpuHotRemoveEnabled() cdktf.IResolvable
	CpuPerformanceCountersEnabled() cdktf.IResolvable
	DatastoreId() *string
	SetDatastoreId(val *string)
	DatastoreIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeploymentOption() *string
	SetDeploymentOption(val *string)
	DeploymentOptionInput() *string
	DiskProvisioning() *string
	SetDiskProvisioning(val *string)
	DiskProvisioningInput() *string
	EnableHiddenProperties() interface{}
	SetEnableHiddenProperties(val interface{})
	EnableHiddenPropertiesInput() interface{}
	Firmware() *string
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
	GuestId() *string
	HostSystemId() *string
	SetHostSystemId(val *string)
	HostSystemIdInput() *string
	Id() *string
	SetId(val *string)
	IdeControllerCount() *float64
	IdInput() *string
	IpAllocationPolicy() *string
	SetIpAllocationPolicy(val *string)
	IpAllocationPolicyInput() *string
	IpProtocol() *string
	SetIpProtocol(val *string)
	IpProtocolInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalOvfPath() *string
	SetLocalOvfPath(val *string)
	LocalOvfPathInput() *string
	Memory() *float64
	MemoryHotAddEnabled() cdktf.IResolvable
	Name() *string
	SetName(val *string)
	NameInput() *string
	NestedHvEnabled() cdktf.IResolvable
	// The tree node.
	Node() constructs.Node
	NumCoresPerSocket() *float64
	NumCpus() *float64
	OvfNetworkMap() *map[string]*string
	SetOvfNetworkMap(val *map[string]*string)
	OvfNetworkMapInput() *map[string]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RemoteOvfUrl() *string
	SetRemoteOvfUrl(val *string)
	RemoteOvfUrlInput() *string
	ResourcePoolId() *string
	SetResourcePoolId(val *string)
	ResourcePoolIdInput() *string
	SataControllerCount() *float64
	ScsiControllerCount() *float64
	ScsiType() *string
	SwapPlacementPolicy() *string
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
	ResetAllowUnverifiedSslCert()
	ResetDatastoreId()
	ResetDeploymentOption()
	ResetDiskProvisioning()
	ResetEnableHiddenProperties()
	ResetFolder()
	ResetId()
	ResetIpAllocationPolicy()
	ResetIpProtocol()
	ResetLocalOvfPath()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOvfNetworkMap()
	ResetRemoteOvfUrl()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataVsphereOvfVmTemplate
type jsiiProxy_DataVsphereOvfVmTemplate struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) AllowUnverifiedSslCert() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUnverifiedSslCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) AllowUnverifiedSslCertInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUnverifiedSslCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) AlternateGuestName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alternateGuestName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) Annotation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) CpuHotAddEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"cpuHotAddEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) CpuHotRemoveEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"cpuHotRemoveEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) CpuPerformanceCountersEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"cpuPerformanceCountersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) DatastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) DatastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) DeploymentOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) DeploymentOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) DiskProvisioning() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskProvisioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) DiskProvisioningInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskProvisioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) EnableHiddenProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHiddenProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) EnableHiddenPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHiddenPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) Firmware() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firmware",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) Folder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) FolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) GuestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) HostSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) HostSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) IdeControllerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ideControllerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) IpAllocationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAllocationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) IpAllocationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAllocationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) IpProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) IpProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) LocalOvfPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localOvfPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) LocalOvfPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localOvfPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) MemoryHotAddEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"memoryHotAddEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) NestedHvEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"nestedHvEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) NumCoresPerSocket() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCoresPerSocket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) NumCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) OvfNetworkMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"ovfNetworkMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) OvfNetworkMapInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"ovfNetworkMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) RemoteOvfUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteOvfUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) RemoteOvfUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteOvfUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) ResourcePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) ResourcePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) SataControllerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sataControllerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) ScsiControllerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scsiControllerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) ScsiType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scsiType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) SwapPlacementPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"swapPlacementPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.12.0/docs/data-sources/ovf_vm_template vsphere_ovf_vm_template} Data Source.
func NewDataVsphereOvfVmTemplate(scope constructs.Construct, id *string, config *DataVsphereOvfVmTemplateConfig) DataVsphereOvfVmTemplate {
	_init_.Initialize()

	if err := validateNewDataVsphereOvfVmTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVsphereOvfVmTemplate{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.dataVsphereOvfVmTemplate.DataVsphereOvfVmTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.12.0/docs/data-sources/ovf_vm_template vsphere_ovf_vm_template} Data Source.
func NewDataVsphereOvfVmTemplate_Override(d DataVsphereOvfVmTemplate, scope constructs.Construct, id *string, config *DataVsphereOvfVmTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.dataVsphereOvfVmTemplate.DataVsphereOvfVmTemplate",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetAllowUnverifiedSslCert(val interface{}) {
	if err := j.validateSetAllowUnverifiedSslCertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowUnverifiedSslCert",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetDatastoreId(val *string) {
	if err := j.validateSetDatastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datastoreId",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetDeploymentOption(val *string) {
	if err := j.validateSetDeploymentOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentOption",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetDiskProvisioning(val *string) {
	if err := j.validateSetDiskProvisioningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskProvisioning",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetEnableHiddenProperties(val interface{}) {
	if err := j.validateSetEnableHiddenPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHiddenProperties",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetFolder(val *string) {
	if err := j.validateSetFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"folder",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetHostSystemId(val *string) {
	if err := j.validateSetHostSystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostSystemId",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetIpAllocationPolicy(val *string) {
	if err := j.validateSetIpAllocationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAllocationPolicy",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetIpProtocol(val *string) {
	if err := j.validateSetIpProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipProtocol",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetLocalOvfPath(val *string) {
	if err := j.validateSetLocalOvfPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localOvfPath",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetOvfNetworkMap(val *map[string]*string) {
	if err := j.validateSetOvfNetworkMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ovfNetworkMap",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetRemoteOvfUrl(val *string) {
	if err := j.validateSetRemoteOvfUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteOvfUrl",
		val,
	)
}

func (j *jsiiProxy_DataVsphereOvfVmTemplate)SetResourcePoolId(val *string) {
	if err := j.validateSetResourcePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePoolId",
		val,
	)
}

// Generates CDKTF code for importing a DataVsphereOvfVmTemplate resource upon running "cdktf plan <stack-name>".
func DataVsphereOvfVmTemplate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataVsphereOvfVmTemplate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.dataVsphereOvfVmTemplate.DataVsphereOvfVmTemplate",
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
func DataVsphereOvfVmTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVsphereOvfVmTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.dataVsphereOvfVmTemplate.DataVsphereOvfVmTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVsphereOvfVmTemplate_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVsphereOvfVmTemplate_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.dataVsphereOvfVmTemplate.DataVsphereOvfVmTemplate",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVsphereOvfVmTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVsphereOvfVmTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.dataVsphereOvfVmTemplate.DataVsphereOvfVmTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVsphereOvfVmTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vsphere.dataVsphereOvfVmTemplate.DataVsphereOvfVmTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataVsphereOvfVmTemplate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataVsphereOvfVmTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataVsphereOvfVmTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataVsphereOvfVmTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataVsphereOvfVmTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataVsphereOvfVmTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataVsphereOvfVmTemplate) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataVsphereOvfVmTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataVsphereOvfVmTemplate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataVsphereOvfVmTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) ResetAllowUnverifiedSslCert() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowUnverifiedSslCert",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) ResetDatastoreId() {
	_jsii_.InvokeVoid(
		d,
		"resetDatastoreId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) ResetDeploymentOption() {
	_jsii_.InvokeVoid(
		d,
		"resetDeploymentOption",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) ResetDiskProvisioning() {
	_jsii_.InvokeVoid(
		d,
		"resetDiskProvisioning",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) ResetEnableHiddenProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableHiddenProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) ResetFolder() {
	_jsii_.InvokeVoid(
		d,
		"resetFolder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) ResetIpAllocationPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetIpAllocationPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) ResetIpProtocol() {
	_jsii_.InvokeVoid(
		d,
		"resetIpProtocol",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) ResetLocalOvfPath() {
	_jsii_.InvokeVoid(
		d,
		"resetLocalOvfPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) ResetOvfNetworkMap() {
	_jsii_.InvokeVoid(
		d,
		"resetOvfNetworkMap",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) ResetRemoteOvfUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetRemoteOvfUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVsphereOvfVmTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

