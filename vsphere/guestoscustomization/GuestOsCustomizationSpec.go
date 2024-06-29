// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package guestoscustomization


type GuestOsCustomizationSpec struct {
	// The list of DNS servers for a virtual network adapter with a static IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/resources/guest_os_customization#dns_server_list GuestOsCustomization#dns_server_list}
	DnsServerList *[]*string `field:"optional" json:"dnsServerList" yaml:"dnsServerList"`
	// A list of DNS search domains to add to the DNS configuration on the virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/resources/guest_os_customization#dns_suffix_list GuestOsCustomization#dns_suffix_list}
	DnsSuffixList *[]*string `field:"optional" json:"dnsSuffixList" yaml:"dnsSuffixList"`
	// The IPv4 default gateway when using network_interface customization on the virtual machine.
	//
	// This address must be local to a static IPv4 address configured in an interface sub-resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/resources/guest_os_customization#ipv4_gateway GuestOsCustomization#ipv4_gateway}
	Ipv4Gateway *string `field:"optional" json:"ipv4Gateway" yaml:"ipv4Gateway"`
	// The IPv6 default gateway when using network_interface customization on the virtual machine.
	//
	// This address must be local to a static IPv4 address configured in an interface sub-resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/resources/guest_os_customization#ipv6_gateway GuestOsCustomization#ipv6_gateway}
	Ipv6Gateway *string `field:"optional" json:"ipv6Gateway" yaml:"ipv6Gateway"`
	// linux_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/resources/guest_os_customization#linux_options GuestOsCustomization#linux_options}
	LinuxOptions *GuestOsCustomizationSpecLinuxOptions `field:"optional" json:"linuxOptions" yaml:"linuxOptions"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/resources/guest_os_customization#network_interface GuestOsCustomization#network_interface}
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// windows_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/resources/guest_os_customization#windows_options GuestOsCustomization#windows_options}
	WindowsOptions *GuestOsCustomizationSpecWindowsOptions `field:"optional" json:"windowsOptions" yaml:"windowsOptions"`
	// Use this option to specify a windows sysprep file directly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/resources/guest_os_customization#windows_sysprep_text GuestOsCustomization#windows_sysprep_text}
	WindowsSysprepText *string `field:"optional" json:"windowsSysprepText" yaml:"windowsSysprepText"`
}

