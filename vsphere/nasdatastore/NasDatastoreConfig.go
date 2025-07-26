// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nasdatastore

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NasDatastoreConfig struct {
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
	// The managed object IDs of the hosts to mount the datastore on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/nas_datastore#host_system_ids NasDatastore#host_system_ids}
	HostSystemIds *[]*string `field:"required" json:"hostSystemIds" yaml:"hostSystemIds"`
	// The name of the datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/nas_datastore#name NasDatastore#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The hostnames or IP addresses of the remote server or servers.
	//
	// Only one element should be present for NFS v3 but multiple can be present for NFS v4.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/nas_datastore#remote_hosts NasDatastore#remote_hosts}
	RemoteHosts *[]*string `field:"required" json:"remoteHosts" yaml:"remoteHosts"`
	// The remote path of the mount point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/nas_datastore#remote_path NasDatastore#remote_path}
	RemotePath *string `field:"required" json:"remotePath" yaml:"remotePath"`
	// Access mode for the mount point. Can be one of readOnly or readWrite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/nas_datastore#access_mode NasDatastore#access_mode}
	AccessMode *string `field:"optional" json:"accessMode" yaml:"accessMode"`
	// A list of custom attributes to set on this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/nas_datastore#custom_attributes NasDatastore#custom_attributes}
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The managed object ID of the datastore cluster to place the datastore in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/nas_datastore#datastore_cluster_id NasDatastore#datastore_cluster_id}
	DatastoreClusterId *string `field:"optional" json:"datastoreClusterId" yaml:"datastoreClusterId"`
	// The path to the datastore folder to put the datastore in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/nas_datastore#folder NasDatastore#folder}
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/nas_datastore#id NasDatastore#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The security type to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/nas_datastore#security_type NasDatastore#security_type}
	SecurityType *string `field:"optional" json:"securityType" yaml:"securityType"`
	// A list of tag IDs to apply to this object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/nas_datastore#tags NasDatastore#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// The type of NAS volume. Can be one of NFS (to denote v3) or NFS41 (to denote NFS v4.1).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/nas_datastore#type NasDatastore#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

