package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Directory struct {
	Name string `json:"name" validate:"required"`
}

func CreateDirectory(c *gin.Context) {
	var reqBody Directory

	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body: " + err.Error(),
		})

		return
	}

	validate := validator.New()

	err := validate.Struct(reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request: " + err.Error(),
		})

		return
	}

	params := c.Param("path")
	dir := filepath.Join(os.Getenv("HOME_CLOUD_STORAGE"), params)

	if err := os.Mkdir(filepath.Join(dir, reqBody.Name), os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create directory: " + err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"path":    params,
		"message": "Directory created successfully",
	})
}
