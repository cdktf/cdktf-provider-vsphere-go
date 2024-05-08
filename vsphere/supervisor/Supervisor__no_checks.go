// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package supervisor

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Supervisor) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validatePutEgressCidrParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validatePutIngressCidrParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validatePutManagementNetworkParameters(value *SupervisorManagementNetwork) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validatePutNamespaceParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validatePutPodCidrParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_Supervisor) validatePutServiceCidrParameters(value *SupervisorServiceCidr) error {
	return nil
}

func validateSupervisor_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSupervisor_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSupervisor_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSupervisor_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Supervisor) validateSetClusterParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Supervisor) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Supervisor) validateSetContentLibraryParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Supervisor) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Supervisor) validateSetDvsUuidParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Supervisor) validateSetEdgeClusterParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Supervisor) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Supervisor) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Supervisor) validateSetMainDnsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Supervisor) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Supervisor) validateSetSearchDomainsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Supervisor) validateSetSizingHintParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Supervisor) validateSetStoragePolicyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Supervisor) validateSetWorkerDnsParameters(val *[]*string) error {
	return nil
}

func validateNewSupervisorParameters(scope constructs.Construct, id *string, config *SupervisorConfig) error {
	return nil
}

