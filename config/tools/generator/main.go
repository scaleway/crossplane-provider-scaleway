package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/scaleway/provider-scaleway/config/tools"
)

// ProviderGoFilePath is the path to the provider.go file where provider configurations are set.
const ProviderGoFilePath = "config/provider.go"

// ExternalNameGoFilePath is the path to the external_name.go file where external name configurations are defined.
const ExternalNameGoFilePath = "config/external_name.go"

func parseTemplates() (*template.Template, error) {
	templatePath := "config/config_templates.tmpl"
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}
	return tmpl, err
}

func generateConfigFile(resourceConfig tools.ResourceConfig) error {
	dirPath := filepath.Join("config", resourceConfig.PackageName)
	if err := os.MkdirAll(dirPath, 0750); err != nil {
		return err
	}

	filePath := filepath.Join(dirPath, "config.go")
	cleanFilePath := filepath.Clean(filePath)
	// Ensure the filePath starts with the expected directory
	if !strings.HasPrefix(cleanFilePath, "config/") {
		return fmt.Errorf("invalid file path: %s", cleanFilePath)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return createNewConfigFile(filePath, resourceConfig)
	}
	return updateExistingConfigFile(content, filePath, resourceConfig)
}

func createNewConfigFile(filePath string, resourceConfig tools.ResourceConfig) error {
	tmpl, err := parseTemplates()
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	templates := []string{"initialConfigTemplate", "resourceConfigTemplate"}
	for _, tmplName := range templates {
		if err = tmpl.ExecuteTemplate(&buf, tmplName, resourceConfig); err != nil {
			return err
		}
	}

	formattedCode, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("failed to format generated code: %w", err)
	}

	cleanFilePath := filepath.Clean(filePath)
	if !strings.HasPrefix(cleanFilePath, "config/") {
		return fmt.Errorf("invalid file path: %s", cleanFilePath)
	}

	err = os.WriteFile(filePath, formattedCode, 0644)
	if err != nil {
		return err
	}

	return nil
}

func updateExistingConfigFile(content []byte, filePath string, resourceConfig tools.ResourceConfig) error {
	contentStr := string(content)
	configurator := fmt.Sprintf("p.AddResourceConfigurator(\"%s\", func(r *config.Resource) {", resourceConfig.TerraformResourceName)

	if strings.Contains(contentStr, configurator) {
		return nil // Configurator already exists
	}

	tmpl, err := parseTemplates()
	if err != nil {
		return err
	}

	var tmplOutput bytes.Buffer
	if err := tmpl.ExecuteTemplate(&tmplOutput, "updateExistingConfigTemplate", resourceConfig); err != nil {
		return err
	}

	insertionPoint := strings.LastIndex(contentStr, "}")
	if insertionPoint == -1 {
		return errors.New("failed to find insertion point in config.go")
	}

	updatedContent := contentStr[:insertionPoint] + tmplOutput.String() + contentStr[insertionPoint:]

	formattedCode, err := format.Source([]byte(updatedContent))
	if err != nil {
		return fmt.Errorf("failed to format updated code: %w", err)
	}

	return os.WriteFile(filePath, formattedCode, 0600)
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

	formattedCode, err := format.Source([]byte(updatedContent))
	if err != nil {
		return fmt.Errorf("failed to format provider.go: %w", err)
	}

	return os.WriteFile(ProviderGoFilePath, formattedCode, 0600)
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

		formattedCode, err := format.Source([]byte(updatedContent))
		if err != nil {
			return fmt.Errorf("failed to format external_name.go: %w", err)
		}

		// Write the updated content back to the file
		return os.WriteFile(ExternalNameGoFilePath, formattedCode, 0600)
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
