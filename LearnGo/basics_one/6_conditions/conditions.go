package main

import "fmt"

func main() {

	age := 19

	if age >= 18 {
		fmt.Println("Vote....")
	} else {
		fmt.Println("Shorten", 18 - age, "Years")
	}

	var role = "Admin"
	var hasPermissions = true

	if role == "Admin" && hasPermissions {
		fmt.Println("Yep")
	}

	if oneline := 11; oneline < 10 {
		fmt.Println("Yay!!", oneline)
	} else if oneline > 0 {
		fmt.Println("Surprise")
	}

}