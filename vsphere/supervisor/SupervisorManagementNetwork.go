// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package supervisor


type SupervisorManagementNetwork struct {
	// Number of addresses to allocate. Starts from 'starting_address'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/supervisor#address_count Supervisor#address_count}
	AddressCount *float64 `field:"required" json:"addressCount" yaml:"addressCount"`
	// Gateway IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/supervisor#gateway Supervisor#gateway}
	Gateway *string `field:"required" json:"gateway" yaml:"gateway"`
	// ID of the network. (e.g. a distributed port group).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/supervisor#network Supervisor#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Starting address of the management network range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/supervisor#starting_address Supervisor#starting_address}
	StartingAddress *string `field:"required" json:"startingAddress" yaml:"startingAddress"`
	// Subnet mask.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/supervisor#subnet_mask Supervisor#subnet_mask}
	SubnetMask *string `field:"required" json:"subnetMask" yaml:"subnetMask"`
}

