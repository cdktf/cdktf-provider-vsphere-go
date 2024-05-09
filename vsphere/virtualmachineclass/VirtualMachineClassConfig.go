// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachineclass

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineClassConfig struct {
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
	// The number of CPUs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/virtual_machine_class#cpus VirtualMachineClass#cpus}
	Cpus *float64 `field:"required" json:"cpus" yaml:"cpus"`
	// The amount of memory (in MB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/virtual_machine_class#memory VirtualMachineClass#memory}
	Memory *float64 `field:"required" json:"memory" yaml:"memory"`
	// The name of the virtual machine class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/virtual_machine_class#name VirtualMachineClass#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The percentage of the available CPU capacity which will be reserved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/virtual_machine_class#cpu_reservation VirtualMachineClass#cpu_reservation}
	CpuReservation *float64 `field:"optional" json:"cpuReservation" yaml:"cpuReservation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/virtual_machine_class#id VirtualMachineClass#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The percentage of the available memory capacity which will be reserved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/virtual_machine_class#memory_reservation VirtualMachineClass#memory_reservation}
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
	// A comma-separated list of GPU devices.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.1/docs/resources/virtual_machine_class#vgpu_devices VirtualMachineClass#vgpu_devices}
	VgpuDevices *[]*string `field:"optional" json:"vgpuDevices" yaml:"vgpuDevices"`
}

