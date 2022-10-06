package host

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HostConfig struct {
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
	// FQDN or IP address of the host.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/host#hostname Host#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Password of the administration account of the host.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/host#password Host#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Username of the administration account of the host.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/host#username Host#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// ID of the vSphere cluster the host will belong to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/host#cluster Host#cluster}
	Cluster *string `field:"optional" json:"cluster" yaml:"cluster"`
	// Must be set if host is a member of a managed compute_cluster resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/host#cluster_managed Host#cluster_managed}
	ClusterManaged interface{} `field:"optional" json:"clusterManaged" yaml:"clusterManaged"`
	// Set the state of the host. If set to false then the host will be asked to disconnect.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/host#connected Host#connected}
	Connected interface{} `field:"optional" json:"connected" yaml:"connected"`
	// A list of custom attributes to set on this resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/host#custom_attributes Host#custom_attributes}
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// ID of the vSphere datacenter the host will belong to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/host#datacenter Host#datacenter}
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// Force add the host to the vSphere inventory even if it's already managed by a different vCenter Server instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/host#force Host#force}
	Force interface{} `field:"optional" json:"force" yaml:"force"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/host#id Host#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// License key that will be applied to this host.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/host#license Host#license}
	License *string `field:"optional" json:"license" yaml:"license"`
	// Set the host's lockdown status. Default is disabled. Valid options are 'disabled', 'normal', 'strict'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/host#lockdown Host#lockdown}
	Lockdown *string `field:"optional" json:"lockdown" yaml:"lockdown"`
	// Set the host's maintenance mode. Default is false.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/host#maintenance Host#maintenance}
	Maintenance interface{} `field:"optional" json:"maintenance" yaml:"maintenance"`
	// A list of tag IDs to apply to this object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/host#tags Host#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Host's certificate SHA-1 thumbprint. If not set then the CA that signed the host's certificate must be trusted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/host#thumbprint Host#thumbprint}
	Thumbprint *string `field:"optional" json:"thumbprint" yaml:"thumbprint"`
}

