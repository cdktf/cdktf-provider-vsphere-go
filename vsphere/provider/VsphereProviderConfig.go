// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type VsphereProviderConfig struct {
	// The user password for vSphere API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs#password VsphereProvider#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The user name for vSphere API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs#user VsphereProvider#user}
	User *string `field:"required" json:"user" yaml:"user"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs#alias VsphereProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// If set, VMware vSphere client will permit unverifiable SSL certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs#allow_unverified_ssl VsphereProvider#allow_unverified_ssl}
	AllowUnverifiedSsl interface{} `field:"optional" json:"allowUnverifiedSsl" yaml:"allowUnverifiedSsl"`
	// API timeout in minutes (Default: 5).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs#api_timeout VsphereProvider#api_timeout}
	ApiTimeout *float64 `field:"optional" json:"apiTimeout" yaml:"apiTimeout"`
	// govmomi debug.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs#client_debug VsphereProvider#client_debug}
	ClientDebug interface{} `field:"optional" json:"clientDebug" yaml:"clientDebug"`
	// govmomi debug path for debug.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs#client_debug_path VsphereProvider#client_debug_path}
	ClientDebugPath *string `field:"optional" json:"clientDebugPath" yaml:"clientDebugPath"`
	// govmomi debug path for a single run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs#client_debug_path_run VsphereProvider#client_debug_path_run}
	ClientDebugPathRun *string `field:"optional" json:"clientDebugPathRun" yaml:"clientDebugPathRun"`
	// Persist vSphere client sessions to disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs#persist_session VsphereProvider#persist_session}
	PersistSession interface{} `field:"optional" json:"persistSession" yaml:"persistSession"`
	// The directory to save vSphere REST API sessions to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs#rest_session_path VsphereProvider#rest_session_path}
	RestSessionPath *string `field:"optional" json:"restSessionPath" yaml:"restSessionPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs#vcenter_server VsphereProvider#vcenter_server}.
	VcenterServer *string `field:"optional" json:"vcenterServer" yaml:"vcenterServer"`
	// Keep alive interval for the VIM session in minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs#vim_keep_alive VsphereProvider#vim_keep_alive}
	VimKeepAlive *float64 `field:"optional" json:"vimKeepAlive" yaml:"vimKeepAlive"`
	// The directory to save vSphere SOAP API sessions to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs#vim_session_path VsphereProvider#vim_session_path}
	VimSessionPath *string `field:"optional" json:"vimSessionPath" yaml:"vimSessionPath"`
	// The vSphere Server name for vSphere API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.2/docs#vsphere_server VsphereProvider#vsphere_server}
	VsphereServer *string `field:"optional" json:"vsphereServer" yaml:"vsphereServer"`
}

