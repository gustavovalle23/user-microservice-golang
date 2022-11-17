package main

import "fmt"

type User struct {
	Name      string
	BirthDate string
}

type UserRepo interface {
	Create() int16
}

func (u User) Create() int16 {
	fmt.Printf("creating user %s with birthDate: %s\n", u.Name, u.BirthDate)
	return 201
}

func main() {
	user := User{Name: "Gusta", BirthDate: "1999-01-01"}
	status := user.Create()
	fmt.Printf("Status: %d", status)
}
