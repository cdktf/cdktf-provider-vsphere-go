package vnic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v3/jsii"

	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v3/vnic/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VnicIpv6OutputReference interface {
	cdktf.ComplexObject
	Addresses() *[]*string
	SetAddresses(val *[]*string)
	AddressesInput() *[]*string
	Autoconfig() interface{}
	SetAutoconfig(val interface{})
	AutoconfigInput() interface{}
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
	Dhcp() interface{}
	SetDhcp(val interface{})
	DhcpInput() interface{}
	// Experimental.
	Fqn() *string
	Gw() *string
	SetGw(val *string)
	GwInput() *string
	InternalValue() *VnicIpv6
	SetInternalValue(val *VnicIpv6)
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
	ResetAddresses()
	ResetAutoconfig()
	ResetDhcp()
	ResetGw()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VnicIpv6OutputReference
type jsiiProxy_VnicIpv6OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VnicIpv6OutputReference) Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VnicIpv6OutputReference) AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VnicIpv6OutputReference) Autoconfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoconfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VnicIpv6OutputReference) AutoconfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoconfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VnicIpv6OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VnicIpv6OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VnicIpv6OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VnicIpv6OutputReference) Dhcp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VnicIpv6OutputReference) DhcpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VnicIpv6OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VnicIpv6OutputReference) Gw() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VnicIpv6OutputReference) GwInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gwInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VnicIpv6OutputReference) InternalValue() *VnicIpv6 {
	var returns *VnicIpv6
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VnicIpv6OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VnicIpv6OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVnicIpv6OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VnicIpv6OutputReference {
	_init_.Initialize()

	if err := validateNewVnicIpv6OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VnicIpv6OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.vnic.VnicIpv6OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVnicIpv6OutputReference_Override(v VnicIpv6OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.vnic.VnicIpv6OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VnicIpv6OutputReference)SetAddresses(val *[]*string) {
	if err := j.validateSetAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addresses",
		val,
	)
}

func (j *jsiiProxy_VnicIpv6OutputReference)SetAutoconfig(val interface{}) {
	if err := j.validateSetAutoconfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoconfig",
		val,
	)
}

func (j *jsiiProxy_VnicIpv6OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VnicIpv6OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VnicIpv6OutputReference)SetDhcp(val interface{}) {
	if err := j.validateSetDhcpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dhcp",
		val,
	)
}

func (j *jsiiProxy_VnicIpv6OutputReference)SetGw(val *string) {
	if err := j.validateSetGwParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gw",
		val,
	)
}

func (j *jsiiProxy_VnicIpv6OutputReference)SetInternalValue(val *VnicIpv6) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VnicIpv6OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VnicIpv6OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VnicIpv6OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VnicIpv6OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VnicIpv6OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VnicIpv6OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VnicIpv6OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VnicIpv6OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VnicIpv6OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VnicIpv6OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VnicIpv6OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VnicIpv6OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VnicIpv6OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VnicIpv6OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VnicIpv6OutputReference) ResetAddresses() {
	_jsii_.InvokeVoid(
		v,
		"resetAddresses",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VnicIpv6OutputReference) ResetAutoconfig() {
	_jsii_.InvokeVoid(
		v,
		"resetAutoconfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VnicIpv6OutputReference) ResetDhcp() {
	_jsii_.InvokeVoid(
		v,
		"resetDhcp",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VnicIpv6OutputReference) ResetGw() {
	_jsii_.InvokeVoid(
		v,
		"resetGw",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VnicIpv6OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VnicIpv6OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

