package main

import (
	"fmt"

	"github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"
)

func main() {
	user := domain.UserFactory("User 1", "user@gmail.com", "password")
	fmt.Printf("User %s with %d points\n", user.GetName(), user.GetPoins())
}
