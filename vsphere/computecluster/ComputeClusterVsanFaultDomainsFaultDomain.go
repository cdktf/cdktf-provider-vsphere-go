// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computecluster


type ComputeClusterVsanFaultDomainsFaultDomain struct {
	// The managed object IDs of the hosts to put in the fault domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/compute_cluster#host_ids ComputeCluster#host_ids}
	HostIds *[]*string `field:"required" json:"hostIds" yaml:"hostIds"`
	// The name of fault domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/compute_cluster#name ComputeCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

