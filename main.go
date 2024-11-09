
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Serve static files (images)
    router.Static("/images", "./static/images")

    router.LoadHTMLGlob("templates/*.html")

    router.GET("/", mainPage)
    router.GET("/peliculas", peliculasHandler)
    router.GET("/crear-usuario", crearUsuarioHandler)
    router.GET("/comprar", comprarHandler)
    router.GET("/ver-compras", verComprasHandler)
    router.GET("/api/peliculas", apiPeliculasHandler) // API for movies

    router.Run(":8080")
}
