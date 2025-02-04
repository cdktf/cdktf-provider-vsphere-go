// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dpmhostoverride

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DpmHostOverrideConfig struct {
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
	// The managed object ID of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/dpm_host_override#compute_cluster_id DpmHostOverride#compute_cluster_id}
	ComputeClusterId *string `field:"required" json:"computeClusterId" yaml:"computeClusterId"`
	// The managed object ID of the host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/dpm_host_override#host_system_id DpmHostOverride#host_system_id}
	HostSystemId *string `field:"required" json:"hostSystemId" yaml:"hostSystemId"`
	// The automation level for power operations on this host. Can be one of manual or automated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/dpm_host_override#dpm_automation_level DpmHostOverride#dpm_automation_level}
	DpmAutomationLevel *string `field:"optional" json:"dpmAutomationLevel" yaml:"dpmAutomationLevel"`
	// Enable DPM for this host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/dpm_host_override#dpm_enabled DpmHostOverride#dpm_enabled}
	DpmEnabled interface{} `field:"optional" json:"dpmEnabled" yaml:"dpmEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/dpm_host_override#id DpmHostOverride#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

