package distributedportgroup


type DistributedPortGroupVlanRange struct {
	// The minimum VLAN to use in the range.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/distributed_port_group#max_vlan DistributedPortGroup#max_vlan}
	MaxVlan *float64 `field:"required" json:"maxVlan" yaml:"maxVlan"`
	// The minimum VLAN to use in the range.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/distributed_port_group#min_vlan DistributedPortGroup#min_vlan}
	MinVlan *float64 `field:"required" json:"minVlan" yaml:"minVlan"`
}

