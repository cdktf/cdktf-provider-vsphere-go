//go:build no_runtime_type_checking

package host

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_Host) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (h *jsiiProxy_Host) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (h *jsiiProxy_Host) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateHost_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHost_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateHost_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetClusterParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetClusterManagedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetConnectedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetCustomAttributesParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetDatacenterParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetForceParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetHostnameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetLicenseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetLockdownParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetMaintenanceParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetPasswordParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetTagsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetThumbprintParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Host) validateSetUsernameParameters(val *string) error {
	return nil
}

func validateNewHostParameters(scope constructs.Construct, id *string, config *HostConfig) error {
	return nil
}

