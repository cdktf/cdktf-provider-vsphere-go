package computeclustervmaffinityrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeClusterVmAffinityRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/compute_cluster_vm_affinity_rule#compute_cluster_id ComputeClusterVmAffinityRule#compute_cluster_id}
	ComputeClusterId *string `field:"required" json:"computeClusterId" yaml:"computeClusterId"`
	// The unique name of the virtual machine group in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/compute_cluster_vm_affinity_rule#name ComputeClusterVmAffinityRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The UUIDs of the virtual machines to run on the same host together.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/compute_cluster_vm_affinity_rule#virtual_machine_ids ComputeClusterVmAffinityRule#virtual_machine_ids}
	VirtualMachineIds *[]*string `field:"required" json:"virtualMachineIds" yaml:"virtualMachineIds"`
	// Enable this rule in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/compute_cluster_vm_affinity_rule#enabled ComputeClusterVmAffinityRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/compute_cluster_vm_affinity_rule#id ComputeClusterVmAffinityRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When true, prevents any virtual machine operations that may violate this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/compute_cluster_vm_affinity_rule#mandatory ComputeClusterVmAffinityRule#mandatory}
	Mandatory interface{} `field:"optional" json:"mandatory" yaml:"mandatory"`
}

