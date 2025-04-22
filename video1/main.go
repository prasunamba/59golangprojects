package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello", hellohandler)
	router.GET("/", func(c *gin.Context) {
		c.File("static/index.html") // directly serves the HTML file
	})
	// Load HTML templates
	router.LoadHTMLGlob("static/*")

	// Serve the form
	router.GET("/form", func(c *gin.Context) {
		c.HTML(200, "form.html", nil)
	})

	// Handle form submission
	router.POST("/submit", func(c *gin.Context) {
		name := c.PostForm("name")
		c.String(200, "Hello, %s!", name)
	})
	router.Run()

}

func hellohandler(c *gin.Context) {

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<h1>Welcome to my awesome site!</h1>"))

	/* gin.H{} is a shorthand for defining a map[string]interface{}, which is used to pass dynamic data to the template.
	 */
	/* c.HTML(http.StatusOK, "index.html", gin.H{
		"Name": "Prasunamba",
		"age ": "OK",
	}) */
	// c.HTML(http.StatusAccepted, "index.html", nil)
}
