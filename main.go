package main

import (
	"context"
	"xiaoxiaojiqiren/internal/pkg/wsclient"

	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		r := gin.Default()
		r.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"username": "name1",
				"data":     "data1",
			})
		})

		r.POST("/card", func(c *gin.Context) {
			var event struct {
				Challenge string `json:"challenge"`
				Type      string `json:"type"`
				Token     string `json:"token"`
			}
			if err := c.ShouldBindJSON(&event); err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(200, gin.H{
				"challenge": event.Challenge,
			})
		})

		r.Run()
	}()

	err := wsclient.Get().Start(context.Background())
	if err != nil {
		panic(err)
	}
}
