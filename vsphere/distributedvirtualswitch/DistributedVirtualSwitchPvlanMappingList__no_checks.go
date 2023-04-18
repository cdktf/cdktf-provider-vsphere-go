//go:build no_runtime_type_checking

package distributedvirtualswitch

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DistributedVirtualSwitchPvlanMappingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DistributedVirtualSwitchPvlanMappingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDistributedVirtualSwitchPvlanMappingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

