# HTMX Counter

## NOTEs

- Install the [gin-gonic](https://github.com/gin-gonic/gin)
    `go get -u github.com/gin-gonic/gin`
- Create initial main.go file
```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte("Hello World!"))
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```
- Run `go mod tidy`
- Render a template:
    - Create: "index.html"
    ```html
    <!DOCTYPE html>
    <html lang="en">

    <head>
        <title>HTMX Counter</title>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <script src="./htmx-1.9.10.js"></script>
    </head>

    <body>
        <div id="counter">0</div>
        <button hx-target="#counter" hx-post="/decrease">
            Decrease
        </button>
        <button hx-target="#counter" hx-post="/increase">
            Increase
        </button>
    </body>

    </html>
    ```
