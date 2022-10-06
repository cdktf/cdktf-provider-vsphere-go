package datavspheredistributedvirtualswitch

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVsphereDistributedVirtualSwitchConfig struct {
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
	// The name of the distributed virtual switch. This can be a name or path.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/d/distributed_virtual_switch#name DataVsphereDistributedVirtualSwitch#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The managed object ID of the datacenter the DVS is in.
	//
	// This is required if the supplied path is not an absolute path containing a datacenter and there are multiple datacenters in your infrastructure.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/d/distributed_virtual_switch#datacenter_id DataVsphereDistributedVirtualSwitch#datacenter_id}
	DatacenterId *string `field:"optional" json:"datacenterId" yaml:"datacenterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/d/distributed_virtual_switch#id DataVsphereDistributedVirtualSwitch#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

