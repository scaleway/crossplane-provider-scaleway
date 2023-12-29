package tools

type ResourceConfig struct {
	PackageName           string
	ShortGroup            string
	ResourceName          string
	TerraformResourceName string
	Kind                  string
	References            map[string]string
}
