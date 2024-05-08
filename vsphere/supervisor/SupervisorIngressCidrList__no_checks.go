// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package supervisor

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SupervisorIngressCidrList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorIngressCidrList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SupervisorIngressCidrList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SupervisorIngressCidrList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SupervisorIngressCidrList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SupervisorIngressCidrList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SupervisorIngressCidrList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSupervisorIngressCidrListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

