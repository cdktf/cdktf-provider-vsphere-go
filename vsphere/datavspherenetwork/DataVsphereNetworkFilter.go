// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datavspherenetwork


type DataVsphereNetworkFilter struct {
	// The type of the network (e.g., Network, DistributedVirtualPortgroup, OpaqueNetwork).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/data-sources/network#network_type DataVsphereNetwork#network_type}
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
}

