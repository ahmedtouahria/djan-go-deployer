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
	flag.StringVar(&conf, "conf", "deployer.yaml", "Set your YAML configuration file")
	flag.Parse()

	// Check if the configuration file is specified
	if conf == ""{
		log.Fatal("Configuration file must be specified")
	}

	// Install dependencies
	// deployer.DependenciesInstaller()

	// Install PM2 process manager
	//deployer.Pm2ProccessManagerInstaller()
	confYaml := deployer.ReadYamlFile("/home/ahmed/Desktop/projects/GoLang/django_deployer/example.yml")
	fmt.Println(confYaml)
	_, err := fmt.Printf("Hello, %s!\n", conf)
	if err != nil {
		panic(err)
	}

	// Additional logic can be added based on the flag values.
}
