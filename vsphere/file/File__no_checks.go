// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package file

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_File) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (f *jsiiProxy_File) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (f *jsiiProxy_File) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateImportFromParameters(id *string) error {
	return nil
}

func (f *jsiiProxy_File) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (f *jsiiProxy_File) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateFile_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateFile_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFile_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateFile_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_File) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_File) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_File) validateSetCreateDirectoriesParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_File) validateSetDatacenterParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_File) validateSetDatastoreParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_File) validateSetDestinationFileParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_File) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_File) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_File) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_File) validateSetSourceDatacenterParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_File) validateSetSourceDatastoreParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_File) validateSetSourceFileParameters(val *string) error {
	return nil
}

func validateNewFileParameters(scope constructs.Construct, id *string, config *FileConfig) error {
	return nil
}

