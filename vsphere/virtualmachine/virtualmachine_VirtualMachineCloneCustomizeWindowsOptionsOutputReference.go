package virtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v2/jsii"

	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v2/virtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineCloneCustomizeWindowsOptionsOutputReference interface {
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
	InternalValue() *VirtualMachineCloneCustomizeWindowsOptions
	SetInternalValue(val *VirtualMachineCloneCustomizeWindowsOptions)
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

// The jsii proxy struct for VirtualMachineCloneCustomizeWindowsOptionsOutputReference
type jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) AutoLogon() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoLogon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) AutoLogonCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoLogonCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) AutoLogonCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoLogonCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) AutoLogonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoLogonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ComputerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ComputerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) DomainAdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAdminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) DomainAdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAdminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) DomainAdminUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAdminUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) DomainAdminUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAdminUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) FullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) FullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) InternalValue() *VirtualMachineCloneCustomizeWindowsOptions {
	var returns *VirtualMachineCloneCustomizeWindowsOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) JoinDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) JoinDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) OrganizationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) OrganizationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ProductKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ProductKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) RunOnceCommandList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"runOnceCommandList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) RunOnceCommandListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"runOnceCommandListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) TimeZone() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) TimeZoneInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) Workgroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workgroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) WorkgroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workgroupInput",
		&returns,
	)
	return returns
}


func NewVirtualMachineCloneCustomizeWindowsOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualMachineCloneCustomizeWindowsOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewVirtualMachineCloneCustomizeWindowsOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachineCloneCustomizeWindowsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualMachineCloneCustomizeWindowsOptionsOutputReference_Override(v VirtualMachineCloneCustomizeWindowsOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachineCloneCustomizeWindowsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference)SetAdminPassword(val *string) {
	if err := j.validateSetAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference)SetAutoLogon(val interface{}) {
	if err := j.validateSetAutoLogonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoLogon",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference)SetAutoLogonCount(val *float64) {
	if err := j.validateSetAutoLogonCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoLogonCount",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference)SetComputerName(val *string) {
	if err := j.validateSetComputerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computerName",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference)SetDomainAdminPassword(val *string) {
	if err := j.validateSetDomainAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainAdminPassword",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference)SetDomainAdminUser(val *string) {
	if err := j.validateSetDomainAdminUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainAdminUser",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference)SetFullName(val *string) {
	if err := j.validateSetFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullName",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference)SetInternalValue(val *VirtualMachineCloneCustomizeWindowsOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference)SetJoinDomain(val *string) {
	if err := j.validateSetJoinDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"joinDomain",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference)SetOrganizationName(val *string) {
	if err := j.validateSetOrganizationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationName",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference)SetProductKey(val *string) {
	if err := j.validateSetProductKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productKey",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference)SetRunOnceCommandList(val *[]*string) {
	if err := j.validateSetRunOnceCommandListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runOnceCommandList",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference)SetTimeZone(val *float64) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference)SetWorkgroup(val *string) {
	if err := j.validateSetWorkgroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workgroup",
		val,
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ResetAdminPassword() {
	_jsii_.InvokeVoid(
		v,
		"resetAdminPassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ResetAutoLogon() {
	_jsii_.InvokeVoid(
		v,
		"resetAutoLogon",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ResetAutoLogonCount() {
	_jsii_.InvokeVoid(
		v,
		"resetAutoLogonCount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ResetDomainAdminPassword() {
	_jsii_.InvokeVoid(
		v,
		"resetDomainAdminPassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ResetDomainAdminUser() {
	_jsii_.InvokeVoid(
		v,
		"resetDomainAdminUser",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ResetFullName() {
	_jsii_.InvokeVoid(
		v,
		"resetFullName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ResetJoinDomain() {
	_jsii_.InvokeVoid(
		v,
		"resetJoinDomain",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ResetOrganizationName() {
	_jsii_.InvokeVoid(
		v,
		"resetOrganizationName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ResetProductKey() {
	_jsii_.InvokeVoid(
		v,
		"resetProductKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ResetRunOnceCommandList() {
	_jsii_.InvokeVoid(
		v,
		"resetRunOnceCommandList",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ResetWorkgroup() {
	_jsii_.InvokeVoid(
		v,
		"resetWorkgroup",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VirtualMachineCloneCustomizeWindowsOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

