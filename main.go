package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.LoadHTMLFiles("templates/index.htm")

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

	r.Run(":8080")
}
