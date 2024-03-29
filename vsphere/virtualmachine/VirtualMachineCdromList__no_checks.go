// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package virtualmachine

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualMachineCdromList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (v *jsiiProxy_VirtualMachineCdromList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VirtualMachineCdromList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineCdromList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineCdromList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineCdromList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VirtualMachineCdromList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVirtualMachineCdromListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

