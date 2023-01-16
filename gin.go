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
	Date   time    `json:"date"`
}

var albums = []album{
	{ID: "1", Title: "Rock", Artist: "Shabnami S", Price: 50.12},
	{ID: "2", Title: "Pop", Artist: "Farzonai Kh", Price: 60.22},
	{ID: "3", Title: "Jazz", Artist: "Nigina A", Price: 70.36},
	{ID: "4", Title: "Mix", Artist: "Sadriddin N", Price: 800.54},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"code": "404"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsById)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8000")
}
