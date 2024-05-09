// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package guestoscustomization


type GuestOsCustomizationSpecLinuxOptions struct {
	// The domain name for this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/guest_os_customization#domain GuestOsCustomization#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The hostname for this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/guest_os_customization#host_name GuestOsCustomization#host_name}
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Specifies whether or not the hardware clock should be in UTC or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/guest_os_customization#hw_clock_utc GuestOsCustomization#hw_clock_utc}
	HwClockUtc interface{} `field:"optional" json:"hwClockUtc" yaml:"hwClockUtc"`
	// The customization script to run before and or after guest customization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/guest_os_customization#script_text GuestOsCustomization#script_text}
	ScriptText *string `field:"optional" json:"scriptText" yaml:"scriptText"`
	// Customize the time zone on the VM. This should be a time zone-style entry, like America/Los_Angeles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/guest_os_customization#time_zone GuestOsCustomization#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

