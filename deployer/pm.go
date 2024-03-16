//PM: Process manager

package deployer

import (
	"fmt"
	"os"
	"strconv"
)

func CreateBashScript(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0755)
	if err != nil {
		return err
	}
	return nil
}

func CreatePM2App(fileName string, directory string, port int) {
	portStr := strconv.Itoa(port)

	// Bash script content
	scriptContent := `#!/bin/bash
		cd ` + directory + `;
		source env/bin/activate;
		python manage.py runserver 0.0.0.0:` + portStr

	// Create the Bash script file
	err := CreateBashScript(fileName, scriptContent)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("PM2 app script file created successfully:", fileName)
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

