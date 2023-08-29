// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcepool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ResourcePoolConfig struct {
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
	// Name of resource pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/resource_pool#name ResourcePool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the root resource pool of the compute resource the resource pool is in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/resource_pool#parent_resource_pool_id ResourcePool#parent_resource_pool_id}
	ParentResourcePoolId *string `field:"required" json:"parentResourcePoolId" yaml:"parentResourcePoolId"`
	// Determines if the reservation on a resource pool can grow beyond the specified value, if the parent resource pool has unreserved resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/resource_pool#cpu_expandable ResourcePool#cpu_expandable}
	CpuExpandable interface{} `field:"optional" json:"cpuExpandable" yaml:"cpuExpandable"`
	// The utilization of a resource pool will not exceed this limit, even if there are available resources.
	//
	// Set to -1 for unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/resource_pool#cpu_limit ResourcePool#cpu_limit}
	CpuLimit *float64 `field:"optional" json:"cpuLimit" yaml:"cpuLimit"`
	// Amount of CPU (MHz) that is guaranteed available to the resource pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/resource_pool#cpu_reservation ResourcePool#cpu_reservation}
	CpuReservation *float64 `field:"optional" json:"cpuReservation" yaml:"cpuReservation"`
	// The allocation level.
	//
	// The level is a simplified view of shares. Levels map to a pre-determined set of numeric values for shares. Can be one of low, normal, high, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/resource_pool#cpu_share_level ResourcePool#cpu_share_level}
	CpuShareLevel *string `field:"optional" json:"cpuShareLevel" yaml:"cpuShareLevel"`
	// The number of shares allocated.
	//
	// Used to determine resource allocation in case of resource contention. If this is set, cpu_share_level must be custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/resource_pool#cpu_shares ResourcePool#cpu_shares}
	CpuShares *float64 `field:"optional" json:"cpuShares" yaml:"cpuShares"`
	// A list of custom attributes to set on this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/resource_pool#custom_attributes ResourcePool#custom_attributes}
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/resource_pool#id ResourcePool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Determines if the reservation on a resource pool can grow beyond the specified value, if the parent resource pool has unreserved resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/resource_pool#memory_expandable ResourcePool#memory_expandable}
	MemoryExpandable interface{} `field:"optional" json:"memoryExpandable" yaml:"memoryExpandable"`
	// The utilization of a resource pool will not exceed this limit, even if there are available resources.
	//
	// Set to -1 for unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/resource_pool#memory_limit ResourcePool#memory_limit}
	MemoryLimit *float64 `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
	// Amount of memory (MB) that is guaranteed available to the resource pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/resource_pool#memory_reservation ResourcePool#memory_reservation}
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
	// The allocation level.
	//
	// The level is a simplified view of shares. Levels map to a pre-determined set of numeric values for shares. Can be one of low, normal, high, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/resource_pool#memory_share_level ResourcePool#memory_share_level}
	MemoryShareLevel *string `field:"optional" json:"memoryShareLevel" yaml:"memoryShareLevel"`
	// The number of shares allocated.
	//
	// Used to determine resource allocation in case of resource contention. If this is set, memory_share_level must be custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/resource_pool#memory_shares ResourcePool#memory_shares}
	MemoryShares *float64 `field:"optional" json:"memoryShares" yaml:"memoryShares"`
	// Determines if the shares of all descendants of the resource pool are scaled up or down when the shares of the resource pool are scaled up or down.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/resource_pool#scale_descendants_shares ResourcePool#scale_descendants_shares}
	ScaleDescendantsShares *string `field:"optional" json:"scaleDescendantsShares" yaml:"scaleDescendantsShares"`
	// A list of tag IDs to apply to this object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/resource_pool#tags ResourcePool#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

