package main
import (
	"django_deployer/deployer"
	"flag"
	"fmt"
	"log"
)

func main() {
	var conf string

	// Define command-line flags
	flag.StringVar(&conf, "conf", "", "Set your YAML configuration file")
	flag.Parse()

	// Check if the configuration file is specified
	if conf == "" {
		log.Fatal("Configuration file must be specified using the -conf flag")
	}

	// Read the YAML file
	confYaml := deployer.ReadYamlFile(conf)
	fmt.Println(confYaml)

	// Access and print the Environment section
	env, ok := confYaml["envirement"].([]any)
	if !ok {
		log.Fatal("Invalid format for 'envirement' section")
	}
	fmt.Println("Environment:")
	for _, e := range env {
		for k, v := range e.(map[any]any) {
			fmt.Printf("- %s: %s\n", k, v)
		}
	}

	// Access and print the Nginx section
	nginx, ok := confYaml["nginx"].([]any)
	if !ok {
		log.Fatal("Invalid format for 'nginx' section")
	}
	fmt.Println("\nNginx:")
	for _, n := range nginx {
		for k, v := range n.(map[any]any) {
			fmt.Printf("- %s: %s\n", k, v)
		}
	}
}
