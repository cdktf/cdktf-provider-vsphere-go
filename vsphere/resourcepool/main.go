// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcepool

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.resourcePool.ResourcePool",
		reflect.TypeOf((*ResourcePool)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "cpuExpandable", GoGetter: "CpuExpandable"},
			_jsii_.MemberProperty{JsiiProperty: "cpuExpandableInput", GoGetter: "CpuExpandableInput"},
			_jsii_.MemberProperty{JsiiProperty: "cpuLimit", GoGetter: "CpuLimit"},
			_jsii_.MemberProperty{JsiiProperty: "cpuLimitInput", GoGetter: "CpuLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "cpuReservation", GoGetter: "CpuReservation"},
			_jsii_.MemberProperty{JsiiProperty: "cpuReservationInput", GoGetter: "CpuReservationInput"},
			_jsii_.MemberProperty{JsiiProperty: "cpuShareLevel", GoGetter: "CpuShareLevel"},
			_jsii_.MemberProperty{JsiiProperty: "cpuShareLevelInput", GoGetter: "CpuShareLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "cpuShares", GoGetter: "CpuShares"},
			_jsii_.MemberProperty{JsiiProperty: "cpuSharesInput", GoGetter: "CpuSharesInput"},
			_jsii_.MemberProperty{JsiiProperty: "customAttributes", GoGetter: "CustomAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "customAttributesInput", GoGetter: "CustomAttributesInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "memoryExpandable", GoGetter: "MemoryExpandable"},
			_jsii_.MemberProperty{JsiiProperty: "memoryExpandableInput", GoGetter: "MemoryExpandableInput"},
			_jsii_.MemberProperty{JsiiProperty: "memoryLimit", GoGetter: "MemoryLimit"},
			_jsii_.MemberProperty{JsiiProperty: "memoryLimitInput", GoGetter: "MemoryLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "memoryReservation", GoGetter: "MemoryReservation"},
			_jsii_.MemberProperty{JsiiProperty: "memoryReservationInput", GoGetter: "MemoryReservationInput"},
			_jsii_.MemberProperty{JsiiProperty: "memoryShareLevel", GoGetter: "MemoryShareLevel"},
			_jsii_.MemberProperty{JsiiProperty: "memoryShareLevelInput", GoGetter: "MemoryShareLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "memoryShares", GoGetter: "MemoryShares"},
			_jsii_.MemberProperty{JsiiProperty: "memorySharesInput", GoGetter: "MemorySharesInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parentResourcePoolId", GoGetter: "ParentResourcePoolId"},
			_jsii_.MemberProperty{JsiiProperty: "parentResourcePoolIdInput", GoGetter: "ParentResourcePoolIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpuExpandable", GoMethod: "ResetCpuExpandable"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpuLimit", GoMethod: "ResetCpuLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpuReservation", GoMethod: "ResetCpuReservation"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpuShareLevel", GoMethod: "ResetCpuShareLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpuShares", GoMethod: "ResetCpuShares"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomAttributes", GoMethod: "ResetCustomAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemoryExpandable", GoMethod: "ResetMemoryExpandable"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemoryLimit", GoMethod: "ResetMemoryLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemoryReservation", GoMethod: "ResetMemoryReservation"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemoryShareLevel", GoMethod: "ResetMemoryShareLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemoryShares", GoMethod: "ResetMemoryShares"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetScaleDescendantsShares", GoMethod: "ResetScaleDescendantsShares"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberProperty{JsiiProperty: "scaleDescendantsShares", GoGetter: "ScaleDescendantsShares"},
			_jsii_.MemberProperty{JsiiProperty: "scaleDescendantsSharesInput", GoGetter: "ScaleDescendantsSharesInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_ResourcePool{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vsphere.resourcePool.ResourcePoolConfig",
		reflect.TypeOf((*ResourcePoolConfig)(nil)).Elem(),
	)
}
