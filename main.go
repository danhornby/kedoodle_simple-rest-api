package main

import "github.com/gin-gonic/gin"

func setupRouter(r *gin.Engine) {
	r.GET("/", helloWorldHandler)
	r.GET("/status", statusHandler)
}

func helloWorldHandler(c *gin.Context) {
	c.String(200, "Hello World")
}

func statusHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"my-application": []gin.H{
			{
				"version":     "1.0",
				"description": "my-application's description.",
				"sha":         "abc53458585",
			},
		},
	})
}

func main() {
	r := gin.Default()
	setupRouter(r)
}
