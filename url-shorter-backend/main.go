package main

import (
	"log"
	"service/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("func config.LoadConfig err : %v", err.Error())
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(config.PORT)
}
