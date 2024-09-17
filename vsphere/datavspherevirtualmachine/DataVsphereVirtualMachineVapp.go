// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datavspherevirtualmachine


type DataVsphereVirtualMachineVapp struct {
	// A map of customizable vApp properties and their values.
	//
	// Allows customization of VMs cloned from OVF templates which have customizable vApp properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs/data-sources/virtual_machine#properties DataVsphereVirtualMachine#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

