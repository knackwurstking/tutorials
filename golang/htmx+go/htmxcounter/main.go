package main

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
)

var (
	counter = NewCounter(0)
)

func main() {
	r := gin.Default()

	r.Static("/assets", "./assets")

	r.GET("/", func(ctx *gin.Context) {
		//c.Writer.Write([]byte("Hello World!"))
		templ, _ := template.ParseFiles("index.templ")
		templ.ExecuteTemplate(ctx.Writer, "index.templ", Index{
			Value: counter.GetValue(),
		})
	})

	r.POST("/decrease", func(ctx *gin.Context) {
		counter.Decrease()
		templ := template.Must(
			template.New("counter").Parse(
				fmt.Sprintf("<div id=\"counter\">%d</div>", counter.GetValue()),
			),
		)
		templ.ExecuteTemplate(ctx.Writer, "counter", nil)
	})

	r.POST("/increase", func(ctx *gin.Context) {
		counter.Increase()
		templ := template.Must(
			template.New("counter").Parse(
				fmt.Sprintf("<div id=\"counter\">%d</div>", counter.GetValue()),
			),
		)
		templ.ExecuteTemplate(ctx.Writer, "counter", nil)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
