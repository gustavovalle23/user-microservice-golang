package main

import (
	"fmt"

	"github.com/gustavovalle23/user-microservice-golang/entities"
)

func main() {
	user := entities.UserFactory("User 1", "1999-01-01", 0)
	fmt.Printf("User %s with %d points\n", user.GetName(), user.GetPoins())
}
