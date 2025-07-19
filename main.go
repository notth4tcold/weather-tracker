package main

import (
    "fmt"
	"github.com/gin-gonic/gin"
    "net/http"
)

func main() {
	router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, World from Gin!",
        })
    })

	fmt.Println("Server running on http://localhost:8080")
    router.Run(":8080")
}