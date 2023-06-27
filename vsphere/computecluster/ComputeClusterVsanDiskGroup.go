package computecluster


type ComputeClusterVsanDiskGroup struct {
	// Cache disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/compute_cluster#cache ComputeCluster#cache}
	Cache *string `field:"optional" json:"cache" yaml:"cache"`
	// List of storage disks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vsphere/2.4.1/docs/resources/compute_cluster#storage ComputeCluster#storage}
	Storage *[]*string `field:"optional" json:"storage" yaml:"storage"`
}

