package main

import "fmt"

const age = 30

// not_allowed := "Something"

func main() {

	// const name = "Golang"
	// name = "Can't change"

	fmt.Println(age)

	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)

}