package main

import (
	_ "embed"
	"encoding/json"
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

type Metadata struct {
	Version     string `json:"version"`
	Description string `json:"description"`
}

func setMetadata(m []byte) {
	var metadata Metadata
	err := json.Unmarshal(m, &metadata)
	if err != nil {
		log.Fatal(err)
	}
	Version = metadata.Version
	Description = metadata.Description
}

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
	setMetadata(metadataBytes)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
