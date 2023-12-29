package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/scaleway/provider-scaleway/config/tools"
	"github.com/upbound/upjet/pkg/registry"

	"gopkg.in/yaml.v3"
)

// ProviderMetadata represents the structure of provider-metadata.yaml
type ProviderMetadata struct {
	Resources map[string]*registry.Resource `yaml:"resources"`
}

func main() {
	currentProviderMetadata := os.Getenv("CURRENT_METADATA")
	newProviderMetadata := os.Getenv("NEW_METADATA")

	currentResources, err := parseProviderMetadata(currentProviderMetadata)
	if err != nil {
		fmt.Printf("Error parsing current provider metadata: %v\n", err)
		return
	}

	newResources, err := parseProviderMetadata(newProviderMetadata)
	if err != nil {
		fmt.Printf("Error parsing new provider metadata: %v\n", err)
		return
	}

	addedResources := findNewResources(currentResources, newResources)
	fmt.Println("New resources found:")
	resourceConfigs := make([]tools.ResourceConfig, 0, len(addedResources))

	for _, resource := range addedResources {
		references := make(map[string]string)

		for _, example := range resource.Examples {
			for refKey, refValue := range example.References {
				resourceType := strings.Split(refValue, ".")[0]
				references[refKey] = resourceType
			}
		}

		config := tools.ResourceConfig{
			PackageName:           strings.ToLower(strings.ReplaceAll(resource.SubCategory, " ", "")), // Convert "Apple Silicon" to "applesilicon"
			ShortGroup:            strings.ToLower(resource.SubCategory),
			ResourceName:          resource.Title,
			TerraformResourceName: resource.Name,
			Kind:                  resource.Title,
			References:            references,
		}
		resourceConfigs = append(resourceConfigs, config)
		fmt.Println(resource.Name)
	}

	jsonData, err := json.Marshal(resourceConfigs)
	if err != nil {
		fmt.Printf("Error marshaling resource configuration: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}

func parseProviderMetadata(metadata string) (map[string]*registry.Resource, error) {
	var providerMetadata ProviderMetadata
	err := yaml.Unmarshal([]byte(metadata), &providerMetadata)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal provider metadata")
	}
	return providerMetadata.Resources, nil
}

func findNewResources(current, new map[string]*registry.Resource) []*registry.Resource {
	var addedResources []*registry.Resource
	for resourceName, resource := range new {
		if _, exists := current[resourceName]; !exists {
			addedResources = append(addedResources, resource)
		}
	}
	return addedResources
}
