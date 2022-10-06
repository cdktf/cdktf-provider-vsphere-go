//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VsphereProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (v *jsiiProxy_VsphereProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateVsphereProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_VsphereProvider) validateSetAllowUnverifiedSslParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VsphereProvider) validateSetClientDebugParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VsphereProvider) validateSetPersistSessionParameters(val interface{}) error {
	return nil
}

func validateNewVsphereProviderParameters(scope constructs.Construct, id *string, config *VsphereProviderConfig) error {
	return nil
}

