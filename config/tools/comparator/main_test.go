package main

import (
	"testing"
)

// TestParseKindFromResourceName tests the parseKindFromResourceName function.
func TestParseKindFromResourceName(t *testing.T) {
	tests := []struct {
		resourceName string
		expected     string
	}{
		{"scaleway_object_bucket", "Bucket"},         // Normal case
		{"scaleway_flexible_ip", "Ip"},               // Multiple words
		{"simple_resource", "Resource"},              // Single word
		{"", ""},                                     // Empty string
		{"singleword", "Singleword"},                 // No underscores
		{"trailing_underscore_", ""},                 // Trailing underscore
		{"_leading_underscore", "Underscore"},        // Leading underscore
		{"middle__double__underscore", "Underscore"}, // Double underscores
		{"_underscore", "Underscore"},                // Starts with underscore
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
		{"Special Chars *&^%$#", "specialchars*&^%$#"},
		{"MixedCASE Category", "mixedcasecategory"},
		{"  Leading and Trailing  ", "leadingandtrailing"},
		{"Dotted.Category", "dotted.category"},
		{"Hyphenated-Category", "hyphenated-category"},
	}

	for _, tt := range tests {
		t.Run(tt.subCategory, func(t *testing.T) {
			actual := generatePackageName(tt.subCategory)
			if actual != tt.expected {
				t.Errorf("generatePackageName(%s) = %s; expected %s", tt.subCategory, actual, tt.expected)
			}
		})
	}
}
