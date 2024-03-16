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

	data, err := deployer.ReadYamlFile(conf)
	if err != nil {
		panic(err)
	  }
	deployer.Installer() // Install all dependencies

	//creating Database
	//...........
	// creating pm2 scripts 
	//...........
	//creating nginx configurations files
	//...........
	//SSL Certificate
	

	key:="DB_NAME"
	// Get the value for a specific key
	value, found := deployer.GetByKey(data,key)
	if found {
		fmt.Println("DB_NAME:", value)
	} else {
		fmt.Println("Key 'envirement.DB_NAME' not found")
	}

	fmt.Printf("Value for key %q: %v\n", key, value)
}

