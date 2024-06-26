package deployer

import (
	"django_deployer/server"
	"fmt"
)

// DependenciesInstaller is a function that installs Linux os dependencies
func DependenciesInstaller() {
	// Update package information
	if err := server.RunCommand("sudo", "apt", "update"); err != nil {
		fmt.Println("Error updating packages:", err)
		return
	}

	// Install required packages
	if err := server.RunCommand("sudo", "apt", "install", "python3-venv", "python3-dev", "libpq-dev", "postgresql", "postgresql-contrib", "nginx", "curl","nodejs","npm","supervisor"); err != nil {
		fmt.Println("Error installing packages:", err)
		return
	}


	fmt.Println("Commands executed successfully.")
}
// ProjectRequirementInstaller is used to install requirements.txt
func ProjectRequirementInstaller() {
	// Install required packages requirements.txt
	if err := server.RunCommand("pip", "install", "-r", "requirements.txt"); err != nil {
		fmt.Println("Error updating packages:", err)
		return
	}
	fmt.Println("Commands executed successfully.")
}

// Pm2ProccessManagerInstaller is used to install pm2 process manager
func Pm2ProccessManagerInstaller() {
	// Install pm2 using npm
	if err := server.RunCommand("npm", "install", "pm2", "-g"); err != nil {
		fmt.Println("Error installing pm2:", err)
		return
	}
	fmt.Println("Commands executed successfully.")
}



// Installer is used to install all dependencies and requirements
func Installer(){
DependenciesInstaller() // Install dependencies
ProjectRequirementInstaller() // Install requirements.txt
Pm2ProccessManagerInstaller() // Install PM2 
}