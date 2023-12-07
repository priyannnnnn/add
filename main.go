package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type numbers struct {
	A float64 `json:"a" binding:"exists"`
	B float64 `json:"b" binding:"exists"`
}

func calculate(c *gin.Context) {
	var n numbers

	c.BindJSON(&n)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": (n.A + n.B)})
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Addition Service v0.1.0 is running...")
	})

	router.POST("/calculate", calculate)

	router.Run()
}
