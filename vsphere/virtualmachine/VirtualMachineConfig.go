// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineConfig struct {
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
	// The name of this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#name VirtualMachine#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of a resource pool to put the virtual machine in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#resource_pool_id VirtualMachine#resource_pool_id}
	ResourcePoolId *string `field:"required" json:"resourcePoolId" yaml:"resourcePoolId"`
	// The guest name for the operating system when guest_id is otherGuest or otherGuest64.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#alternate_guest_name VirtualMachine#alternate_guest_name}
	AlternateGuestName *string `field:"optional" json:"alternateGuestName" yaml:"alternateGuestName"`
	// User-provided description of the virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#annotation VirtualMachine#annotation}
	Annotation *string `field:"optional" json:"annotation" yaml:"annotation"`
	// The number of milliseconds to wait before starting the boot sequence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#boot_delay VirtualMachine#boot_delay}
	BootDelay *float64 `field:"optional" json:"bootDelay" yaml:"bootDelay"`
	// The number of milliseconds to wait before retrying the boot sequence. This only valid if boot_retry_enabled is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#boot_retry_delay VirtualMachine#boot_retry_delay}
	BootRetryDelay *float64 `field:"optional" json:"bootRetryDelay" yaml:"bootRetryDelay"`
	// If set to true, a virtual machine that fails to boot will try again after the delay defined in boot_retry_delay.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#boot_retry_enabled VirtualMachine#boot_retry_enabled}
	BootRetryEnabled interface{} `field:"optional" json:"bootRetryEnabled" yaml:"bootRetryEnabled"`
	// cdrom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#cdrom VirtualMachine#cdrom}
	Cdrom interface{} `field:"optional" json:"cdrom" yaml:"cdrom"`
	// clone block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#clone VirtualMachine#clone}
	Clone *VirtualMachineClone `field:"optional" json:"clone" yaml:"clone"`
	// Allow CPUs to be added to this virtual machine while it is running.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#cpu_hot_add_enabled VirtualMachine#cpu_hot_add_enabled}
	CpuHotAddEnabled interface{} `field:"optional" json:"cpuHotAddEnabled" yaml:"cpuHotAddEnabled"`
	// Allow CPUs to be added to this virtual machine while it is running.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#cpu_hot_remove_enabled VirtualMachine#cpu_hot_remove_enabled}
	CpuHotRemoveEnabled interface{} `field:"optional" json:"cpuHotRemoveEnabled" yaml:"cpuHotRemoveEnabled"`
	// The maximum amount of memory (in MB) or CPU (in MHz) that this virtual machine can consume, regardless of available resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#cpu_limit VirtualMachine#cpu_limit}
	CpuLimit *float64 `field:"optional" json:"cpuLimit" yaml:"cpuLimit"`
	// Enable CPU performance counters on this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#cpu_performance_counters_enabled VirtualMachine#cpu_performance_counters_enabled}
	CpuPerformanceCountersEnabled interface{} `field:"optional" json:"cpuPerformanceCountersEnabled" yaml:"cpuPerformanceCountersEnabled"`
	// The amount of memory (in MB) or CPU (in MHz) that this virtual machine is guaranteed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#cpu_reservation VirtualMachine#cpu_reservation}
	CpuReservation *float64 `field:"optional" json:"cpuReservation" yaml:"cpuReservation"`
	// The amount of shares to allocate to cpu for a custom share level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#cpu_share_count VirtualMachine#cpu_share_count}
	CpuShareCount *float64 `field:"optional" json:"cpuShareCount" yaml:"cpuShareCount"`
	// The allocation level for cpu resources. Can be one of high, low, normal, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#cpu_share_level VirtualMachine#cpu_share_level}
	CpuShareLevel *string `field:"optional" json:"cpuShareLevel" yaml:"cpuShareLevel"`
	// A list of custom attributes to set on this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#custom_attributes VirtualMachine#custom_attributes}
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The ID of the datacenter where the VM is to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#datacenter_id VirtualMachine#datacenter_id}
	DatacenterId *string `field:"optional" json:"datacenterId" yaml:"datacenterId"`
	// The ID of a datastore cluster to put the virtual machine in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#datastore_cluster_id VirtualMachine#datastore_cluster_id}
	DatastoreClusterId *string `field:"optional" json:"datastoreClusterId" yaml:"datastoreClusterId"`
	// The ID of the virtual machine's datastore.
	//
	// The virtual machine configuration is placed here, along with any virtual disks that are created without datastores.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#datastore_id VirtualMachine#datastore_id}
	DatastoreId *string `field:"optional" json:"datastoreId" yaml:"datastoreId"`
	// disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#disk VirtualMachine#disk}
	Disk interface{} `field:"optional" json:"disk" yaml:"disk"`
	// When the boot type set in firmware is efi, this enables EFI secure boot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#efi_secure_boot_enabled VirtualMachine#efi_secure_boot_enabled}
	EfiSecureBootEnabled interface{} `field:"optional" json:"efiSecureBootEnabled" yaml:"efiSecureBootEnabled"`
	// Expose the UUIDs of attached virtual disks to the virtual machine, allowing access to them in the guest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#enable_disk_uuid VirtualMachine#enable_disk_uuid}
	EnableDiskUuid interface{} `field:"optional" json:"enableDiskUuid" yaml:"enableDiskUuid"`
	// Enable logging on this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#enable_logging VirtualMachine#enable_logging}
	EnableLogging interface{} `field:"optional" json:"enableLogging" yaml:"enableLogging"`
	// The EPT/RVI (hardware memory virtualization) setting for this virtual machine. Can be one of automatic, on, or off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#ept_rvi_mode VirtualMachine#ept_rvi_mode}
	EptRviMode *string `field:"optional" json:"eptRviMode" yaml:"eptRviMode"`
	// Extra configuration data for this virtual machine.
	//
	// Can be used to supply advanced parameters not normally in configuration, such as instance metadata, or configuration data for OVF images.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#extra_config VirtualMachine#extra_config}
	ExtraConfig *map[string]*string `field:"optional" json:"extraConfig" yaml:"extraConfig"`
	// Allow the virtual machine to be rebooted when a change to `extra_config` occurs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#extra_config_reboot_required VirtualMachine#extra_config_reboot_required}
	ExtraConfigRebootRequired interface{} `field:"optional" json:"extraConfigRebootRequired" yaml:"extraConfigRebootRequired"`
	// The firmware interface to use on the virtual machine. Can be one of bios or efi.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#firmware VirtualMachine#firmware}
	Firmware *string `field:"optional" json:"firmware" yaml:"firmware"`
	// The name of the folder to locate the virtual machine in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#folder VirtualMachine#folder}
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// Set to true to force power-off a virtual machine if a graceful guest shutdown failed for a necessary operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#force_power_off VirtualMachine#force_power_off}
	ForcePowerOff interface{} `field:"optional" json:"forcePowerOff" yaml:"forcePowerOff"`
	// The guest ID for the operating system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#guest_id VirtualMachine#guest_id}
	GuestId *string `field:"optional" json:"guestId" yaml:"guestId"`
	// The hardware version for the virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#hardware_version VirtualMachine#hardware_version}
	HardwareVersion *float64 `field:"optional" json:"hardwareVersion" yaml:"hardwareVersion"`
	// The ID of an optional host system to pin the virtual machine to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#host_system_id VirtualMachine#host_system_id}
	HostSystemId *string `field:"optional" json:"hostSystemId" yaml:"hostSystemId"`
	// The (non-nested) hardware virtualization setting for this virtual machine. Can be one of hvAuto, hvOn, or hvOff.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#hv_mode VirtualMachine#hv_mode}
	HvMode *string `field:"optional" json:"hvMode" yaml:"hvMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#id VirtualMachine#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The number of IDE controllers that Terraform manages on this virtual machine.
	//
	// This directly affects the amount of disks you can add to the virtual machine and the maximum disk unit number. Note that lowering this value does not remove controllers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#ide_controller_count VirtualMachine#ide_controller_count}
	IdeControllerCount *float64 `field:"optional" json:"ideControllerCount" yaml:"ideControllerCount"`
	// List of IP addresses and CIDR networks to ignore while waiting for an IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#ignored_guest_ips VirtualMachine#ignored_guest_ips}
	IgnoredGuestIps *[]*string `field:"optional" json:"ignoredGuestIps" yaml:"ignoredGuestIps"`
	// Controls the scheduling delay of the virtual machine.
	//
	// Use a higher sensitivity for applications that require lower latency, such as VOIP, media player applications, or applications that require frequent access to mouse or keyboard devices. Can be one of low, normal, medium, or high.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#latency_sensitivity VirtualMachine#latency_sensitivity}
	LatencySensitivity *string `field:"optional" json:"latencySensitivity" yaml:"latencySensitivity"`
	// The size of the virtual machine's memory, in MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#memory VirtualMachine#memory}
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// Allow memory to be added to this virtual machine while it is running.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#memory_hot_add_enabled VirtualMachine#memory_hot_add_enabled}
	MemoryHotAddEnabled interface{} `field:"optional" json:"memoryHotAddEnabled" yaml:"memoryHotAddEnabled"`
	// The maximum amount of memory (in MB) or CPU (in MHz) that this virtual machine can consume, regardless of available resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#memory_limit VirtualMachine#memory_limit}
	MemoryLimit *float64 `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
	// The amount of memory (in MB) or CPU (in MHz) that this virtual machine is guaranteed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#memory_reservation VirtualMachine#memory_reservation}
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
	// If set true, memory resource reservation for this virtual machine will always be equal to the virtual machine's memory size;increases in memory size will be rejected when a corresponding reservation increase is not possible. This feature may only be enabled if it is currently possible to reserve all of the virtual machine's memory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#memory_reservation_locked_to_max VirtualMachine#memory_reservation_locked_to_max}
	MemoryReservationLockedToMax interface{} `field:"optional" json:"memoryReservationLockedToMax" yaml:"memoryReservationLockedToMax"`
	// The amount of shares to allocate to memory for a custom share level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#memory_share_count VirtualMachine#memory_share_count}
	MemoryShareCount *float64 `field:"optional" json:"memoryShareCount" yaml:"memoryShareCount"`
	// The allocation level for memory resources. Can be one of high, low, normal, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#memory_share_level VirtualMachine#memory_share_level}
	MemoryShareLevel *string `field:"optional" json:"memoryShareLevel" yaml:"memoryShareLevel"`
	// The amount of time, in minutes, to wait for a vMotion operation to complete before failing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#migrate_wait_timeout VirtualMachine#migrate_wait_timeout}
	MigrateWaitTimeout *float64 `field:"optional" json:"migrateWaitTimeout" yaml:"migrateWaitTimeout"`
	// Enable nested hardware virtualization on this virtual machine, facilitating nested virtualization in the guest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#nested_hv_enabled VirtualMachine#nested_hv_enabled}
	NestedHvEnabled interface{} `field:"optional" json:"nestedHvEnabled" yaml:"nestedHvEnabled"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#network_interface VirtualMachine#network_interface}
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// The number of cores to distribute amongst the CPUs in this virtual machine.
	//
	// If specified, the value supplied to num_cpus must be evenly divisible by this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#num_cores_per_socket VirtualMachine#num_cores_per_socket}
	NumCoresPerSocket *float64 `field:"optional" json:"numCoresPerSocket" yaml:"numCoresPerSocket"`
	// The number of virtual processors to assign to this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#num_cpus VirtualMachine#num_cpus}
	NumCpus *float64 `field:"optional" json:"numCpus" yaml:"numCpus"`
	// The number of NVMe controllers that Terraform manages on this virtual machine.
	//
	// This directly affects the amount of disks you can add to the virtual machine and the maximum disk unit number. Note that lowering this value does not remove controllers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#nvme_controller_count VirtualMachine#nvme_controller_count}
	NvmeControllerCount *float64 `field:"optional" json:"nvmeControllerCount" yaml:"nvmeControllerCount"`
	// ovf_deploy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#ovf_deploy VirtualMachine#ovf_deploy}
	OvfDeploy *VirtualMachineOvfDeploy `field:"optional" json:"ovfDeploy" yaml:"ovfDeploy"`
	// A list of PCI passthrough devices.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#pci_device_id VirtualMachine#pci_device_id}
	PciDeviceId *[]*string `field:"optional" json:"pciDeviceId" yaml:"pciDeviceId"`
	// The amount of time, in seconds, that we will be trying to power on a VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#poweron_timeout VirtualMachine#poweron_timeout}
	PoweronTimeout *float64 `field:"optional" json:"poweronTimeout" yaml:"poweronTimeout"`
	// Triggers replacement of resource whenever it changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#replace_trigger VirtualMachine#replace_trigger}
	ReplaceTrigger *string `field:"optional" json:"replaceTrigger" yaml:"replaceTrigger"`
	// Enable the run of scripts after virtual machine power-on when VMware Tools is installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#run_tools_scripts_after_power_on VirtualMachine#run_tools_scripts_after_power_on}
	RunToolsScriptsAfterPowerOn interface{} `field:"optional" json:"runToolsScriptsAfterPowerOn" yaml:"runToolsScriptsAfterPowerOn"`
	// Enable the run of scripts after virtual machine resume when when VMware Tools is installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#run_tools_scripts_after_resume VirtualMachine#run_tools_scripts_after_resume}
	RunToolsScriptsAfterResume interface{} `field:"optional" json:"runToolsScriptsAfterResume" yaml:"runToolsScriptsAfterResume"`
	// Enable the run of scripts before guest operating system reboot when VMware Tools is installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#run_tools_scripts_before_guest_reboot VirtualMachine#run_tools_scripts_before_guest_reboot}
	RunToolsScriptsBeforeGuestReboot interface{} `field:"optional" json:"runToolsScriptsBeforeGuestReboot" yaml:"runToolsScriptsBeforeGuestReboot"`
	// Enable the run of scripts before guest operating system shutdown when VMware Tools is installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#run_tools_scripts_before_guest_shutdown VirtualMachine#run_tools_scripts_before_guest_shutdown}
	RunToolsScriptsBeforeGuestShutdown interface{} `field:"optional" json:"runToolsScriptsBeforeGuestShutdown" yaml:"runToolsScriptsBeforeGuestShutdown"`
	// Enable the run of scripts before guest operating system standby when VMware Tools is installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#run_tools_scripts_before_guest_standby VirtualMachine#run_tools_scripts_before_guest_standby}
	RunToolsScriptsBeforeGuestStandby interface{} `field:"optional" json:"runToolsScriptsBeforeGuestStandby" yaml:"runToolsScriptsBeforeGuestStandby"`
	// The number of SATA controllers that Terraform manages on this virtual machine.
	//
	// This directly affects the amount of disks you can add to the virtual machine and the maximum disk unit number. Note that lowering this value does not remove controllers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#sata_controller_count VirtualMachine#sata_controller_count}
	SataControllerCount *float64 `field:"optional" json:"sataControllerCount" yaml:"sataControllerCount"`
	// Mode for sharing the SCSI bus. The modes are physicalSharing, virtualSharing, and noSharing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#scsi_bus_sharing VirtualMachine#scsi_bus_sharing}
	ScsiBusSharing *string `field:"optional" json:"scsiBusSharing" yaml:"scsiBusSharing"`
	// The number of SCSI controllers that Terraform manages on this virtual machine.
	//
	// This directly affects the amount of disks you can add to the virtual machine and the maximum disk unit number. Note that lowering this value does not remove controllers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#scsi_controller_count VirtualMachine#scsi_controller_count}
	ScsiControllerCount *float64 `field:"optional" json:"scsiControllerCount" yaml:"scsiControllerCount"`
	// The type of SCSI bus this virtual machine will have. Can be one of lsilogic, lsilogic-sas or pvscsi.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#scsi_type VirtualMachine#scsi_type}
	ScsiType *string `field:"optional" json:"scsiType" yaml:"scsiType"`
	// The amount of time, in minutes, to wait for shutdown when making necessary updates to the virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#shutdown_wait_timeout VirtualMachine#shutdown_wait_timeout}
	ShutdownWaitTimeout *float64 `field:"optional" json:"shutdownWaitTimeout" yaml:"shutdownWaitTimeout"`
	// The ID of the storage policy to assign to the virtual machine home directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#storage_policy_id VirtualMachine#storage_policy_id}
	StoragePolicyId *string `field:"optional" json:"storagePolicyId" yaml:"storagePolicyId"`
	// The swap file placement policy for this virtual machine. Can be one of inherit, hostLocal, or vmDirectory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#swap_placement_policy VirtualMachine#swap_placement_policy}
	SwapPlacementPolicy *string `field:"optional" json:"swapPlacementPolicy" yaml:"swapPlacementPolicy"`
	// Enable guest clock synchronization with the host.
	//
	// On vSphere 7.0 U1 and above, with only this setting the clock is synchronized on startup and resume. Requires VMware Tools to be installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#sync_time_with_host VirtualMachine#sync_time_with_host}
	SyncTimeWithHost interface{} `field:"optional" json:"syncTimeWithHost" yaml:"syncTimeWithHost"`
	// Enable periodic clock synchronization with the host.
	//
	// Supported only on vSphere 7.0 U1 and above. On prior versions setting `sync_time_with_host` is enough for periodic synchronization. Requires VMware Tools to be installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#sync_time_with_host_periodically VirtualMachine#sync_time_with_host_periodically}
	SyncTimeWithHostPeriodically interface{} `field:"optional" json:"syncTimeWithHostPeriodically" yaml:"syncTimeWithHostPeriodically"`
	// A list of tag IDs to apply to this object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#tags VirtualMachine#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Set the upgrade policy for VMware Tools. Can be one of `manual` or `upgradeAtPowerCycle`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#tools_upgrade_policy VirtualMachine#tools_upgrade_policy}
	ToolsUpgradePolicy *string `field:"optional" json:"toolsUpgradePolicy" yaml:"toolsUpgradePolicy"`
	// vapp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#vapp VirtualMachine#vapp}
	Vapp *VirtualMachineVapp `field:"optional" json:"vapp" yaml:"vapp"`
	// Flag to specify if Virtualization-based security is enabled for this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#vbs_enabled VirtualMachine#vbs_enabled}
	VbsEnabled interface{} `field:"optional" json:"vbsEnabled" yaml:"vbsEnabled"`
	// vtpm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#vtpm VirtualMachine#vtpm}
	Vtpm *VirtualMachineVtpm `field:"optional" json:"vtpm" yaml:"vtpm"`
	// Flag to specify if I/O MMU virtualization, also called Intel Virtualization Technology for Directed I/O (VT-d) and AMD I/O Virtualization (AMD-Vi or IOMMU), is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#vvtd_enabled VirtualMachine#vvtd_enabled}
	VvtdEnabled interface{} `field:"optional" json:"vvtdEnabled" yaml:"vvtdEnabled"`
	// The amount of time, in minutes, to wait for an available IP address on this virtual machine.
	//
	// A value less than 1 disables the waiter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#wait_for_guest_ip_timeout VirtualMachine#wait_for_guest_ip_timeout}
	WaitForGuestIpTimeout *float64 `field:"optional" json:"waitForGuestIpTimeout" yaml:"waitForGuestIpTimeout"`
	// Controls whether or not the guest network waiter waits for a routable address.
	//
	// When false, the waiter does not wait for a default gateway, nor are IP addresses checked against any discovered default gateways as part of its success criteria.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#wait_for_guest_net_routable VirtualMachine#wait_for_guest_net_routable}
	WaitForGuestNetRoutable interface{} `field:"optional" json:"waitForGuestNetRoutable" yaml:"waitForGuestNetRoutable"`
	// The amount of time, in minutes, to wait for an available IP address on this virtual machine.
	//
	// A value less than 1 disables the waiter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.14.2/docs/resources/virtual_machine#wait_for_guest_net_timeout VirtualMachine#wait_for_guest_net_timeout}
	WaitForGuestNetTimeout *float64 `field:"optional" json:"waitForGuestNetTimeout" yaml:"waitForGuestNetTimeout"`
}

