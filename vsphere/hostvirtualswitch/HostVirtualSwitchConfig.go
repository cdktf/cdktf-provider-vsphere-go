// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hostvirtualswitch

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HostVirtualSwitchConfig struct {
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
	// List of active network adapters used for load balancing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#active_nics HostVirtualSwitch#active_nics}
	ActiveNics *[]*string `field:"required" json:"activeNics" yaml:"activeNics"`
	// The managed object ID of the host to set the virtual switch up on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#host_system_id HostVirtualSwitch#host_system_id}
	HostSystemId *string `field:"required" json:"hostSystemId" yaml:"hostSystemId"`
	// The name of the virtual switch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#name HostVirtualSwitch#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of network adapters to bind to this virtual switch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#network_adapters HostVirtualSwitch#network_adapters}
	NetworkAdapters *[]*string `field:"required" json:"networkAdapters" yaml:"networkAdapters"`
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than that of its own.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#allow_forged_transmits HostVirtualSwitch#allow_forged_transmits}
	AllowForgedTransmits interface{} `field:"optional" json:"allowForgedTransmits" yaml:"allowForgedTransmits"`
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#allow_mac_changes HostVirtualSwitch#allow_mac_changes}
	AllowMacChanges interface{} `field:"optional" json:"allowMacChanges" yaml:"allowMacChanges"`
	// Enable promiscuous mode on the network.
	//
	// This flag indicates whether or not all traffic is seen on a given port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#allow_promiscuous HostVirtualSwitch#allow_promiscuous}
	AllowPromiscuous interface{} `field:"optional" json:"allowPromiscuous" yaml:"allowPromiscuous"`
	// Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#beacon_interval HostVirtualSwitch#beacon_interval}
	BeaconInterval *float64 `field:"optional" json:"beaconInterval" yaml:"beaconInterval"`
	// Enable beacon probing.
	//
	// Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#check_beacon HostVirtualSwitch#check_beacon}
	CheckBeacon interface{} `field:"optional" json:"checkBeacon" yaml:"checkBeacon"`
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#failback HostVirtualSwitch#failback}
	Failback interface{} `field:"optional" json:"failback" yaml:"failback"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#id HostVirtualSwitch#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#link_discovery_operation HostVirtualSwitch#link_discovery_operation}
	LinkDiscoveryOperation *string `field:"optional" json:"linkDiscoveryOperation" yaml:"linkDiscoveryOperation"`
	// The discovery protocol type. Valid values are cdp and lldp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#link_discovery_protocol HostVirtualSwitch#link_discovery_protocol}
	LinkDiscoveryProtocol *string `field:"optional" json:"linkDiscoveryProtocol" yaml:"linkDiscoveryProtocol"`
	// The maximum transmission unit (MTU) of the virtual switch in bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#mtu HostVirtualSwitch#mtu}
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#notify_switches HostVirtualSwitch#notify_switches}
	NotifySwitches interface{} `field:"optional" json:"notifySwitches" yaml:"notifySwitches"`
	// The number of ports that this virtual switch is configured to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#number_of_ports HostVirtualSwitch#number_of_ports}
	NumberOfPorts *float64 `field:"optional" json:"numberOfPorts" yaml:"numberOfPorts"`
	// The average bandwidth in bits per second if traffic shaping is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#shaping_average_bandwidth HostVirtualSwitch#shaping_average_bandwidth}
	ShapingAverageBandwidth *float64 `field:"optional" json:"shapingAverageBandwidth" yaml:"shapingAverageBandwidth"`
	// The maximum burst size allowed in bytes if traffic shaping is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#shaping_burst_size HostVirtualSwitch#shaping_burst_size}
	ShapingBurstSize *float64 `field:"optional" json:"shapingBurstSize" yaml:"shapingBurstSize"`
	// Enable traffic shaping on this virtual switch or port group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#shaping_enabled HostVirtualSwitch#shaping_enabled}
	ShapingEnabled interface{} `field:"optional" json:"shapingEnabled" yaml:"shapingEnabled"`
	// The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#shaping_peak_bandwidth HostVirtualSwitch#shaping_peak_bandwidth}
	ShapingPeakBandwidth *float64 `field:"optional" json:"shapingPeakBandwidth" yaml:"shapingPeakBandwidth"`
	// List of standby network adapters used for failover.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#standby_nics HostVirtualSwitch#standby_nics}
	StandbyNics *[]*string `field:"optional" json:"standbyNics" yaml:"standbyNics"`
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or failover_explicit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/host_virtual_switch#teaming_policy HostVirtualSwitch#teaming_policy}
	TeamingPolicy *string `field:"optional" json:"teamingPolicy" yaml:"teamingPolicy"`
}

