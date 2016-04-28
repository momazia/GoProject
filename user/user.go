package user

import ()

type User struct {
	Id string
}

// Creates a default user which is not loggedIn
func CreateUser() User {
	return User{}
}
