package virtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v2/jsii"

	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v2/virtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineCloneCustomizeOutputReference interface {
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
	InternalValue() *VirtualMachineCloneCustomize
	SetInternalValue(val *VirtualMachineCloneCustomize)
	Ipv4Gateway() *string
	SetIpv4Gateway(val *string)
	Ipv4GatewayInput() *string
	Ipv6Gateway() *string
	SetIpv6Gateway(val *string)
	Ipv6GatewayInput() *string
	LinuxOptions() VirtualMachineCloneCustomizeLinuxOptionsOutputReference
	LinuxOptionsInput() *VirtualMachineCloneCustomizeLinuxOptions
	NetworkInterface() VirtualMachineCloneCustomizeNetworkInterfaceList
	NetworkInterfaceInput() interface{}
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
	WindowsOptions() VirtualMachineCloneCustomizeWindowsOptionsOutputReference
	WindowsOptionsInput() *VirtualMachineCloneCustomizeWindowsOptions
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
	PutLinuxOptions(value *VirtualMachineCloneCustomizeLinuxOptions)
	PutNetworkInterface(value interface{})
	PutWindowsOptions(value *VirtualMachineCloneCustomizeWindowsOptions)
	ResetDnsServerList()
	ResetDnsSuffixList()
	ResetIpv4Gateway()
	ResetIpv6Gateway()
	ResetLinuxOptions()
	ResetNetworkInterface()
	ResetTimeout()
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

// The jsii proxy struct for VirtualMachineCloneCustomizeOutputReference
type jsiiProxy_VirtualMachineCloneCustomizeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) DnsServerList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServerList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) DnsServerListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServerListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) DnsSuffixList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsSuffixList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) DnsSuffixListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsSuffixListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) InternalValue() *VirtualMachineCloneCustomize {
	var returns *VirtualMachineCloneCustomize
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) Ipv4Gateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4Gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) Ipv4GatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4GatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) Ipv6Gateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) Ipv6GatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6GatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) LinuxOptions() VirtualMachineCloneCustomizeLinuxOptionsOutputReference {
	var returns VirtualMachineCloneCustomizeLinuxOptionsOutputReference
	_jsii_.Get(
		j,
		"linuxOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) LinuxOptionsInput() *VirtualMachineCloneCustomizeLinuxOptions {
	var returns *VirtualMachineCloneCustomizeLinuxOptions
	_jsii_.Get(
		j,
		"linuxOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) NetworkInterface() VirtualMachineCloneCustomizeNetworkInterfaceList {
	var returns VirtualMachineCloneCustomizeNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) WindowsOptions() VirtualMachineCloneCustomizeWindowsOptionsOutputReference {
	var returns VirtualMachineCloneCustomizeWindowsOptionsOutputReference
	_jsii_.Get(
		j,
		"windowsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) WindowsOptionsInput() *VirtualMachineCloneCustomizeWindowsOptions {
	var returns *VirtualMachineCloneCustomizeWindowsOptions
	_jsii_.Get(
		j,
		"windowsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) WindowsSysprepText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowsSysprepText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) WindowsSysprepTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowsSysprepTextInput",
		&returns,
	)
	return returns
}


func NewVirtualMachineCloneCustomizeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualMachineCloneCustomizeOutputReference {
	_init_.Initialize()

	if err := validateNewVirtualMachineCloneCustomizeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachineCloneCustomizeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachineCloneCustomizeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualMachineCloneCustomizeOutputReference_Override(v VirtualMachineCloneCustomizeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachineCloneCustomizeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference)SetDnsServerList(val *[]*string) {
	if err := j.validateSetDnsServerListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServerList",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference)SetDnsSuffixList(val *[]*string) {
	if err := j.validateSetDnsSuffixListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsSuffixList",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference)SetInternalValue(val *VirtualMachineCloneCustomize) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference)SetIpv4Gateway(val *string) {
	if err := j.validateSetIpv4GatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Gateway",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference)SetIpv6Gateway(val *string) {
	if err := j.validateSetIpv6GatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Gateway",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeOutputReference)SetWindowsSysprepText(val *string) {
	if err := j.validateSetWindowsSysprepTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowsSysprepText",
		val,
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) PutLinuxOptions(value *VirtualMachineCloneCustomizeLinuxOptions) {
	if err := v.validatePutLinuxOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putLinuxOptions",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) PutNetworkInterface(value interface{}) {
	if err := v.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) PutWindowsOptions(value *VirtualMachineCloneCustomizeWindowsOptions) {
	if err := v.validatePutWindowsOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putWindowsOptions",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) ResetDnsServerList() {
	_jsii_.InvokeVoid(
		v,
		"resetDnsServerList",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) ResetDnsSuffixList() {
	_jsii_.InvokeVoid(
		v,
		"resetDnsSuffixList",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) ResetIpv4Gateway() {
	_jsii_.InvokeVoid(
		v,
		"resetIpv4Gateway",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) ResetIpv6Gateway() {
	_jsii_.InvokeVoid(
		v,
		"resetIpv6Gateway",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) ResetLinuxOptions() {
	_jsii_.InvokeVoid(
		v,
		"resetLinuxOptions",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		v,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeout",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) ResetWindowsOptions() {
	_jsii_.InvokeVoid(
		v,
		"resetWindowsOptions",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) ResetWindowsSysprepText() {
	_jsii_.InvokeVoid(
		v,
		"resetWindowsSysprepText",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

