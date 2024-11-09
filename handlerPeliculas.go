
package main

import (
    "github.com/gin-gonic/gin"
)

func peliculasHandler(c *gin.Context) {
    c.HTML(200, "peliculas.html", nil)
}
