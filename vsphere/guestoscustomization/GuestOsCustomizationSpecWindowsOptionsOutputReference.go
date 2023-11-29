// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package guestoscustomization

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v8/jsii"

	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v8/guestoscustomization/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GuestOsCustomizationSpecWindowsOptionsOutputReference interface {
	cdktf.ComplexObject
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
	AutoLogon() interface{}
	SetAutoLogon(val interface{})
	AutoLogonCount() *float64
	SetAutoLogonCount(val *float64)
	AutoLogonCountInput() *float64
	AutoLogonInput() interface{}
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
	ComputerName() *string
	SetComputerName(val *string)
	ComputerNameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DomainAdminPassword() *string
	SetDomainAdminPassword(val *string)
	DomainAdminPasswordInput() *string
	DomainAdminUser() *string
	SetDomainAdminUser(val *string)
	DomainAdminUserInput() *string
	// Experimental.
	Fqn() *string
	FullName() *string
	SetFullName(val *string)
	FullNameInput() *string
	InternalValue() *GuestOsCustomizationSpecWindowsOptions
	SetInternalValue(val *GuestOsCustomizationSpecWindowsOptions)
	JoinDomain() *string
	SetJoinDomain(val *string)
	JoinDomainInput() *string
	OrganizationName() *string
	SetOrganizationName(val *string)
	OrganizationNameInput() *string
	ProductKey() *string
	SetProductKey(val *string)
	ProductKeyInput() *string
	RunOnceCommandList() *[]*string
	SetRunOnceCommandList(val *[]*string)
	RunOnceCommandListInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() *float64
	SetTimeZone(val *float64)
	TimeZoneInput() *float64
	Workgroup() *string
	SetWorkgroup(val *string)
	WorkgroupInput() *string
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
	ResetAdminPassword()
	ResetAutoLogon()
	ResetAutoLogonCount()
	ResetDomainAdminPassword()
	ResetDomainAdminUser()
	ResetFullName()
	ResetJoinDomain()
	ResetOrganizationName()
	ResetProductKey()
	ResetRunOnceCommandList()
	ResetTimeZone()
	ResetWorkgroup()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GuestOsCustomizationSpecWindowsOptionsOutputReference
type jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) AutoLogon() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoLogon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) AutoLogonCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoLogonCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) AutoLogonCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoLogonCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) AutoLogonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoLogonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ComputerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ComputerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) DomainAdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAdminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) DomainAdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAdminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) DomainAdminUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAdminUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) DomainAdminUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAdminUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) FullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) FullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) InternalValue() *GuestOsCustomizationSpecWindowsOptions {
	var returns *GuestOsCustomizationSpecWindowsOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) JoinDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) JoinDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) OrganizationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) OrganizationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ProductKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ProductKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) RunOnceCommandList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"runOnceCommandList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) RunOnceCommandListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"runOnceCommandListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) TimeZone() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) TimeZoneInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) Workgroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workgroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) WorkgroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workgroupInput",
		&returns,
	)
	return returns
}


func NewGuestOsCustomizationSpecWindowsOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GuestOsCustomizationSpecWindowsOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGuestOsCustomizationSpecWindowsOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.guestOsCustomization.GuestOsCustomizationSpecWindowsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGuestOsCustomizationSpecWindowsOptionsOutputReference_Override(g GuestOsCustomizationSpecWindowsOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.guestOsCustomization.GuestOsCustomizationSpecWindowsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference)SetAdminPassword(val *string) {
	if err := j.validateSetAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference)SetAutoLogon(val interface{}) {
	if err := j.validateSetAutoLogonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoLogon",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference)SetAutoLogonCount(val *float64) {
	if err := j.validateSetAutoLogonCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoLogonCount",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference)SetComputerName(val *string) {
	if err := j.validateSetComputerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computerName",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference)SetDomainAdminPassword(val *string) {
	if err := j.validateSetDomainAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainAdminPassword",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference)SetDomainAdminUser(val *string) {
	if err := j.validateSetDomainAdminUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainAdminUser",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference)SetFullName(val *string) {
	if err := j.validateSetFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullName",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference)SetInternalValue(val *GuestOsCustomizationSpecWindowsOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference)SetJoinDomain(val *string) {
	if err := j.validateSetJoinDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"joinDomain",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference)SetOrganizationName(val *string) {
	if err := j.validateSetOrganizationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationName",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference)SetProductKey(val *string) {
	if err := j.validateSetProductKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productKey",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference)SetRunOnceCommandList(val *[]*string) {
	if err := j.validateSetRunOnceCommandListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runOnceCommandList",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference)SetTimeZone(val *float64) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (j *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference)SetWorkgroup(val *string) {
	if err := j.validateSetWorkgroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workgroup",
		val,
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ResetAdminPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetAdminPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ResetAutoLogon() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoLogon",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ResetAutoLogonCount() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoLogonCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ResetDomainAdminPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetDomainAdminPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ResetDomainAdminUser() {
	_jsii_.InvokeVoid(
		g,
		"resetDomainAdminUser",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ResetFullName() {
	_jsii_.InvokeVoid(
		g,
		"resetFullName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ResetJoinDomain() {
	_jsii_.InvokeVoid(
		g,
		"resetJoinDomain",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ResetOrganizationName() {
	_jsii_.InvokeVoid(
		g,
		"resetOrganizationName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ResetProductKey() {
	_jsii_.InvokeVoid(
		g,
		"resetProductKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ResetRunOnceCommandList() {
	_jsii_.InvokeVoid(
		g,
		"resetRunOnceCommandList",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ResetWorkgroup() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkgroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GuestOsCustomizationSpecWindowsOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

