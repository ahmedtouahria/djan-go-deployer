package deployer

import (
	//"fmt"
	"log"
)

// Config represents the structure of the YAML configuration file
type Config struct {
	Envirement map[string]string `yaml:"envirement"`
	Nginx       map[string]string `yaml:"nginx"`
	// Add more fields as needed based on your actual YAML structure
}

// NewConfig creates a new Config instance with default values
func NewConfig() *Config {
	return &Config{
		Envirement: make(map[string]string),
		Nginx:       make(map[string]string),
		// Initialize other fields with default values
	}
}

// LoadConfig loads the configuration from a YAML file
func LoadConfig(filePath string) (*Config, error) {
	// Implement your YAML file reading logic here and unmarshal it into the Config struct
	// Use a YAML parsing library like gopkg.in/yaml.v2

	// For example:
	// yamlFile, err := ioutil.ReadFile(filePath)
	// if err != nil {
	// 	return nil, err
	// }
	// var config Config
	// err = yaml.Unmarshal(yamlFile, &config)
	// if err != nil {
	// 	return nil, err
	// }
	// return &config, nil

	// For the purpose of this example, return a new Config with default values
	return NewConfig(), nil
}

// HandleError handles the error based on the configuration
func HandleError(config *Config, err error) {
	if err != nil {
		log.Println("Error occurred:", err)
		// Implement your error handling logic based on the configuration here
		// For example, you can use the configuration values to log, notify, or perform other actions
	}
}

