// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package supervisor

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SupervisorPodCidrList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorPodCidrList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SupervisorPodCidrList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SupervisorPodCidrList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SupervisorPodCidrList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SupervisorPodCidrList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SupervisorPodCidrList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSupervisorPodCidrListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

