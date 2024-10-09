// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package supervisor


type SupervisorNamespace struct {
	// The name of the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#name Supervisor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of content libraries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#content_libraries Supervisor#content_libraries}
	ContentLibraries *[]*string `field:"optional" json:"contentLibraries" yaml:"contentLibraries"`
	// A list of virtual machine classes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/supervisor#vm_classes Supervisor#vm_classes}
	VmClasses *[]*string `field:"optional" json:"vmClasses" yaml:"vmClasses"`
}

