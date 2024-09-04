// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package host

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HostServicesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (h *jsiiProxy_HostServicesList) validateGetParameters(index *float64) error {
	return nil
}

func (h *jsiiProxy_HostServicesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_HostServicesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_HostServicesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HostServicesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_HostServicesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewHostServicesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

