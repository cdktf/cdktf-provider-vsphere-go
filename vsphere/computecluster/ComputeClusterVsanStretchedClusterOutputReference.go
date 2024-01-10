// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v9/jsii"

	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v9/computecluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeClusterVsanStretchedClusterOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeClusterVsanStretchedCluster
	SetInternalValue(val *ComputeClusterVsanStretchedCluster)
	PreferredFaultDomainHostIds() *[]*string
	SetPreferredFaultDomainHostIds(val *[]*string)
	PreferredFaultDomainHostIdsInput() *[]*string
	PreferredFaultDomainName() *string
	SetPreferredFaultDomainName(val *string)
	PreferredFaultDomainNameInput() *string
	SecondaryFaultDomainHostIds() *[]*string
	SetSecondaryFaultDomainHostIds(val *[]*string)
	SecondaryFaultDomainHostIdsInput() *[]*string
	SecondaryFaultDomainName() *string
	SetSecondaryFaultDomainName(val *string)
	SecondaryFaultDomainNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WitnessNode() *string
	SetWitnessNode(val *string)
	WitnessNodeInput() *string
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
	ResetPreferredFaultDomainName()
	ResetSecondaryFaultDomainName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeClusterVsanStretchedClusterOutputReference
type jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) InternalValue() *ComputeClusterVsanStretchedCluster {
	var returns *ComputeClusterVsanStretchedCluster
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) PreferredFaultDomainHostIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredFaultDomainHostIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) PreferredFaultDomainHostIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredFaultDomainHostIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) PreferredFaultDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredFaultDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) PreferredFaultDomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredFaultDomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) SecondaryFaultDomainHostIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secondaryFaultDomainHostIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) SecondaryFaultDomainHostIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secondaryFaultDomainHostIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) SecondaryFaultDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryFaultDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) SecondaryFaultDomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryFaultDomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) WitnessNode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"witnessNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) WitnessNodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"witnessNodeInput",
		&returns,
	)
	return returns
}


func NewComputeClusterVsanStretchedClusterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeClusterVsanStretchedClusterOutputReference {
	_init_.Initialize()

	if err := validateNewComputeClusterVsanStretchedClusterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.computeCluster.ComputeClusterVsanStretchedClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeClusterVsanStretchedClusterOutputReference_Override(c ComputeClusterVsanStretchedClusterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.computeCluster.ComputeClusterVsanStretchedClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference)SetInternalValue(val *ComputeClusterVsanStretchedCluster) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference)SetPreferredFaultDomainHostIds(val *[]*string) {
	if err := j.validateSetPreferredFaultDomainHostIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredFaultDomainHostIds",
		val,
	)
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference)SetPreferredFaultDomainName(val *string) {
	if err := j.validateSetPreferredFaultDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredFaultDomainName",
		val,
	)
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference)SetSecondaryFaultDomainHostIds(val *[]*string) {
	if err := j.validateSetSecondaryFaultDomainHostIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryFaultDomainHostIds",
		val,
	)
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference)SetSecondaryFaultDomainName(val *string) {
	if err := j.validateSetSecondaryFaultDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryFaultDomainName",
		val,
	)
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference)SetWitnessNode(val *string) {
	if err := j.validateSetWitnessNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"witnessNode",
		val,
	)
}

func (c *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) ResetPreferredFaultDomainName() {
	_jsii_.InvokeVoid(
		c,
		"resetPreferredFaultDomainName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) ResetSecondaryFaultDomainName() {
	_jsii_.InvokeVoid(
		c,
		"resetSecondaryFaultDomainName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeClusterVsanStretchedClusterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

