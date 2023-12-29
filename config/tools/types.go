package tools

// ResourceConfig contains configuration information for a Terraform resource.
type ResourceConfig struct {
	PackageName           string
	ShortGroup            string
	ResourceName          string
	TerraformResourceName string
	Kind                  string
	References            map[string]string
}
