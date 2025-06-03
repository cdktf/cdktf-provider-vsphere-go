// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package contentlibrary


type ContentLibrarySubscription struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/content_library#authentication_method ContentLibrary#authentication_method}.
	AuthenticationMethod *string `field:"optional" json:"authenticationMethod" yaml:"authenticationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/content_library#automatic_sync ContentLibrary#automatic_sync}.
	AutomaticSync interface{} `field:"optional" json:"automaticSync" yaml:"automaticSync"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/content_library#on_demand ContentLibrary#on_demand}.
	OnDemand interface{} `field:"optional" json:"onDemand" yaml:"onDemand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/content_library#password ContentLibrary#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/content_library#subscription_url ContentLibrary#subscription_url}.
	SubscriptionUrl *string `field:"optional" json:"subscriptionUrl" yaml:"subscriptionUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/content_library#username ContentLibrary#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

