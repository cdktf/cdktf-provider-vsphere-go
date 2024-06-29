// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datavspherevirtualmachine

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVsphereVirtualMachineConfig struct {
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
	// The guest name for the operating system when guest_id is otherGuest or otherGuest64.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#alternate_guest_name DataVsphereVirtualMachine#alternate_guest_name}
	AlternateGuestName *string `field:"optional" json:"alternateGuestName" yaml:"alternateGuestName"`
	// User-provided description of the virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#annotation DataVsphereVirtualMachine#annotation}
	Annotation *string `field:"optional" json:"annotation" yaml:"annotation"`
	// The number of milliseconds to wait before starting the boot sequence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#boot_delay DataVsphereVirtualMachine#boot_delay}
	BootDelay *float64 `field:"optional" json:"bootDelay" yaml:"bootDelay"`
	// The number of milliseconds to wait before retrying the boot sequence. This only valid if boot_retry_enabled is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#boot_retry_delay DataVsphereVirtualMachine#boot_retry_delay}
	BootRetryDelay *float64 `field:"optional" json:"bootRetryDelay" yaml:"bootRetryDelay"`
	// If set to true, a virtual machine that fails to boot will try again after the delay defined in boot_retry_delay.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#boot_retry_enabled DataVsphereVirtualMachine#boot_retry_enabled}
	BootRetryEnabled interface{} `field:"optional" json:"bootRetryEnabled" yaml:"bootRetryEnabled"`
	// Allow CPUs to be added to this virtual machine while it is running.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#cpu_hot_add_enabled DataVsphereVirtualMachine#cpu_hot_add_enabled}
	CpuHotAddEnabled interface{} `field:"optional" json:"cpuHotAddEnabled" yaml:"cpuHotAddEnabled"`
	// Allow CPUs to be added to this virtual machine while it is running.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#cpu_hot_remove_enabled DataVsphereVirtualMachine#cpu_hot_remove_enabled}
	CpuHotRemoveEnabled interface{} `field:"optional" json:"cpuHotRemoveEnabled" yaml:"cpuHotRemoveEnabled"`
	// The maximum amount of memory (in MB) or CPU (in MHz) that this virtual machine can consume, regardless of available resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#cpu_limit DataVsphereVirtualMachine#cpu_limit}
	CpuLimit *float64 `field:"optional" json:"cpuLimit" yaml:"cpuLimit"`
	// Enable CPU performance counters on this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#cpu_performance_counters_enabled DataVsphereVirtualMachine#cpu_performance_counters_enabled}
	CpuPerformanceCountersEnabled interface{} `field:"optional" json:"cpuPerformanceCountersEnabled" yaml:"cpuPerformanceCountersEnabled"`
	// The amount of memory (in MB) or CPU (in MHz) that this virtual machine is guaranteed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#cpu_reservation DataVsphereVirtualMachine#cpu_reservation}
	CpuReservation *float64 `field:"optional" json:"cpuReservation" yaml:"cpuReservation"`
	// The amount of shares to allocate to cpu for a custom share level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#cpu_share_count DataVsphereVirtualMachine#cpu_share_count}
	CpuShareCount *float64 `field:"optional" json:"cpuShareCount" yaml:"cpuShareCount"`
	// The allocation level for cpu resources. Can be one of high, low, normal, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#cpu_share_level DataVsphereVirtualMachine#cpu_share_level}
	CpuShareLevel *string `field:"optional" json:"cpuShareLevel" yaml:"cpuShareLevel"`
	// The managed object ID of the datacenter the virtual machine is in.
	//
	// This is not required when using ESXi directly, or if there is only one datacenter in your infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#datacenter_id DataVsphereVirtualMachine#datacenter_id}
	DatacenterId *string `field:"optional" json:"datacenterId" yaml:"datacenterId"`
	// When the boot type set in firmware is efi, this enables EFI secure boot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#efi_secure_boot_enabled DataVsphereVirtualMachine#efi_secure_boot_enabled}
	EfiSecureBootEnabled interface{} `field:"optional" json:"efiSecureBootEnabled" yaml:"efiSecureBootEnabled"`
	// Expose the UUIDs of attached virtual disks to the virtual machine, allowing access to them in the guest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#enable_disk_uuid DataVsphereVirtualMachine#enable_disk_uuid}
	EnableDiskUuid interface{} `field:"optional" json:"enableDiskUuid" yaml:"enableDiskUuid"`
	// Enable logging on this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#enable_logging DataVsphereVirtualMachine#enable_logging}
	EnableLogging interface{} `field:"optional" json:"enableLogging" yaml:"enableLogging"`
	// The EPT/RVI (hardware memory virtualization) setting for this virtual machine. Can be one of automatic, on, or off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#ept_rvi_mode DataVsphereVirtualMachine#ept_rvi_mode}
	EptRviMode *string `field:"optional" json:"eptRviMode" yaml:"eptRviMode"`
	// Extra configuration data for this virtual machine.
	//
	// Can be used to supply advanced parameters not normally in configuration, such as instance metadata, or configuration data for OVF images.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#extra_config DataVsphereVirtualMachine#extra_config}
	ExtraConfig *map[string]*string `field:"optional" json:"extraConfig" yaml:"extraConfig"`
	// Allow the virtual machine to be rebooted when a change to `extra_config` occurs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#extra_config_reboot_required DataVsphereVirtualMachine#extra_config_reboot_required}
	ExtraConfigRebootRequired interface{} `field:"optional" json:"extraConfigRebootRequired" yaml:"extraConfigRebootRequired"`
	// The firmware interface to use on the virtual machine. Can be one of bios or efi.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#firmware DataVsphereVirtualMachine#firmware}
	Firmware *string `field:"optional" json:"firmware" yaml:"firmware"`
	// The name of the folder the virtual machine is in.
	//
	// Allows distinguishing virtual machines with the same name in different folder paths
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#folder DataVsphereVirtualMachine#folder}
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// The guest ID for the operating system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#guest_id DataVsphereVirtualMachine#guest_id}
	GuestId *string `field:"optional" json:"guestId" yaml:"guestId"`
	// The hardware version for the virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#hardware_version DataVsphereVirtualMachine#hardware_version}
	HardwareVersion *float64 `field:"optional" json:"hardwareVersion" yaml:"hardwareVersion"`
	// The (non-nested) hardware virtualization setting for this virtual machine. Can be one of hvAuto, hvOn, or hvOff.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#hv_mode DataVsphereVirtualMachine#hv_mode}
	HvMode *string `field:"optional" json:"hvMode" yaml:"hvMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#id DataVsphereVirtualMachine#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The number of IDE controllers to scan for disk sizes and controller types on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#ide_controller_scan_count DataVsphereVirtualMachine#ide_controller_scan_count}
	IdeControllerScanCount *float64 `field:"optional" json:"ideControllerScanCount" yaml:"ideControllerScanCount"`
	// Controls the scheduling delay of the virtual machine.
	//
	// Use a higher sensitivity for applications that require lower latency, such as VOIP, media player applications, or applications that require frequent access to mouse or keyboard devices. Can be one of low, normal, medium, or high.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#latency_sensitivity DataVsphereVirtualMachine#latency_sensitivity}
	LatencySensitivity *string `field:"optional" json:"latencySensitivity" yaml:"latencySensitivity"`
	// The size of the virtual machine's memory, in MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#memory DataVsphereVirtualMachine#memory}
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// Allow memory to be added to this virtual machine while it is running.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#memory_hot_add_enabled DataVsphereVirtualMachine#memory_hot_add_enabled}
	MemoryHotAddEnabled interface{} `field:"optional" json:"memoryHotAddEnabled" yaml:"memoryHotAddEnabled"`
	// The maximum amount of memory (in MB) or CPU (in MHz) that this virtual machine can consume, regardless of available resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#memory_limit DataVsphereVirtualMachine#memory_limit}
	MemoryLimit *float64 `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
	// The amount of memory (in MB) or CPU (in MHz) that this virtual machine is guaranteed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#memory_reservation DataVsphereVirtualMachine#memory_reservation}
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
	// If set true, memory resource reservation for this virtual machine will always be equal to the virtual machine's memory size;increases in memory size will be rejected when a corresponding reservation increase is not possible. This feature may only be enabled if it is currently possible to reserve all of the virtual machine's memory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#memory_reservation_locked_to_max DataVsphereVirtualMachine#memory_reservation_locked_to_max}
	MemoryReservationLockedToMax interface{} `field:"optional" json:"memoryReservationLockedToMax" yaml:"memoryReservationLockedToMax"`
	// The amount of shares to allocate to memory for a custom share level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#memory_share_count DataVsphereVirtualMachine#memory_share_count}
	MemoryShareCount *float64 `field:"optional" json:"memoryShareCount" yaml:"memoryShareCount"`
	// The allocation level for memory resources. Can be one of high, low, normal, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#memory_share_level DataVsphereVirtualMachine#memory_share_level}
	MemoryShareLevel *string `field:"optional" json:"memoryShareLevel" yaml:"memoryShareLevel"`
	// The machine object ID from VMware vSphere.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#moid DataVsphereVirtualMachine#moid}
	Moid *string `field:"optional" json:"moid" yaml:"moid"`
	// The name of this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#name DataVsphereVirtualMachine#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Enable nested hardware virtualization on this virtual machine, facilitating nested virtualization in the guest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#nested_hv_enabled DataVsphereVirtualMachine#nested_hv_enabled}
	NestedHvEnabled interface{} `field:"optional" json:"nestedHvEnabled" yaml:"nestedHvEnabled"`
	// The number of cores to distribute amongst the CPUs in this virtual machine.
	//
	// If specified, the value supplied to num_cpus must be evenly divisible by this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#num_cores_per_socket DataVsphereVirtualMachine#num_cores_per_socket}
	NumCoresPerSocket *float64 `field:"optional" json:"numCoresPerSocket" yaml:"numCoresPerSocket"`
	// The number of virtual processors to assign to this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#num_cpus DataVsphereVirtualMachine#num_cpus}
	NumCpus *float64 `field:"optional" json:"numCpus" yaml:"numCpus"`
	// Triggers replacement of resource whenever it changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#replace_trigger DataVsphereVirtualMachine#replace_trigger}
	ReplaceTrigger *string `field:"optional" json:"replaceTrigger" yaml:"replaceTrigger"`
	// Enable the run of scripts after virtual machine power-on when VMware Tools is installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#run_tools_scripts_after_power_on DataVsphereVirtualMachine#run_tools_scripts_after_power_on}
	RunToolsScriptsAfterPowerOn interface{} `field:"optional" json:"runToolsScriptsAfterPowerOn" yaml:"runToolsScriptsAfterPowerOn"`
	// Enable the run of scripts after virtual machine resume when when VMware Tools is installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#run_tools_scripts_after_resume DataVsphereVirtualMachine#run_tools_scripts_after_resume}
	RunToolsScriptsAfterResume interface{} `field:"optional" json:"runToolsScriptsAfterResume" yaml:"runToolsScriptsAfterResume"`
	// Enable the run of scripts before guest operating system reboot when VMware Tools is installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#run_tools_scripts_before_guest_reboot DataVsphereVirtualMachine#run_tools_scripts_before_guest_reboot}
	RunToolsScriptsBeforeGuestReboot interface{} `field:"optional" json:"runToolsScriptsBeforeGuestReboot" yaml:"runToolsScriptsBeforeGuestReboot"`
	// Enable the run of scripts before guest operating system shutdown when VMware Tools is installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#run_tools_scripts_before_guest_shutdown DataVsphereVirtualMachine#run_tools_scripts_before_guest_shutdown}
	RunToolsScriptsBeforeGuestShutdown interface{} `field:"optional" json:"runToolsScriptsBeforeGuestShutdown" yaml:"runToolsScriptsBeforeGuestShutdown"`
	// Enable the run of scripts before guest operating system standby when VMware Tools is installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#run_tools_scripts_before_guest_standby DataVsphereVirtualMachine#run_tools_scripts_before_guest_standby}
	RunToolsScriptsBeforeGuestStandby interface{} `field:"optional" json:"runToolsScriptsBeforeGuestStandby" yaml:"runToolsScriptsBeforeGuestStandby"`
	// The number of SATA controllers to scan for disk sizes and controller types on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#sata_controller_scan_count DataVsphereVirtualMachine#sata_controller_scan_count}
	SataControllerScanCount *float64 `field:"optional" json:"sataControllerScanCount" yaml:"sataControllerScanCount"`
	// The number of SCSI controllers to scan for disk sizes and controller types on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#scsi_controller_scan_count DataVsphereVirtualMachine#scsi_controller_scan_count}
	ScsiControllerScanCount *float64 `field:"optional" json:"scsiControllerScanCount" yaml:"scsiControllerScanCount"`
	// The ID of the storage policy to assign to the virtual machine home directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#storage_policy_id DataVsphereVirtualMachine#storage_policy_id}
	StoragePolicyId *string `field:"optional" json:"storagePolicyId" yaml:"storagePolicyId"`
	// The swap file placement policy for this virtual machine. Can be one of inherit, hostLocal, or vmDirectory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#swap_placement_policy DataVsphereVirtualMachine#swap_placement_policy}
	SwapPlacementPolicy *string `field:"optional" json:"swapPlacementPolicy" yaml:"swapPlacementPolicy"`
	// Enable guest clock synchronization with the host.
	//
	// On vSphere 7.0 U1 and above, with only this setting the clock is synchronized on startup and resume. Requires VMware Tools to be installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#sync_time_with_host DataVsphereVirtualMachine#sync_time_with_host}
	SyncTimeWithHost interface{} `field:"optional" json:"syncTimeWithHost" yaml:"syncTimeWithHost"`
	// Enable periodic clock synchronization with the host.
	//
	// Supported only on vSphere 7.0 U1 and above. On prior versions setting `sync_time_with_host` is enough for periodic synchronization. Requires VMware Tools to be installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#sync_time_with_host_periodically DataVsphereVirtualMachine#sync_time_with_host_periodically}
	SyncTimeWithHostPeriodically interface{} `field:"optional" json:"syncTimeWithHostPeriodically" yaml:"syncTimeWithHostPeriodically"`
	// Set the upgrade policy for VMware Tools. Can be one of `manual` or `upgradeAtPowerCycle`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#tools_upgrade_policy DataVsphereVirtualMachine#tools_upgrade_policy}
	ToolsUpgradePolicy *string `field:"optional" json:"toolsUpgradePolicy" yaml:"toolsUpgradePolicy"`
	// The UUID of the virtual machine. Also exposed as the ID of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#uuid DataVsphereVirtualMachine#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
	// vapp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#vapp DataVsphereVirtualMachine#vapp}
	Vapp *DataVsphereVirtualMachineVapp `field:"optional" json:"vapp" yaml:"vapp"`
	// Flag to specify if Virtualization-based security is enabled for this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#vbs_enabled DataVsphereVirtualMachine#vbs_enabled}
	VbsEnabled interface{} `field:"optional" json:"vbsEnabled" yaml:"vbsEnabled"`
	// Flag to specify if I/O MMU virtualization, also called Intel Virtualization Technology for Directed I/O (VT-d) and AMD I/O Virtualization (AMD-Vi or IOMMU), is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.2/docs/data-sources/virtual_machine#vvtd_enabled DataVsphereVirtualMachine#vvtd_enabled}
	VvtdEnabled interface{} `field:"optional" json:"vvtdEnabled" yaml:"vvtdEnabled"`
}

