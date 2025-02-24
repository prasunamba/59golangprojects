package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("static/*") // Load HTML static (load all files in "static/" directory)

	router.GET("/hello", handleGet)
	router.Run()

}
func handleGet(c *gin.Context) {
	/*
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<h1>Welcome to my awesome site!</h1>"))
	*/
	/* gin.H{} is a shorthand for defining a map[string]interface{}, which is used to pass dynamic data to the template.
	 */
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Name": "Prasunamba",
		"age ": "OK",
	})

}
