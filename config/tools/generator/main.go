package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/scaleway/provider-scaleway/config/tools"
)

const initialConfigTemplate = `package {{ .PackageName }}

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "{{ .ShortGroup }}"
`

const resourceConfigTemplate = `
// Configure adds configurations for {{ .ResourceName }} resource.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("{{ .TerraformResourceName }}", func(r *config.Resource) {
        r.ExternalName = config.IdentifierFromProvider
        r.ShortGroup = shortGroup
        r.Kind = "{{ .Kind }}"
        {{ range $key, $value := .References }}
        r.References["{{$key}}"] = config.Reference{
            Type: "{{$value}}",
        }
        {{ end }}
    })
}
`

const (
	ProviderGoFilePath     = "config/provider.go"
	ExternalNameGoFilePath = "config/external_name.go"
)

func generateConfigFile(resourceConfig tools.ResourceConfig) error {
	dirPath := filepath.Join("config", resourceConfig.PackageName)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return err
	}

	filePath := filepath.Join(dirPath, "config.go")

	// Check if config.go already exists
	if content, err := os.ReadFile(filePath); err == nil {
		contentStr := string(content)
		configurator := fmt.Sprintf("p.AddResourceConfigurator(\"%s\", func(r *config.Resource) {", resourceConfig.TerraformResourceName)

		// Check if the configurator already exists
		if !strings.Contains(contentStr, configurator) {
			tmpl, err := template.New("config").Parse(`
            p.AddResourceConfigurator("{{ .TerraformResourceName }}", func(r *config.Resource) {
                r.ExternalName = config.IdentifierFromProvider
                r.ShortGroup = "{{ .ShortGroup }}"
                r.Kind = "{{ .Kind }}"
            })
            `)
			if err != nil {
				return err
			}

			var tmplOutput bytes.Buffer
			if err := tmpl.Execute(&tmplOutput, resourceConfig); err != nil {
				return err
			}

			// Append the new configurator before the last closing brace '}'
			insertionPoint := strings.LastIndex(contentStr, "}")
			if insertionPoint == -1 {
				return errors.New("failed to find insertion point in config.go")
			}
			updatedContent := contentStr[:insertionPoint] + tmplOutput.String() + contentStr[insertionPoint:]
			return os.WriteFile(filePath, []byte(updatedContent), 0644)
		}
	} else {
		// Use initial template for new file
		tmpl, err := template.New("config").Parse(initialConfigTemplate + resourceConfigTemplate)
		if err != nil {
			return err
		}
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		return tmpl.Execute(file, resourceConfig)
	}
	return nil
}

func updateProviderGo(resourceConfig tools.ResourceConfig) error {
	content, err := os.ReadFile(ProviderGoFilePath)
	if err != nil {
		return err
	}

	contentStr := string(content)

	existingLine := fmt.Sprintf("\t%s.Configure,\n", resourceConfig.PackageName)
	if strings.Contains(contentStr, existingLine) {
		fmt.Printf("Configuration for %s already exists in provider.go\n", resourceConfig.PackageName)
		return nil // Skip insertion as it already exists
	}

	// Find the last occurrence of ".Configure,"
	insertionPoint := strings.LastIndex(contentStr, ".Configure,")
	if insertionPoint == -1 {
		return errors.New("failed to find insertion point in provider.go")
	}

	lineEnd := strings.Index(contentStr[insertionPoint:], "\n")
	if lineEnd == -1 {
		return errors.New("failed to find the end of the line for insertion point in provider.go")
	}

	// Calculate the actual insertion point in the content
	actualInsertionPoint := insertionPoint + lineEnd + 1
	newLine := existingLine

	// Insert the new line at the correct position
	updatedContent := contentStr[:actualInsertionPoint] + newLine + contentStr[actualInsertionPoint:]

	return os.WriteFile(ProviderGoFilePath, []byte(updatedContent), 0644)
}

func updateExternalNameGo(resourceConfig tools.ResourceConfig) error {
	fileContent, err := os.ReadFile(ExternalNameGoFilePath)
	if err != nil {
		return err
	}

	contentStr := string(fileContent)
	resourceName := fmt.Sprintf("\"%s\": config.NameAsIdentifier,", resourceConfig.TerraformResourceName)

	// Check if the resource already exists in the file
	if !strings.Contains(contentStr, resourceName) {
		// Find the insertion point, which is before the last entry in the map
		insertionPoint := strings.LastIndex(contentStr, "\t\"scaleway_")
		if insertionPoint == -1 {
			return errors.New("failed to find insertion point in external_name.go")
		}
		newLine := fmt.Sprintf("\t%s\n", resourceName)

		// Insert the new line at the correct position
		updatedContent := contentStr[:insertionPoint] + newLine + contentStr[insertionPoint:]

		// Write the updated content back to the file
		return os.WriteFile(ExternalNameGoFilePath, []byte(updatedContent), 0644)
	}

	return nil
}

func main() {
	jsonData, err := io.ReadAll(os.Stdin)

	if err != nil {
		fmt.Printf("Error reading JSON input: %v\n", err)
		return
	}

	var resourceConfigs []tools.ResourceConfig
	if err := json.Unmarshal(jsonData, &resourceConfigs); err != nil {
		fmt.Printf("Error parsing JSON input: %v\n", err)
		return
	}

	for _, config := range resourceConfigs {
		if err := generateConfigFile(config); err != nil {
			fmt.Printf("Error generating config file for %s: %v\n", config.ResourceName, err)
			continue
		}

		if err := updateProviderGo(config); err != nil {
			fmt.Printf("Error updating provider.go for %s: %v\n", config.ResourceName, err)
			continue
		}

		if err := updateExternalNameGo(config); err != nil {
			fmt.Printf("Error updating external_name.go for %s: %v\n", config.ResourceName, err)
			continue
		}
	}
}
