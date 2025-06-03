// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package entitypermissions


type EntityPermissionsPermissions struct {
	// Whether user_or_group field refers to a user or a group. True for a group and false for a user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/entity_permissions#is_group EntityPermissions#is_group}
	IsGroup interface{} `field:"required" json:"isGroup" yaml:"isGroup"`
	// Whether or not this permission propagates down the hierarchy to sub-entities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/entity_permissions#propagate EntityPermissions#propagate}
	Propagate interface{} `field:"required" json:"propagate" yaml:"propagate"`
	// Reference to the role providing the access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/entity_permissions#role_id EntityPermissions#role_id}
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// User or group receiving access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/entity_permissions#user_or_group EntityPermissions#user_or_group}
	UserOrGroup *string `field:"required" json:"userOrGroup" yaml:"userOrGroup"`
}

