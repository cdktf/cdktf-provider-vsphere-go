// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package guestoscustomization

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v10/jsii"

	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v10/guestoscustomization/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GuestOsCustomizationSpecNetworkInterfaceOutputReference interface {
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

// The jsii proxy struct for GuestOsCustomizationSpecNetworkInterfaceOutputReference
type jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) DnsDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) DnsDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) DnsServerList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServerList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) DnsServerListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServerListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) Ipv4Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) Ipv4AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) Ipv4Netmask() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4Netmask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) Ipv4NetmaskInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4NetmaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) Ipv6Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) Ipv6AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) Ipv6Netmask() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6Netmask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) Ipv6NetmaskInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6NetmaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGuestOsCustomizationSpecNetworkInterfaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GuestOsCustomizationSpecNetworkInterfaceOutputReference {
	_init_.Initialize()

	if err := validateNewGuestOsCustomizationSpecNetworkInterfaceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.guestOsCustomization.GuestOsCustomizationSpecNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGuestOsCustomizationSpecNetworkInterfaceOutputReference_Override(g GuestOsCustomizationSpecNetworkInterfaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.guestOsCustomization.GuestOsCustomizationSpecNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference)SetDnsDomain(val *string) {
	if err := j.validateSetDnsDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsDomain",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference)SetDnsServerList(val *[]*string) {
	if err := j.validateSetDnsServerListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServerList",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference)SetIpv4Address(val *string) {
	if err := j.validateSetIpv4AddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Address",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference)SetIpv4Netmask(val *float64) {
	if err := j.validateSetIpv4NetmaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Netmask",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference)SetIpv6Address(val *string) {
	if err := j.validateSetIpv6AddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Address",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference)SetIpv6Netmask(val *float64) {
	if err := j.validateSetIpv6NetmaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Netmask",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) ResetDnsDomain() {
	_jsii_.InvokeVoid(
		g,
		"resetDnsDomain",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) ResetDnsServerList() {
	_jsii_.InvokeVoid(
		g,
		"resetDnsServerList",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) ResetIpv4Address() {
	_jsii_.InvokeVoid(
		g,
		"resetIpv4Address",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) ResetIpv4Netmask() {
	_jsii_.InvokeVoid(
		g,
		"resetIpv4Netmask",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) ResetIpv6Address() {
	_jsii_.InvokeVoid(
		g,
		"resetIpv6Address",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) ResetIpv6Netmask() {
	_jsii_.InvokeVoid(
		g,
		"resetIpv6Netmask",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

