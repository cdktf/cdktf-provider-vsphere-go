// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine


type VirtualMachineCloneCustomizeWindowsOptions struct {
	// The host name for this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/virtual_machine#computer_name VirtualMachine#computer_name}
	ComputerName *string `field:"required" json:"computerName" yaml:"computerName"`
	// The new administrator password for this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/virtual_machine#admin_password VirtualMachine#admin_password}
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// Specifies whether or not the VM automatically logs on as Administrator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/virtual_machine#auto_logon VirtualMachine#auto_logon}
	AutoLogon interface{} `field:"optional" json:"autoLogon" yaml:"autoLogon"`
	// Specifies how many times the VM should auto-logon the Administrator account when auto_logon is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/virtual_machine#auto_logon_count VirtualMachine#auto_logon_count}
	AutoLogonCount *float64 `field:"optional" json:"autoLogonCount" yaml:"autoLogonCount"`
	// The password of the domain administrator used to join this virtual machine to the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/virtual_machine#domain_admin_password VirtualMachine#domain_admin_password}
	DomainAdminPassword *string `field:"optional" json:"domainAdminPassword" yaml:"domainAdminPassword"`
	// The user account of the domain administrator used to join this virtual machine to the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/virtual_machine#domain_admin_user VirtualMachine#domain_admin_user}
	DomainAdminUser *string `field:"optional" json:"domainAdminUser" yaml:"domainAdminUser"`
	// The MachineObjectOU which specifies the full LDAP path name of the OU to which the virtual machine belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/virtual_machine#domain_ou VirtualMachine#domain_ou}
	DomainOu *string `field:"optional" json:"domainOu" yaml:"domainOu"`
	// The full name of the user of this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/virtual_machine#full_name VirtualMachine#full_name}
	FullName *string `field:"optional" json:"fullName" yaml:"fullName"`
	// The domain that the virtual machine should join.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/virtual_machine#join_domain VirtualMachine#join_domain}
	JoinDomain *string `field:"optional" json:"joinDomain" yaml:"joinDomain"`
	// The organization name this virtual machine is being installed for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/virtual_machine#organization_name VirtualMachine#organization_name}
	OrganizationName *string `field:"optional" json:"organizationName" yaml:"organizationName"`
	// The product key for this virtual machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/virtual_machine#product_key VirtualMachine#product_key}
	ProductKey *string `field:"optional" json:"productKey" yaml:"productKey"`
	// A list of commands to run at first user logon, after guest customization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/virtual_machine#run_once_command_list VirtualMachine#run_once_command_list}
	RunOnceCommandList *[]*string `field:"optional" json:"runOnceCommandList" yaml:"runOnceCommandList"`
	// The new time zone for the virtual machine. This is a sysprep-dictated timezone code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/virtual_machine#time_zone VirtualMachine#time_zone}
	TimeZone *float64 `field:"optional" json:"timeZone" yaml:"timeZone"`
	// The workgroup for this virtual machine if not joining a domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vmware/vsphere/2.13.0/docs/resources/virtual_machine#workgroup VirtualMachine#workgroup}
	Workgroup *string `field:"optional" json:"workgroup" yaml:"workgroup"`
}

