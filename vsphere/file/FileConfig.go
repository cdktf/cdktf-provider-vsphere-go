// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package file

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FileConfig struct {
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
	// The name of the datastore to which to upload the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/file#datastore File#datastore}
	Datastore *string `field:"required" json:"datastore" yaml:"datastore"`
	// The path to where the file should be uploaded or copied to on the destination datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/file#destination_file File#destination_file}
	DestinationFile *string `field:"required" json:"destinationFile" yaml:"destinationFile"`
	// The path to the file being uploaded from or copied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/file#source_file File#source_file}
	SourceFile *string `field:"required" json:"sourceFile" yaml:"sourceFile"`
	// Specifies whether to create the parent directories of the destination file if they do not exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/file#create_directories File#create_directories}
	CreateDirectories interface{} `field:"optional" json:"createDirectories" yaml:"createDirectories"`
	// The name of a datacenter to which the file will be uploaded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/file#datacenter File#datacenter}
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/file#id File#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of a datacenter from which the file will be copied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/file#source_datacenter File#source_datacenter}
	SourceDatacenter *string `field:"optional" json:"sourceDatacenter" yaml:"sourceDatacenter"`
	// The name of the datastore from which file will be copied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/file#source_datastore File#source_datastore}
	SourceDatastore *string `field:"optional" json:"sourceDatastore" yaml:"sourceDatastore"`
}

