// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagedrsvmoverride

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageDrsVmOverrideConfig struct {
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
	// The managed object ID of the datastore cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/storage_drs_vm_override#datastore_cluster_id StorageDrsVmOverride#datastore_cluster_id}
	DatastoreClusterId *string `field:"required" json:"datastoreClusterId" yaml:"datastoreClusterId"`
	// The managed object ID of the virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/storage_drs_vm_override#virtual_machine_id StorageDrsVmOverride#virtual_machine_id}
	VirtualMachineId *string `field:"required" json:"virtualMachineId" yaml:"virtualMachineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/storage_drs_vm_override#id StorageDrsVmOverride#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Overrides any Storage DRS automation levels for this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/storage_drs_vm_override#sdrs_automation_level StorageDrsVmOverride#sdrs_automation_level}
	SdrsAutomationLevel *string `field:"optional" json:"sdrsAutomationLevel" yaml:"sdrsAutomationLevel"`
	// Overrides the default Storage DRS setting for this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/storage_drs_vm_override#sdrs_enabled StorageDrsVmOverride#sdrs_enabled}
	SdrsEnabled *string `field:"optional" json:"sdrsEnabled" yaml:"sdrsEnabled"`
	// Overrides the intra-VM affinity setting for this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/storage_drs_vm_override#sdrs_intra_vm_affinity StorageDrsVmOverride#sdrs_intra_vm_affinity}
	SdrsIntraVmAffinity *string `field:"optional" json:"sdrsIntraVmAffinity" yaml:"sdrsIntraVmAffinity"`
}

