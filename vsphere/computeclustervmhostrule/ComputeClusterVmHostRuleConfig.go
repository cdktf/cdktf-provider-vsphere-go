// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeclustervmhostrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeClusterVmHostRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/compute_cluster_vm_host_rule#compute_cluster_id ComputeClusterVmHostRule#compute_cluster_id}
	ComputeClusterId *string `field:"required" json:"computeClusterId" yaml:"computeClusterId"`
	// The unique name of the virtual machine group in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/compute_cluster_vm_host_rule#name ComputeClusterVmHostRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the virtual machine group to use with this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/compute_cluster_vm_host_rule#vm_group_name ComputeClusterVmHostRule#vm_group_name}
	VmGroupName *string `field:"required" json:"vmGroupName" yaml:"vmGroupName"`
	// When this field is used, virtual machines defined in vm_group_name will be run on the hosts defined in this host group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/compute_cluster_vm_host_rule#affinity_host_group_name ComputeClusterVmHostRule#affinity_host_group_name}
	AffinityHostGroupName *string `field:"optional" json:"affinityHostGroupName" yaml:"affinityHostGroupName"`
	// When this field is used, virtual machines defined in vm_group_name will not be run on the hosts defined in this host group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/compute_cluster_vm_host_rule#anti_affinity_host_group_name ComputeClusterVmHostRule#anti_affinity_host_group_name}
	AntiAffinityHostGroupName *string `field:"optional" json:"antiAffinityHostGroupName" yaml:"antiAffinityHostGroupName"`
	// Enable this rule in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/compute_cluster_vm_host_rule#enabled ComputeClusterVmHostRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/compute_cluster_vm_host_rule#id ComputeClusterVmHostRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When true, prevents any virtual machine operations that may violate this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.1/docs/resources/compute_cluster_vm_host_rule#mandatory ComputeClusterVmHostRule#mandatory}
	Mandatory interface{} `field:"optional" json:"mandatory" yaml:"mandatory"`
}

