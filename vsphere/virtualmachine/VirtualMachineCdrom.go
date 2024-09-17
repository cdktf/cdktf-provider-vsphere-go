// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine


type VirtualMachineCdrom struct {
	// Indicates whether the device should be mapped to a remote client device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs/resources/virtual_machine#client_device VirtualMachine#client_device}
	ClientDevice interface{} `field:"optional" json:"clientDevice" yaml:"clientDevice"`
	// The datastore ID the ISO is located on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs/resources/virtual_machine#datastore_id VirtualMachine#datastore_id}
	DatastoreId *string `field:"optional" json:"datastoreId" yaml:"datastoreId"`
	// The path to the ISO file on the datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs/resources/virtual_machine#path VirtualMachine#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

