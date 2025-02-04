// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vappentity

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VappEntityConfig struct {
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
	// Managed object ID of the vApp container the entity is a member of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/vapp_entity#container_id VappEntity#container_id}
	ContainerId *string `field:"required" json:"containerId" yaml:"containerId"`
	// Managed object ID of the entity to power on or power off.
	//
	// This can be a virtual machine or a vApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/vapp_entity#target_id VappEntity#target_id}
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	// A list of custom attributes to set on this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/vapp_entity#custom_attributes VappEntity#custom_attributes}
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/vapp_entity#id VappEntity#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// How to start the entity.
	//
	// Valid settings are none or powerOn. If set to none, then the entity does not participate in auto-start.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/vapp_entity#start_action VappEntity#start_action}
	StartAction *string `field:"optional" json:"startAction" yaml:"startAction"`
	// Delay in seconds before continuing with the next entity in the order of entities to be started.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/vapp_entity#start_delay VappEntity#start_delay}
	StartDelay *float64 `field:"optional" json:"startDelay" yaml:"startDelay"`
	// Order to start and stop target in vApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/vapp_entity#start_order VappEntity#start_order}
	StartOrder *float64 `field:"optional" json:"startOrder" yaml:"startOrder"`
	// Defines the stop action for the entity.
	//
	// Can be set to none, powerOff, guestShutdown, or suspend. If set to none, then the entity does not participate in auto-stop.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/vapp_entity#stop_action VappEntity#stop_action}
	StopAction *string `field:"optional" json:"stopAction" yaml:"stopAction"`
	// Delay in seconds before continuing with the next entity in the order of entities to be stopped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/vapp_entity#stop_delay VappEntity#stop_delay}
	StopDelay *float64 `field:"optional" json:"stopDelay" yaml:"stopDelay"`
	// A list of tag IDs to apply to this object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/vapp_entity#tags VappEntity#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Determines if the VM should be marked as being started when VMware Tools are ready instead of waiting for start_delay.
	//
	// This property has no effect for vApps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.11.1/docs/resources/vapp_entity#wait_for_guest VappEntity#wait_for_guest}
	WaitForGuest interface{} `field:"optional" json:"waitForGuest" yaml:"waitForGuest"`
}

