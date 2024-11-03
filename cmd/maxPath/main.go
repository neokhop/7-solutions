package main

import (
	"challenge/services"
	"fmt"
)

func main() {
	result := services.MaxPath("files/hard.json")
	fmt.Println("output =", result)
}
