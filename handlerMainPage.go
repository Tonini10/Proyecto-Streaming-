
package main

import (
    "github.com/gin-gonic/gin"
)

func mainPage(c *gin.Context) {
    c.HTML(200, "index.html", nil)
}
