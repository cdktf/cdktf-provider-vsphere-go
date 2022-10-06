package tagcategory

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TagCategoryConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Object types to which this category's tags can be attached.
	//
	// Valid types include: Folder, ClusterComputeResource, Datacenter, Datastore, StoragePod, DistributedVirtualPortgroup, DistributedVirtualSwitch, VmwareDistributedVirtualSwitch, HostSystem, com.vmware.content.Library, com.vmware.content.library.Item, HostNetwork, Network, OpaqueNetwork, ResourcePool, VirtualApp, VirtualMachine.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/tag_category#associable_types TagCategory#associable_types}
	AssociableTypes *[]*string `field:"required" json:"associableTypes" yaml:"associableTypes"`
	// The associated cardinality of the category.
	//
	// Can be one of SINGLE (object can only be assigned one tag in this category) or MULTIPLE (object can be assigned multiple tags in this category).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/tag_category#cardinality TagCategory#cardinality}
	Cardinality *string `field:"required" json:"cardinality" yaml:"cardinality"`
	// The display name of the category.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/tag_category#name TagCategory#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the category.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/tag_category#description TagCategory#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/tag_category#id TagCategory#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

