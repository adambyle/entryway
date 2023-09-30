package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	r.Static("/static", "./static")

	r.GET("/login", getLoginPage)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

func getLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "universal.html", gin.H{
		"title":   "Login",
		"content": "loginContent",
	})
}
