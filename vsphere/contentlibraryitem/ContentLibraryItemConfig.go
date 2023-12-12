// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package contentlibraryitem

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContentLibraryItemConfig struct {
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
	// ID of the content library to contain item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.1/docs/resources/content_library_item#library_id ContentLibraryItem#library_id}
	LibraryId *string `field:"required" json:"libraryId" yaml:"libraryId"`
	// The name of the content library item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.1/docs/resources/content_library_item#name ContentLibraryItem#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional description of the content library item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.1/docs/resources/content_library_item#description ContentLibraryItem#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// ID of source VM of content library item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.1/docs/resources/content_library_item#file_url ContentLibraryItem#file_url}
	FileUrl *string `field:"optional" json:"fileUrl" yaml:"fileUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.1/docs/resources/content_library_item#id ContentLibraryItem#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The managed object ID of an existing VM to be cloned to the content library.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.1/docs/resources/content_library_item#source_uuid ContentLibraryItem#source_uuid}
	SourceUuid *string `field:"optional" json:"sourceUuid" yaml:"sourceUuid"`
	// Type of content library item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.1/docs/resources/content_library_item#type ContentLibraryItem#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

