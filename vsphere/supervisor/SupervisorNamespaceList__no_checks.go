// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package supervisor

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SupervisorNamespaceList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorNamespaceList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SupervisorNamespaceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SupervisorNamespaceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SupervisorNamespaceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SupervisorNamespaceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SupervisorNamespaceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSupervisorNamespaceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

