package main

import "github.com/gin-gonic/gin"

func setupRouter(r *gin.Engine) {
	r.GET("/", helloWorldHandler)
}

func helloWorldHandler(c *gin.Context) {
	c.String(200, "Hello World")
}

func main() {
	r := gin.Default()
	setupRouter(r)
}
