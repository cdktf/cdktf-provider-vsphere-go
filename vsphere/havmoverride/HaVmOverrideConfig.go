// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package havmoverride

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HaVmOverrideConfig struct {
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
	// The managed object ID of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/ha_vm_override#compute_cluster_id HaVmOverride#compute_cluster_id}
	ComputeClusterId *string `field:"required" json:"computeClusterId" yaml:"computeClusterId"`
	// The managed object ID of the virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/ha_vm_override#virtual_machine_id HaVmOverride#virtual_machine_id}
	VirtualMachineId *string `field:"required" json:"virtualMachineId" yaml:"virtualMachineId"`
	// Controls the action to take on this virtual machine if an APD status on an affected datastore clears in the middle of an APD event.
	//
	// Can be one of useClusterDefault, none or reset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/ha_vm_override#ha_datastore_apd_recovery_action HaVmOverride#ha_datastore_apd_recovery_action}
	HaDatastoreApdRecoveryAction *string `field:"optional" json:"haDatastoreApdRecoveryAction" yaml:"haDatastoreApdRecoveryAction"`
	// Controls the action to take on this virtual machine when the cluster has detected loss to all paths to a relevant datastore.
	//
	// Can be one of clusterDefault, disabled, warning, restartConservative, or restartAggressive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/ha_vm_override#ha_datastore_apd_response HaVmOverride#ha_datastore_apd_response}
	HaDatastoreApdResponse *string `field:"optional" json:"haDatastoreApdResponse" yaml:"haDatastoreApdResponse"`
	// Controls the delay in seconds to wait after an APD timeout event to execute the response action defined in ha_datastore_apd_response.
	//
	// Specify -1 to use the cluster setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/ha_vm_override#ha_datastore_apd_response_delay HaVmOverride#ha_datastore_apd_response_delay}
	HaDatastoreApdResponseDelay *float64 `field:"optional" json:"haDatastoreApdResponseDelay" yaml:"haDatastoreApdResponseDelay"`
	// Controls the action to take on this virtual machine when the cluster has detected a permanent device loss to a relevant datastore.
	//
	// Can be one of clusterDefault, disabled, warning, or restartAggressive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/ha_vm_override#ha_datastore_pdl_response HaVmOverride#ha_datastore_pdl_response}
	HaDatastorePdlResponse *string `field:"optional" json:"haDatastorePdlResponse" yaml:"haDatastorePdlResponse"`
	// The action to take on this virtual machine when a host is isolated from the rest of the cluster.
	//
	// Can be one of clusterIsolationResponse, none, powerOff, or shutdown.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/ha_vm_override#ha_host_isolation_response HaVmOverride#ha_host_isolation_response}
	HaHostIsolationResponse *string `field:"optional" json:"haHostIsolationResponse" yaml:"haHostIsolationResponse"`
	// If a heartbeat from this virtual machine is not received within this configured interval, the virtual machine is marked as failed.
	//
	// The value is in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/ha_vm_override#ha_vm_failure_interval HaVmOverride#ha_vm_failure_interval}
	HaVmFailureInterval *float64 `field:"optional" json:"haVmFailureInterval" yaml:"haVmFailureInterval"`
	// The length of the reset window in which ha_vm_maximum_resets can operate.
	//
	// When this window expires, no more resets are attempted regardless of the setting configured in ha_vm_maximum_resets. -1 means no window, meaning an unlimited reset time is allotted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/ha_vm_override#ha_vm_maximum_failure_window HaVmOverride#ha_vm_maximum_failure_window}
	HaVmMaximumFailureWindow *float64 `field:"optional" json:"haVmMaximumFailureWindow" yaml:"haVmMaximumFailureWindow"`
	// The maximum number of resets that HA will perform to this virtual machine when responding to a failure event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/ha_vm_override#ha_vm_maximum_resets HaVmOverride#ha_vm_maximum_resets}
	HaVmMaximumResets *float64 `field:"optional" json:"haVmMaximumResets" yaml:"haVmMaximumResets"`
	// The time, in seconds, that HA waits after powering on this virtual machine before monitoring for heartbeats.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/ha_vm_override#ha_vm_minimum_uptime HaVmOverride#ha_vm_minimum_uptime}
	HaVmMinimumUptime *float64 `field:"optional" json:"haVmMinimumUptime" yaml:"haVmMinimumUptime"`
	// The type of virtual machine monitoring to use for this virtual machine.
	//
	// Can be one of vmMonitoringDisabled, vmMonitoringOnly, or vmAndAppMonitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/ha_vm_override#ha_vm_monitoring HaVmOverride#ha_vm_monitoring}
	HaVmMonitoring *string `field:"optional" json:"haVmMonitoring" yaml:"haVmMonitoring"`
	// Determines whether or not the cluster's default settings or the VM override settings specified in this resource are used for virtual machine monitoring.
	//
	// The default is true (use cluster defaults) - set to false to have overrides take effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/ha_vm_override#ha_vm_monitoring_use_cluster_defaults HaVmOverride#ha_vm_monitoring_use_cluster_defaults}
	HaVmMonitoringUseClusterDefaults interface{} `field:"optional" json:"haVmMonitoringUseClusterDefaults" yaml:"haVmMonitoringUseClusterDefaults"`
	// The restart priority for this virtual machine when vSphere detects a host failure.
	//
	// Can be one of clusterRestartPriority, lowest, low, medium, high, or highest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/ha_vm_override#ha_vm_restart_priority HaVmOverride#ha_vm_restart_priority}
	HaVmRestartPriority *string `field:"optional" json:"haVmRestartPriority" yaml:"haVmRestartPriority"`
	// The maximum time, in seconds, that vSphere HA will wait for the virtual machine to be ready.
	//
	// Use -1 to use the cluster default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/ha_vm_override#ha_vm_restart_timeout HaVmOverride#ha_vm_restart_timeout}
	HaVmRestartTimeout *float64 `field:"optional" json:"haVmRestartTimeout" yaml:"haVmRestartTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.2/docs/resources/ha_vm_override#id HaVmOverride#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

