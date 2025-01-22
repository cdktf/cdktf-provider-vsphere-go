// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine


type VirtualMachineCloneCustomize struct {
	// The list of DNS servers for a virtual network adapter with a static IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/resources/virtual_machine#dns_server_list VirtualMachine#dns_server_list}
	DnsServerList *[]*string `field:"optional" json:"dnsServerList" yaml:"dnsServerList"`
	// A list of DNS search domains to add to the DNS configuration on the virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/resources/virtual_machine#dns_suffix_list VirtualMachine#dns_suffix_list}
	DnsSuffixList *[]*string `field:"optional" json:"dnsSuffixList" yaml:"dnsSuffixList"`
	// The IPv4 default gateway when using network_interface customization on the virtual machine.
	//
	// This address must be local to a static IPv4 address configured in an interface sub-resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/resources/virtual_machine#ipv4_gateway VirtualMachine#ipv4_gateway}
	Ipv4Gateway *string `field:"optional" json:"ipv4Gateway" yaml:"ipv4Gateway"`
	// The IPv6 default gateway when using network_interface customization on the virtual machine.
	//
	// This address must be local to a static IPv4 address configured in an interface sub-resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/resources/virtual_machine#ipv6_gateway VirtualMachine#ipv6_gateway}
	Ipv6Gateway *string `field:"optional" json:"ipv6Gateway" yaml:"ipv6Gateway"`
	// linux_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/resources/virtual_machine#linux_options VirtualMachine#linux_options}
	LinuxOptions *VirtualMachineCloneCustomizeLinuxOptions `field:"optional" json:"linuxOptions" yaml:"linuxOptions"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/resources/virtual_machine#network_interface VirtualMachine#network_interface}
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// The amount of time, in minutes, to wait for guest OS customization to complete before returning with an error.
	//
	// Setting this value to 0 or a negative value skips the waiter. Default: 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/resources/virtual_machine#timeout VirtualMachine#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// windows_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/resources/virtual_machine#windows_options VirtualMachine#windows_options}
	WindowsOptions *VirtualMachineCloneCustomizeWindowsOptions `field:"optional" json:"windowsOptions" yaml:"windowsOptions"`
	// Use this option to specify a windows sysprep file directly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/resources/virtual_machine#windows_sysprep_text VirtualMachine#windows_sysprep_text}
	WindowsSysprepText *string `field:"optional" json:"windowsSysprepText" yaml:"windowsSysprepText"`
}

