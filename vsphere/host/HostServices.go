// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package host


type HostServices struct {
	// ntpd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/host#ntpd Host#ntpd}
	Ntpd *HostServicesNtpd `field:"optional" json:"ntpd" yaml:"ntpd"`
}

