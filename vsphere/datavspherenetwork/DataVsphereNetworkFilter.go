// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datavspherenetwork


type DataVsphereNetworkFilter struct {
	// The type of the network (e.g., Network, DistributedVirtualPortgroup, OpaqueNetwork).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/data-sources/network#network_type DataVsphereNetwork#network_type}
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
}

