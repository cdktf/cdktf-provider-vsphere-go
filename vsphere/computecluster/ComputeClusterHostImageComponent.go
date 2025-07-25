// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computecluster


type ComputeClusterHostImageComponent struct {
	// The identifier for the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/compute_cluster#key ComputeCluster#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The version to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/compute_cluster#version ComputeCluster#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

