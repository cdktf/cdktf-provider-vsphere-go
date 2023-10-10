// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine


type VirtualMachineNetworkInterface struct {
	// The ID of the network to connect this network interface to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.0/docs/resources/virtual_machine#network_id VirtualMachine#network_id}
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// The controller type. Can be one of e1000, e1000e, vmxnet3, or vrdma.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.0/docs/resources/virtual_machine#adapter_type VirtualMachine#adapter_type}
	AdapterType *string `field:"optional" json:"adapterType" yaml:"adapterType"`
	// The upper bandwidth limit of this network interface, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.0/docs/resources/virtual_machine#bandwidth_limit VirtualMachine#bandwidth_limit}
	BandwidthLimit *float64 `field:"optional" json:"bandwidthLimit" yaml:"bandwidthLimit"`
	// The bandwidth reservation of this network interface, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.0/docs/resources/virtual_machine#bandwidth_reservation VirtualMachine#bandwidth_reservation}
	BandwidthReservation *float64 `field:"optional" json:"bandwidthReservation" yaml:"bandwidthReservation"`
	// The share count for this network interface when the share level is custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.0/docs/resources/virtual_machine#bandwidth_share_count VirtualMachine#bandwidth_share_count}
	BandwidthShareCount *float64 `field:"optional" json:"bandwidthShareCount" yaml:"bandwidthShareCount"`
	// The bandwidth share allocation level for this interface. Can be one of low, normal, high, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.0/docs/resources/virtual_machine#bandwidth_share_level VirtualMachine#bandwidth_share_level}
	BandwidthShareLevel *string `field:"optional" json:"bandwidthShareLevel" yaml:"bandwidthShareLevel"`
	// The MAC address of this network interface. Can only be manually set if use_static_mac is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.0/docs/resources/virtual_machine#mac_address VirtualMachine#mac_address}
	MacAddress *string `field:"optional" json:"macAddress" yaml:"macAddress"`
	// Mapping of network interface to OVF network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.0/docs/resources/virtual_machine#ovf_mapping VirtualMachine#ovf_mapping}
	OvfMapping *string `field:"optional" json:"ovfMapping" yaml:"ovfMapping"`
	// If true, the mac_address field is treated as a static MAC address and set accordingly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.0/docs/resources/virtual_machine#use_static_mac VirtualMachine#use_static_mac}
	UseStaticMac interface{} `field:"optional" json:"useStaticMac" yaml:"useStaticMac"`
}

