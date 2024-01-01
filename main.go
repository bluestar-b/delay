package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	configParser := GetConfigParser()
	Host := configParser.Get("HOST_NAME")
	Port := configParser.Get("HOST_PORT")
	DataDir := configParser.Get("DATA_DIR")
	if Host == "" {
		Host = "localhost"
	}

	if Port == "" {
		Port = "5000"
	}

	address := fmt.Sprintf("%s:%s", Host, Port)
	log.Printf("Server running on %s\n", address)
	log.Printf("Data Directory is %s\n", DataDir)
	router := gin.Default()

	router.GET("/delay", MediaProxyHandler)
	router.GET("/api/info", InfoHandler)

	address = fmt.Sprintf("%s:%s", Host, Port)
	router.Run(address)
}
