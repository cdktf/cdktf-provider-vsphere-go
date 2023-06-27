package distributedportgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DistributedPortGroupConfig struct {
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
	// The UUID of the DVS to attach this port group to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#distributed_virtual_switch_uuid DistributedPortGroup#distributed_virtual_switch_uuid}
	DistributedVirtualSwitchUuid *string `field:"required" json:"distributedVirtualSwitchUuid" yaml:"distributedVirtualSwitchUuid"`
	// The name of the portgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#name DistributedPortGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#active_uplinks DistributedPortGroup#active_uplinks}
	ActiveUplinks *[]*string `field:"optional" json:"activeUplinks" yaml:"activeUplinks"`
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than that of its own.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#allow_forged_transmits DistributedPortGroup#allow_forged_transmits}
	AllowForgedTransmits interface{} `field:"optional" json:"allowForgedTransmits" yaml:"allowForgedTransmits"`
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#allow_mac_changes DistributedPortGroup#allow_mac_changes}
	AllowMacChanges interface{} `field:"optional" json:"allowMacChanges" yaml:"allowMacChanges"`
	// Enable promiscuous mode on the network.
	//
	// This flag indicates whether or not all traffic is seen on a given port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#allow_promiscuous DistributedPortGroup#allow_promiscuous}
	AllowPromiscuous interface{} `field:"optional" json:"allowPromiscuous" yaml:"allowPromiscuous"`
	// Auto-expands the port group beyond the port count configured in number_of_ports when necessary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#auto_expand DistributedPortGroup#auto_expand}
	AutoExpand interface{} `field:"optional" json:"autoExpand" yaml:"autoExpand"`
	// Indicates whether to block all ports by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#block_all_ports DistributedPortGroup#block_all_ports}
	BlockAllPorts interface{} `field:"optional" json:"blockAllPorts" yaml:"blockAllPorts"`
	// Allow the blocked setting of an individual port to override the setting in the portgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#block_override_allowed DistributedPortGroup#block_override_allowed}
	BlockOverrideAllowed interface{} `field:"optional" json:"blockOverrideAllowed" yaml:"blockOverrideAllowed"`
	// Enable beacon probing on the ports this policy applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#check_beacon DistributedPortGroup#check_beacon}
	CheckBeacon interface{} `field:"optional" json:"checkBeacon" yaml:"checkBeacon"`
	// A list of custom attributes to set on this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#custom_attributes DistributedPortGroup#custom_attributes}
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The description of the portgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#description DistributedPortGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Allow VMDirectPath Gen2 on the ports this policy applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#directpath_gen2_allowed DistributedPortGroup#directpath_gen2_allowed}
	DirectpathGen2Allowed interface{} `field:"optional" json:"directpathGen2Allowed" yaml:"directpathGen2Allowed"`
	// The average egress bandwidth in bits per second if egress shaping is enabled on the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#egress_shaping_average_bandwidth DistributedPortGroup#egress_shaping_average_bandwidth}
	EgressShapingAverageBandwidth *float64 `field:"optional" json:"egressShapingAverageBandwidth" yaml:"egressShapingAverageBandwidth"`
	// The maximum egress burst size allowed in bytes if egress shaping is enabled on the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#egress_shaping_burst_size DistributedPortGroup#egress_shaping_burst_size}
	EgressShapingBurstSize *float64 `field:"optional" json:"egressShapingBurstSize" yaml:"egressShapingBurstSize"`
	// True if the traffic shaper is enabled for egress traffic on the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#egress_shaping_enabled DistributedPortGroup#egress_shaping_enabled}
	EgressShapingEnabled interface{} `field:"optional" json:"egressShapingEnabled" yaml:"egressShapingEnabled"`
	// The peak egress bandwidth during bursts in bits per second if egress traffic shaping is enabled on the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#egress_shaping_peak_bandwidth DistributedPortGroup#egress_shaping_peak_bandwidth}
	EgressShapingPeakBandwidth *float64 `field:"optional" json:"egressShapingPeakBandwidth" yaml:"egressShapingPeakBandwidth"`
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#failback DistributedPortGroup#failback}
	Failback interface{} `field:"optional" json:"failback" yaml:"failback"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#id DistributedPortGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The average ingress bandwidth in bits per second if ingress shaping is enabled on the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#ingress_shaping_average_bandwidth DistributedPortGroup#ingress_shaping_average_bandwidth}
	IngressShapingAverageBandwidth *float64 `field:"optional" json:"ingressShapingAverageBandwidth" yaml:"ingressShapingAverageBandwidth"`
	// The maximum ingress burst size allowed in bytes if ingress shaping is enabled on the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#ingress_shaping_burst_size DistributedPortGroup#ingress_shaping_burst_size}
	IngressShapingBurstSize *float64 `field:"optional" json:"ingressShapingBurstSize" yaml:"ingressShapingBurstSize"`
	// True if the traffic shaper is enabled for ingress traffic on the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#ingress_shaping_enabled DistributedPortGroup#ingress_shaping_enabled}
	IngressShapingEnabled interface{} `field:"optional" json:"ingressShapingEnabled" yaml:"ingressShapingEnabled"`
	// The peak ingress bandwidth during bursts in bits per second if ingress traffic shaping is enabled on the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#ingress_shaping_peak_bandwidth DistributedPortGroup#ingress_shaping_peak_bandwidth}
	IngressShapingPeakBandwidth *float64 `field:"optional" json:"ingressShapingPeakBandwidth" yaml:"ingressShapingPeakBandwidth"`
	// Whether or not to enable LACP on all uplink ports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#lacp_enabled DistributedPortGroup#lacp_enabled}
	LacpEnabled interface{} `field:"optional" json:"lacpEnabled" yaml:"lacpEnabled"`
	// The uplink LACP mode to use. Can be one of active or passive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#lacp_mode DistributedPortGroup#lacp_mode}
	LacpMode *string `field:"optional" json:"lacpMode" yaml:"lacpMode"`
	// Allow a live port to be moved in and out of the portgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#live_port_moving_allowed DistributedPortGroup#live_port_moving_allowed}
	LivePortMovingAllowed interface{} `field:"optional" json:"livePortMovingAllowed" yaml:"livePortMovingAllowed"`
	// Indicates whether to enable netflow on all ports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#netflow_enabled DistributedPortGroup#netflow_enabled}
	NetflowEnabled interface{} `field:"optional" json:"netflowEnabled" yaml:"netflowEnabled"`
	// Allow the enabling or disabling of Netflow on a port, contrary to the policy in the portgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#netflow_override_allowed DistributedPortGroup#netflow_override_allowed}
	NetflowOverrideAllowed interface{} `field:"optional" json:"netflowOverrideAllowed" yaml:"netflowOverrideAllowed"`
	// The key of a network resource pool to associate with this portgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#network_resource_pool_key DistributedPortGroup#network_resource_pool_key}
	NetworkResourcePoolKey *string `field:"optional" json:"networkResourcePoolKey" yaml:"networkResourcePoolKey"`
	// Allow the network resource pool of an individual port to override the setting in the portgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#network_resource_pool_override_allowed DistributedPortGroup#network_resource_pool_override_allowed}
	NetworkResourcePoolOverrideAllowed interface{} `field:"optional" json:"networkResourcePoolOverrideAllowed" yaml:"networkResourcePoolOverrideAllowed"`
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#notify_switches DistributedPortGroup#notify_switches}
	NotifySwitches interface{} `field:"optional" json:"notifySwitches" yaml:"notifySwitches"`
	// The number of ports in this portgroup. The DVS will expand and shrink by modifying this setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#number_of_ports DistributedPortGroup#number_of_ports}
	NumberOfPorts *float64 `field:"optional" json:"numberOfPorts" yaml:"numberOfPorts"`
	// Reset the setting of any ports in this portgroup back to the default setting when the port disconnects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#port_config_reset_at_disconnect DistributedPortGroup#port_config_reset_at_disconnect}
	PortConfigResetAtDisconnect interface{} `field:"optional" json:"portConfigResetAtDisconnect" yaml:"portConfigResetAtDisconnect"`
	// A template string to use when creating ports in the portgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#port_name_format DistributedPortGroup#port_name_format}
	PortNameFormat *string `field:"optional" json:"portNameFormat" yaml:"portNameFormat"`
	// The secondary VLAN ID for this port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#port_private_secondary_vlan_id DistributedPortGroup#port_private_secondary_vlan_id}
	PortPrivateSecondaryVlanId *float64 `field:"optional" json:"portPrivateSecondaryVlanId" yaml:"portPrivateSecondaryVlanId"`
	// Allow security policy settings on a port to override those on the portgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#security_policy_override_allowed DistributedPortGroup#security_policy_override_allowed}
	SecurityPolicyOverrideAllowed interface{} `field:"optional" json:"securityPolicyOverrideAllowed" yaml:"securityPolicyOverrideAllowed"`
	// Allow the traffic shaping policies of an individual port to override the settings in the portgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#shaping_override_allowed DistributedPortGroup#shaping_override_allowed}
	ShapingOverrideAllowed interface{} `field:"optional" json:"shapingOverrideAllowed" yaml:"shapingOverrideAllowed"`
	// List of standby uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#standby_uplinks DistributedPortGroup#standby_uplinks}
	StandbyUplinks *[]*string `field:"optional" json:"standbyUplinks" yaml:"standbyUplinks"`
	// A list of tag IDs to apply to this object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#tags DistributedPortGroup#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, failover_explicit, or loadbalance_loadbased.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#teaming_policy DistributedPortGroup#teaming_policy}
	TeamingPolicy *string `field:"optional" json:"teamingPolicy" yaml:"teamingPolicy"`
	// Allow any filter policies set on the individual port to override those in the portgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#traffic_filter_override_allowed DistributedPortGroup#traffic_filter_override_allowed}
	TrafficFilterOverrideAllowed interface{} `field:"optional" json:"trafficFilterOverrideAllowed" yaml:"trafficFilterOverrideAllowed"`
	// If true, a copy of packets sent to the switch will always be forwarded to an uplink in addition to the regular packet forwarded done by the switch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#tx_uplink DistributedPortGroup#tx_uplink}
	TxUplink interface{} `field:"optional" json:"txUplink" yaml:"txUplink"`
	// The type of portgroup. Can be one of earlyBinding (static) or ephemeral.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#type DistributedPortGroup#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Allow the uplink teaming policies on a port to override those on the portgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#uplink_teaming_override_allowed DistributedPortGroup#uplink_teaming_override_allowed}
	UplinkTeamingOverrideAllowed interface{} `field:"optional" json:"uplinkTeamingOverrideAllowed" yaml:"uplinkTeamingOverrideAllowed"`
	// The VLAN ID for single VLAN mode. 0 denotes no VLAN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#vlan_id DistributedPortGroup#vlan_id}
	VlanId *float64 `field:"optional" json:"vlanId" yaml:"vlanId"`
	// Allow the VLAN configuration on a port to override those on the portgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#vlan_override_allowed DistributedPortGroup#vlan_override_allowed}
	VlanOverrideAllowed interface{} `field:"optional" json:"vlanOverrideAllowed" yaml:"vlanOverrideAllowed"`
	// vlan_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/distributed_port_group#vlan_range DistributedPortGroup#vlan_range}
	VlanRange interface{} `field:"optional" json:"vlanRange" yaml:"vlanRange"`
}

