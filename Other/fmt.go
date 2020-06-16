package main

import (
	"fmt"
	"os"
)

// Intialize Constants
const (
	text = "Hello"
	number = 9
)

func main() {
	s := "World" // Short hand
	var blank string // Empty 
	var word, numeric = ".", 99 // Multiply
	switch os.Args[1] {
	case "Print":
		// Prints using the defualt formats
		fmt.Print(text, s, word, number, blank, numeric,"\nNew Line.")
	case "Println":
		// Prints using default formats, spaces between variables
		fmt.Println(text, s, word, number, blank, numeric,"\nNew Line.")
	case "Printf":
		// Prints using format supplied in the string
		fmt.Printf("%s %s %s %d %s %d %s", text, s, word, number, blank, numeric,"\nNew Line.")
	case "Fprintf":

		// Create a text file
		f, err := os.Create("output.txt")
		if err != nil {
			fmt.Println(err)
					f.Close()
			return
		}
		// Print to a file using format supplied in the string
		fmt.Fprintf(f,"%s %s %s %d %s %d %s", text, s, word, number, blank, numeric,"\nNew Line.")
	
	}
}
