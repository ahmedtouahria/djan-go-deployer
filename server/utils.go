package server

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)



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
		e := fmt.Errorf("Unmarshal: %w", err)
		return nil, e
	}
	
	return obj, nil
  }


// GetByKey retrieves a specific key from the parsed YAML data

func GetByKey(data map[string]interface{}, key string) (string, bool) {
    // Handle nested keys with recursion
    parts := strings.Split(key, ".")
    value, ok := data[parts[0]]
    if !ok {
        return "", false
    }
    
    if len(parts) == 1 {
        // Convert the value to string
        stringValue, ok := value.(string)
        if !ok {
            // If the value is not a string, return an empty string
            return "", false
        }
        return stringValue, true
    }
    
    // Recursively navigate through nested maps
    m, ok := value.(map[string]interface{})
    if !ok {
        return "", false // Key points to a non-map value
    }
    
    return GetByKey(m, strings.Join(parts[1:], "."))
}


func CreateFile(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0755)
	if err != nil {
		return err
	}
	return nil
}