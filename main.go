package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	configParser := GetConfigParser()

	Host := getString(configParser.Get("HOST_NAME"), "localhost")
	Port := getString(configParser.Get("HOST_PORT"), "5000")
	DataDir := getString(configParser.Get("DATA_DIR"), "")

	address := fmt.Sprintf("%s:%s", Host, Port)
	log.Printf("Server Running on %s\n", address)
	if DataDir != "" {
		log.Printf("Data Directory is %s\n", DataDir)
	}
	origins := getStringSlice(configParser.Get("ALLOW_ORIGINS"))
	if len(origins) > 0 {
		log.Printf("Allowed Origins: %s", strings.Join(origins, ", "))
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  origins,
		AllowMethods:  []string{"GET"},
		AllowHeaders:  []string{"Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	router.GET("/delay", MediaProxyHandler)
	router.GET("/api/info", InfoHandler)

	router.Run(address)
}

func getString(value interface{}, defaultValue string) string {
	if str, ok := value.(string); ok && str != "" {
		return str
	}
	return defaultValue
}

func getStringSlice(value interface{}) []string {
	if slice, ok := value.([]string); ok {
		return slice
	}
	return nil
}
