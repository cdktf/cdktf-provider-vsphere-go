// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine


type VirtualMachineVapp struct {
	// A map of customizable vApp properties and their values.
	//
	// Allows customization of VMs cloned from OVF templates which have customizable vApp properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.0/docs/resources/virtual_machine#properties VirtualMachine#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

