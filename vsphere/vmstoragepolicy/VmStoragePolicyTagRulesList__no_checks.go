// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package vmstoragepolicy

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VmStoragePolicyTagRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VmStoragePolicyTagRulesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VmStoragePolicyTagRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VmStoragePolicyTagRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VmStoragePolicyTagRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VmStoragePolicyTagRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVmStoragePolicyTagRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

