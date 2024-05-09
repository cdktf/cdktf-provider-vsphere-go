// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package supervisor


type SupervisorPodCidr struct {
	// Network address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/supervisor#address Supervisor#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// Subnet prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/supervisor#prefix Supervisor#prefix}
	Prefix *float64 `field:"required" json:"prefix" yaml:"prefix"`
}

