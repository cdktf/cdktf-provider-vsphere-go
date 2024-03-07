// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hostportgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HostPortGroupConfig struct {
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
	// The managed object ID of the host to set the virtual switch up on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/host_port_group#host_system_id HostPortGroup#host_system_id}
	HostSystemId *string `field:"required" json:"hostSystemId" yaml:"hostSystemId"`
	// The name of the port group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/host_port_group#name HostPortGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the virtual switch to bind this port group to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/host_port_group#virtual_switch_name HostPortGroup#virtual_switch_name}
	VirtualSwitchName *string `field:"required" json:"virtualSwitchName" yaml:"virtualSwitchName"`
	// List of active network adapters used for load balancing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/host_port_group#active_nics HostPortGroup#active_nics}
	ActiveNics *[]*string `field:"optional" json:"activeNics" yaml:"activeNics"`
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than that of its own.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/host_port_group#allow_forged_transmits HostPortGroup#allow_forged_transmits}
	AllowForgedTransmits interface{} `field:"optional" json:"allowForgedTransmits" yaml:"allowForgedTransmits"`
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/host_port_group#allow_mac_changes HostPortGroup#allow_mac_changes}
	AllowMacChanges interface{} `field:"optional" json:"allowMacChanges" yaml:"allowMacChanges"`
	// Enable promiscuous mode on the network.
	//
	// This flag indicates whether or not all traffic is seen on a given port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/host_port_group#allow_promiscuous HostPortGroup#allow_promiscuous}
	AllowPromiscuous interface{} `field:"optional" json:"allowPromiscuous" yaml:"allowPromiscuous"`
	// Enable beacon probing.
	//
	// Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/host_port_group#check_beacon HostPortGroup#check_beacon}
	CheckBeacon interface{} `field:"optional" json:"checkBeacon" yaml:"checkBeacon"`
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/host_port_group#failback HostPortGroup#failback}
	Failback interface{} `field:"optional" json:"failback" yaml:"failback"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/host_port_group#id HostPortGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/host_port_group#notify_switches HostPortGroup#notify_switches}
	NotifySwitches interface{} `field:"optional" json:"notifySwitches" yaml:"notifySwitches"`
	// The average bandwidth in bits per second if traffic shaping is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/host_port_group#shaping_average_bandwidth HostPortGroup#shaping_average_bandwidth}
	ShapingAverageBandwidth *float64 `field:"optional" json:"shapingAverageBandwidth" yaml:"shapingAverageBandwidth"`
	// The maximum burst size allowed in bytes if traffic shaping is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/host_port_group#shaping_burst_size HostPortGroup#shaping_burst_size}
	ShapingBurstSize *float64 `field:"optional" json:"shapingBurstSize" yaml:"shapingBurstSize"`
	// Enable traffic shaping on this virtual switch or port group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/host_port_group#shaping_enabled HostPortGroup#shaping_enabled}
	ShapingEnabled interface{} `field:"optional" json:"shapingEnabled" yaml:"shapingEnabled"`
	// The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/host_port_group#shaping_peak_bandwidth HostPortGroup#shaping_peak_bandwidth}
	ShapingPeakBandwidth *float64 `field:"optional" json:"shapingPeakBandwidth" yaml:"shapingPeakBandwidth"`
	// List of standby network adapters used for failover.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/host_port_group#standby_nics HostPortGroup#standby_nics}
	StandbyNics *[]*string `field:"optional" json:"standbyNics" yaml:"standbyNics"`
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or failover_explicit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/host_port_group#teaming_policy HostPortGroup#teaming_policy}
	TeamingPolicy *string `field:"optional" json:"teamingPolicy" yaml:"teamingPolicy"`
	// The VLAN ID/trunk mode for this port group.
	//
	// An ID of 0 denotes no tagging, an ID of 1-4094 tags with the specific ID, and an ID of 4095 enables trunk mode, allowing the guest to manage its own tagging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.7.0/docs/resources/host_port_group#vlan_id HostPortGroup#vlan_id}
	VlanId *float64 `field:"optional" json:"vlanId" yaml:"vlanId"`
}

