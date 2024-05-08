// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package supervisor

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SupervisorEgressCidrList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorEgressCidrList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SupervisorEgressCidrList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SupervisorEgressCidrList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SupervisorEgressCidrList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SupervisorEgressCidrList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SupervisorEgressCidrList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSupervisorEgressCidrListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

