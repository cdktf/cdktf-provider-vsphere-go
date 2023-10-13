// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package distributedvirtualswitch

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DistributedVirtualSwitchConfig struct {
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
	// The ID of the datacenter to create this virtual switch in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#datacenter_id DistributedVirtualSwitch#datacenter_id}
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// The name for the DVS. Must be unique in the folder that it is being created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#name DistributedVirtualSwitch#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#active_uplinks DistributedVirtualSwitch#active_uplinks}
	ActiveUplinks *[]*string `field:"optional" json:"activeUplinks" yaml:"activeUplinks"`
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than that of its own.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#allow_forged_transmits DistributedVirtualSwitch#allow_forged_transmits}
	AllowForgedTransmits interface{} `field:"optional" json:"allowForgedTransmits" yaml:"allowForgedTransmits"`
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#allow_mac_changes DistributedVirtualSwitch#allow_mac_changes}
	AllowMacChanges interface{} `field:"optional" json:"allowMacChanges" yaml:"allowMacChanges"`
	// Enable promiscuous mode on the network.
	//
	// This flag indicates whether or not all traffic is seen on a given port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#allow_promiscuous DistributedVirtualSwitch#allow_promiscuous}
	AllowPromiscuous interface{} `field:"optional" json:"allowPromiscuous" yaml:"allowPromiscuous"`
	// The maximum allowed usage for the backupNfc traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#backupnfc_maximum_mbit DistributedVirtualSwitch#backupnfc_maximum_mbit}
	BackupnfcMaximumMbit *float64 `field:"optional" json:"backupnfcMaximumMbit" yaml:"backupnfcMaximumMbit"`
	// The amount of guaranteed bandwidth for the backupNfc traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#backupnfc_reservation_mbit DistributedVirtualSwitch#backupnfc_reservation_mbit}
	BackupnfcReservationMbit *float64 `field:"optional" json:"backupnfcReservationMbit" yaml:"backupnfcReservationMbit"`
	// The amount of shares to allocate to the backupNfc traffic class for a custom share level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#backupnfc_share_count DistributedVirtualSwitch#backupnfc_share_count}
	BackupnfcShareCount *float64 `field:"optional" json:"backupnfcShareCount" yaml:"backupnfcShareCount"`
	// The allocation level for the backupNfc traffic class. Can be one of high, low, normal, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#backupnfc_share_level DistributedVirtualSwitch#backupnfc_share_level}
	BackupnfcShareLevel *string `field:"optional" json:"backupnfcShareLevel" yaml:"backupnfcShareLevel"`
	// Indicates whether to block all ports by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#block_all_ports DistributedVirtualSwitch#block_all_ports}
	BlockAllPorts interface{} `field:"optional" json:"blockAllPorts" yaml:"blockAllPorts"`
	// Enable beacon probing on the ports this policy applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#check_beacon DistributedVirtualSwitch#check_beacon}
	CheckBeacon interface{} `field:"optional" json:"checkBeacon" yaml:"checkBeacon"`
	// The contact detail for this DVS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#contact_detail DistributedVirtualSwitch#contact_detail}
	ContactDetail *string `field:"optional" json:"contactDetail" yaml:"contactDetail"`
	// The contact name for this DVS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#contact_name DistributedVirtualSwitch#contact_name}
	ContactName *string `field:"optional" json:"contactName" yaml:"contactName"`
	// A list of custom attributes to set on this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#custom_attributes DistributedVirtualSwitch#custom_attributes}
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The description of the DVS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#description DistributedVirtualSwitch#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Allow VMDirectPath Gen2 on the ports this policy applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#directpath_gen2_allowed DistributedVirtualSwitch#directpath_gen2_allowed}
	DirectpathGen2Allowed interface{} `field:"optional" json:"directpathGen2Allowed" yaml:"directpathGen2Allowed"`
	// The average egress bandwidth in bits per second if egress shaping is enabled on the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#egress_shaping_average_bandwidth DistributedVirtualSwitch#egress_shaping_average_bandwidth}
	EgressShapingAverageBandwidth *float64 `field:"optional" json:"egressShapingAverageBandwidth" yaml:"egressShapingAverageBandwidth"`
	// The maximum egress burst size allowed in bytes if egress shaping is enabled on the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#egress_shaping_burst_size DistributedVirtualSwitch#egress_shaping_burst_size}
	EgressShapingBurstSize *float64 `field:"optional" json:"egressShapingBurstSize" yaml:"egressShapingBurstSize"`
	// True if the traffic shaper is enabled for egress traffic on the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#egress_shaping_enabled DistributedVirtualSwitch#egress_shaping_enabled}
	EgressShapingEnabled interface{} `field:"optional" json:"egressShapingEnabled" yaml:"egressShapingEnabled"`
	// The peak egress bandwidth during bursts in bits per second if egress traffic shaping is enabled on the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#egress_shaping_peak_bandwidth DistributedVirtualSwitch#egress_shaping_peak_bandwidth}
	EgressShapingPeakBandwidth *float64 `field:"optional" json:"egressShapingPeakBandwidth" yaml:"egressShapingPeakBandwidth"`
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#failback DistributedVirtualSwitch#failback}
	Failback interface{} `field:"optional" json:"failback" yaml:"failback"`
	// The maximum allowed usage for the faultTolerance traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#faulttolerance_maximum_mbit DistributedVirtualSwitch#faulttolerance_maximum_mbit}
	FaulttoleranceMaximumMbit *float64 `field:"optional" json:"faulttoleranceMaximumMbit" yaml:"faulttoleranceMaximumMbit"`
	// The amount of guaranteed bandwidth for the faultTolerance traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#faulttolerance_reservation_mbit DistributedVirtualSwitch#faulttolerance_reservation_mbit}
	FaulttoleranceReservationMbit *float64 `field:"optional" json:"faulttoleranceReservationMbit" yaml:"faulttoleranceReservationMbit"`
	// The amount of shares to allocate to the faultTolerance traffic class for a custom share level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#faulttolerance_share_count DistributedVirtualSwitch#faulttolerance_share_count}
	FaulttoleranceShareCount *float64 `field:"optional" json:"faulttoleranceShareCount" yaml:"faulttoleranceShareCount"`
	// The allocation level for the faultTolerance traffic class. Can be one of high, low, normal, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#faulttolerance_share_level DistributedVirtualSwitch#faulttolerance_share_level}
	FaulttoleranceShareLevel *string `field:"optional" json:"faulttoleranceShareLevel" yaml:"faulttoleranceShareLevel"`
	// The folder to create this virtual switch in, relative to the datacenter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#folder DistributedVirtualSwitch#folder}
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// The maximum allowed usage for the hbr traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#hbr_maximum_mbit DistributedVirtualSwitch#hbr_maximum_mbit}
	HbrMaximumMbit *float64 `field:"optional" json:"hbrMaximumMbit" yaml:"hbrMaximumMbit"`
	// The amount of guaranteed bandwidth for the hbr traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#hbr_reservation_mbit DistributedVirtualSwitch#hbr_reservation_mbit}
	HbrReservationMbit *float64 `field:"optional" json:"hbrReservationMbit" yaml:"hbrReservationMbit"`
	// The amount of shares to allocate to the hbr traffic class for a custom share level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#hbr_share_count DistributedVirtualSwitch#hbr_share_count}
	HbrShareCount *float64 `field:"optional" json:"hbrShareCount" yaml:"hbrShareCount"`
	// The allocation level for the hbr traffic class. Can be one of high, low, normal, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#hbr_share_level DistributedVirtualSwitch#hbr_share_level}
	HbrShareLevel *string `field:"optional" json:"hbrShareLevel" yaml:"hbrShareLevel"`
	// host block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#host DistributedVirtualSwitch#host}
	Host interface{} `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#id DistributedVirtualSwitch#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether to ignore existing PVLAN mappings not managed by this resource. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#ignore_other_pvlan_mappings DistributedVirtualSwitch#ignore_other_pvlan_mappings}
	IgnoreOtherPvlanMappings interface{} `field:"optional" json:"ignoreOtherPvlanMappings" yaml:"ignoreOtherPvlanMappings"`
	// The average ingress bandwidth in bits per second if ingress shaping is enabled on the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#ingress_shaping_average_bandwidth DistributedVirtualSwitch#ingress_shaping_average_bandwidth}
	IngressShapingAverageBandwidth *float64 `field:"optional" json:"ingressShapingAverageBandwidth" yaml:"ingressShapingAverageBandwidth"`
	// The maximum ingress burst size allowed in bytes if ingress shaping is enabled on the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#ingress_shaping_burst_size DistributedVirtualSwitch#ingress_shaping_burst_size}
	IngressShapingBurstSize *float64 `field:"optional" json:"ingressShapingBurstSize" yaml:"ingressShapingBurstSize"`
	// True if the traffic shaper is enabled for ingress traffic on the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#ingress_shaping_enabled DistributedVirtualSwitch#ingress_shaping_enabled}
	IngressShapingEnabled interface{} `field:"optional" json:"ingressShapingEnabled" yaml:"ingressShapingEnabled"`
	// The peak ingress bandwidth during bursts in bits per second if ingress traffic shaping is enabled on the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#ingress_shaping_peak_bandwidth DistributedVirtualSwitch#ingress_shaping_peak_bandwidth}
	IngressShapingPeakBandwidth *float64 `field:"optional" json:"ingressShapingPeakBandwidth" yaml:"ingressShapingPeakBandwidth"`
	// The IPv4 address of the switch.
	//
	// This can be used to see the DVS as a unique device with NetFlow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#ipv4_address DistributedVirtualSwitch#ipv4_address}
	Ipv4Address *string `field:"optional" json:"ipv4Address" yaml:"ipv4Address"`
	// The maximum allowed usage for the iSCSI traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#iscsi_maximum_mbit DistributedVirtualSwitch#iscsi_maximum_mbit}
	IscsiMaximumMbit *float64 `field:"optional" json:"iscsiMaximumMbit" yaml:"iscsiMaximumMbit"`
	// The amount of guaranteed bandwidth for the iSCSI traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#iscsi_reservation_mbit DistributedVirtualSwitch#iscsi_reservation_mbit}
	IscsiReservationMbit *float64 `field:"optional" json:"iscsiReservationMbit" yaml:"iscsiReservationMbit"`
	// The amount of shares to allocate to the iSCSI traffic class for a custom share level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#iscsi_share_count DistributedVirtualSwitch#iscsi_share_count}
	IscsiShareCount *float64 `field:"optional" json:"iscsiShareCount" yaml:"iscsiShareCount"`
	// The allocation level for the iSCSI traffic class. Can be one of high, low, normal, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#iscsi_share_level DistributedVirtualSwitch#iscsi_share_level}
	IscsiShareLevel *string `field:"optional" json:"iscsiShareLevel" yaml:"iscsiShareLevel"`
	// The Link Aggregation Control Protocol group version in the switch. Can be one of singleLag or multipleLag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#lacp_api_version DistributedVirtualSwitch#lacp_api_version}
	LacpApiVersion *string `field:"optional" json:"lacpApiVersion" yaml:"lacpApiVersion"`
	// Whether or not to enable LACP on all uplink ports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#lacp_enabled DistributedVirtualSwitch#lacp_enabled}
	LacpEnabled interface{} `field:"optional" json:"lacpEnabled" yaml:"lacpEnabled"`
	// The uplink LACP mode to use. Can be one of active or passive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#lacp_mode DistributedVirtualSwitch#lacp_mode}
	LacpMode *string `field:"optional" json:"lacpMode" yaml:"lacpMode"`
	// Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#link_discovery_operation DistributedVirtualSwitch#link_discovery_operation}
	LinkDiscoveryOperation *string `field:"optional" json:"linkDiscoveryOperation" yaml:"linkDiscoveryOperation"`
	// The discovery protocol type. Valid values are cdp and lldp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#link_discovery_protocol DistributedVirtualSwitch#link_discovery_protocol}
	LinkDiscoveryProtocol *string `field:"optional" json:"linkDiscoveryProtocol" yaml:"linkDiscoveryProtocol"`
	// The maximum allowed usage for the management traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#management_maximum_mbit DistributedVirtualSwitch#management_maximum_mbit}
	ManagementMaximumMbit *float64 `field:"optional" json:"managementMaximumMbit" yaml:"managementMaximumMbit"`
	// The amount of guaranteed bandwidth for the management traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#management_reservation_mbit DistributedVirtualSwitch#management_reservation_mbit}
	ManagementReservationMbit *float64 `field:"optional" json:"managementReservationMbit" yaml:"managementReservationMbit"`
	// The amount of shares to allocate to the management traffic class for a custom share level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#management_share_count DistributedVirtualSwitch#management_share_count}
	ManagementShareCount *float64 `field:"optional" json:"managementShareCount" yaml:"managementShareCount"`
	// The allocation level for the management traffic class. Can be one of high, low, normal, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#management_share_level DistributedVirtualSwitch#management_share_level}
	ManagementShareLevel *string `field:"optional" json:"managementShareLevel" yaml:"managementShareLevel"`
	// The maximum MTU on the switch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#max_mtu DistributedVirtualSwitch#max_mtu}
	MaxMtu *float64 `field:"optional" json:"maxMtu" yaml:"maxMtu"`
	// The multicast filtering mode on the switch. Can be one of legacyFiltering, or snooping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#multicast_filtering_mode DistributedVirtualSwitch#multicast_filtering_mode}
	MulticastFilteringMode *string `field:"optional" json:"multicastFilteringMode" yaml:"multicastFilteringMode"`
	// The number of seconds after which active flows are forced to be exported to the collector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#netflow_active_flow_timeout DistributedVirtualSwitch#netflow_active_flow_timeout}
	NetflowActiveFlowTimeout *float64 `field:"optional" json:"netflowActiveFlowTimeout" yaml:"netflowActiveFlowTimeout"`
	// IP address for the netflow collector, using IPv4 or IPv6.
	//
	// IPv6 is supported in vSphere Distributed Switch Version 6.0 or later.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#netflow_collector_ip_address DistributedVirtualSwitch#netflow_collector_ip_address}
	NetflowCollectorIpAddress *string `field:"optional" json:"netflowCollectorIpAddress" yaml:"netflowCollectorIpAddress"`
	// The port for the netflow collector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#netflow_collector_port DistributedVirtualSwitch#netflow_collector_port}
	NetflowCollectorPort *float64 `field:"optional" json:"netflowCollectorPort" yaml:"netflowCollectorPort"`
	// Indicates whether to enable netflow on all ports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#netflow_enabled DistributedVirtualSwitch#netflow_enabled}
	NetflowEnabled interface{} `field:"optional" json:"netflowEnabled" yaml:"netflowEnabled"`
	// The number of seconds after which idle flows are forced to be exported to the collector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#netflow_idle_flow_timeout DistributedVirtualSwitch#netflow_idle_flow_timeout}
	NetflowIdleFlowTimeout *float64 `field:"optional" json:"netflowIdleFlowTimeout" yaml:"netflowIdleFlowTimeout"`
	// Whether to limit analysis to traffic that has both source and destination served by the same host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#netflow_internal_flows_only DistributedVirtualSwitch#netflow_internal_flows_only}
	NetflowInternalFlowsOnly interface{} `field:"optional" json:"netflowInternalFlowsOnly" yaml:"netflowInternalFlowsOnly"`
	// The observation Domain ID for the netflow collector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#netflow_observation_domain_id DistributedVirtualSwitch#netflow_observation_domain_id}
	NetflowObservationDomainId *float64 `field:"optional" json:"netflowObservationDomainId" yaml:"netflowObservationDomainId"`
	// The ratio of total number of packets to the number of packets analyzed.
	//
	// Set to 0 to disable sampling, meaning that all packets are analyzed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#netflow_sampling_rate DistributedVirtualSwitch#netflow_sampling_rate}
	NetflowSamplingRate *float64 `field:"optional" json:"netflowSamplingRate" yaml:"netflowSamplingRate"`
	// Whether or not to enable network resource control, enabling advanced traffic shaping and resource control features.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#network_resource_control_enabled DistributedVirtualSwitch#network_resource_control_enabled}
	NetworkResourceControlEnabled interface{} `field:"optional" json:"networkResourceControlEnabled" yaml:"networkResourceControlEnabled"`
	// The network I/O control version to use. Can be one of version2 or version3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#network_resource_control_version DistributedVirtualSwitch#network_resource_control_version}
	NetworkResourceControlVersion *string `field:"optional" json:"networkResourceControlVersion" yaml:"networkResourceControlVersion"`
	// The maximum allowed usage for the nfs traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#nfs_maximum_mbit DistributedVirtualSwitch#nfs_maximum_mbit}
	NfsMaximumMbit *float64 `field:"optional" json:"nfsMaximumMbit" yaml:"nfsMaximumMbit"`
	// The amount of guaranteed bandwidth for the nfs traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#nfs_reservation_mbit DistributedVirtualSwitch#nfs_reservation_mbit}
	NfsReservationMbit *float64 `field:"optional" json:"nfsReservationMbit" yaml:"nfsReservationMbit"`
	// The amount of shares to allocate to the nfs traffic class for a custom share level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#nfs_share_count DistributedVirtualSwitch#nfs_share_count}
	NfsShareCount *float64 `field:"optional" json:"nfsShareCount" yaml:"nfsShareCount"`
	// The allocation level for the nfs traffic class. Can be one of high, low, normal, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#nfs_share_level DistributedVirtualSwitch#nfs_share_level}
	NfsShareLevel *string `field:"optional" json:"nfsShareLevel" yaml:"nfsShareLevel"`
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#notify_switches DistributedVirtualSwitch#notify_switches}
	NotifySwitches interface{} `field:"optional" json:"notifySwitches" yaml:"notifySwitches"`
	// The secondary VLAN ID for this port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#port_private_secondary_vlan_id DistributedVirtualSwitch#port_private_secondary_vlan_id}
	PortPrivateSecondaryVlanId *float64 `field:"optional" json:"portPrivateSecondaryVlanId" yaml:"portPrivateSecondaryVlanId"`
	// pvlan_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#pvlan_mapping DistributedVirtualSwitch#pvlan_mapping}
	PvlanMapping interface{} `field:"optional" json:"pvlanMapping" yaml:"pvlanMapping"`
	// List of standby uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#standby_uplinks DistributedVirtualSwitch#standby_uplinks}
	StandbyUplinks *[]*string `field:"optional" json:"standbyUplinks" yaml:"standbyUplinks"`
	// A list of tag IDs to apply to this object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#tags DistributedVirtualSwitch#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, failover_explicit, or loadbalance_loadbased.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#teaming_policy DistributedVirtualSwitch#teaming_policy}
	TeamingPolicy *string `field:"optional" json:"teamingPolicy" yaml:"teamingPolicy"`
	// If true, a copy of packets sent to the switch will always be forwarded to an uplink in addition to the regular packet forwarded done by the switch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#tx_uplink DistributedVirtualSwitch#tx_uplink}
	TxUplink interface{} `field:"optional" json:"txUplink" yaml:"txUplink"`
	// A list of uplink ports.
	//
	// The contents of this list control both the uplink count and names of the uplinks on the DVS across hosts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#uplinks DistributedVirtualSwitch#uplinks}
	Uplinks *[]*string `field:"optional" json:"uplinks" yaml:"uplinks"`
	// The maximum allowed usage for the vdp traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#vdp_maximum_mbit DistributedVirtualSwitch#vdp_maximum_mbit}
	VdpMaximumMbit *float64 `field:"optional" json:"vdpMaximumMbit" yaml:"vdpMaximumMbit"`
	// The amount of guaranteed bandwidth for the vdp traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#vdp_reservation_mbit DistributedVirtualSwitch#vdp_reservation_mbit}
	VdpReservationMbit *float64 `field:"optional" json:"vdpReservationMbit" yaml:"vdpReservationMbit"`
	// The amount of shares to allocate to the vdp traffic class for a custom share level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#vdp_share_count DistributedVirtualSwitch#vdp_share_count}
	VdpShareCount *float64 `field:"optional" json:"vdpShareCount" yaml:"vdpShareCount"`
	// The allocation level for the vdp traffic class. Can be one of high, low, normal, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#vdp_share_level DistributedVirtualSwitch#vdp_share_level}
	VdpShareLevel *string `field:"optional" json:"vdpShareLevel" yaml:"vdpShareLevel"`
	// The version of this virtual switch.
	//
	// Allowed versions are 8.0.0, 7.0.3, 7.0.2, 7.0.0, 6.6.0, 6.5.0, 6.0.0, 5.5.0, 5.1.0, and 5.0.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#version DistributedVirtualSwitch#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
	// The maximum allowed usage for the virtualMachine traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#virtualmachine_maximum_mbit DistributedVirtualSwitch#virtualmachine_maximum_mbit}
	VirtualmachineMaximumMbit *float64 `field:"optional" json:"virtualmachineMaximumMbit" yaml:"virtualmachineMaximumMbit"`
	// The amount of guaranteed bandwidth for the virtualMachine traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#virtualmachine_reservation_mbit DistributedVirtualSwitch#virtualmachine_reservation_mbit}
	VirtualmachineReservationMbit *float64 `field:"optional" json:"virtualmachineReservationMbit" yaml:"virtualmachineReservationMbit"`
	// The amount of shares to allocate to the virtualMachine traffic class for a custom share level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#virtualmachine_share_count DistributedVirtualSwitch#virtualmachine_share_count}
	VirtualmachineShareCount *float64 `field:"optional" json:"virtualmachineShareCount" yaml:"virtualmachineShareCount"`
	// The allocation level for the virtualMachine traffic class. Can be one of high, low, normal, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#virtualmachine_share_level DistributedVirtualSwitch#virtualmachine_share_level}
	VirtualmachineShareLevel *string `field:"optional" json:"virtualmachineShareLevel" yaml:"virtualmachineShareLevel"`
	// The VLAN ID for single VLAN mode. 0 denotes no VLAN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#vlan_id DistributedVirtualSwitch#vlan_id}
	VlanId *float64 `field:"optional" json:"vlanId" yaml:"vlanId"`
	// vlan_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#vlan_range DistributedVirtualSwitch#vlan_range}
	VlanRange interface{} `field:"optional" json:"vlanRange" yaml:"vlanRange"`
	// The maximum allowed usage for the vmotion traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#vmotion_maximum_mbit DistributedVirtualSwitch#vmotion_maximum_mbit}
	VmotionMaximumMbit *float64 `field:"optional" json:"vmotionMaximumMbit" yaml:"vmotionMaximumMbit"`
	// The amount of guaranteed bandwidth for the vmotion traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#vmotion_reservation_mbit DistributedVirtualSwitch#vmotion_reservation_mbit}
	VmotionReservationMbit *float64 `field:"optional" json:"vmotionReservationMbit" yaml:"vmotionReservationMbit"`
	// The amount of shares to allocate to the vmotion traffic class for a custom share level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#vmotion_share_count DistributedVirtualSwitch#vmotion_share_count}
	VmotionShareCount *float64 `field:"optional" json:"vmotionShareCount" yaml:"vmotionShareCount"`
	// The allocation level for the vmotion traffic class. Can be one of high, low, normal, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#vmotion_share_level DistributedVirtualSwitch#vmotion_share_level}
	VmotionShareLevel *string `field:"optional" json:"vmotionShareLevel" yaml:"vmotionShareLevel"`
	// The maximum allowed usage for the vsan traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#vsan_maximum_mbit DistributedVirtualSwitch#vsan_maximum_mbit}
	VsanMaximumMbit *float64 `field:"optional" json:"vsanMaximumMbit" yaml:"vsanMaximumMbit"`
	// The amount of guaranteed bandwidth for the vsan traffic class, in Mbits/sec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#vsan_reservation_mbit DistributedVirtualSwitch#vsan_reservation_mbit}
	VsanReservationMbit *float64 `field:"optional" json:"vsanReservationMbit" yaml:"vsanReservationMbit"`
	// The amount of shares to allocate to the vsan traffic class for a custom share level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#vsan_share_count DistributedVirtualSwitch#vsan_share_count}
	VsanShareCount *float64 `field:"optional" json:"vsanShareCount" yaml:"vsanShareCount"`
	// The allocation level for the vsan traffic class. Can be one of high, low, normal, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.5.1/docs/resources/distributed_virtual_switch#vsan_share_level DistributedVirtualSwitch#vsan_share_level}
	VsanShareLevel *string `field:"optional" json:"vsanShareLevel" yaml:"vsanShareLevel"`
}

