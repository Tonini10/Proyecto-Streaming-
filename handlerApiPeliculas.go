
package main

import (
    "github.com/gin-gonic/gin"
)

type Pelicula struct {
    ID     int    `json:"id"`
    Titulo string `json:"titulo"`
    Genero string `json:"genero"`
    Anio   int    `json:"anio"`
    Imagen string `json:"imagen"` // Field for image URL
}

func apiPeliculasHandler(c *gin.Context) {
    peliculas := []Pelicula{
        {1, "El Padrino", "Drama", 1972, "/images/el_padrino.jpg"},
        {2, "Pulp Fiction", "Crimen", 1994, "/images/pulp_fiction.jpg"},
        {3, "El Caballero de la Noche", "Acción", 2008, "/images/el_caballero_de_la_noche.jpg"},
        {4, "Forrest Gump", "Drama", 1994, "/images/forrest_gump.jpg"},
        {5, "Inception", "Ciencia Ficción", 2010, "/images/inception.jpg"},
        {6, "Matrix", "Ciencia Ficción", 1999, "/images/matrix.jpg"},
        {7, "El Señor de los Anillos: La Comunidad del Anillo", "Fantasía", 2001, "/images/senor_de_los_anillos.jpg"},
        {8, "La Lista de Schindler", "Histórico", 1993, "/images/la_lista_de_schindler.jpg"},
        {9, "Gladiador", "Acción", 2000, "/images/gladiador.jpg"},
        {10, "Titanic", "Romance", 1997, "/images/titanic.jpg"},
        {11, "Jurassic Park", "Aventura", 1993, "/images/jurassic_park.jpg"},
        {12, "El Resplandor", "Terror", 1980, "/images/el_resplandor.jpg"},
        {13, "Los Infiltrados", "Crimen", 2006, "/images/los_infiltrados.jpg"},
        {14, "Memento", "Misterio", 2000, "/images/memento.jpg"},
        {15, "El Gran Lebowski", "Comedia", 1998, "/images/el_gran_lebowski.jpg"},
        {16, "Avatar", "Ciencia Ficción", 2009, "/images/avatar.jpg"},
        {17, "Amélie", "Romance", 2001, "/images/amelie.jpg"},
        {18, "Rocky", "Deporte", 1976, "/images/rocky.jpg"},
        {19, "Django Unchained", "Western", 2012, "/images/django_unchained.jpg"},
        {20, "La La Land", "Musical", 2016, "/images/la_la_land.jpg"},
        {21, "Harry Potter y la Piedra Filosofal", "Fantasía", 2001, "/images/harry_potter.jpg"},
        {22, "Casino Royale", "Acción", 2006, "/images/casino_royale.jpg"},
        {23, "El Origen", "Ciencia Ficción", 2010, "/images/el_origen.jpg"},
        {24, "Interestelar", "Ciencia Ficción", 2014, "/images/interestelar.jpg"},
        {25, "Toy Story", "Animación", 1995, "/images/toy_story.jpg"},
    }

    c.JSON(200, peliculas)
}
