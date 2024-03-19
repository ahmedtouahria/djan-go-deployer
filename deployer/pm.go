//PM: Process manager

package deployer

import (
	"django_deployer/server"
	"fmt"
	"os"
)

func CreateBashScript(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0755)
	if err != nil {
		return err
	}
	return nil
}

func CreatePM2App(conf string) error {
	data, err := server.ReadYamlFile(conf)
	if err != nil {
		panic(err)
	  }
	port,ok:= server.GetByKey(data,"DJANGO_PORT")
	if !ok {
		fmt.Println("Error:", err)
		return fmt.Errorf("Error:","DJANGO_PORT is not set")
	}

	directory,ok:= server.GetByKey(data,"DJANGO_DIRECTORY")
	if !ok {
		fmt.Println("Error:", err)
		return fmt.Errorf("Error:","DJANGO_DIRECTORY is not set")
	}
	fileName,ok:= server.GetByKey(data,"PROJECT_NAME")
	if !ok {
		fmt.Println("Error:", err)
		return fmt.Errorf("Error:","PROJECT_NAME is not set")

	}
	// Bash script content
	scriptContent := `#!/bin/bash
		cd ` + directory + `;
		source env/bin/activate;
		python manage.py runserver 0.0.0.0:` + port

	// Create the Bash script file
	err = CreateBashScript(fileName, scriptContent)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	fmt.Println("PM2 app script file created successfully:", fileName)
	return err
}

func CreatePM2ChannelsApp(fileName string, directory string) {
	// Bash script content
	scriptContent := `#!/bin/bash
		cd ` + directory + `;
		source env/bin/activate;
		daphne -b 0.0.0.0 -p 8001 core.asgi:application
				`
	// Create the Bash script file
	err := CreateBashScript(fileName, scriptContent)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("PM2 app script file for channels created successfully:", fileName)
}

