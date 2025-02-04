package auth

func LoginWithCredentials(username, password string) bool {
	return username == "admin" && password == "admin"
}