// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine


type VirtualMachineCloneCustomizeNetworkInterface struct {
	// A DNS search domain to add to the DNS configuration on the virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/virtual_machine#dns_domain VirtualMachine#dns_domain}
	DnsDomain *string `field:"optional" json:"dnsDomain" yaml:"dnsDomain"`
	// Network-interface specific DNS settings for Windows operating systems. Ignored on Linux.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/virtual_machine#dns_server_list VirtualMachine#dns_server_list}
	DnsServerList *[]*string `field:"optional" json:"dnsServerList" yaml:"dnsServerList"`
	// The IPv4 address assigned to this network adapter. If left blank, DHCP is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/virtual_machine#ipv4_address VirtualMachine#ipv4_address}
	Ipv4Address *string `field:"optional" json:"ipv4Address" yaml:"ipv4Address"`
	// The IPv4 CIDR netmask for the supplied IP address. Ignored if DHCP is selected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/virtual_machine#ipv4_netmask VirtualMachine#ipv4_netmask}
	Ipv4Netmask *float64 `field:"optional" json:"ipv4Netmask" yaml:"ipv4Netmask"`
	// The IPv6 address assigned to this network adapter. If left blank, default auto-configuration is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/virtual_machine#ipv6_address VirtualMachine#ipv6_address}
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
	// The IPv6 CIDR netmask for the supplied IP address. Ignored if auto-configuration is selected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/virtual_machine#ipv6_netmask VirtualMachine#ipv6_netmask}
	Ipv6Netmask *float64 `field:"optional" json:"ipv6Netmask" yaml:"ipv6Netmask"`
}

