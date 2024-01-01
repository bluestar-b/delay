package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	mediaPathTemplate = "%s/%s/%s"
	downloadMutex     sync.Mutex
)

func MediaProxyHandler(c *gin.Context) {

	mediaURL := c.Query("origin")
	if mediaURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'origin' parameter"})
		return
	}

	domain := extractDomain(mediaURL)
	if domain == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to extract domain"})
		return
	}
	configParser := GetConfigParser()
	dataDir := configParser.Get("DATA_DIR")
	if dataDir == "" {
		dataDir = "./data"
	}

	domainDir := fmt.Sprintf("%s/%s", dataDir, domain)

	mediaPath := fmt.Sprintf(mediaPathTemplate, dataDir, domain, filepath.Base(mediaURL))

	downloadMutex.Lock()
	defer downloadMutex.Unlock()

	if _, err := os.Stat(mediaPath); os.IsNotExist(err) {
		fmt.Printf("Media for domain '%s' not found, downloading...\n", domain)

		if err := os.MkdirAll(domainDir, 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create domain directory"})
			return
		}

		err := downloadMedia(mediaURL, mediaPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to download media"})
			return
		}
		fmt.Printf("Media downloaded successfully for domain '%s'.\n", domain)
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.File(mediaPath)
}
