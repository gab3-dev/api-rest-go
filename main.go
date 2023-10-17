package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Function used to reslice a slice when removing an element
func RemoveIndex(s []album, index int) []album {
	return append(s[:index], s[index+1:]...)
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbum create an album in the slice and return the slice(albums) as json
func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}

// dropAlbum delete an album from the slice and return the new slice as json
func dropAlbum(c *gin.Context) {
	var deleteAlbum album

	if err := c.BindJSON(&deleteAlbum); err != nil {
		return
	}

	for i, a := range albums {
		if a.ID == deleteAlbum.ID {
			albums = RemoveIndex(albums, i)
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Error": "album not found when trying to delete"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbum)
	router.DELETE("/albums", dropAlbum)

	router.Run("localhost:8080")
}
