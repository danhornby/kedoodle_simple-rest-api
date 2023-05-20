package main

import "github.com/gin-gonic/gin"

var (
	Version     = "dev"
	Description = "my-application's description."
	Commit      = "unknown"
)

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
				"version":     Version,
				"description": Description,
				"sha":         Commit,
			},
		},
	})
}

func main() {
	r := gin.Default()
	setupRouter(r)

	r.Run()
}
