// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datavspherevmfsdisks

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVsphereVmfsDisksConfig struct {
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
	// The managed object ID of the host to search for disks on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/data-sources/vmfs_disks#host_system_id DataVsphereVmfsDisks#host_system_id}
	HostSystemId *string `field:"required" json:"hostSystemId" yaml:"hostSystemId"`
	// A regular expression to filter the disks against. Only disks with canonical names that match will be included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/data-sources/vmfs_disks#filter DataVsphereVmfsDisks#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/data-sources/vmfs_disks#id DataVsphereVmfsDisks#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Rescan the system for disks before querying. This may lengthen the time it takes to gather information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.0/docs/data-sources/vmfs_disks#rescan DataVsphereVmfsDisks#rescan}
	Rescan interface{} `field:"optional" json:"rescan" yaml:"rescan"`
}

