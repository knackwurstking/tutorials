# HTMX Counter

## NOTEs

- Install the [gin-gonic](https://github.com/gin-gonic/gin)
    `go get -u github.com/gin-gonic/gin`
- Create initial main.go file
```golang
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
