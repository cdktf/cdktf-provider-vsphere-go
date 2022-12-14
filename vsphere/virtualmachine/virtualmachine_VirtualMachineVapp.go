package virtualmachine


type VirtualMachineVapp struct {
	// A map of customizable vApp properties and their values.
	//
	// Allows customization of VMs cloned from OVF templates which have customizable vApp properties.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/virtual_machine#properties VirtualMachine#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

