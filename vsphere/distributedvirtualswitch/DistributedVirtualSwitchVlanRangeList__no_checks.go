// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package distributedvirtualswitch

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DistributedVirtualSwitchVlanRangeList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DistributedVirtualSwitchVlanRangeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DistributedVirtualSwitchVlanRangeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DistributedVirtualSwitchVlanRangeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DistributedVirtualSwitchVlanRangeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DistributedVirtualSwitchVlanRangeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDistributedVirtualSwitchVlanRangeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

