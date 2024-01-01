package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	mediaPathTemplate = "./data/%s/%s"
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

	domainDir := fmt.Sprintf("./data/%s", domain)

	mediaPath := fmt.Sprintf(mediaPathTemplate, domain, filepath.Base(mediaURL))

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

func downloadMedia(mediaURL, mediaPath string) error {
	response, err := http.Get(mediaURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(mediaPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func extractDomain(url string) string {
	parts := strings.Split(url, "//")
	if len(parts) >= 2 {
		domain := strings.Split(parts[1], "/")[0]
		return domain
	}
	return ""
}
