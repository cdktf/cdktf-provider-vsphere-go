// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package contentlibrary

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v11/jsii"

	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v11/contentlibrary/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContentLibrarySubscriptionOutputReference interface {
	cdktf.ComplexObject
	AuthenticationMethod() *string
	SetAuthenticationMethod(val *string)
	AuthenticationMethodInput() *string
	AutomaticSync() interface{}
	SetAutomaticSync(val interface{})
	AutomaticSyncInput() interface{}
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
	// Experimental.
	Fqn() *string
	InternalValue() *ContentLibrarySubscription
	SetInternalValue(val *ContentLibrarySubscription)
	OnDemand() interface{}
	SetOnDemand(val interface{})
	OnDemandInput() interface{}
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	SubscriptionUrl() *string
	SetSubscriptionUrl(val *string)
	SubscriptionUrlInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	ResetAuthenticationMethod()
	ResetAutomaticSync()
	ResetOnDemand()
	ResetPassword()
	ResetSubscriptionUrl()
	ResetUsername()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContentLibrarySubscriptionOutputReference
type jsiiProxy_ContentLibrarySubscriptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) AuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) AuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) AutomaticSync() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticSync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) AutomaticSyncInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticSyncInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) InternalValue() *ContentLibrarySubscription {
	var returns *ContentLibrarySubscription
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) OnDemand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onDemand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) OnDemandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onDemandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) SubscriptionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) SubscriptionUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewContentLibrarySubscriptionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContentLibrarySubscriptionOutputReference {
	_init_.Initialize()

	if err := validateNewContentLibrarySubscriptionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContentLibrarySubscriptionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.contentLibrary.ContentLibrarySubscriptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContentLibrarySubscriptionOutputReference_Override(c ContentLibrarySubscriptionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.contentLibrary.ContentLibrarySubscriptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference)SetAuthenticationMethod(val *string) {
	if err := j.validateSetAuthenticationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationMethod",
		val,
	)
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference)SetAutomaticSync(val interface{}) {
	if err := j.validateSetAutomaticSyncParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticSync",
		val,
	)
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference)SetInternalValue(val *ContentLibrarySubscription) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference)SetOnDemand(val interface{}) {
	if err := j.validateSetOnDemandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemand",
		val,
	)
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference)SetSubscriptionUrl(val *string) {
	if err := j.validateSetSubscriptionUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscriptionUrl",
		val,
	)
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContentLibrarySubscriptionOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) ResetAuthenticationMethod() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthenticationMethod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) ResetAutomaticSync() {
	_jsii_.InvokeVoid(
		c,
		"resetAutomaticSync",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) ResetOnDemand() {
	_jsii_.InvokeVoid(
		c,
		"resetOnDemand",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		c,
		"resetPassword",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) ResetSubscriptionUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetSubscriptionUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		c,
		"resetUsername",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContentLibrarySubscriptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

