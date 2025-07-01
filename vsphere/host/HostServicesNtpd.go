// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package host


type HostServicesNtpd struct {
	// Whether the NTP service is enabled. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.0/docs/resources/host#enabled Host#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.0/docs/resources/host#ntp_servers Host#ntp_servers}.
	NtpServers *[]*string `field:"optional" json:"ntpServers" yaml:"ntpServers"`
	// The policy for the NTP service.
	//
	// Valid values are 'Start and stop with host', 'Start and stop manually', 'Start and stop with port usage'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.0/docs/resources/host#policy Host#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

