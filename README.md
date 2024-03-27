# djan-go-deployer

The `djan-go-deployer` package allows you to deploy any Django project using a simple YAML configuration file. This README provides an overview of the package and instructions on how to use it effectively.

## Installation

To install the `djan-go-deployer` package, you can use `go get`:

```bash
go get github.com/ahmedtouahria/djan-go-deployer
```

## Usage

The package provides functionality to deploy Django projects using a YAML configuration file. Here's how you can create a YAML file to specify your deployment configuration:

```yaml
# Nginx configuration
DB_NAME: Fatal
DB_USER: postgre
DB_PASSWORD: "1954"
DB_HOST: "127.0.0.1"
DB_PORT: test
DJANGO_SECRET: secret
ALLOWED_HOSTS: 127.0.0.1,localhost

SSL: "false"
SECURE: "false" # Be careful! This option will implement all security practices and settings in your web server configurations
NGINX_PORT: "80" # Nginx port
DJANGO_PORT: "8000" # Django process port
CHANNELS_PORT: "8001" # Django process port

DJANGO_DIRECTORY: /home/user/project # Django directory, must contain a requirements.txt file
INCLUDE_CHANNELS: "true" # Indicates if you're using channels in your Django project
CHANNELS_HOSTS: localhost
INCLUDE_CELERY: "false"
PROJECT_NAME: django
```

Ensure that you have the necessary permissions and dependencies installed to deploy the Django project successfully.

## Package Functions

### CreateNGINXConf

The `CreateNGINXConf` function generates an NGINX configuration based on the provided NGINXConfig struct.

```go
func CreateNGINXConf(config NGINXConfig) (string, error)
```

### CreateFile

The `CreateFile` function creates a file at a specified path with the given content.

```go
func CreateFile(filename string, content string) error
```

### RestartNginxServer

The `RestartNginxServer` function restarts the NGINX server.

```go
func RestartNginxServer() error
```

## Example

Here's an example of how you can use the `djan-go-deployer` package:

```go
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
	fmt.Println(data)
	if err != nil {
		panic(err)
	  }

	deployer.Installer() // Install all dependencies
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

```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
