package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vsphere.provider.VsphereProvider",
		reflect.TypeOf((*VsphereProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowUnverifiedSsl", GoGetter: "AllowUnverifiedSsl"},
			_jsii_.MemberProperty{JsiiProperty: "allowUnverifiedSslInput", GoGetter: "AllowUnverifiedSslInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiTimeout", GoGetter: "ApiTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "apiTimeoutInput", GoGetter: "ApiTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientDebug", GoGetter: "ClientDebug"},
			_jsii_.MemberProperty{JsiiProperty: "clientDebugInput", GoGetter: "ClientDebugInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientDebugPath", GoGetter: "ClientDebugPath"},
			_jsii_.MemberProperty{JsiiProperty: "clientDebugPathInput", GoGetter: "ClientDebugPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientDebugPathRun", GoGetter: "ClientDebugPathRun"},
			_jsii_.MemberProperty{JsiiProperty: "clientDebugPathRunInput", GoGetter: "ClientDebugPathRunInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "passwordInput", GoGetter: "PasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "persistSession", GoGetter: "PersistSession"},
			_jsii_.MemberProperty{JsiiProperty: "persistSessionInput", GoGetter: "PersistSessionInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowUnverifiedSsl", GoMethod: "ResetAllowUnverifiedSsl"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiTimeout", GoMethod: "ResetApiTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientDebug", GoMethod: "ResetClientDebug"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientDebugPath", GoMethod: "ResetClientDebugPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientDebugPathRun", GoMethod: "ResetClientDebugPathRun"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPersistSession", GoMethod: "ResetPersistSession"},
			_jsii_.MemberMethod{JsiiMethod: "resetRestSessionPath", GoMethod: "ResetRestSessionPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetVcenterServer", GoMethod: "ResetVcenterServer"},
			_jsii_.MemberMethod{JsiiMethod: "resetVimKeepAlive", GoMethod: "ResetVimKeepAlive"},
			_jsii_.MemberMethod{JsiiMethod: "resetVimSessionPath", GoMethod: "ResetVimSessionPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetVsphereServer", GoMethod: "ResetVsphereServer"},
			_jsii_.MemberProperty{JsiiProperty: "restSessionPath", GoGetter: "RestSessionPath"},
			_jsii_.MemberProperty{JsiiProperty: "restSessionPathInput", GoGetter: "RestSessionPathInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
			_jsii_.MemberProperty{JsiiProperty: "userInput", GoGetter: "UserInput"},
			_jsii_.MemberProperty{JsiiProperty: "vcenterServer", GoGetter: "VcenterServer"},
			_jsii_.MemberProperty{JsiiProperty: "vcenterServerInput", GoGetter: "VcenterServerInput"},
			_jsii_.MemberProperty{JsiiProperty: "vimKeepAlive", GoGetter: "VimKeepAlive"},
			_jsii_.MemberProperty{JsiiProperty: "vimKeepAliveInput", GoGetter: "VimKeepAliveInput"},
			_jsii_.MemberProperty{JsiiProperty: "vimSessionPath", GoGetter: "VimSessionPath"},
			_jsii_.MemberProperty{JsiiProperty: "vimSessionPathInput", GoGetter: "VimSessionPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "vsphereServer", GoGetter: "VsphereServer"},
			_jsii_.MemberProperty{JsiiProperty: "vsphereServerInput", GoGetter: "VsphereServerInput"},
		},
		func() interface{} {
			j := jsiiProxy_VsphereProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vsphere.provider.VsphereProviderConfig",
		reflect.TypeOf((*VsphereProviderConfig)(nil)).Elem(),
	)
}
