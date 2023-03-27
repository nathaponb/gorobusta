package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathaponb/robusta-gosrv/data"
)

type RequestAuthPayload struct {
	Username string `json:"username,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

func (app *Config) login(c *gin.Context) {

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

	// retrieve db user
	user, err := app.Repo.GetByUsername(req.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "username or password is incorrect!",
			"data":    nil,
		})
		return
	}

	// compare hash
	hashed := app.hasher(req.Password)
	if hashed != user.Password {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "username or password is incorrect!",
			"data":    nil,
		})
		return
	}

	// everything is fine
	c.AbortWithStatusJSON(http.StatusAccepted, gin.H{
		"error":   false,
		"message": "gotcha bro",
		"data":    nil,
	})

}

func (app *Config) register(c *gin.Context) {

	// unmarshal json payload
	var req data.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "invalid credentials",
			"data":    nil,
		})
		return
	}

	// hash password before do db transaction
	hashed := app.hasher(req.Password)
	req.Password = hashed

	// save to db
	err := app.Repo.Register(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "internal error",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error":   false,
		"message": "successfully registed user",
		"data":    nil,
	})
}
