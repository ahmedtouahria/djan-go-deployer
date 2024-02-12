package main

import (
	"flag"
	"fmt"
	//"os"
)

func main() {
	var name string
	var age int
	var greet bool

	// Define command-line flags
	flag.StringVar(&name, "name", "Guest", "Specify a name")
	flag.IntVar(&age, "age", 0, "Specify an age")
	flag.BoolVar(&greet, "greet", false, "Enable greeting message")

	// Parse command-line arguments
	flag.Parse()

	// Perform actions based on flags
	if greet {
		_, err := fmt.Printf("Hello, %s!\n", name)
		if err != nil {
			
			
		}
		if age > 0 {
			_, err := fmt.Printf("You are %d years old.\n", age)
			if err != nil {
				panic(err)
			}
		}
	} else {
		fmt.Println("No greeting enabled.")
	}

	// Additional logic can be added based on the flag values.
}
