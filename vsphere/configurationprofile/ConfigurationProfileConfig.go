// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configurationprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigurationProfileConfig struct {
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
	// The identifier of the cluster that will be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.15.0/docs/resources/configuration_profile#cluster_id ConfigurationProfile#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The configuration JSON.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.15.0/docs/resources/configuration_profile#configuration ConfigurationProfile#configuration}
	Configuration *string `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.15.0/docs/resources/configuration_profile#id ConfigurationProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The identifier of the host to use as a source of the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.15.0/docs/resources/configuration_profile#reference_host_id ConfigurationProfile#reference_host_id}
	ReferenceHostId *string `field:"optional" json:"referenceHostId" yaml:"referenceHostId"`
}

