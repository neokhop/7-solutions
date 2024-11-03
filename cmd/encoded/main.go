package main

import (
	"challenge/services"
	"fmt"
)

func main() {
	var encoded string
	fmt.Print("Enter the encoded string: ")
	fmt.Scan(&encoded)

	result := services.Encoded(encoded)
	fmt.Println("Result:", result)
}
