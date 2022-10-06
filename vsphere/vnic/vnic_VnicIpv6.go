package vnic


type VnicIpv6 struct {
	// List of IPv6 addresses.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/vnic#addresses Vnic#addresses}
	Addresses *[]*string `field:"optional" json:"addresses" yaml:"addresses"`
	// Use IPv6 Autoconfiguration (RFC2462).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/vnic#autoconfig Vnic#autoconfig}
	Autoconfig interface{} `field:"optional" json:"autoconfig" yaml:"autoconfig"`
	// Use DHCP to configure the interface's IPv4 stack.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/vnic#dhcp Vnic#dhcp}
	Dhcp interface{} `field:"optional" json:"dhcp" yaml:"dhcp"`
	// IP address of the default gateway, if DHCP or autoconfig is not set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/vnic#gw Vnic#gw}
	Gw *string `field:"optional" json:"gw" yaml:"gw"`
}

