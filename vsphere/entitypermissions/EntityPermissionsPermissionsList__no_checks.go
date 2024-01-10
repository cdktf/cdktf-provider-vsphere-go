// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package entitypermissions

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EntityPermissionsPermissionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EntityPermissionsPermissionsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EntityPermissionsPermissionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EntityPermissionsPermissionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EntityPermissionsPermissionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EntityPermissionsPermissionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EntityPermissionsPermissionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEntityPermissionsPermissionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

