package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/assets", "./assets")

	r.GET("/", func(c *gin.Context) {
		//c.Writer.Write([]byte("Hello World!"))
		templ, _ := template.ParseFiles("index.html")
		templ.ExecuteTemplate(c.Writer, "index.html", nil)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
