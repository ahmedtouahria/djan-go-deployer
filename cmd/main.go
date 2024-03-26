package main

import (
	"django_deployer/deployer"
	"django_deployer/server"
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

	data, err := server.ReadYamlFile(conf)
	if err != nil {
		panic(err)
	  }

	//deployer.Installer() // Install all dependencies
	// Create database
	mydatabase,found:=server.GetByKey(data,"DB_NAME")

	if !found{
		fmt.Println("Could not find database")
	}

	myuser,found:=server.GetByKey(data,"DB_USER")
	if !found{
		fmt.Println("Could not find DATABASE user")
	}

	mypassword,found:=server.GetByKey(data,"DB_PASSWORD")
	if !found{
		fmt.Println("Could not find DATABASE password")
	}

	database :=deployer.NewDatabaseBuilder().
	WithName(mydatabase).
	WithUsername(myuser).
	WithPassword(mypassword).
	WithHost("localhost").
	WithPort("5432")
	err=database.Build()

	if err != nil {
		panic(err)
	}
	// creating pm2 scripts 

	err = deployer.CreatePM2App(conf)
	if err != nil {
		panic(err)
	}
	// restarting pm2 processes
	err = deployer.RestartPm2Process()
	if err != nil {
		panic(err)
	}
	//creating nginx configurations files
	err,content:=deployer.CreateNGINXConf(data)
	if err != nil {
		panic(err)
	}
	projectName,found:=server.GetByKey(data,"PROJECT_NAME")
	if !found {
		fmt.Println("Could not find projectName")
		panic(found)

	}
	
	err=deployer.BuildNginxFile(content,projectName)
	if err != nil {
		panic(err)
	}

	// restart nginx server
	err = deployer.RerstartNginxServer()
	if err != nil {
		panic(err)
	}
	//SSL Certificate
	err = deployer.SetSSLCert()
	if err != nil {
		panic(err)
	}
}

