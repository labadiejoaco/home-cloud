package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	params := c.Param("path")
	dir := filepath.Join(os.Getenv("HOME_CLOUD_STORAGE"), params)

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to proccess the request: " + err.Error(),
		})

		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No files found in the request",
		})

		return
	}

	for _, file := range files {
		dst := filepath.Join(dir, file.Filename)

		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save uploaded file: " + err.Error(),
			})

			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"path":    params,
		"message": "Files uploaded successfully",
	})
}
