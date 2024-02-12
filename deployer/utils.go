package deployer

import (
	"fmt"
	"os"
	"os/exec"
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
func ReadYamlFile(configFile string) map[string]any{
	obj := make(map[string]any)
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		_, err = fmt.Printf("yamlFile.Get err #%v ", err)
		if err != nil {
			panic(err)
		}
	}
	err = yaml.Unmarshal(yamlFile, obj)
	if err != nil {
		_, err = fmt.Printf("Unmarshal: %v", err)
		if err != nil {
			panic(err)
		}
	}

	return obj
}
