package main

import (
	"DEVGIN/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "5000"
	config.Load()
	router := gin.Default()

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
