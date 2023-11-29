// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachine


type VirtualMachineOvfDeploy struct {
	// Allow unverified ssl certificates while deploying ovf/ova from url.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.0/docs/resources/virtual_machine#allow_unverified_ssl_cert VirtualMachine#allow_unverified_ssl_cert}
	AllowUnverifiedSslCert interface{} `field:"optional" json:"allowUnverifiedSslCert" yaml:"allowUnverifiedSslCert"`
	// The Deployment option to be chosen. If empty, the default option is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.0/docs/resources/virtual_machine#deployment_option VirtualMachine#deployment_option}
	DeploymentOption *string `field:"optional" json:"deploymentOption" yaml:"deploymentOption"`
	// An optional disk provisioning.
	//
	// If set, all the disks in the deployed ovf will have the same specified disk type (e.g., thin provisioned).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.0/docs/resources/virtual_machine#disk_provisioning VirtualMachine#disk_provisioning}
	DiskProvisioning *string `field:"optional" json:"diskProvisioning" yaml:"diskProvisioning"`
	// Allow properties with ovf:userConfigurable=false to be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.0/docs/resources/virtual_machine#enable_hidden_properties VirtualMachine#enable_hidden_properties}
	EnableHiddenProperties interface{} `field:"optional" json:"enableHiddenProperties" yaml:"enableHiddenProperties"`
	// The IP allocation policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.0/docs/resources/virtual_machine#ip_allocation_policy VirtualMachine#ip_allocation_policy}
	IpAllocationPolicy *string `field:"optional" json:"ipAllocationPolicy" yaml:"ipAllocationPolicy"`
	// The IP protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.0/docs/resources/virtual_machine#ip_protocol VirtualMachine#ip_protocol}
	IpProtocol *string `field:"optional" json:"ipProtocol" yaml:"ipProtocol"`
	// The absolute path to the ovf/ova file in the local system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.0/docs/resources/virtual_machine#local_ovf_path VirtualMachine#local_ovf_path}
	LocalOvfPath *string `field:"optional" json:"localOvfPath" yaml:"localOvfPath"`
	// The mapping of name of network identifiers from the ovf descriptor to network UUID in the VI infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.0/docs/resources/virtual_machine#ovf_network_map VirtualMachine#ovf_network_map}
	OvfNetworkMap *map[string]*string `field:"optional" json:"ovfNetworkMap" yaml:"ovfNetworkMap"`
	// URL to the remote ovf/ova file to be deployed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.6.0/docs/resources/virtual_machine#remote_ovf_url VirtualMachine#remote_ovf_url}
	RemoteOvfUrl *string `field:"optional" json:"remoteOvfUrl" yaml:"remoteOvfUrl"`
}

