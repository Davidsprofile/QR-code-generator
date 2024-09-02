package main

import (
	"net/http"
	"os" // Import os package for environment variables

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func main() {
	r := gin.Default()

	// Ensure the path matches the file extension in your project
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.htm", nil)
	})

	r.POST("/generate", func(c *gin.Context) {
		text := c.PostForm("text")
		qrCode, err := qrcode.Encode(text, qrcode.Medium, 256)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to generate QR code")
			return
		}
		c.Data(http.StatusOK, "image/png", qrCode)
	})

	// Get the port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
