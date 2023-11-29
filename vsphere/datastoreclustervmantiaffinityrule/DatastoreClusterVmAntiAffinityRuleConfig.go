// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastoreclustervmantiaffinityrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatastoreClusterVmAntiAffinityRuleConfig struct {
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
	// The managed object ID of the datastore cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.0/docs/resources/datastore_cluster_vm_anti_affinity_rule#datastore_cluster_id DatastoreClusterVmAntiAffinityRule#datastore_cluster_id}
	DatastoreClusterId *string `field:"required" json:"datastoreClusterId" yaml:"datastoreClusterId"`
	// The unique name of the virtual machine group in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.0/docs/resources/datastore_cluster_vm_anti_affinity_rule#name DatastoreClusterVmAntiAffinityRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The UUIDs of the virtual machines to run on different datastores from each other.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.0/docs/resources/datastore_cluster_vm_anti_affinity_rule#virtual_machine_ids DatastoreClusterVmAntiAffinityRule#virtual_machine_ids}
	VirtualMachineIds *[]*string `field:"required" json:"virtualMachineIds" yaml:"virtualMachineIds"`
	// Enable this rule in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.0/docs/resources/datastore_cluster_vm_anti_affinity_rule#enabled DatastoreClusterVmAntiAffinityRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.0/docs/resources/datastore_cluster_vm_anti_affinity_rule#id DatastoreClusterVmAntiAffinityRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When true, prevents any virtual machine operations that may violate this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.0/docs/resources/datastore_cluster_vm_anti_affinity_rule#mandatory DatastoreClusterVmAntiAffinityRule#mandatory}
	Mandatory interface{} `field:"optional" json:"mandatory" yaml:"mandatory"`
}

