// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computecluster


type ComputeClusterVsanFaultDomains struct {
	// fault_domain block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/compute_cluster#fault_domain ComputeCluster#fault_domain}
	FaultDomain interface{} `field:"optional" json:"faultDomain" yaml:"faultDomain"`
}

