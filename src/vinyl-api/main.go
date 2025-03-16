package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "El vicio y la fama", Artist: "Gera MX", Price: 56.99},
	{ID: "2", Title: "Los Niños Grandes no Juegan", Artist: "Gera MX", Price: 45.99},
	{ID: "3", Title: "Precipicio", Artist: "Gera MX", Price: 26.99},
}

// controlador
func getAlbums(c *gin.Context) {
	//responde al cliente con la data serializada y el estatus HTTP
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// carga los datos a la estructura json --> struct
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// agregamos la estructura al slice
	albums = append(albums, newAlbum)

	// devuelve el http status y la estructura serializada struct --> Json
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "albúm no encontrado"})
}

func main() {
	// Inatancia de la herramienta
	router := gin.Default()

	//End points
	router.GET("/albums", getAlbums)

	router.GET("/albums/:id", getAlbumsById)

	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
