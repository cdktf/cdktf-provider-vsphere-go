// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v8/jsii"

	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v8/virtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineCloneCustomizeNetworkInterfaceOutputReference interface {
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
	DnsDomain() *string
	SetDnsDomain(val *string)
	DnsDomainInput() *string
	DnsServerList() *[]*string
	SetDnsServerList(val *[]*string)
	DnsServerListInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ipv4Address() *string
	SetIpv4Address(val *string)
	Ipv4AddressInput() *string
	Ipv4Netmask() *float64
	SetIpv4Netmask(val *float64)
	Ipv4NetmaskInput() *float64
	Ipv6Address() *string
	SetIpv6Address(val *string)
	Ipv6AddressInput() *string
	Ipv6Netmask() *float64
	SetIpv6Netmask(val *float64)
	Ipv6NetmaskInput() *float64
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
	ResetDnsDomain()
	ResetDnsServerList()
	ResetIpv4Address()
	ResetIpv4Netmask()
	ResetIpv6Address()
	ResetIpv6Netmask()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualMachineCloneCustomizeNetworkInterfaceOutputReference
type jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) DnsDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) DnsDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) DnsServerList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServerList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) DnsServerListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServerListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) Ipv4Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) Ipv4AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) Ipv4Netmask() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4Netmask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) Ipv4NetmaskInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4NetmaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) Ipv6Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) Ipv6AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) Ipv6Netmask() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6Netmask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) Ipv6NetmaskInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6NetmaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualMachineCloneCustomizeNetworkInterfaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VirtualMachineCloneCustomizeNetworkInterfaceOutputReference {
	_init_.Initialize()

	if err := validateNewVirtualMachineCloneCustomizeNetworkInterfaceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachineCloneCustomizeNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVirtualMachineCloneCustomizeNetworkInterfaceOutputReference_Override(v VirtualMachineCloneCustomizeNetworkInterfaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachineCloneCustomizeNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference)SetDnsDomain(val *string) {
	if err := j.validateSetDnsDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsDomain",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference)SetDnsServerList(val *[]*string) {
	if err := j.validateSetDnsServerListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServerList",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference)SetIpv4Address(val *string) {
	if err := j.validateSetIpv4AddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Address",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference)SetIpv4Netmask(val *float64) {
	if err := j.validateSetIpv4NetmaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Netmask",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference)SetIpv6Address(val *string) {
	if err := j.validateSetIpv6AddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Address",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference)SetIpv6Netmask(val *float64) {
	if err := j.validateSetIpv6NetmaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Netmask",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) ResetDnsDomain() {
	_jsii_.InvokeVoid(
		v,
		"resetDnsDomain",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) ResetDnsServerList() {
	_jsii_.InvokeVoid(
		v,
		"resetDnsServerList",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) ResetIpv4Address() {
	_jsii_.InvokeVoid(
		v,
		"resetIpv4Address",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) ResetIpv4Netmask() {
	_jsii_.InvokeVoid(
		v,
		"resetIpv4Netmask",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) ResetIpv6Address() {
	_jsii_.InvokeVoid(
		v,
		"resetIpv6Address",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) ResetIpv6Netmask() {
	_jsii_.InvokeVoid(
		v,
		"resetIpv6Netmask",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

