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
	c.String(200, "OK")
}

func main() {
	r := gin.Default()
	setupRouter(r)
}
