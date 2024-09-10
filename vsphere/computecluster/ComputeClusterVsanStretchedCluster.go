// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computecluster


type ComputeClusterVsanStretchedCluster struct {
	// The managed object IDs of the hosts to put in the first fault domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.1/docs/resources/compute_cluster#preferred_fault_domain_host_ids ComputeCluster#preferred_fault_domain_host_ids}
	PreferredFaultDomainHostIds *[]*string `field:"required" json:"preferredFaultDomainHostIds" yaml:"preferredFaultDomainHostIds"`
	// The managed object IDs of the hosts to put in the second fault domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.1/docs/resources/compute_cluster#secondary_fault_domain_host_ids ComputeCluster#secondary_fault_domain_host_ids}
	SecondaryFaultDomainHostIds *[]*string `field:"required" json:"secondaryFaultDomainHostIds" yaml:"secondaryFaultDomainHostIds"`
	// The managed object IDs of the host selected as witness node when enable stretched cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.1/docs/resources/compute_cluster#witness_node ComputeCluster#witness_node}
	WitnessNode *string `field:"required" json:"witnessNode" yaml:"witnessNode"`
	// The name of prepferred fault domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.1/docs/resources/compute_cluster#preferred_fault_domain_name ComputeCluster#preferred_fault_domain_name}
	PreferredFaultDomainName *string `field:"optional" json:"preferredFaultDomainName" yaml:"preferredFaultDomainName"`
	// The name of secondary fault domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.1/docs/resources/compute_cluster#secondary_fault_domain_name ComputeCluster#secondary_fault_domain_name}
	SecondaryFaultDomainName *string `field:"optional" json:"secondaryFaultDomainName" yaml:"secondaryFaultDomainName"`
}

