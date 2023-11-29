// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package guestoscustomization

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v8/jsii"

	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v8/guestoscustomization/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GuestOsCustomizationSpecOutputReference interface {
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
	DnsServerList() *[]*string
	SetDnsServerList(val *[]*string)
	DnsServerListInput() *[]*string
	DnsSuffixList() *[]*string
	SetDnsSuffixList(val *[]*string)
	DnsSuffixListInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GuestOsCustomizationSpec
	SetInternalValue(val *GuestOsCustomizationSpec)
	Ipv4Gateway() *string
	SetIpv4Gateway(val *string)
	Ipv4GatewayInput() *string
	Ipv6Gateway() *string
	SetIpv6Gateway(val *string)
	Ipv6GatewayInput() *string
	LinuxOptions() GuestOsCustomizationSpecLinuxOptionsOutputReference
	LinuxOptionsInput() *GuestOsCustomizationSpecLinuxOptions
	NetworkInterface() GuestOsCustomizationSpecNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WindowsOptions() GuestOsCustomizationSpecWindowsOptionsOutputReference
	WindowsOptionsInput() *GuestOsCustomizationSpecWindowsOptions
	WindowsSysprepText() *string
	SetWindowsSysprepText(val *string)
	WindowsSysprepTextInput() *string
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
	PutLinuxOptions(value *GuestOsCustomizationSpecLinuxOptions)
	PutNetworkInterface(value interface{})
	PutWindowsOptions(value *GuestOsCustomizationSpecWindowsOptions)
	ResetDnsServerList()
	ResetDnsSuffixList()
	ResetIpv4Gateway()
	ResetIpv6Gateway()
	ResetLinuxOptions()
	ResetNetworkInterface()
	ResetWindowsOptions()
	ResetWindowsSysprepText()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GuestOsCustomizationSpecOutputReference
type jsiiProxy_GuestOsCustomizationSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) DnsServerList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServerList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) DnsServerListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServerListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) DnsSuffixList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsSuffixList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) DnsSuffixListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsSuffixListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) InternalValue() *GuestOsCustomizationSpec {
	var returns *GuestOsCustomizationSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) Ipv4Gateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4Gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) Ipv4GatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4GatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) Ipv6Gateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) Ipv6GatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6GatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) LinuxOptions() GuestOsCustomizationSpecLinuxOptionsOutputReference {
	var returns GuestOsCustomizationSpecLinuxOptionsOutputReference
	_jsii_.Get(
		j,
		"linuxOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) LinuxOptionsInput() *GuestOsCustomizationSpecLinuxOptions {
	var returns *GuestOsCustomizationSpecLinuxOptions
	_jsii_.Get(
		j,
		"linuxOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) NetworkInterface() GuestOsCustomizationSpecNetworkInterfaceList {
	var returns GuestOsCustomizationSpecNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) WindowsOptions() GuestOsCustomizationSpecWindowsOptionsOutputReference {
	var returns GuestOsCustomizationSpecWindowsOptionsOutputReference
	_jsii_.Get(
		j,
		"windowsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) WindowsOptionsInput() *GuestOsCustomizationSpecWindowsOptions {
	var returns *GuestOsCustomizationSpecWindowsOptions
	_jsii_.Get(
		j,
		"windowsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) WindowsSysprepText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowsSysprepText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference) WindowsSysprepTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowsSysprepTextInput",
		&returns,
	)
	return returns
}


func NewGuestOsCustomizationSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GuestOsCustomizationSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGuestOsCustomizationSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GuestOsCustomizationSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.guestOsCustomization.GuestOsCustomizationSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGuestOsCustomizationSpecOutputReference_Override(g GuestOsCustomizationSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.guestOsCustomization.GuestOsCustomizationSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference)SetDnsServerList(val *[]*string) {
	if err := j.validateSetDnsServerListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServerList",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference)SetDnsSuffixList(val *[]*string) {
	if err := j.validateSetDnsSuffixListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsSuffixList",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference)SetInternalValue(val *GuestOsCustomizationSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference)SetIpv4Gateway(val *string) {
	if err := j.validateSetIpv4GatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Gateway",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference)SetIpv6Gateway(val *string) {
	if err := j.validateSetIpv6GatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Gateway",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecOutputReference)SetWindowsSysprepText(val *string) {
	if err := j.validateSetWindowsSysprepTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowsSysprepText",
		val,
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) PutLinuxOptions(value *GuestOsCustomizationSpecLinuxOptions) {
	if err := g.validatePutLinuxOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLinuxOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) PutNetworkInterface(value interface{}) {
	if err := g.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) PutWindowsOptions(value *GuestOsCustomizationSpecWindowsOptions) {
	if err := g.validatePutWindowsOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWindowsOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) ResetDnsServerList() {
	_jsii_.InvokeVoid(
		g,
		"resetDnsServerList",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) ResetDnsSuffixList() {
	_jsii_.InvokeVoid(
		g,
		"resetDnsSuffixList",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) ResetIpv4Gateway() {
	_jsii_.InvokeVoid(
		g,
		"resetIpv4Gateway",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) ResetIpv6Gateway() {
	_jsii_.InvokeVoid(
		g,
		"resetIpv6Gateway",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) ResetLinuxOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetLinuxOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) ResetWindowsOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetWindowsOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) ResetWindowsSysprepText() {
	_jsii_.InvokeVoid(
		g,
		"resetWindowsSysprepText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GuestOsCustomizationSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

