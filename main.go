package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"
)

func main() {
	address := domain.Address{
		Street: "Random street",
		City:   "Random City",
		State:  "Random State",
	}
	birthDate := domain.NewDate(1990, time.January, 1)
	user, err := domain.NewUser("Test User", "password", "exmaple@example.com", "123456789", address, birthDate)

	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	userJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling user data to JSON:", err)
		return
	}

	fmt.Println(string(userJSON))
}
