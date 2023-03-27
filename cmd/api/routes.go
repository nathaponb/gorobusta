package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *Config) routes() http.Handler {
	ngin := gin.New()
	ngin.Use(gin.Logger())
	ngin.Use(gin.Recovery())

	ngin.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	ngin.POST("/login", app.authenticate)

	return ngin
}
