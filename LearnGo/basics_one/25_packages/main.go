package main

import (
	"fmt"

	"github.com/ashparshp/learngo/25_packages/auth"
	"github.com/ashparshp/learngo/25_packages/user"
	"github.com/fatih/color"
)

func main() {

	authData := auth.LoginWithCredentials("admin", "notadmin")
	fmt.Println("Welcome to the main function", authData)

	sessionData := auth.GetSession()
	fmt.Println("Session data is", sessionData)

	user := user.User{
		Username: "admin",
		Password: "admin",
		Email:    "user@user.com",
	}

	fmt.Println("User data is", user)

	color.Red("This is a red text")

}