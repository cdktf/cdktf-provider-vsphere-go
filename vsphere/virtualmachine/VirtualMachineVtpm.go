// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine


type VirtualMachineVtpm struct {
	// The version of the TPM device. Default is 2.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#version VirtualMachine#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

