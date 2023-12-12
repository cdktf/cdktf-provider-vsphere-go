// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datavspherehostpcidevice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVsphereHostPciDeviceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The Managed Object ID of the host system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.1/docs/data-sources/host_pci_device#host_id DataVsphereHostPciDevice#host_id}
	HostId *string `field:"required" json:"hostId" yaml:"hostId"`
	// The hexadecimal value of the PCI device's class ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.1/docs/data-sources/host_pci_device#class_id DataVsphereHostPciDevice#class_id}
	ClassId *string `field:"optional" json:"classId" yaml:"classId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.1/docs/data-sources/host_pci_device#id DataVsphereHostPciDevice#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A regular expression used to match the PCI device name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.1/docs/data-sources/host_pci_device#name_regex DataVsphereHostPciDevice#name_regex}
	NameRegex *string `field:"optional" json:"nameRegex" yaml:"nameRegex"`
	// The hexadecimal value of the PCI device's vendor ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.1/docs/data-sources/host_pci_device#vendor_id DataVsphereHostPciDevice#vendor_id}
	VendorId *string `field:"optional" json:"vendorId" yaml:"vendorId"`
}

