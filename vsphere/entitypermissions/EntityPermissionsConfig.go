// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package entitypermissions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EntityPermissionsConfig struct {
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
	// The managed object id or uuid of the entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.3/docs/resources/entity_permissions#entity_id EntityPermissions#entity_id}
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// The entity managed object type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.3/docs/resources/entity_permissions#entity_type EntityPermissions#entity_type}
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
	// permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.3/docs/resources/entity_permissions#permissions EntityPermissions#permissions}
	Permissions interface{} `field:"required" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.8.3/docs/resources/entity_permissions#id EntityPermissions#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

