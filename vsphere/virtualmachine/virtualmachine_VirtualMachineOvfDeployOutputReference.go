package virtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v2/jsii"

	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v2/virtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineOvfDeployOutputReference interface {
	cdktf.ComplexObject
	AllowUnverifiedSslCert() interface{}
	SetAllowUnverifiedSslCert(val interface{})
	AllowUnverifiedSslCertInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeploymentOption() *string
	SetDeploymentOption(val *string)
	DeploymentOptionInput() *string
	DiskProvisioning() *string
	SetDiskProvisioning(val *string)
	DiskProvisioningInput() *string
	EnableHiddenProperties() interface{}
	SetEnableHiddenProperties(val interface{})
	EnableHiddenPropertiesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *VirtualMachineOvfDeploy
	SetInternalValue(val *VirtualMachineOvfDeploy)
	IpAllocationPolicy() *string
	SetIpAllocationPolicy(val *string)
	IpAllocationPolicyInput() *string
	IpProtocol() *string
	SetIpProtocol(val *string)
	IpProtocolInput() *string
	LocalOvfPath() *string
	SetLocalOvfPath(val *string)
	LocalOvfPathInput() *string
	OvfNetworkMap() *map[string]*string
	SetOvfNetworkMap(val *map[string]*string)
	OvfNetworkMapInput() *map[string]*string
	RemoteOvfUrl() *string
	SetRemoteOvfUrl(val *string)
	RemoteOvfUrlInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAllowUnverifiedSslCert()
	ResetDeploymentOption()
	ResetDiskProvisioning()
	ResetEnableHiddenProperties()
	ResetIpAllocationPolicy()
	ResetIpProtocol()
	ResetLocalOvfPath()
	ResetOvfNetworkMap()
	ResetRemoteOvfUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualMachineOvfDeployOutputReference
type jsiiProxy_VirtualMachineOvfDeployOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) AllowUnverifiedSslCert() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUnverifiedSslCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) AllowUnverifiedSslCertInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUnverifiedSslCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) DeploymentOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) DeploymentOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) DiskProvisioning() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskProvisioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) DiskProvisioningInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskProvisioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) EnableHiddenProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHiddenProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) EnableHiddenPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHiddenPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) InternalValue() *VirtualMachineOvfDeploy {
	var returns *VirtualMachineOvfDeploy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) IpAllocationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAllocationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) IpAllocationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAllocationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) IpProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) IpProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) LocalOvfPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localOvfPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) LocalOvfPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localOvfPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) OvfNetworkMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"ovfNetworkMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) OvfNetworkMapInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"ovfNetworkMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) RemoteOvfUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteOvfUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) RemoteOvfUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteOvfUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualMachineOvfDeployOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualMachineOvfDeployOutputReference {
	_init_.Initialize()

	if err := validateNewVirtualMachineOvfDeployOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachineOvfDeployOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachineOvfDeployOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualMachineOvfDeployOutputReference_Override(v VirtualMachineOvfDeployOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachineOvfDeployOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference)SetAllowUnverifiedSslCert(val interface{}) {
	if err := j.validateSetAllowUnverifiedSslCertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowUnverifiedSslCert",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference)SetDeploymentOption(val *string) {
	if err := j.validateSetDeploymentOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentOption",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference)SetDiskProvisioning(val *string) {
	if err := j.validateSetDiskProvisioningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskProvisioning",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference)SetEnableHiddenProperties(val interface{}) {
	if err := j.validateSetEnableHiddenPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHiddenProperties",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference)SetInternalValue(val *VirtualMachineOvfDeploy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference)SetIpAllocationPolicy(val *string) {
	if err := j.validateSetIpAllocationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAllocationPolicy",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference)SetIpProtocol(val *string) {
	if err := j.validateSetIpProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipProtocol",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference)SetLocalOvfPath(val *string) {
	if err := j.validateSetLocalOvfPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localOvfPath",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference)SetOvfNetworkMap(val *map[string]*string) {
	if err := j.validateSetOvfNetworkMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ovfNetworkMap",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference)SetRemoteOvfUrl(val *string) {
	if err := j.validateSetRemoteOvfUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteOvfUrl",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineOvfDeployOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) ResetAllowUnverifiedSslCert() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowUnverifiedSslCert",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) ResetDeploymentOption() {
	_jsii_.InvokeVoid(
		v,
		"resetDeploymentOption",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) ResetDiskProvisioning() {
	_jsii_.InvokeVoid(
		v,
		"resetDiskProvisioning",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) ResetEnableHiddenProperties() {
	_jsii_.InvokeVoid(
		v,
		"resetEnableHiddenProperties",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) ResetIpAllocationPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetIpAllocationPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) ResetIpProtocol() {
	_jsii_.InvokeVoid(
		v,
		"resetIpProtocol",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) ResetLocalOvfPath() {
	_jsii_.InvokeVoid(
		v,
		"resetLocalOvfPath",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) ResetOvfNetworkMap() {
	_jsii_.InvokeVoid(
		v,
		"resetOvfNetworkMap",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) ResetRemoteOvfUrl() {
	_jsii_.InvokeVoid(
		v,
		"resetRemoteOvfUrl",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineOvfDeployOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

