// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualdisk

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualDiskConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.0/docs/resources/virtual_disk#datastore VirtualDisk#datastore}.
	Datastore *string `field:"required" json:"datastore" yaml:"datastore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.0/docs/resources/virtual_disk#size VirtualDisk#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.0/docs/resources/virtual_disk#vmdk_path VirtualDisk#vmdk_path}.
	VmdkPath *string `field:"required" json:"vmdkPath" yaml:"vmdkPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.0/docs/resources/virtual_disk#adapter_type VirtualDisk#adapter_type}.
	AdapterType *string `field:"optional" json:"adapterType" yaml:"adapterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.0/docs/resources/virtual_disk#create_directories VirtualDisk#create_directories}.
	CreateDirectories interface{} `field:"optional" json:"createDirectories" yaml:"createDirectories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.0/docs/resources/virtual_disk#datacenter VirtualDisk#datacenter}.
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.0/docs/resources/virtual_disk#id VirtualDisk#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.0/docs/resources/virtual_disk#type VirtualDisk#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

