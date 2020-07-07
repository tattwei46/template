package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func add(c *gin.Context) {
	var req request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BAD_REQUEST"})
		return
	}

	c.JSON(http.StatusOK, req)
}
