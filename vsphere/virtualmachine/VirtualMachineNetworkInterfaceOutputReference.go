// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v11/jsii"

	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v11/virtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineNetworkInterfaceOutputReference interface {
	cdktf.ComplexObject
	AdapterType() *string
	SetAdapterType(val *string)
	AdapterTypeInput() *string
	BandwidthLimit() *float64
	SetBandwidthLimit(val *float64)
	BandwidthLimitInput() *float64
	BandwidthReservation() *float64
	SetBandwidthReservation(val *float64)
	BandwidthReservationInput() *float64
	BandwidthShareCount() *float64
	SetBandwidthShareCount(val *float64)
	BandwidthShareCountInput() *float64
	BandwidthShareLevel() *string
	SetBandwidthShareLevel(val *string)
	BandwidthShareLevelInput() *string
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
	DeviceAddress() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *float64
	MacAddress() *string
	SetMacAddress(val *string)
	MacAddressInput() *string
	NetworkId() *string
	SetNetworkId(val *string)
	NetworkIdInput() *string
	OvfMapping() *string
	SetOvfMapping(val *string)
	OvfMappingInput() *string
	PhysicalFunction() *string
	SetPhysicalFunction(val *string)
	PhysicalFunctionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseStaticMac() interface{}
	SetUseStaticMac(val interface{})
	UseStaticMacInput() interface{}
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
	ResetAdapterType()
	ResetBandwidthLimit()
	ResetBandwidthReservation()
	ResetBandwidthShareCount()
	ResetBandwidthShareLevel()
	ResetMacAddress()
	ResetOvfMapping()
	ResetPhysicalFunction()
	ResetUseStaticMac()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualMachineNetworkInterfaceOutputReference
type jsiiProxy_VirtualMachineNetworkInterfaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) AdapterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adapterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) AdapterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adapterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) BandwidthLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) BandwidthLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) BandwidthReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) BandwidthReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) BandwidthShareCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthShareCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) BandwidthShareCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthShareCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) BandwidthShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) BandwidthShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) DeviceAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) Key() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) MacAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) MacAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) NetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) OvfMapping() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ovfMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) OvfMappingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ovfMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) PhysicalFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) PhysicalFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) UseStaticMac() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useStaticMac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) UseStaticMacInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useStaticMacInput",
		&returns,
	)
	return returns
}


func NewVirtualMachineNetworkInterfaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VirtualMachineNetworkInterfaceOutputReference {
	_init_.Initialize()

	if err := validateNewVirtualMachineNetworkInterfaceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachineNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachineNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVirtualMachineNetworkInterfaceOutputReference_Override(v VirtualMachineNetworkInterfaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachineNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference)SetAdapterType(val *string) {
	if err := j.validateSetAdapterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adapterType",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference)SetBandwidthLimit(val *float64) {
	if err := j.validateSetBandwidthLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthLimit",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference)SetBandwidthReservation(val *float64) {
	if err := j.validateSetBandwidthReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthReservation",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference)SetBandwidthShareCount(val *float64) {
	if err := j.validateSetBandwidthShareCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthShareCount",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference)SetBandwidthShareLevel(val *string) {
	if err := j.validateSetBandwidthShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthShareLevel",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference)SetMacAddress(val *string) {
	if err := j.validateSetMacAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"macAddress",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference)SetNetworkId(val *string) {
	if err := j.validateSetNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference)SetOvfMapping(val *string) {
	if err := j.validateSetOvfMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ovfMapping",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference)SetPhysicalFunction(val *string) {
	if err := j.validateSetPhysicalFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"physicalFunction",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference)SetUseStaticMac(val interface{}) {
	if err := j.validateSetUseStaticMacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useStaticMac",
		val,
	)
}

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) ResetAdapterType() {
	_jsii_.InvokeVoid(
		v,
		"resetAdapterType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) ResetBandwidthLimit() {
	_jsii_.InvokeVoid(
		v,
		"resetBandwidthLimit",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) ResetBandwidthReservation() {
	_jsii_.InvokeVoid(
		v,
		"resetBandwidthReservation",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) ResetBandwidthShareCount() {
	_jsii_.InvokeVoid(
		v,
		"resetBandwidthShareCount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) ResetBandwidthShareLevel() {
	_jsii_.InvokeVoid(
		v,
		"resetBandwidthShareLevel",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) ResetMacAddress() {
	_jsii_.InvokeVoid(
		v,
		"resetMacAddress",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) ResetOvfMapping() {
	_jsii_.InvokeVoid(
		v,
		"resetOvfMapping",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) ResetPhysicalFunction() {
	_jsii_.InvokeVoid(
		v,
		"resetPhysicalFunction",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) ResetUseStaticMac() {
	_jsii_.InvokeVoid(
		v,
		"resetUseStaticMac",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VirtualMachineNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

