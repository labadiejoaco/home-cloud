package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type fileDirectoryListing struct {
	Files       []string `json:"files"`
	Directories []string `json:"directories"`
}

func List(c *gin.Context) {
	params := c.Param("path")
	dir := filepath.Join(os.Getenv("HOME_CLOUD_STORAGE"), params)

	f, err := os.Open(dir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to open directory: " + err.Error(),
		})

		return
	}
	defer f.Close()

	files, err := f.Readdir(0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read directory content: " + err.Error(),
		})

		return
	}

	listing := fileDirectoryListing{
		Files:       []string{},
		Directories: []string{},
	}

	for _, v := range files {
		if v.IsDir() {
			listing.Directories = append(listing.Directories, v.Name())
		} else {
			listing.Files = append(listing.Files, v.Name())
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"path":    params,
		"content": listing,
	})
}
