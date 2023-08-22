package vappcontainer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VappContainerConfig struct {
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
	// The name of the vApp container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/vapp_container#name VappContainer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The managed object ID of the parent resource pool or the compute resource the vApp container is in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/vapp_container#parent_resource_pool_id VappContainer#parent_resource_pool_id}
	ParentResourcePoolId *string `field:"required" json:"parentResourcePoolId" yaml:"parentResourcePoolId"`
	// Determines if the reservation on a vApp container can grow beyond the specified value, if the parent resource pool has unreserved resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/vapp_container#cpu_expandable VappContainer#cpu_expandable}
	CpuExpandable interface{} `field:"optional" json:"cpuExpandable" yaml:"cpuExpandable"`
	// The utilization of a vApp container will not exceed this limit, even if there are available resources.
	//
	// Set to -1 for unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/vapp_container#cpu_limit VappContainer#cpu_limit}
	CpuLimit *float64 `field:"optional" json:"cpuLimit" yaml:"cpuLimit"`
	// Amount of CPU (MHz) that is guaranteed available to the vApp container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/vapp_container#cpu_reservation VappContainer#cpu_reservation}
	CpuReservation *float64 `field:"optional" json:"cpuReservation" yaml:"cpuReservation"`
	// The allocation level.
	//
	// The level is a simplified view of shares. Levels map to a pre-determined set of numeric values for shares. Can be one of low, normal, high, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/vapp_container#cpu_share_level VappContainer#cpu_share_level}
	CpuShareLevel *string `field:"optional" json:"cpuShareLevel" yaml:"cpuShareLevel"`
	// The number of shares allocated.
	//
	// Used to determine resource allocation in case of resource contention. If this is set, cpu_share_level must be custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/vapp_container#cpu_shares VappContainer#cpu_shares}
	CpuShares *float64 `field:"optional" json:"cpuShares" yaml:"cpuShares"`
	// A list of custom attributes to set on this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/vapp_container#custom_attributes VappContainer#custom_attributes}
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/vapp_container#id VappContainer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Determines if the reservation on a vApp container can grow beyond the specified value, if the parent resource pool has unreserved resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/vapp_container#memory_expandable VappContainer#memory_expandable}
	MemoryExpandable interface{} `field:"optional" json:"memoryExpandable" yaml:"memoryExpandable"`
	// The utilization of a vApp container will not exceed this limit, even if there are available resources.
	//
	// Set to -1 for unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/vapp_container#memory_limit VappContainer#memory_limit}
	MemoryLimit *float64 `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
	// Amount of memory (MB) that is guaranteed available to the vApp container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/vapp_container#memory_reservation VappContainer#memory_reservation}
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
	// The allocation level.
	//
	// The level is a simplified view of shares. Levels map to a pre-determined set of numeric values for shares. Can be one of low, normal, high, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/vapp_container#memory_share_level VappContainer#memory_share_level}
	MemoryShareLevel *string `field:"optional" json:"memoryShareLevel" yaml:"memoryShareLevel"`
	// The number of shares allocated.
	//
	// Used to determine resource allocation in case of resource contention. If this is set, memory_share_level must be custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/vapp_container#memory_shares VappContainer#memory_shares}
	MemoryShares *float64 `field:"optional" json:"memoryShares" yaml:"memoryShares"`
	// The ID of the parent VM folder.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/vapp_container#parent_folder_id VappContainer#parent_folder_id}
	ParentFolderId *string `field:"optional" json:"parentFolderId" yaml:"parentFolderId"`
	// A list of tag IDs to apply to this object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/vapp_container#tags VappContainer#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

