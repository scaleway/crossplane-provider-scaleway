package main

import (
	"go/token"
	"testing"
)

// TestParseKindFromResourceName tests the parseKindFromResourceName function.
func TestParseKindFromResourceName(t *testing.T) {
	tests := []struct {
		resourceName string
		expected     string
	}{
		{"scaleway_object_bucket", "Bucket"},
		{"scaleway_flexible_ip", "Ip"},
		{"simple_resource", "Resource"},
		{"", ""},
		{"singleword", "Singleword"},
		{"trailing_underscore_", ""},
		{"_leading_underscore", "Underscore"},
		{"middle__double__underscore", "Underscore"},
		{"_underscore", "Underscore"},
	}

	for _, tt := range tests {
		t.Run(tt.resourceName, func(t *testing.T) {
			actual := parseKindFromResourceName(tt.resourceName)
			if actual != tt.expected {
				t.Errorf("parseKindFromResourceName(%s) = %s; expected %s", tt.resourceName, actual, tt.expected)
			}
		})
	}
}

// TestGeneratePackageName tests the generatePackageName function.
func TestGeneratePackageName(t *testing.T) {
	tests := []struct {
		subCategory string
		expected    string
	}{
		{"Apple Silicon", "applesilicon"},
		{"Virtual Machines", "virtualmachines"},
		{" ", ""},
		{"", ""},
		{"Special Chars *&^%$#", "specialchars"},
		{"*&^%$#", ""},
		{"Database V2", "databasev2"},
		{"MixedCASE Category", "mixedcasecategory"},
		{"  Leading and Trailing  ", "leadingandtrailing"},
		{"Dotted.Category", "dottedcategory"},
		{"Hyphenated-Category", "hyphenatedcategory"},
	}

	for _, tt := range tests {
		t.Run(tt.subCategory, func(t *testing.T) {
			actual := generatePackageName(tt.subCategory)
			if actual != tt.expected {
				t.Errorf("generatePackageName(%s) = %s; expected %s", tt.subCategory, actual, tt.expected)
			}
			if actual != "" && !token.IsIdentifier(actual) {
				t.Errorf("generatePackageName(%s) returned an invalid Go package identifier %q", tt.subCategory, actual)
			}
		})
	}
}
