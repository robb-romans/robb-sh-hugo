//go:build ignore
// +build ignore

// Sync-tool-versions is a Go program to sync .tool-versions with the [build.environment] section in netlify.toml.
// Place in your project root and run manually or from a pre-commit hook.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// parseToolVersions extracts tool versions from .tool-versions content
func parseToolVersions(content string, toolKeys []string) map[string]string {
	tools := make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(content))

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			key := strings.ToUpper(parts[0]) + "_VERSION"
			tools[key] = parts[1]
		}
	}
	return tools
}

// syncTomlContent processes netlify.toml content and syncs version variables
func syncTomlContent(tomlContent string, tools map[string]string, toolKeys []string) (string, []string) {
	var out bytes.Buffer
	var warnings []string
	s := bufio.NewScanner(strings.NewReader(tomlContent))
	inEnv := false
	envBlockCount := 0
	versionWritten := false

	// Regex to match any of our version variables
	reVersionVar := regexp.MustCompile(`^\s*(NODE_VERSION|GO_VERSION|HUGO_VERSION)\s*=`)

	for s.Scan() {
		line := s.Text()
		trimmedLine := strings.TrimSpace(line)

		if strings.HasPrefix(line, "[build.environment]") {
			inEnv = true
			envBlockCount++
			versionWritten = false
			out.WriteString(line + "\n")
			continue
		}

		if inEnv && strings.HasPrefix(trimmedLine, "[") {
			// Exiting [build.environment] section
			// Write versions if we haven't already
			if !versionWritten {
				for _, k := range toolKeys {
					if v, ok := tools[k]; ok {
						out.WriteString(fmt.Sprintf("%s = \"%s\"\n", k, v))
					}
				}
			}
			inEnv = false
			out.WriteString(line + "\n")
			continue
		}

		if inEnv {
			// Skip all existing version variable lines
			if reVersionVar.MatchString(line) {
				// If this is the first version variable we encounter, write all our versions
				if !versionWritten {
					for _, k := range toolKeys {
						if v, ok := tools[k]; ok {
							out.WriteString(fmt.Sprintf("%s = \"%s\"\n", k, v))
						}
					}
					versionWritten = true
				}
				continue // Skip this line
			}
		}

		out.WriteString(line + "\n")
	}
	// If still in env block at EOF, write versions if not written
	if inEnv && !versionWritten {
		for _, k := range toolKeys {
			if v, ok := tools[k]; ok {
				out.WriteString(fmt.Sprintf("%s = \"%s\"\n", k, v))
			}
		}
	}

	// Add block if it didn't exist
	if envBlockCount == 0 {
		out.WriteString("\n[build.environment]\n")
		for _, k := range toolKeys {
			if v, ok := tools[k]; ok {
				out.WriteString(fmt.Sprintf("%s = \"%s\"\n", k, v))
			}
		}
	}
	if envBlockCount > 1 {
		warnings = append(warnings, "Multiple [build.environment] blocks found")
	}

	return out.String(), warnings
}

// syncToolVersions performs the complete sync operation
func syncToolVersions(toolVersionsPath, netlifyTomlPath string, stderr io.Writer) error {
	toolKeys := []string{"NODE_VERSION", "GO_VERSION", "HUGO_VERSION"}

	// Read .tool-versions
	toolContent, err := os.ReadFile(toolVersionsPath)
	if err != nil {
		return fmt.Errorf("reading %s: %w", toolVersionsPath, err)
	}

	tools := parseToolVersions(string(toolContent), toolKeys)

	// Warn if any tool version is missing
	for _, k := range toolKeys {
		if _, ok := tools[k]; !ok {
			fmt.Fprintf(stderr, "Warning: %s not found in %s\n", k, toolVersionsPath)
		}
	}

	// Read netlify.toml
	tomlContent, err := os.ReadFile(netlifyTomlPath)
	if err != nil {
		return fmt.Errorf("reading %s: %w", netlifyTomlPath, err)
	}

	// Process content
	result, warnings := syncTomlContent(string(tomlContent), tools, toolKeys)

	// Print warnings
	for _, warning := range warnings {
		fmt.Fprintf(stderr, "Warning: %s\n", warning)
	}

	// Write result atomically
	tmpPath := netlifyTomlPath + ".tmp"
	if err := os.WriteFile(tmpPath, []byte(result), 0644); err != nil {
		return fmt.Errorf("writing temp file: %w", err)
	}
	if err := os.Rename(tmpPath, netlifyTomlPath); err != nil {
		return fmt.Errorf("renaming temp file: %w", err)
	}

	return nil
}

func main() {
	err := syncToolVersions(".tool-versions", "netlify.toml", os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✔ Synced .tool-versions to netlify.toml")
}
