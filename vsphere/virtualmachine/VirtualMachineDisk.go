// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine


type VirtualMachineDisk struct {
	// A unique label for this disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/virtual_machine#label VirtualMachine#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// If this is true, the disk is attached instead of created. Implies keep_on_remove.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/virtual_machine#attach VirtualMachine#attach}
	Attach interface{} `field:"optional" json:"attach" yaml:"attach"`
	// The type of controller the disk should be connected to. Must be 'scsi', 'sata', or 'ide'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/virtual_machine#controller_type VirtualMachine#controller_type}
	ControllerType *string `field:"optional" json:"controllerType" yaml:"controllerType"`
	// The datastore ID for this virtual disk, if different than the virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/virtual_machine#datastore_id VirtualMachine#datastore_id}
	DatastoreId *string `field:"optional" json:"datastoreId" yaml:"datastoreId"`
	// The mode of this this virtual disk for purposes of writes and snapshotting.
	//
	// Can be one of append, independent_nonpersistent, independent_persistent, nonpersistent, persistent, or undoable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/virtual_machine#disk_mode VirtualMachine#disk_mode}
	DiskMode *string `field:"optional" json:"diskMode" yaml:"diskMode"`
	// The sharing mode of this virtual disk. Can be one of sharingMultiWriter or sharingNone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/virtual_machine#disk_sharing VirtualMachine#disk_sharing}
	DiskSharing *string `field:"optional" json:"diskSharing" yaml:"diskSharing"`
	// The virtual disk file zeroing policy when thin_provision is not true.
	//
	// The default is false, which lazily-zeros the disk, speeding up thick-provisioned disk creation time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/virtual_machine#eagerly_scrub VirtualMachine#eagerly_scrub}
	EagerlyScrub interface{} `field:"optional" json:"eagerlyScrub" yaml:"eagerlyScrub"`
	// The upper limit of IOPS that this disk can use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/virtual_machine#io_limit VirtualMachine#io_limit}
	IoLimit *float64 `field:"optional" json:"ioLimit" yaml:"ioLimit"`
	// The I/O guarantee that this disk has, in IOPS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/virtual_machine#io_reservation VirtualMachine#io_reservation}
	IoReservation *float64 `field:"optional" json:"ioReservation" yaml:"ioReservation"`
	// The share count for this disk when the share level is custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/virtual_machine#io_share_count VirtualMachine#io_share_count}
	IoShareCount *float64 `field:"optional" json:"ioShareCount" yaml:"ioShareCount"`
	// The share allocation level for this disk. Can be one of low, normal, high, or custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/virtual_machine#io_share_level VirtualMachine#io_share_level}
	IoShareLevel *string `field:"optional" json:"ioShareLevel" yaml:"ioShareLevel"`
	// Set to true to keep the underlying VMDK file when removing this virtual disk from configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/virtual_machine#keep_on_remove VirtualMachine#keep_on_remove}
	KeepOnRemove interface{} `field:"optional" json:"keepOnRemove" yaml:"keepOnRemove"`
	// The full path of the virtual disk.
	//
	// This can only be provided if attach is set to true, otherwise it is a read-only value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/virtual_machine#path VirtualMachine#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The size of the disk, in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/virtual_machine#size VirtualMachine#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The ID of the storage policy to assign to the virtual disk in VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/virtual_machine#storage_policy_id VirtualMachine#storage_policy_id}
	StoragePolicyId *string `field:"optional" json:"storagePolicyId" yaml:"storagePolicyId"`
	// If true, this disk is thin provisioned, with space for the file being allocated on an as-needed basis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/virtual_machine#thin_provisioned VirtualMachine#thin_provisioned}
	ThinProvisioned interface{} `field:"optional" json:"thinProvisioned" yaml:"thinProvisioned"`
	// The unique device number for this disk.
	//
	// This number determines where on the SCSI bus this device will be attached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/virtual_machine#unit_number VirtualMachine#unit_number}
	UnitNumber *float64 `field:"optional" json:"unitNumber" yaml:"unitNumber"`
	// If true, writes for this disk are sent directly to the filesystem immediately instead of being buffered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.0/docs/resources/virtual_machine#write_through VirtualMachine#write_through}
	WriteThrough interface{} `field:"optional" json:"writeThrough" yaml:"writeThrough"`
}

