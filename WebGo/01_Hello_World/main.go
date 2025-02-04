package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// HelloWorld is a simple HTTP handler function that writes a response to the client
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Bytes written: %d\n", n)
}

// HomePage is a simple HTTP handler function that writes a response to the client
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home Page...")
}

// AboutPage is a simple HTTP handler function that writes a response to the client
func AboutPage(w http.ResponseWriter, r *http.Request) {
	addtwo := addValues(2, 2)
	fmt.Fprintf(w, "This is the About Page... and 2 + 2 is: %d", addtwo)
}

// addValues is a simple function that adds two integers and returns the sum
func addValues(x, y int) int {
	sum := x + y
	return sum
}

// main is the entry point for the application
func main() {
	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/home", HomePage)
	http.HandleFunc("/about", AboutPage)

	fmt.Printf("Starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}