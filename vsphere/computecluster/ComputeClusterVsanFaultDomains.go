// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computecluster


type ComputeClusterVsanFaultDomains struct {
	// fault_domain block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs/resources/compute_cluster#fault_domain ComputeCluster#fault_domain}
	FaultDomain interface{} `field:"optional" json:"faultDomain" yaml:"faultDomain"`
}

