// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package guestoscustomization


type GuestOsCustomizationSpecWindowsOptions struct {
	// The host name for this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/guest_os_customization#computer_name GuestOsCustomization#computer_name}
	ComputerName *string `field:"required" json:"computerName" yaml:"computerName"`
	// The new administrator password for this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/guest_os_customization#admin_password GuestOsCustomization#admin_password}
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// Specifies whether or not the VM automatically logs on as Administrator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/guest_os_customization#auto_logon GuestOsCustomization#auto_logon}
	AutoLogon interface{} `field:"optional" json:"autoLogon" yaml:"autoLogon"`
	// Specifies how many times the VM should auto-logon the Administrator account when auto_logon is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/guest_os_customization#auto_logon_count GuestOsCustomization#auto_logon_count}
	AutoLogonCount *float64 `field:"optional" json:"autoLogonCount" yaml:"autoLogonCount"`
	// The password of the domain administrator used to join this virtual machine to the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/guest_os_customization#domain_admin_password GuestOsCustomization#domain_admin_password}
	DomainAdminPassword *string `field:"optional" json:"domainAdminPassword" yaml:"domainAdminPassword"`
	// The user account of the domain administrator used to join this virtual machine to the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/guest_os_customization#domain_admin_user GuestOsCustomization#domain_admin_user}
	DomainAdminUser *string `field:"optional" json:"domainAdminUser" yaml:"domainAdminUser"`
	// The MachineObjectOU which specifies the full LDAP path name of the OU to which the virtual machine belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/guest_os_customization#domain_ou GuestOsCustomization#domain_ou}
	DomainOu *string `field:"optional" json:"domainOu" yaml:"domainOu"`
	// The full name of the user of this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/guest_os_customization#full_name GuestOsCustomization#full_name}
	FullName *string `field:"optional" json:"fullName" yaml:"fullName"`
	// The domain that the virtual machine should join.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/guest_os_customization#join_domain GuestOsCustomization#join_domain}
	JoinDomain *string `field:"optional" json:"joinDomain" yaml:"joinDomain"`
	// The organization name this virtual machine is being installed for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/guest_os_customization#organization_name GuestOsCustomization#organization_name}
	OrganizationName *string `field:"optional" json:"organizationName" yaml:"organizationName"`
	// The product key for this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/guest_os_customization#product_key GuestOsCustomization#product_key}
	ProductKey *string `field:"optional" json:"productKey" yaml:"productKey"`
	// A list of commands to run at first user logon, after guest customization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/guest_os_customization#run_once_command_list GuestOsCustomization#run_once_command_list}
	RunOnceCommandList *[]*string `field:"optional" json:"runOnceCommandList" yaml:"runOnceCommandList"`
	// The new time zone for the virtual machine. This is a sysprep-dictated timezone code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/guest_os_customization#time_zone GuestOsCustomization#time_zone}
	TimeZone *float64 `field:"optional" json:"timeZone" yaml:"timeZone"`
	// The workgroup for this virtual machine if not joining a domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.9.3/docs/resources/guest_os_customization#workgroup GuestOsCustomization#workgroup}
	Workgroup *string `field:"optional" json:"workgroup" yaml:"workgroup"`
}

