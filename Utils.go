package main

import (
	"io"
	"net/http"
	"os"
	"strings"
)

// A function to extract domain from link
func extractDomain(url string) string {
	parts := strings.Split(url, "//")
	if len(parts) >= 2 {
		domain := strings.Split(parts[1], "/")[0]
		return domain
	}
	return ""
}

// Function to download media to server
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
