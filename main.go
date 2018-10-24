package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

// Redirect all requests to GCP API
func Redirect() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL
		url.Host = "gcp-api.quizzpy.app"
		url.Scheme = "https"
		c.Redirect(http.StatusMovedPermanently, url.String())
	}
}


func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(Redirect())

	router.Run(":" + port)
}
