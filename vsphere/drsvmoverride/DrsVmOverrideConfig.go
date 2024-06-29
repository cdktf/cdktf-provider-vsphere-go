// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package drsvmoverride

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DrsVmOverrideConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/resources/drs_vm_override#compute_cluster_id DrsVmOverride#compute_cluster_id}
	ComputeClusterId *string `field:"required" json:"computeClusterId" yaml:"computeClusterId"`
	// The managed object ID of the virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/resources/drs_vm_override#virtual_machine_id DrsVmOverride#virtual_machine_id}
	VirtualMachineId *string `field:"required" json:"virtualMachineId" yaml:"virtualMachineId"`
	// The automation level for this virtual machine in the cluster. Can be one of manual, partiallyAutomated, or fullyAutomated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/resources/drs_vm_override#drs_automation_level DrsVmOverride#drs_automation_level}
	DrsAutomationLevel *string `field:"optional" json:"drsAutomationLevel" yaml:"drsAutomationLevel"`
	// Enable DRS for this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/resources/drs_vm_override#drs_enabled DrsVmOverride#drs_enabled}
	DrsEnabled interface{} `field:"optional" json:"drsEnabled" yaml:"drsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/resources/drs_vm_override#id DrsVmOverride#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

