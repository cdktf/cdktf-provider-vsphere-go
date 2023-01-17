package virtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v3/jsii"

	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v3/virtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineDiskOutputReference interface {
	cdktf.ComplexObject
	Attach() interface{}
	SetAttach(val interface{})
	AttachInput() interface{}
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
	ControllerType() *string
	SetControllerType(val *string)
	ControllerTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatastoreId() *string
	SetDatastoreId(val *string)
	DatastoreIdInput() *string
	DeviceAddress() *string
	DiskMode() *string
	SetDiskMode(val *string)
	DiskModeInput() *string
	DiskSharing() *string
	SetDiskSharing(val *string)
	DiskSharingInput() *string
	EagerlyScrub() interface{}
	SetEagerlyScrub(val interface{})
	EagerlyScrubInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IoLimit() *float64
	SetIoLimit(val *float64)
	IoLimitInput() *float64
	IoReservation() *float64
	SetIoReservation(val *float64)
	IoReservationInput() *float64
	IoShareCount() *float64
	SetIoShareCount(val *float64)
	IoShareCountInput() *float64
	IoShareLevel() *string
	SetIoShareLevel(val *string)
	IoShareLevelInput() *string
	KeepOnRemove() interface{}
	SetKeepOnRemove(val interface{})
	KeepOnRemoveInput() interface{}
	Key() *float64
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	Size() *float64
	SetSize(val *float64)
	SizeInput() *float64
	StoragePolicyId() *string
	SetStoragePolicyId(val *string)
	StoragePolicyIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThinProvisioned() interface{}
	SetThinProvisioned(val interface{})
	ThinProvisionedInput() interface{}
	UnitNumber() *float64
	SetUnitNumber(val *float64)
	UnitNumberInput() *float64
	Uuid() *string
	WriteThrough() interface{}
	SetWriteThrough(val interface{})
	WriteThroughInput() interface{}
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
	ResetAttach()
	ResetControllerType()
	ResetDatastoreId()
	ResetDiskMode()
	ResetDiskSharing()
	ResetEagerlyScrub()
	ResetIoLimit()
	ResetIoReservation()
	ResetIoShareCount()
	ResetIoShareLevel()
	ResetKeepOnRemove()
	ResetPath()
	ResetSize()
	ResetStoragePolicyId()
	ResetThinProvisioned()
	ResetUnitNumber()
	ResetWriteThrough()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualMachineDiskOutputReference
type jsiiProxy_VirtualMachineDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) Attach() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) AttachInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attachInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) ControllerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controllerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) ControllerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controllerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) DatastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) DatastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) DeviceAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) DiskMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) DiskModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) DiskSharing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskSharing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) DiskSharingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskSharingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) EagerlyScrub() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eagerlyScrub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) EagerlyScrubInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eagerlyScrubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) IoLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ioLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) IoLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ioLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) IoReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ioReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) IoReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ioReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) IoShareCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ioShareCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) IoShareCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ioShareCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) IoShareLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ioShareLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) IoShareLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ioShareLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) KeepOnRemove() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepOnRemove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) KeepOnRemoveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepOnRemoveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) Key() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) SizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) StoragePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) StoragePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) ThinProvisioned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thinProvisioned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) ThinProvisionedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thinProvisionedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) UnitNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unitNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) UnitNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unitNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) WriteThrough() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeThrough",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference) WriteThroughInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeThroughInput",
		&returns,
	)
	return returns
}


func NewVirtualMachineDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VirtualMachineDiskOutputReference {
	_init_.Initialize()

	if err := validateNewVirtualMachineDiskOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachineDiskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachineDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVirtualMachineDiskOutputReference_Override(v VirtualMachineDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.virtualMachine.VirtualMachineDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetAttach(val interface{}) {
	if err := j.validateSetAttachParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attach",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetControllerType(val *string) {
	if err := j.validateSetControllerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controllerType",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetDatastoreId(val *string) {
	if err := j.validateSetDatastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datastoreId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetDiskMode(val *string) {
	if err := j.validateSetDiskModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskMode",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetDiskSharing(val *string) {
	if err := j.validateSetDiskSharingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSharing",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetEagerlyScrub(val interface{}) {
	if err := j.validateSetEagerlyScrubParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eagerlyScrub",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetIoLimit(val *float64) {
	if err := j.validateSetIoLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ioLimit",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetIoReservation(val *float64) {
	if err := j.validateSetIoReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ioReservation",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetIoShareCount(val *float64) {
	if err := j.validateSetIoShareCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ioShareCount",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetIoShareLevel(val *string) {
	if err := j.validateSetIoShareLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ioShareLevel",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetKeepOnRemove(val interface{}) {
	if err := j.validateSetKeepOnRemoveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepOnRemove",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetSize(val *float64) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetStoragePolicyId(val *string) {
	if err := j.validateSetStoragePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storagePolicyId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetThinProvisioned(val interface{}) {
	if err := j.validateSetThinProvisionedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thinProvisioned",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetUnitNumber(val *float64) {
	if err := j.validateSetUnitNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unitNumber",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineDiskOutputReference)SetWriteThrough(val interface{}) {
	if err := j.validateSetWriteThroughParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeThrough",
		val,
	)
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualMachineDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualMachineDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualMachineDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualMachineDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualMachineDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualMachineDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualMachineDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualMachineDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ResetAttach() {
	_jsii_.InvokeVoid(
		v,
		"resetAttach",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ResetControllerType() {
	_jsii_.InvokeVoid(
		v,
		"resetControllerType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ResetDatastoreId() {
	_jsii_.InvokeVoid(
		v,
		"resetDatastoreId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ResetDiskMode() {
	_jsii_.InvokeVoid(
		v,
		"resetDiskMode",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ResetDiskSharing() {
	_jsii_.InvokeVoid(
		v,
		"resetDiskSharing",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ResetEagerlyScrub() {
	_jsii_.InvokeVoid(
		v,
		"resetEagerlyScrub",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ResetIoLimit() {
	_jsii_.InvokeVoid(
		v,
		"resetIoLimit",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ResetIoReservation() {
	_jsii_.InvokeVoid(
		v,
		"resetIoReservation",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ResetIoShareCount() {
	_jsii_.InvokeVoid(
		v,
		"resetIoShareCount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ResetIoShareLevel() {
	_jsii_.InvokeVoid(
		v,
		"resetIoShareLevel",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ResetKeepOnRemove() {
	_jsii_.InvokeVoid(
		v,
		"resetKeepOnRemove",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		v,
		"resetPath",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ResetSize() {
	_jsii_.InvokeVoid(
		v,
		"resetSize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ResetStoragePolicyId() {
	_jsii_.InvokeVoid(
		v,
		"resetStoragePolicyId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ResetThinProvisioned() {
	_jsii_.InvokeVoid(
		v,
		"resetThinProvisioned",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ResetUnitNumber() {
	_jsii_.InvokeVoid(
		v,
		"resetUnitNumber",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ResetWriteThrough() {
	_jsii_.InvokeVoid(
		v,
		"resetWriteThrough",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineDiskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VirtualMachineDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

