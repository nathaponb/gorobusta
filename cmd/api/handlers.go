package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestAuthPayload struct {
	Username string `json:"username,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

func (app *Config) authenticate(c *gin.Context) {

	// unmarshal json payload
	var req RequestAuthPayload
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "invalid credentials",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"error":   false,
		"message": "gotcha! :)",
		"data":    nil,
	})

}

func (app *Config) register(c *gin.Context) {}
