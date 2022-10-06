package distributedvirtualswitch


type DistributedVirtualSwitchVlanRange struct {
	// The minimum VLAN to use in the range.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/distributed_virtual_switch#max_vlan DistributedVirtualSwitch#max_vlan}
	MaxVlan *float64 `field:"required" json:"maxVlan" yaml:"maxVlan"`
	// The minimum VLAN to use in the range.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/distributed_virtual_switch#min_vlan DistributedVirtualSwitch#min_vlan}
	MinVlan *float64 `field:"required" json:"minVlan" yaml:"minVlan"`
}

