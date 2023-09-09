// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vnic


type VnicIpv4 struct {
	// Use DHCP to configure the interface's IPv4 stack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.3/docs/resources/vnic#dhcp Vnic#dhcp}
	Dhcp interface{} `field:"optional" json:"dhcp" yaml:"dhcp"`
	// IP address of the default gateway, if DHCP is not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.3/docs/resources/vnic#gw Vnic#gw}
	Gw *string `field:"optional" json:"gw" yaml:"gw"`
	// address of the interface, if DHCP is not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.3/docs/resources/vnic#ip Vnic#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// netmask of the interface, if DHCP is not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.3/docs/resources/vnic#netmask Vnic#netmask}
	Netmask *string `field:"optional" json:"netmask" yaml:"netmask"`
}

