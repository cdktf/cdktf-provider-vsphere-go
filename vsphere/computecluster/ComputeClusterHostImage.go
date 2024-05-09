// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computecluster


type ComputeClusterHostImage struct {
	// component block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/compute_cluster#component ComputeCluster#component}
	Component interface{} `field:"optional" json:"component" yaml:"component"`
	// The ESXi version which the image is based on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/compute_cluster#esx_version ComputeCluster#esx_version}
	EsxVersion *string `field:"optional" json:"esxVersion" yaml:"esxVersion"`
}

