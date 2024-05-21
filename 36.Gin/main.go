package main

import (
	"gin_/postgres"
	"log"
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
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	
	db := postgres.DB{}

	err := db.InserIntoAllbum(newAlbum.ID,newAlbum.Title,newAlbum.Artist,newAlbum.Price)
	if err != nil {
		log.Fatal(err)
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func Delete(c *gin.Context){
	id := c.Param("id")
	for i,al := range albums {
		if al.ID == id {
			db := postgres.DB{}
			albums = append(albums[:i],albums[i+1:]... )
			c.IndentedJSON(200,al)
			err := db.Delete(al.ID)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}
}

func Update(c *gin.Context){
	id := c.Param("id")
	for i,al := range albums{
		if al.ID == id {
			db := postgres.DB{}

			err  := db.Delete(al.ID)
			if err != nil {
				log.Fatal(err)
			}
			albums = append(albums[:i],albums[i+1:]... )
			var newAll album
			err = c.BindJSON(&newAll)
			
			if err!= nil{
				panic(err)
			}
			newAll.ID = al.ID
			err  = db.InserIntoAllbum(newAll.ID,newAll.Title,newAll.Artist,newAll.Price)
			if err != nil {
				log.Fatal(err)
			}
			albums = append(albums, newAll)
			c.IndentedJSON(http.StatusAccepted,newAll)
			return
		}
		
	}
}


func ReadALlAllmus() error{
	db := postgres.DB{}
	err := db.Connect()
	if err != nil {
		return err
	}

	query := 
	`
	SELECT * FROM album;
	`
	rows,err := db.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	for  rows.Next() {
		var al album

		err = rows.Scan(&al.ID,&al.Title,&al.Artist,&al.Price)
		if err !=  nil {
			return err
		}
		albums = append(albums, al)
	}
	return nil
}

func main() {

	err := ReadALlAllmus()
	if err != nil {
		log.Fatal(err)
	}
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.POST("/albums", postAlbums)
	router.PUT("/albums/:id",Update)
	router.DELETE("/albums/:id",Delete)

    router.Run("localhost:8080")
}