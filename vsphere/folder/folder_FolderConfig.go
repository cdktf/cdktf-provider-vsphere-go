package folder

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FolderConfig struct {
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
	// The path of the folder and any parents, relative to the datacenter and folder type being defined.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/folder#path Folder#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// The type of the folder.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/folder#type Folder#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A list of custom attributes to set on this resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/folder#custom_attributes Folder#custom_attributes}
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The ID of the datacenter. Can be ignored if creating a datacenter folder, otherwise required.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/folder#datacenter_id Folder#datacenter_id}
	DatacenterId *string `field:"optional" json:"datacenterId" yaml:"datacenterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/folder#id Folder#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A list of tag IDs to apply to this object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/folder#tags Folder#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

