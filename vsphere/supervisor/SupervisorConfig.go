// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package supervisor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SupervisorConfig struct {
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
	// ID of the vSphere cluster on which workload management will be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#cluster Supervisor#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// ID of the subscribed content library.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#content_library Supervisor#content_library}
	ContentLibrary *string `field:"required" json:"contentLibrary" yaml:"contentLibrary"`
	// The UUID (not ID) of the distributed switch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#dvs_uuid Supervisor#dvs_uuid}
	DvsUuid *string `field:"required" json:"dvsUuid" yaml:"dvsUuid"`
	// ID of the NSX Edge Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#edge_cluster Supervisor#edge_cluster}
	EdgeCluster *string `field:"required" json:"edgeCluster" yaml:"edgeCluster"`
	// egress_cidr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#egress_cidr Supervisor#egress_cidr}
	EgressCidr interface{} `field:"required" json:"egressCidr" yaml:"egressCidr"`
	// ingress_cidr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#ingress_cidr Supervisor#ingress_cidr}
	IngressCidr interface{} `field:"required" json:"ingressCidr" yaml:"ingressCidr"`
	// List of DNS servers to use on the Kubernetes API server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#main_dns Supervisor#main_dns}
	MainDns *[]*string `field:"required" json:"mainDns" yaml:"mainDns"`
	// management_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#management_network Supervisor#management_network}
	ManagementNetwork *SupervisorManagementNetwork `field:"required" json:"managementNetwork" yaml:"managementNetwork"`
	// pod_cidr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#pod_cidr Supervisor#pod_cidr}
	PodCidr interface{} `field:"required" json:"podCidr" yaml:"podCidr"`
	// List of DNS search domains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#search_domains Supervisor#search_domains}
	SearchDomains *[]*string `field:"required" json:"searchDomains" yaml:"searchDomains"`
	// service_cidr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#service_cidr Supervisor#service_cidr}
	ServiceCidr *SupervisorServiceCidr `field:"required" json:"serviceCidr" yaml:"serviceCidr"`
	// Size of the Kubernetes API server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#sizing_hint Supervisor#sizing_hint}
	SizingHint *string `field:"required" json:"sizingHint" yaml:"sizingHint"`
	// The name of a storage policy associated with the datastore where the container images will be stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#storage_policy Supervisor#storage_policy}
	StoragePolicy *string `field:"required" json:"storagePolicy" yaml:"storagePolicy"`
	// List of DNS servers to use on the worker nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#worker_dns Supervisor#worker_dns}
	WorkerDns *[]*string `field:"required" json:"workerDns" yaml:"workerDns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#id Supervisor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// namespace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#namespace Supervisor#namespace}
	Namespace interface{} `field:"optional" json:"namespace" yaml:"namespace"`
}

