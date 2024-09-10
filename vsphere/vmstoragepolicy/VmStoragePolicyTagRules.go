// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmstoragepolicy


type VmStoragePolicyTagRules struct {
	// The tag category to select the tags from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.1/docs/resources/vm_storage_policy#tag_category VmStoragePolicy#tag_category}
	TagCategory *string `field:"required" json:"tagCategory" yaml:"tagCategory"`
	// The tags to use for creating a tag-based vm placement rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.1/docs/resources/vm_storage_policy#tags VmStoragePolicy#tags}
	Tags *[]*string `field:"required" json:"tags" yaml:"tags"`
	// Whether to include or exclude datastores tagged with the provided tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.1/docs/resources/vm_storage_policy#include_datastores_with_tags VmStoragePolicy#include_datastores_with_tags}
	IncludeDatastoresWithTags interface{} `field:"optional" json:"includeDatastoresWithTags" yaml:"includeDatastoresWithTags"`
}

