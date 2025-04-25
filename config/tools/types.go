package tools

// ResourceConfig contains configuration information for a Terraform resource.
type ResourceConfig struct {
	PackageName           string            `json:"package_name"`
	ShortGroup            string            `json:"short_group"`
	ResourceName          string            `json:"resource_name"`
	TerraformResourceName string            `json:"terraform_resource_name"`
	Kind                  string            `json:"kind"`
	References            map[string]string `json:"references"`
}
