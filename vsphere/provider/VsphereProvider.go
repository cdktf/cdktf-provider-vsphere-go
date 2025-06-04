// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vsphere-go/vsphere/v11/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs vsphere}.
type VsphereProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AllowUnverifiedSsl() interface{}
	SetAllowUnverifiedSsl(val interface{})
	AllowUnverifiedSslInput() interface{}
	ApiTimeout() *float64
	SetApiTimeout(val *float64)
	ApiTimeoutInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientDebug() interface{}
	SetClientDebug(val interface{})
	ClientDebugInput() interface{}
	ClientDebugPath() *string
	SetClientDebugPath(val *string)
	ClientDebugPathInput() *string
	ClientDebugPathRun() *string
	SetClientDebugPathRun(val *string)
	ClientDebugPathRunInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PersistSession() interface{}
	SetPersistSession(val interface{})
	PersistSessionInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	RestSessionPath() *string
	SetRestSessionPath(val *string)
	RestSessionPathInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	User() *string
	SetUser(val *string)
	UserInput() *string
	VcenterServer() *string
	SetVcenterServer(val *string)
	VcenterServerInput() *string
	VimKeepAlive() *float64
	SetVimKeepAlive(val *float64)
	VimKeepAliveInput() *float64
	VimSessionPath() *string
	SetVimSessionPath(val *string)
	VimSessionPathInput() *string
	VsphereServer() *string
	SetVsphereServer(val *string)
	VsphereServerInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetAllowUnverifiedSsl()
	ResetApiTimeout()
	ResetClientDebug()
	ResetClientDebugPath()
	ResetClientDebugPathRun()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPersistSession()
	ResetRestSessionPath()
	ResetVcenterServer()
	ResetVimKeepAlive()
	ResetVimSessionPath()
	ResetVsphereServer()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VsphereProvider
type jsiiProxy_VsphereProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_VsphereProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) AllowUnverifiedSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUnverifiedSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) AllowUnverifiedSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUnverifiedSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) ApiTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"apiTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) ApiTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"apiTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) ClientDebug() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientDebug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) ClientDebugInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientDebugInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) ClientDebugPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientDebugPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) ClientDebugPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientDebugPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) ClientDebugPathRun() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientDebugPathRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) ClientDebugPathRunInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientDebugPathRunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) PersistSession() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistSession",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) PersistSessionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistSessionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) RestSessionPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restSessionPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) RestSessionPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restSessionPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) VcenterServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vcenterServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) VcenterServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vcenterServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) VimKeepAlive() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vimKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) VimKeepAliveInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vimKeepAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) VimSessionPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vimSessionPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) VimSessionPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vimSessionPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) VsphereServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vsphereServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsphereProvider) VsphereServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vsphereServerInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs vsphere} Resource.
func NewVsphereProvider(scope constructs.Construct, id *string, config *VsphereProviderConfig) VsphereProvider {
	_init_.Initialize()

	if err := validateNewVsphereProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VsphereProvider{}

	_jsii_.Create(
		"@cdktf/provider-vsphere.provider.VsphereProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs vsphere} Resource.
func NewVsphereProvider_Override(v VsphereProvider, scope constructs.Construct, id *string, config *VsphereProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vsphere.provider.VsphereProvider",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VsphereProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_VsphereProvider)SetAllowUnverifiedSsl(val interface{}) {
	if err := j.validateSetAllowUnverifiedSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowUnverifiedSsl",
		val,
	)
}

func (j *jsiiProxy_VsphereProvider)SetApiTimeout(val *float64) {
	_jsii_.Set(
		j,
		"apiTimeout",
		val,
	)
}

func (j *jsiiProxy_VsphereProvider)SetClientDebug(val interface{}) {
	if err := j.validateSetClientDebugParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientDebug",
		val,
	)
}

func (j *jsiiProxy_VsphereProvider)SetClientDebugPath(val *string) {
	_jsii_.Set(
		j,
		"clientDebugPath",
		val,
	)
}

func (j *jsiiProxy_VsphereProvider)SetClientDebugPathRun(val *string) {
	_jsii_.Set(
		j,
		"clientDebugPathRun",
		val,
	)
}

func (j *jsiiProxy_VsphereProvider)SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_VsphereProvider)SetPersistSession(val interface{}) {
	if err := j.validateSetPersistSessionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistSession",
		val,
	)
}

func (j *jsiiProxy_VsphereProvider)SetRestSessionPath(val *string) {
	_jsii_.Set(
		j,
		"restSessionPath",
		val,
	)
}

func (j *jsiiProxy_VsphereProvider)SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (j *jsiiProxy_VsphereProvider)SetVcenterServer(val *string) {
	_jsii_.Set(
		j,
		"vcenterServer",
		val,
	)
}

func (j *jsiiProxy_VsphereProvider)SetVimKeepAlive(val *float64) {
	_jsii_.Set(
		j,
		"vimKeepAlive",
		val,
	)
}

func (j *jsiiProxy_VsphereProvider)SetVimSessionPath(val *string) {
	_jsii_.Set(
		j,
		"vimSessionPath",
		val,
	)
}

func (j *jsiiProxy_VsphereProvider)SetVsphereServer(val *string) {
	_jsii_.Set(
		j,
		"vsphereServer",
		val,
	)
}

// Generates CDKTF code for importing a VsphereProvider resource upon running "cdktf plan <stack-name>".
func VsphereProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVsphereProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.provider.VsphereProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func VsphereProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVsphereProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.provider.VsphereProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VsphereProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVsphereProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.provider.VsphereProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VsphereProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVsphereProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vsphere.provider.VsphereProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VsphereProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vsphere.provider.VsphereProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VsphereProvider) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VsphereProvider) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VsphereProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		v,
		"resetAlias",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsphereProvider) ResetAllowUnverifiedSsl() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowUnverifiedSsl",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsphereProvider) ResetApiTimeout() {
	_jsii_.InvokeVoid(
		v,
		"resetApiTimeout",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsphereProvider) ResetClientDebug() {
	_jsii_.InvokeVoid(
		v,
		"resetClientDebug",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsphereProvider) ResetClientDebugPath() {
	_jsii_.InvokeVoid(
		v,
		"resetClientDebugPath",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsphereProvider) ResetClientDebugPathRun() {
	_jsii_.InvokeVoid(
		v,
		"resetClientDebugPathRun",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsphereProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsphereProvider) ResetPersistSession() {
	_jsii_.InvokeVoid(
		v,
		"resetPersistSession",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsphereProvider) ResetRestSessionPath() {
	_jsii_.InvokeVoid(
		v,
		"resetRestSessionPath",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsphereProvider) ResetVcenterServer() {
	_jsii_.InvokeVoid(
		v,
		"resetVcenterServer",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsphereProvider) ResetVimKeepAlive() {
	_jsii_.InvokeVoid(
		v,
		"resetVimKeepAlive",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsphereProvider) ResetVimSessionPath() {
	_jsii_.InvokeVoid(
		v,
		"resetVimSessionPath",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsphereProvider) ResetVsphereServer() {
	_jsii_.InvokeVoid(
		v,
		"resetVsphereServer",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsphereProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VsphereProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VsphereProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VsphereProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VsphereProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VsphereProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

