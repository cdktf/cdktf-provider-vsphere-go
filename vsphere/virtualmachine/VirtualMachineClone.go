// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine


type VirtualMachineClone struct {
	// The UUID of the source virtual machine or template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/virtual_machine#template_uuid VirtualMachine#template_uuid}
	TemplateUuid *string `field:"required" json:"templateUuid" yaml:"templateUuid"`
	// customize block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/virtual_machine#customize VirtualMachine#customize}
	Customize *VirtualMachineCloneCustomize `field:"optional" json:"customize" yaml:"customize"`
	// Whether or not to create a linked clone when cloning.
	//
	// When this option is used, the source VM must have a single snapshot associated with it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/virtual_machine#linked_clone VirtualMachine#linked_clone}
	LinkedClone interface{} `field:"optional" json:"linkedClone" yaml:"linkedClone"`
	// Mapping of ovf networks to the networks to use in vSphere.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/virtual_machine#ovf_network_map VirtualMachine#ovf_network_map}
	OvfNetworkMap *map[string]*string `field:"optional" json:"ovfNetworkMap" yaml:"ovfNetworkMap"`
	// Mapping of ovf storage to the datastores to use in vSphere.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/virtual_machine#ovf_storage_map VirtualMachine#ovf_storage_map}
	OvfStorageMap *map[string]*string `field:"optional" json:"ovfStorageMap" yaml:"ovfStorageMap"`
	// The timeout, in minutes, to wait for the virtual machine clone to complete.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/virtual_machine#timeout VirtualMachine#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

