// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package distributedvirtualswitch


type DistributedVirtualSwitchHost struct {
	// The managed object ID of the host this specification applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/distributed_virtual_switch#host_system_id DistributedVirtualSwitch#host_system_id}
	HostSystemId *string `field:"required" json:"hostSystemId" yaml:"hostSystemId"`
	// Name of the physical NIC to be added to the proxy switch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/distributed_virtual_switch#devices DistributedVirtualSwitch#devices}
	Devices *[]*string `field:"optional" json:"devices" yaml:"devices"`
}

