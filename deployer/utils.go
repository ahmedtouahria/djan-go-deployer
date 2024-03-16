package deployer

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

// runCommand is a function that runs an os command
func RunCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
// ReadYamlFile is a function that reads a YAML file
func ReadYamlFile(configFile string) (map[string]any, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
	  return nil, fmt.Errorf("ReadFile: %w", err)
	}
	
	obj := make(map[string]any)
	err = yaml.Unmarshal(data, obj)
	if err != nil {
	  return nil, fmt.Errorf("Unmarshal: %w", err)
	}
	
	return obj, nil
  }


// GetByKey retrieves a specific key from the parsed YAML data
func GetByKey(data map[string]any, key string) (interface{}, bool) {
	// Handle nested keys with recursion
	parts := strings.Split(key, ".")
	value, ok := data[parts[0]]
	if !ok {
	  return nil, false
	}
	
	if len(parts) == 1 {
	  return value, true
	}
	
	// Recursively navigate through nested maps
	m, ok := value.(map[string]any)
	if !ok {
	  return nil, false // Key points to a non-map value
	}
	
	return GetByKey(m, strings.Join(parts[1:], "."))
  }
  