// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vappentity

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.vappEntity.VappEntity",
		reflect.TypeOf((*VappEntity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "containerId", GoGetter: "ContainerId"},
			_jsii_.MemberProperty{JsiiProperty: "containerIdInput", GoGetter: "ContainerIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomAttributes", GoMethod: "ResetCustomAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetStartAction", GoMethod: "ResetStartAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetStartDelay", GoMethod: "ResetStartDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resetStartOrder", GoMethod: "ResetStartOrder"},
			_jsii_.MemberMethod{JsiiMethod: "resetStopAction", GoMethod: "ResetStopAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetStopDelay", GoMethod: "ResetStopDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetWaitForGuest", GoMethod: "ResetWaitForGuest"},
			_jsii_.MemberProperty{JsiiProperty: "startAction", GoGetter: "StartAction"},
			_jsii_.MemberProperty{JsiiProperty: "startActionInput", GoGetter: "StartActionInput"},
			_jsii_.MemberProperty{JsiiProperty: "startDelay", GoGetter: "StartDelay"},
			_jsii_.MemberProperty{JsiiProperty: "startDelayInput", GoGetter: "StartDelayInput"},
			_jsii_.MemberProperty{JsiiProperty: "startOrder", GoGetter: "StartOrder"},
			_jsii_.MemberProperty{JsiiProperty: "startOrderInput", GoGetter: "StartOrderInput"},
			_jsii_.MemberProperty{JsiiProperty: "stopAction", GoGetter: "StopAction"},
			_jsii_.MemberProperty{JsiiProperty: "stopActionInput", GoGetter: "StopActionInput"},
			_jsii_.MemberProperty{JsiiProperty: "stopDelay", GoGetter: "StopDelay"},
			_jsii_.MemberProperty{JsiiProperty: "stopDelayInput", GoGetter: "StopDelayInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetId", GoGetter: "TargetId"},
			_jsii_.MemberProperty{JsiiProperty: "targetIdInput", GoGetter: "TargetIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "waitForGuest", GoGetter: "WaitForGuest"},
			_jsii_.MemberProperty{JsiiProperty: "waitForGuestInput", GoGetter: "WaitForGuestInput"},
		},
		func() interface{} {
			j := jsiiProxy_VappEntity{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vsphere.vappEntity.VappEntityConfig",
		reflect.TypeOf((*VappEntityConfig)(nil)).Elem(),
	)
}
