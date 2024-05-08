// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastorecluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatastoreClusterConfig struct {
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
	// The managed object ID of the datacenter to put the datastore cluster in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#datacenter_id DatastoreCluster#datacenter_id}
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Name for the new storage pod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#name DatastoreCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of custom attributes to set on this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#custom_attributes DatastoreCluster#custom_attributes}
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The name of the folder to locate the datastore cluster in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#folder DatastoreCluster#folder}
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#id DatastoreCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Advanced configuration options for storage DRS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_advanced_options DatastoreCluster#sdrs_advanced_options}
	SdrsAdvancedOptions *map[string]*string `field:"optional" json:"sdrsAdvancedOptions" yaml:"sdrsAdvancedOptions"`
	// The default automation level for all virtual machines in this storage cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_automation_level DatastoreCluster#sdrs_automation_level}
	SdrsAutomationLevel *string `field:"optional" json:"sdrsAutomationLevel" yaml:"sdrsAutomationLevel"`
	// When true, storage DRS keeps VMDKs for individual VMs on the same datastore by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_default_intra_vm_affinity DatastoreCluster#sdrs_default_intra_vm_affinity}
	SdrsDefaultIntraVmAffinity interface{} `field:"optional" json:"sdrsDefaultIntraVmAffinity" yaml:"sdrsDefaultIntraVmAffinity"`
	// Enable storage DRS for this datastore cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_enabled DatastoreCluster#sdrs_enabled}
	SdrsEnabled interface{} `field:"optional" json:"sdrsEnabled" yaml:"sdrsEnabled"`
	// The threshold, in GB, that storage DRS uses to make decisions to migrate VMs out of a datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_free_space_threshold DatastoreCluster#sdrs_free_space_threshold}
	SdrsFreeSpaceThreshold *float64 `field:"optional" json:"sdrsFreeSpaceThreshold" yaml:"sdrsFreeSpaceThreshold"`
	// The free space threshold to use.
	//
	// When set to utilization, drs_space_utilization_threshold is used, and when set to freeSpace, drs_free_space_threshold is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_free_space_threshold_mode DatastoreCluster#sdrs_free_space_threshold_mode}
	SdrsFreeSpaceThresholdMode *string `field:"optional" json:"sdrsFreeSpaceThresholdMode" yaml:"sdrsFreeSpaceThresholdMode"`
	// The threshold, in percent, of difference between space utilization in datastores before storage DRS makes decisions to balance the space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_free_space_utilization_difference DatastoreCluster#sdrs_free_space_utilization_difference}
	SdrsFreeSpaceUtilizationDifference *float64 `field:"optional" json:"sdrsFreeSpaceUtilizationDifference" yaml:"sdrsFreeSpaceUtilizationDifference"`
	// Overrides the default automation settings when correcting I/O load imbalances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_io_balance_automation_level DatastoreCluster#sdrs_io_balance_automation_level}
	SdrsIoBalanceAutomationLevel *string `field:"optional" json:"sdrsIoBalanceAutomationLevel" yaml:"sdrsIoBalanceAutomationLevel"`
	// The I/O latency threshold, in milliseconds, that storage DRS uses to make recommendations to move disks from this datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_io_latency_threshold DatastoreCluster#sdrs_io_latency_threshold}
	SdrsIoLatencyThreshold *float64 `field:"optional" json:"sdrsIoLatencyThreshold" yaml:"sdrsIoLatencyThreshold"`
	// Enable I/O load balancing for this datastore cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_io_load_balance_enabled DatastoreCluster#sdrs_io_load_balance_enabled}
	SdrsIoLoadBalanceEnabled interface{} `field:"optional" json:"sdrsIoLoadBalanceEnabled" yaml:"sdrsIoLoadBalanceEnabled"`
	// The difference between load in datastores in the cluster before storage DRS makes recommendations to balance the load.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_io_load_imbalance_threshold DatastoreCluster#sdrs_io_load_imbalance_threshold}
	SdrsIoLoadImbalanceThreshold *float64 `field:"optional" json:"sdrsIoLoadImbalanceThreshold" yaml:"sdrsIoLoadImbalanceThreshold"`
	// The threshold of reservable IOPS of all virtual machines on the datastore before storage DRS makes recommendations to move VMs off of a datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_io_reservable_iops_threshold DatastoreCluster#sdrs_io_reservable_iops_threshold}
	SdrsIoReservableIopsThreshold *float64 `field:"optional" json:"sdrsIoReservableIopsThreshold" yaml:"sdrsIoReservableIopsThreshold"`
	// The threshold, in percent, of actual estimated performance of the datastore (in IOPS) that storage DRS uses to make recommendations to move VMs off of a datastore when the total reservable IOPS exceeds the threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_io_reservable_percent_threshold DatastoreCluster#sdrs_io_reservable_percent_threshold}
	SdrsIoReservablePercentThreshold *float64 `field:"optional" json:"sdrsIoReservablePercentThreshold" yaml:"sdrsIoReservablePercentThreshold"`
	// The reservable IOPS threshold to use, percent in the event of automatic, or manual threshold in the event of manual.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_io_reservable_threshold_mode DatastoreCluster#sdrs_io_reservable_threshold_mode}
	SdrsIoReservableThresholdMode *string `field:"optional" json:"sdrsIoReservableThresholdMode" yaml:"sdrsIoReservableThresholdMode"`
	// The storage DRS poll interval, in minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_load_balance_interval DatastoreCluster#sdrs_load_balance_interval}
	SdrsLoadBalanceInterval *float64 `field:"optional" json:"sdrsLoadBalanceInterval" yaml:"sdrsLoadBalanceInterval"`
	// Overrides the default automation settings when correcting storage and VM policy violations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_policy_enforcement_automation_level DatastoreCluster#sdrs_policy_enforcement_automation_level}
	SdrsPolicyEnforcementAutomationLevel *string `field:"optional" json:"sdrsPolicyEnforcementAutomationLevel" yaml:"sdrsPolicyEnforcementAutomationLevel"`
	// Overrides the default automation settings when correcting affinity rule violations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_rule_enforcement_automation_level DatastoreCluster#sdrs_rule_enforcement_automation_level}
	SdrsRuleEnforcementAutomationLevel *string `field:"optional" json:"sdrsRuleEnforcementAutomationLevel" yaml:"sdrsRuleEnforcementAutomationLevel"`
	// Overrides the default automation settings when correcting disk space imbalances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_space_balance_automation_level DatastoreCluster#sdrs_space_balance_automation_level}
	SdrsSpaceBalanceAutomationLevel *string `field:"optional" json:"sdrsSpaceBalanceAutomationLevel" yaml:"sdrsSpaceBalanceAutomationLevel"`
	// The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_space_utilization_threshold DatastoreCluster#sdrs_space_utilization_threshold}
	SdrsSpaceUtilizationThreshold *float64 `field:"optional" json:"sdrsSpaceUtilizationThreshold" yaml:"sdrsSpaceUtilizationThreshold"`
	// Overrides the default automation settings when generating recommendations for datastore evacuation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#sdrs_vm_evacuation_automation_level DatastoreCluster#sdrs_vm_evacuation_automation_level}
	SdrsVmEvacuationAutomationLevel *string `field:"optional" json:"sdrsVmEvacuationAutomationLevel" yaml:"sdrsVmEvacuationAutomationLevel"`
	// A list of tag IDs to apply to this object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.0/docs/resources/datastore_cluster#tags DatastoreCluster#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

