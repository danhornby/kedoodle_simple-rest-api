package main

import (
	_ "embed"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	Version     = "No version provided"
	Description = "No description provided"
	Commit      = "No commit provided"
)

//go:embed metadata.json
var metadataBytes []byte

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

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
