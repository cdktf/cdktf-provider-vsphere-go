package computecluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeClusterConfig struct {
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
	// The managed object ID of the datacenter to put the cluster in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#datacenter_id ComputeCluster#datacenter_id}
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// Name for the new cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#name ComputeCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of custom attributes to set on this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#custom_attributes ComputeCluster#custom_attributes}
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The automation level for host power operations in this cluster. Can be one of manual or automated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#dpm_automation_level ComputeCluster#dpm_automation_level}
	DpmAutomationLevel *string `field:"optional" json:"dpmAutomationLevel" yaml:"dpmAutomationLevel"`
	// Enable DPM support for DRS.
	//
	// This allows you to dynamically control the power of hosts depending on the needs of virtual machines in the cluster. Requires that DRS be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#dpm_enabled ComputeCluster#dpm_enabled}
	DpmEnabled interface{} `field:"optional" json:"dpmEnabled" yaml:"dpmEnabled"`
	// A value between 1 and 5 indicating the threshold of load within the cluster that influences host power operations.
	//
	// This affects both power on and power off operations - a lower setting will tolerate more of a surplus/deficit than a higher setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#dpm_threshold ComputeCluster#dpm_threshold}
	DpmThreshold *float64 `field:"optional" json:"dpmThreshold" yaml:"dpmThreshold"`
	// Advanced configuration options for DRS and DPM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#drs_advanced_options ComputeCluster#drs_advanced_options}
	DrsAdvancedOptions *map[string]*string `field:"optional" json:"drsAdvancedOptions" yaml:"drsAdvancedOptions"`
	// The default automation level for all virtual machines in this cluster. Can be one of manual, partiallyAutomated, or fullyAutomated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#drs_automation_level ComputeCluster#drs_automation_level}
	DrsAutomationLevel *string `field:"optional" json:"drsAutomationLevel" yaml:"drsAutomationLevel"`
	// Enable DRS for this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#drs_enabled ComputeCluster#drs_enabled}
	DrsEnabled interface{} `field:"optional" json:"drsEnabled" yaml:"drsEnabled"`
	// When true, enables DRS to use data from vRealize Operations Manager to make proactive DRS recommendations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#drs_enable_predictive_drs ComputeCluster#drs_enable_predictive_drs}
	DrsEnablePredictiveDrs interface{} `field:"optional" json:"drsEnablePredictiveDrs" yaml:"drsEnablePredictiveDrs"`
	// When true, allows individual VM overrides within this cluster to be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#drs_enable_vm_overrides ComputeCluster#drs_enable_vm_overrides}
	DrsEnableVmOverrides interface{} `field:"optional" json:"drsEnableVmOverrides" yaml:"drsEnableVmOverrides"`
	// A value between 1 and 5 indicating the threshold of imbalance tolerated between hosts.
	//
	// A lower setting will tolerate more imbalance while a higher setting will tolerate less.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#drs_migration_threshold ComputeCluster#drs_migration_threshold}
	DrsMigrationThreshold *float64 `field:"optional" json:"drsMigrationThreshold" yaml:"drsMigrationThreshold"`
	// Enable scalable shares for all descendants of this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#drs_scale_descendants_shares ComputeCluster#drs_scale_descendants_shares}
	DrsScaleDescendantsShares *string `field:"optional" json:"drsScaleDescendantsShares" yaml:"drsScaleDescendantsShares"`
	// The name of the folder to locate the cluster in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#folder ComputeCluster#folder}
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// Force removal of all hosts in the cluster during destroy and make them standalone hosts.
	//
	// Use of this flag mainly exists for testing and is not recommended in normal use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#force_evacuate_on_destroy ComputeCluster#force_evacuate_on_destroy}
	ForceEvacuateOnDestroy interface{} `field:"optional" json:"forceEvacuateOnDestroy" yaml:"forceEvacuateOnDestroy"`
	// When ha_admission_control_policy is failoverHosts, this defines the managed object IDs of hosts to use as dedicated failover hosts.
	//
	// These hosts are kept as available as possible - admission control will block access to the host, and DRS will ignore the host when making recommendations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_admission_control_failover_host_system_ids ComputeCluster#ha_admission_control_failover_host_system_ids}
	HaAdmissionControlFailoverHostSystemIds *[]*string `field:"optional" json:"haAdmissionControlFailoverHostSystemIds" yaml:"haAdmissionControlFailoverHostSystemIds"`
	// The maximum number of failed hosts that admission control tolerates when making decisions on whether to permit virtual machine operations.
	//
	// The maximum is one less than the number of hosts in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_admission_control_host_failure_tolerance ComputeCluster#ha_admission_control_host_failure_tolerance}
	HaAdmissionControlHostFailureTolerance *float64 `field:"optional" json:"haAdmissionControlHostFailureTolerance" yaml:"haAdmissionControlHostFailureTolerance"`
	// The percentage of resource reduction that a cluster of VMs can tolerate in case of a failover.
	//
	// A value of 0 produces warnings only, whereas a value of 100 disables the setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_admission_control_performance_tolerance ComputeCluster#ha_admission_control_performance_tolerance}
	HaAdmissionControlPerformanceTolerance *float64 `field:"optional" json:"haAdmissionControlPerformanceTolerance" yaml:"haAdmissionControlPerformanceTolerance"`
	// The type of admission control policy to use with vSphere HA, which controls whether or not specific VM operations are permitted in the cluster in order to protect the reliability of the cluster.
	//
	// Can be one of resourcePercentage, slotPolicy, failoverHosts, or disabled. Note that disabling admission control is not recommended and can lead to service issues.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_admission_control_policy ComputeCluster#ha_admission_control_policy}
	HaAdmissionControlPolicy *string `field:"optional" json:"haAdmissionControlPolicy" yaml:"haAdmissionControlPolicy"`
	// When ha_admission_control_policy is resourcePercentage, automatically determine available resource percentages by subtracting the average number of host resources represented by the ha_admission_control_host_failure_tolerance setting from the total amount of resources in the cluster.
	//
	// Disable to supply user-defined values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_admission_control_resource_percentage_auto_compute ComputeCluster#ha_admission_control_resource_percentage_auto_compute}
	HaAdmissionControlResourcePercentageAutoCompute interface{} `field:"optional" json:"haAdmissionControlResourcePercentageAutoCompute" yaml:"haAdmissionControlResourcePercentageAutoCompute"`
	// When ha_admission_control_policy is resourcePercentage, this controls the user-defined percentage of CPU resources in the cluster to reserve for failover.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_admission_control_resource_percentage_cpu ComputeCluster#ha_admission_control_resource_percentage_cpu}
	HaAdmissionControlResourcePercentageCpu *float64 `field:"optional" json:"haAdmissionControlResourcePercentageCpu" yaml:"haAdmissionControlResourcePercentageCpu"`
	// When ha_admission_control_policy is resourcePercentage, this controls the user-defined percentage of memory resources in the cluster to reserve for failover.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_admission_control_resource_percentage_memory ComputeCluster#ha_admission_control_resource_percentage_memory}
	HaAdmissionControlResourcePercentageMemory *float64 `field:"optional" json:"haAdmissionControlResourcePercentageMemory" yaml:"haAdmissionControlResourcePercentageMemory"`
	// When ha_admission_control_policy is slotPolicy, this controls the user-defined CPU slot size, in MHz.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_admission_control_slot_policy_explicit_cpu ComputeCluster#ha_admission_control_slot_policy_explicit_cpu}
	HaAdmissionControlSlotPolicyExplicitCpu *float64 `field:"optional" json:"haAdmissionControlSlotPolicyExplicitCpu" yaml:"haAdmissionControlSlotPolicyExplicitCpu"`
	// When ha_admission_control_policy is slotPolicy, this controls the user-defined memory slot size, in MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_admission_control_slot_policy_explicit_memory ComputeCluster#ha_admission_control_slot_policy_explicit_memory}
	HaAdmissionControlSlotPolicyExplicitMemory *float64 `field:"optional" json:"haAdmissionControlSlotPolicyExplicitMemory" yaml:"haAdmissionControlSlotPolicyExplicitMemory"`
	// When ha_admission_control_policy is slotPolicy, this setting controls whether or not you wish to supply explicit values to CPU and memory slot sizes.
	//
	// The default is to gather a automatic average based on all powered-on virtual machines currently in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_admission_control_slot_policy_use_explicit_size ComputeCluster#ha_admission_control_slot_policy_use_explicit_size}
	HaAdmissionControlSlotPolicyUseExplicitSize interface{} `field:"optional" json:"haAdmissionControlSlotPolicyUseExplicitSize" yaml:"haAdmissionControlSlotPolicyUseExplicitSize"`
	// Advanced configuration options for vSphere HA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_advanced_options ComputeCluster#ha_advanced_options}
	HaAdvancedOptions *map[string]*string `field:"optional" json:"haAdvancedOptions" yaml:"haAdvancedOptions"`
	// When ha_vm_component_protection is enabled, controls the action to take on virtual machines if an APD status on an affected datastore clears in the middle of an APD event.
	//
	// Can be one of none or reset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_datastore_apd_recovery_action ComputeCluster#ha_datastore_apd_recovery_action}
	HaDatastoreApdRecoveryAction *string `field:"optional" json:"haDatastoreApdRecoveryAction" yaml:"haDatastoreApdRecoveryAction"`
	// When ha_vm_component_protection is enabled, controls the action to take on virtual machines when the cluster has detected loss to all paths to a relevant datastore.
	//
	// Can be one of disabled, warning, restartConservative, or restartAggressive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_datastore_apd_response ComputeCluster#ha_datastore_apd_response}
	HaDatastoreApdResponse *string `field:"optional" json:"haDatastoreApdResponse" yaml:"haDatastoreApdResponse"`
	// When ha_vm_component_protection is enabled, controls the delay in seconds to wait after an APD timeout event to execute the response action defined in ha_datastore_apd_response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_datastore_apd_response_delay ComputeCluster#ha_datastore_apd_response_delay}
	HaDatastoreApdResponseDelay *float64 `field:"optional" json:"haDatastoreApdResponseDelay" yaml:"haDatastoreApdResponseDelay"`
	// When ha_vm_component_protection is enabled, controls the action to take on virtual machines when the cluster has detected a permanent device loss to a relevant datastore.
	//
	// Can be one of disabled, warning, or restartAggressive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_datastore_pdl_response ComputeCluster#ha_datastore_pdl_response}
	HaDatastorePdlResponse *string `field:"optional" json:"haDatastorePdlResponse" yaml:"haDatastorePdlResponse"`
	// Enable vSphere HA for this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_enabled ComputeCluster#ha_enabled}
	HaEnabled interface{} `field:"optional" json:"haEnabled" yaml:"haEnabled"`
	// The list of managed object IDs for preferred datastores to use for HA heartbeating.
	//
	// This setting is only useful when ha_heartbeat_datastore_policy is set to either userSelectedDs or allFeasibleDsWithUserPreference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_heartbeat_datastore_ids ComputeCluster#ha_heartbeat_datastore_ids}
	HaHeartbeatDatastoreIds *[]*string `field:"optional" json:"haHeartbeatDatastoreIds" yaml:"haHeartbeatDatastoreIds"`
	// The selection policy for HA heartbeat datastores. Can be one of allFeasibleDs, userSelectedDs, or allFeasibleDsWithUserPreference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_heartbeat_datastore_policy ComputeCluster#ha_heartbeat_datastore_policy}
	HaHeartbeatDatastorePolicy *string `field:"optional" json:"haHeartbeatDatastorePolicy" yaml:"haHeartbeatDatastorePolicy"`
	// The action to take on virtual machines when a host has detected that it has been isolated from the rest of the cluster.
	//
	// Can be one of none, powerOff, or shutdown.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_host_isolation_response ComputeCluster#ha_host_isolation_response}
	HaHostIsolationResponse *string `field:"optional" json:"haHostIsolationResponse" yaml:"haHostIsolationResponse"`
	// Global setting that controls whether vSphere HA remediates VMs on host failure. Can be one of enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_host_monitoring ComputeCluster#ha_host_monitoring}
	HaHostMonitoring *string `field:"optional" json:"haHostMonitoring" yaml:"haHostMonitoring"`
	// Controls vSphere VM component protection for virtual machines in this cluster.
	//
	// This allows vSphere HA to react to failures between hosts and specific virtual machine components, such as datastores. Can be one of enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_vm_component_protection ComputeCluster#ha_vm_component_protection}
	HaVmComponentProtection *string `field:"optional" json:"haVmComponentProtection" yaml:"haVmComponentProtection"`
	// The condition used to determine whether or not VMs in a certain restart priority class are online, allowing HA to move on to restarting VMs on the next priority.
	//
	// Can be one of none, poweredOn, guestHbStatusGreen, or appHbStatusGreen.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_vm_dependency_restart_condition ComputeCluster#ha_vm_dependency_restart_condition}
	HaVmDependencyRestartCondition *string `field:"optional" json:"haVmDependencyRestartCondition" yaml:"haVmDependencyRestartCondition"`
	// If a heartbeat from a virtual machine is not received within this configured interval, the virtual machine is marked as failed.
	//
	// The value is in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_vm_failure_interval ComputeCluster#ha_vm_failure_interval}
	HaVmFailureInterval *float64 `field:"optional" json:"haVmFailureInterval" yaml:"haVmFailureInterval"`
	// The length of the reset window in which ha_vm_maximum_resets can operate.
	//
	// When this window expires, no more resets are attempted regardless of the setting configured in ha_vm_maximum_resets. -1 means no window, meaning an unlimited reset time is allotted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_vm_maximum_failure_window ComputeCluster#ha_vm_maximum_failure_window}
	HaVmMaximumFailureWindow *float64 `field:"optional" json:"haVmMaximumFailureWindow" yaml:"haVmMaximumFailureWindow"`
	// The maximum number of resets that HA will perform to a virtual machine when responding to a failure event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_vm_maximum_resets ComputeCluster#ha_vm_maximum_resets}
	HaVmMaximumResets *float64 `field:"optional" json:"haVmMaximumResets" yaml:"haVmMaximumResets"`
	// The time, in seconds, that HA waits after powering on a virtual machine before monitoring for heartbeats.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_vm_minimum_uptime ComputeCluster#ha_vm_minimum_uptime}
	HaVmMinimumUptime *float64 `field:"optional" json:"haVmMinimumUptime" yaml:"haVmMinimumUptime"`
	// The type of virtual machine monitoring to use when HA is enabled in the cluster.
	//
	// Can be one of vmMonitoringDisabled, vmMonitoringOnly, or vmAndAppMonitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_vm_monitoring ComputeCluster#ha_vm_monitoring}
	HaVmMonitoring *string `field:"optional" json:"haVmMonitoring" yaml:"haVmMonitoring"`
	// Additional delay in seconds after ready condition is met. A VM is considered ready at this point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_vm_restart_additional_delay ComputeCluster#ha_vm_restart_additional_delay}
	HaVmRestartAdditionalDelay *float64 `field:"optional" json:"haVmRestartAdditionalDelay" yaml:"haVmRestartAdditionalDelay"`
	// The default restart priority for affected VMs when vSphere detects a host failure.
	//
	// Can be one of lowest, low, medium, high, or highest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_vm_restart_priority ComputeCluster#ha_vm_restart_priority}
	HaVmRestartPriority *string `field:"optional" json:"haVmRestartPriority" yaml:"haVmRestartPriority"`
	// The maximum time, in seconds, that vSphere HA will wait for virtual machines in one priority to be ready before proceeding with the next priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#ha_vm_restart_timeout ComputeCluster#ha_vm_restart_timeout}
	HaVmRestartTimeout *float64 `field:"optional" json:"haVmRestartTimeout" yaml:"haVmRestartTimeout"`
	// The timeout for each host maintenance mode operation when removing hosts from a cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#host_cluster_exit_timeout ComputeCluster#host_cluster_exit_timeout}
	HostClusterExitTimeout *float64 `field:"optional" json:"hostClusterExitTimeout" yaml:"hostClusterExitTimeout"`
	// Must be set if cluster enrollment is managed from host resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#host_managed ComputeCluster#host_managed}
	HostManaged interface{} `field:"optional" json:"hostManaged" yaml:"hostManaged"`
	// The managed object IDs of the hosts to put in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#host_system_ids ComputeCluster#host_system_ids}
	HostSystemIds *[]*string `field:"optional" json:"hostSystemIds" yaml:"hostSystemIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#id ComputeCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The DRS behavior for proactive HA recommendations. Can be one of Automated or Manual.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#proactive_ha_automation_level ComputeCluster#proactive_ha_automation_level}
	ProactiveHaAutomationLevel *string `field:"optional" json:"proactiveHaAutomationLevel" yaml:"proactiveHaAutomationLevel"`
	// Enables proactive HA, allowing for vSphere to get HA data from external providers and use DRS to perform remediation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#proactive_ha_enabled ComputeCluster#proactive_ha_enabled}
	ProactiveHaEnabled interface{} `field:"optional" json:"proactiveHaEnabled" yaml:"proactiveHaEnabled"`
	// The configured remediation for moderately degraded hosts.
	//
	// Can be one of MaintenanceMode or QuarantineMode. Note that this cannot be set to MaintenanceMode when proactive_ha_severe_remediation is set to QuarantineMode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#proactive_ha_moderate_remediation ComputeCluster#proactive_ha_moderate_remediation}
	ProactiveHaModerateRemediation *string `field:"optional" json:"proactiveHaModerateRemediation" yaml:"proactiveHaModerateRemediation"`
	// The list of IDs for health update providers configured for this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#proactive_ha_provider_ids ComputeCluster#proactive_ha_provider_ids}
	ProactiveHaProviderIds *[]*string `field:"optional" json:"proactiveHaProviderIds" yaml:"proactiveHaProviderIds"`
	// The configured remediation for severely degraded hosts.
	//
	// Can be one of MaintenanceMode or QuarantineMode. Note that this cannot be set to QuarantineMode when proactive_ha_moderate_remediation is set to MaintenanceMode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#proactive_ha_severe_remediation ComputeCluster#proactive_ha_severe_remediation}
	ProactiveHaSevereRemediation *string `field:"optional" json:"proactiveHaSevereRemediation" yaml:"proactiveHaSevereRemediation"`
	// A list of tag IDs to apply to this object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#tags ComputeCluster#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Whether the vSAN compression service is enabled for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#vsan_compression_enabled ComputeCluster#vsan_compression_enabled}
	VsanCompressionEnabled interface{} `field:"optional" json:"vsanCompressionEnabled" yaml:"vsanCompressionEnabled"`
	// Whether the vSAN deduplication service is enabled for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#vsan_dedup_enabled ComputeCluster#vsan_dedup_enabled}
	VsanDedupEnabled interface{} `field:"optional" json:"vsanDedupEnabled" yaml:"vsanDedupEnabled"`
	// vsan_disk_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#vsan_disk_group ComputeCluster#vsan_disk_group}
	VsanDiskGroup interface{} `field:"optional" json:"vsanDiskGroup" yaml:"vsanDiskGroup"`
	// Whether the vSAN data-in-transit encryption is enabled for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#vsan_dit_encryption_enabled ComputeCluster#vsan_dit_encryption_enabled}
	VsanDitEncryptionEnabled interface{} `field:"optional" json:"vsanDitEncryptionEnabled" yaml:"vsanDitEncryptionEnabled"`
	// When vsan_dit_encryption_enabled is enabled, sets the rekey interval of data-in-transit encryption (in minutes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#vsan_dit_rekey_interval ComputeCluster#vsan_dit_rekey_interval}
	VsanDitRekeyInterval *float64 `field:"optional" json:"vsanDitRekeyInterval" yaml:"vsanDitRekeyInterval"`
	// Whether the vSAN service is enabled for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#vsan_enabled ComputeCluster#vsan_enabled}
	VsanEnabled interface{} `field:"optional" json:"vsanEnabled" yaml:"vsanEnabled"`
	// Whether the vSAN network diagnostic mode is enabled for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#vsan_network_diagnostic_mode_enabled ComputeCluster#vsan_network_diagnostic_mode_enabled}
	VsanNetworkDiagnosticModeEnabled interface{} `field:"optional" json:"vsanNetworkDiagnosticModeEnabled" yaml:"vsanNetworkDiagnosticModeEnabled"`
	// Whether the vSAN performance service is enabled for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#vsan_performance_enabled ComputeCluster#vsan_performance_enabled}
	VsanPerformanceEnabled interface{} `field:"optional" json:"vsanPerformanceEnabled" yaml:"vsanPerformanceEnabled"`
	// The managed object IDs of the vSAN datastore to be mounted on the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#vsan_remote_datastore_ids ComputeCluster#vsan_remote_datastore_ids}
	VsanRemoteDatastoreIds *[]*string `field:"optional" json:"vsanRemoteDatastoreIds" yaml:"vsanRemoteDatastoreIds"`
	// Whether the vSAN unmap service is enabled for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#vsan_unmap_enabled ComputeCluster#vsan_unmap_enabled}
	VsanUnmapEnabled interface{} `field:"optional" json:"vsanUnmapEnabled" yaml:"vsanUnmapEnabled"`
	// Whether the vSAN verbose mode is enabled for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.3.1/docs/resources/compute_cluster#vsan_verbose_mode_enabled ComputeCluster#vsan_verbose_mode_enabled}
	VsanVerboseModeEnabled interface{} `field:"optional" json:"vsanVerboseModeEnabled" yaml:"vsanVerboseModeEnabled"`
}

