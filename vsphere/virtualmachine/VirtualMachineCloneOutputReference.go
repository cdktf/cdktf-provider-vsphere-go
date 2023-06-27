package virtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v6/jsii"

	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v6/virtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineCloneOutputReference interface {
	cdktf.ComplexObject
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
	Customize() VirtualMachineCloneCustomizeOutputReference
	CustomizeInput() *VirtualMachineCloneCustomize
	// Experimental.
	Fqn() *string
	InternalValue() *VirtualMachineClone
	SetInternalValue(val *VirtualMachineClone)
	LinkedClone() interface{}
	SetLinkedClone(val interface{})
	LinkedCloneInput() interface{}
	OvfNetworkMap() *map[string]*string
	SetOvfNetworkMap(val *map[string]*string)
	OvfNetworkMapInput() *map[string]*string
	OvfStorageMap() *map[string]*string
	SetOvfStorageMap(val *map[string]*string)
	OvfStorageMapInput() *map[string]*string
	TemplateUuid() *string
	SetTemplateUuid(val *string)
	TemplateUuidInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
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
	PutCustomize(value *VirtualMachineCloneCustomize)
	ResetCustomize()
	ResetLinkedClone()
	ResetOvfNetworkMap()
	ResetOvfStorageMap()
	ResetTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualMachineCloneOutputReference
type jsiiProxy_VirtualMachineCloneOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) Customize() VirtualMachineCloneCustomizeOutputReference {
	var returns VirtualMachineCloneCustomizeOutputReference
	_jsii_.Get(
		j,
		"customize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) CustomizeInput() *VirtualMachineCloneCustomize {
	var returns *VirtualMachineCloneCustomize
	_jsii_.Get(
		j,
		"customizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) InternalValue() *VirtualMachineClone {
	var returns *VirtualMachineClone
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) LinkedClone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linkedClone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) LinkedCloneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linkedCloneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) OvfNetworkMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"ovfNetworkMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) OvfNetworkMapInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"ovfNetworkMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) OvfStorageMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"ovfStorageMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) OvfStorageMapInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"ovfStorageMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) TemplateUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) TemplateUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}


func NewVirtualMachineCloneOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualMachineCloneOutputReference {
	_init_.Initialize()

	if err := validateNewVirtualMachineCloneOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachineCloneOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachineCloneOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualMachineCloneOutputReference_Override(v VirtualMachineCloneOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachineCloneOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference)SetInternalValue(val *VirtualMachineClone) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference)SetLinkedClone(val interface{}) {
	if err := j.validateSetLinkedCloneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkedClone",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference)SetOvfNetworkMap(val *map[string]*string) {
	if err := j.validateSetOvfNetworkMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ovfNetworkMap",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference)SetOvfStorageMap(val *map[string]*string) {
	if err := j.validateSetOvfStorageMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ovfStorageMap",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference)SetTemplateUuid(val *string) {
	if err := j.validateSetTemplateUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateUuid",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneOutputReference)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (v *jsiiProxy_VirtualMachineCloneOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineCloneOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualMachineCloneOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineCloneOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualMachineCloneOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualMachineCloneOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualMachineCloneOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualMachineCloneOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualMachineCloneOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualMachineCloneOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualMachineCloneOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineCloneOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineCloneOutputReference) PutCustomize(value *VirtualMachineCloneCustomize) {
	if err := v.validatePutCustomizeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putCustomize",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineCloneOutputReference) ResetCustomize() {
	_jsii_.InvokeVoid(
		v,
		"resetCustomize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneOutputReference) ResetLinkedClone() {
	_jsii_.InvokeVoid(
		v,
		"resetLinkedClone",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneOutputReference) ResetOvfNetworkMap() {
	_jsii_.InvokeVoid(
		v,
		"resetOvfNetworkMap",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneOutputReference) ResetOvfStorageMap() {
	_jsii_.InvokeVoid(
		v,
		"resetOvfStorageMap",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeout",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VirtualMachineCloneOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

