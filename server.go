package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const defaultPort = 8080

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = strconv.Itoa(defaultPort)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Error: Invalid PORT value")
		return
	}

	fmt.Println(port)

}
