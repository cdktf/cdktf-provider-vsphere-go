// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datavsphereovfvmtemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVsphereOvfVmTemplateConfig struct {
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
	// The ID of an optional host system to pin the virtual machine to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/data-sources/ovf_vm_template#host_system_id DataVsphereOvfVmTemplate#host_system_id}
	HostSystemId *string `field:"required" json:"hostSystemId" yaml:"hostSystemId"`
	// Name of the virtual machine to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/data-sources/ovf_vm_template#name DataVsphereOvfVmTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of a resource pool to put the virtual machine in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/data-sources/ovf_vm_template#resource_pool_id DataVsphereOvfVmTemplate#resource_pool_id}
	ResourcePoolId *string `field:"required" json:"resourcePoolId" yaml:"resourcePoolId"`
	// Allow unverified ssl certificates while deploying ovf/ova from url.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/data-sources/ovf_vm_template#allow_unverified_ssl_cert DataVsphereOvfVmTemplate#allow_unverified_ssl_cert}
	AllowUnverifiedSslCert interface{} `field:"optional" json:"allowUnverifiedSslCert" yaml:"allowUnverifiedSslCert"`
	// The ID of the virtual machine's datastore.
	//
	// The virtual machine configuration is placed here, along with any virtual disks that are created without datastores.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/data-sources/ovf_vm_template#datastore_id DataVsphereOvfVmTemplate#datastore_id}
	DatastoreId *string `field:"optional" json:"datastoreId" yaml:"datastoreId"`
	// The Deployment option to be chosen. If empty, the default option is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/data-sources/ovf_vm_template#deployment_option DataVsphereOvfVmTemplate#deployment_option}
	DeploymentOption *string `field:"optional" json:"deploymentOption" yaml:"deploymentOption"`
	// An optional disk provisioning.
	//
	// If set, all the disks in the deployed ovf will have the same specified disk type (e.g., thin provisioned).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/data-sources/ovf_vm_template#disk_provisioning DataVsphereOvfVmTemplate#disk_provisioning}
	DiskProvisioning *string `field:"optional" json:"diskProvisioning" yaml:"diskProvisioning"`
	// Allow properties with ovf:userConfigurable=false to be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/data-sources/ovf_vm_template#enable_hidden_properties DataVsphereOvfVmTemplate#enable_hidden_properties}
	EnableHiddenProperties interface{} `field:"optional" json:"enableHiddenProperties" yaml:"enableHiddenProperties"`
	// The name of the folder to locate the virtual machine in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/data-sources/ovf_vm_template#folder DataVsphereOvfVmTemplate#folder}
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/data-sources/ovf_vm_template#id DataVsphereOvfVmTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The IP allocation policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/data-sources/ovf_vm_template#ip_allocation_policy DataVsphereOvfVmTemplate#ip_allocation_policy}
	IpAllocationPolicy *string `field:"optional" json:"ipAllocationPolicy" yaml:"ipAllocationPolicy"`
	// The IP protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/data-sources/ovf_vm_template#ip_protocol DataVsphereOvfVmTemplate#ip_protocol}
	IpProtocol *string `field:"optional" json:"ipProtocol" yaml:"ipProtocol"`
	// The absolute path to the ovf/ova file in the local system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/data-sources/ovf_vm_template#local_ovf_path DataVsphereOvfVmTemplate#local_ovf_path}
	LocalOvfPath *string `field:"optional" json:"localOvfPath" yaml:"localOvfPath"`
	// The mapping of name of network identifiers from the ovf descriptor to network UUID in the VI infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/data-sources/ovf_vm_template#ovf_network_map DataVsphereOvfVmTemplate#ovf_network_map}
	OvfNetworkMap *map[string]*string `field:"optional" json:"ovfNetworkMap" yaml:"ovfNetworkMap"`
	// URL to the remote ovf/ova file to be deployed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/data-sources/ovf_vm_template#remote_ovf_url DataVsphereOvfVmTemplate#remote_ovf_url}
	RemoteOvfUrl *string `field:"optional" json:"remoteOvfUrl" yaml:"remoteOvfUrl"`
}

