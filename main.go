package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/delay", MediaProxyHandler)
	router.GET("/api/info", InfoHandler)
	router.Run(":4000")
}
