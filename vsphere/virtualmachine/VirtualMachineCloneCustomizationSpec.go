// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine


type VirtualMachineCloneCustomizationSpec struct {
	// The unique identifier of the customization specification is its name and is unique per vCenter Server instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/virtual_machine#id VirtualMachine#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The amount of time, in minutes, to wait for guest OS customization to complete before returning with an error.
	//
	// Setting this value to 0 or a negative value skips the waiter. Default: 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/virtual_machine#timeout VirtualMachine#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

