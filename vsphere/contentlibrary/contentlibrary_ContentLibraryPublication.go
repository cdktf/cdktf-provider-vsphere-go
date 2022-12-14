package contentlibrary


type ContentLibraryPublication struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/content_library#authentication_method ContentLibrary#authentication_method}.
	AuthenticationMethod *string `field:"optional" json:"authenticationMethod" yaml:"authenticationMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/content_library#password ContentLibrary#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/content_library#published ContentLibrary#published}.
	Published interface{} `field:"optional" json:"published" yaml:"published"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vsphere/r/content_library#username ContentLibrary#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

