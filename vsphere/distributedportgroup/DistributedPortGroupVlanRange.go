// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package distributedportgroup


type DistributedPortGroupVlanRange struct {
	// The minimum VLAN to use in the range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_port_group#max_vlan DistributedPortGroup#max_vlan}
	MaxVlan *float64 `field:"required" json:"maxVlan" yaml:"maxVlan"`
	// The minimum VLAN to use in the range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_port_group#min_vlan DistributedPortGroup#min_vlan}
	MinVlan *float64 `field:"required" json:"minVlan" yaml:"minVlan"`
}

