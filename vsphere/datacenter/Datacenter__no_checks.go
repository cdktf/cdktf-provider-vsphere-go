// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datacenter

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Datacenter) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (d *jsiiProxy_Datacenter) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_Datacenter) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datacenter) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datacenter) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datacenter) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datacenter) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datacenter) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datacenter) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datacenter) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datacenter) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datacenter) validateImportFromParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_Datacenter) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datacenter) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_Datacenter) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (d *jsiiProxy_Datacenter) validateMoveToIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_Datacenter) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateDatacenter_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateDatacenter_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatacenter_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDatacenter_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Datacenter) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Datacenter) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Datacenter) validateSetCustomAttributesParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Datacenter) validateSetFolderParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Datacenter) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Datacenter) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Datacenter) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Datacenter) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Datacenter) validateSetTagsParameters(val *[]*string) error {
	return nil
}

func validateNewDatacenterParameters(scope constructs.Construct, id *string, config *DatacenterConfig) error {
	return nil
}

