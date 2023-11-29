// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package guestoscustomization

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GuestOsCustomizationSpecNetworkInterfaceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGuestOsCustomizationSpecNetworkInterfaceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

