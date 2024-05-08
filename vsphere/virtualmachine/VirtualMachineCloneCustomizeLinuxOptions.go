// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine


type VirtualMachineCloneCustomizeLinuxOptions struct {
	// The domain name for this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/virtual_machine#domain VirtualMachine#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The hostname for this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/virtual_machine#host_name VirtualMachine#host_name}
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Specifies whether or not the hardware clock should be in UTC or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/virtual_machine#hw_clock_utc VirtualMachine#hw_clock_utc}
	HwClockUtc interface{} `field:"optional" json:"hwClockUtc" yaml:"hwClockUtc"`
	// The customization script to run before and or after guest customization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/virtual_machine#script_text VirtualMachine#script_text}
	ScriptText *string `field:"optional" json:"scriptText" yaml:"scriptText"`
	// Customize the time zone on the VM. This should be a time zone-style entry, like America/Los_Angeles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/virtual_machine#time_zone VirtualMachine#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

