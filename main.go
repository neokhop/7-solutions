package main

import (
	"challenge/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	beefPath := router.Group("/beef")
	beefPath.GET("/summary", services.Summary)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error running server:", err)
	}

	fmt.Println("Server running on port 8080")
}
