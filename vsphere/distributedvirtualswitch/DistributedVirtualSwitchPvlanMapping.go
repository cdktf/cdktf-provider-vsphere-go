// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package distributedvirtualswitch


type DistributedVirtualSwitchPvlanMapping struct {
	// The primary VLAN ID.
	//
	// The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/distributed_virtual_switch#primary_vlan_id DistributedVirtualSwitch#primary_vlan_id}
	PrimaryVlanId *float64 `field:"required" json:"primaryVlanId" yaml:"primaryVlanId"`
	// The private VLAN type. Valid values are promiscuous, community and isolated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/distributed_virtual_switch#pvlan_type DistributedVirtualSwitch#pvlan_type}
	PvlanType *string `field:"required" json:"pvlanType" yaml:"pvlanType"`
	// The secondary VLAN ID.
	//
	// The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/distributed_virtual_switch#secondary_vlan_id DistributedVirtualSwitch#secondary_vlan_id}
	SecondaryVlanId *float64 `field:"required" json:"secondaryVlanId" yaml:"secondaryVlanId"`
}

