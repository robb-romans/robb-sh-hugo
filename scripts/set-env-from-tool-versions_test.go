//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestParseToolVersions(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]string
	}{
		{
			name: "standard format",
			input: `node 22.18.0
go 1.24.6
hugo 0.148.2`,
			expected: map[string]string{
				"NODE_VERSION": "22.18.0",
				"GO_VERSION":   "1.24.6",
				"HUGO_VERSION": "0.148.2",
			},
		},
		{
			name: "with comments and empty lines",
			input: `# This is a comment
node 22.18.0

go 1.24.6
# Another comment
hugo 0.148.2`,
			expected: map[string]string{
				"NODE_VERSION": "22.18.0",
				"GO_VERSION":   "1.24.6",
				"HUGO_VERSION": "0.148.2",
			},
		},
		{
			name: "extra whitespace",
			input: `  node   22.18.0  
    go	1.24.6
hugo	0.148.2   `,
			expected: map[string]string{
				"NODE_VERSION": "22.18.0",
				"GO_VERSION":   "1.24.6",
				"HUGO_VERSION": "0.148.2",
			},
		},
	}

	toolKeys := []string{"NODE_VERSION", "GO_VERSION", "HUGO_VERSION"}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseToolVersions(tt.input, toolKeys)

			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d tools, got %d", len(tt.expected), len(result))
			}

			for k, v := range tt.expected {
				if result[k] != v {
					t.Errorf("Expected %s = %s, got %s", k, v, result[k])
				}
			}
		})
	}
}

func TestSyncTomlContent(t *testing.T) {
	toolKeys := []string{"NODE_VERSION", "GO_VERSION", "HUGO_VERSION"}
	tools := map[string]string{
		"NODE_VERSION": "22.18.0",
		"GO_VERSION":   "1.24.6",
		"HUGO_VERSION": "0.148.2",
	}

	tests := []struct {
		name         string
		inputToml    string
		expectedToml string
		expectWarn   bool
	}{
		{
			name: "basic sync with existing section",
			inputToml: `[build]
command = "npm install && npm run build"

[build.environment]
MISE_TRUSTED = "1"

[context.deploy-preview]
command = "npm install && npm run build:preview"`,
			expectedToml: `[build]
command = "npm install && npm run build"

[build.environment]
MISE_TRUSTED = "1"

NODE_VERSION = "22.18.0"
GO_VERSION = "1.24.6"
HUGO_VERSION = "0.148.2"
[context.deploy-preview]
command = "npm install && npm run build:preview"`,
		},
		{
			name: "remove duplicates",
			inputToml: `[build.environment]
NODE_VERSION = "20.0.0"
GO_VERSION = "1.20.0"
HUGO_VERSION = "0.100.0"
MISE_TRUSTED = "1"
NODE_VERSION = "21.0.0"
GO_VERSION = "1.21.0"
HUGO_VERSION = "0.110.0"

[context.deploy-preview]
command = "test"`,
			expectedToml: `[build.environment]
NODE_VERSION = "22.18.0"
GO_VERSION = "1.24.6"
HUGO_VERSION = "0.148.2"
MISE_TRUSTED = "1"

[context.deploy-preview]
command = "test"`,
		},
		{
			name: "missing build.environment section",
			inputToml: `[build]
command = "npm install && npm run build"

[context.deploy-preview]
command = "test"`,
			expectedToml: `[build]
command = "npm install && npm run build"

[context.deploy-preview]
command = "test"

[build.environment]
NODE_VERSION = "22.18.0"
GO_VERSION = "1.24.6"
HUGO_VERSION = "0.148.2"`,
		},
		{
			name: "build.environment at end of file",
			inputToml: `[build]
command = "npm install && npm run build"

[build.environment]
MISE_TRUSTED = "1"`,
			expectedToml: `[build]
command = "npm install && npm run build"

[build.environment]
MISE_TRUSTED = "1"
NODE_VERSION = "22.18.0"
GO_VERSION = "1.24.6"
HUGO_VERSION = "0.148.2"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, warnings := syncTomlContent(tt.inputToml, tools, toolKeys)

			// Normalize line endings and trim trailing whitespace
			expected := strings.TrimSpace(strings.ReplaceAll(tt.expectedToml, "\r\n", "\n"))
			actual := strings.TrimSpace(strings.ReplaceAll(result, "\r\n", "\n"))

			if expected != actual {
				t.Errorf("Expected:\n%q\n\nActual:\n%q", expected, actual)
			}

			if tt.expectWarn && len(warnings) == 0 {
				t.Error("Expected warnings but got none")
			}
			if !tt.expectWarn && len(warnings) > 0 {
				t.Errorf("Unexpected warnings: %v", warnings)
			}
		})
	}
}

func TestSyncToolVersionsIntegration(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "netlify-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Save current working directory
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(oldWd)

	// Change to temp directory
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatal(err)
	}

	// Write test files
	toolVersions := `node 22.18.0
go 1.24.6
hugo 0.148.2`

	inputToml := `[build]
command = "npm install && npm run build"

[build.environment]
MISE_TRUSTED = "1"

[context.deploy-preview]
command = "test"`

	expectedToml := `[build]
command = "npm install && npm run build"

[build.environment]
MISE_TRUSTED = "1"

NODE_VERSION = "22.18.0"
GO_VERSION = "1.24.6"
HUGO_VERSION = "0.148.2"
[context.deploy-preview]
command = "test"`

	err = os.WriteFile(".tool-versions", []byte(toolVersions), 0644)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile("netlify.toml", []byte(inputToml), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Capture stderr
	var stderr bytes.Buffer

	// Run sync
	err = syncToolVersions(".tool-versions", "netlify.toml", &stderr)
	if err != nil {
		t.Fatal(err)
	}

	// Read result
	result, err := os.ReadFile("netlify.toml")
	if err != nil {
		t.Fatal(err)
	}

	// Compare - normalize and trim whitespace
	expected := strings.TrimSpace(strings.ReplaceAll(expectedToml, "\r\n", "\n"))
	actual := strings.TrimSpace(strings.ReplaceAll(string(result), "\r\n", "\n"))

	if expected != actual {
		t.Errorf("Expected:\n%q\n\nActual:\n%q", expected, actual)
	}

	// Test idempotency - run again
	err = syncToolVersions(".tool-versions", "netlify.toml", &stderr)
	if err != nil {
		t.Fatal(err)
	}

	result2, err := os.ReadFile("netlify.toml")
	if err != nil {
		t.Fatal(err)
	}

	if string(result) != string(result2) {
		t.Error("Second run produced different result - not idempotent")
	}
}
