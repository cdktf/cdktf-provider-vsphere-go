// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vnic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VnicConfig struct {
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
	// ESX host the interface belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/vnic#host Vnic#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Key of the distributed portgroup the nic will connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/vnic#distributed_port_group Vnic#distributed_port_group}
	DistributedPortGroup *string `field:"optional" json:"distributedPortGroup" yaml:"distributedPortGroup"`
	// UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/vnic#distributed_switch_port Vnic#distributed_switch_port}
	DistributedSwitchPort *string `field:"optional" json:"distributedSwitchPort" yaml:"distributedSwitchPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/vnic#id Vnic#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ipv4 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/vnic#ipv4 Vnic#ipv4}
	Ipv4 *VnicIpv4 `field:"optional" json:"ipv4" yaml:"ipv4"`
	// ipv6 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/vnic#ipv6 Vnic#ipv6}
	Ipv6 *VnicIpv6 `field:"optional" json:"ipv6" yaml:"ipv6"`
	// MAC address of the interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/vnic#mac Vnic#mac}
	Mac *string `field:"optional" json:"mac" yaml:"mac"`
	// MTU of the interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/vnic#mtu Vnic#mtu}
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// TCP/IP stack setting for this interface. Possible values are 'defaultTcpipStack', 'vmotion', 'provisioning'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/vnic#netstack Vnic#netstack}
	Netstack *string `field:"optional" json:"netstack" yaml:"netstack"`
	// portgroup to attach the nic to. Do not set if you set distributed_switch_port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/vnic#portgroup Vnic#portgroup}
	Portgroup *string `field:"optional" json:"portgroup" yaml:"portgroup"`
	// Enabled services setting for this interface. Current possible values are 'vmotion', 'management' and 'vsan'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/vnic#services Vnic#services}
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
}

