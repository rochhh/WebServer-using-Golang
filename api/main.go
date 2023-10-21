package main

import (
	"fmt"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)


type album struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums( c *gin.Context ){
	c.IndentedJSON(http.StatusOK , albums)
}


func postAlbums( c *gin.Context ){

	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	} 

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated , newAlbum)

}


func getAlbumById(c *gin.Context){
	id := c.Param("id")

	for _ , a := range albums{
		if a.ID == id{
			c.IndentedJSON(http.StatusOK , a)
		}
	}

	c.IndentedJSON(http.StatusNotFound , gin.H{"message" : "album not found !"})

}

func getAlbmsByArtist( name string ) ( []album , error ) {  

	var albums []album

	rows , err := db.Query("SELECT * FROM album WHERE artist = ? " , name) 			// prevents sql injection !

	if err != nil{
		return nil , fmt.Errorf("albumsByArtist %q : %v " , name , err)
	}

	defer rows.Close()

	for rows.Next(){
		var alb album 
		

		if err := rows.Scan(&alb.ID , &alb.Title , &alb.Artist , &alb.Price); err != nil {
			return nil , fmt.Errorf("albumsByArtist %q : %v " , name , err)
		}

		albums = append(albums, alb)

	}

	if err := rows.Err(); err != nil {
		return nil , fmt.Errorf("albumsByArtist %q : %v " , name , err)
	}

	return albums , nil
}


func getAlbumsByID( id int64 ) ( album , error ){
	var alb album

	row := db.QueryRow("SELECT * FROM album WHERE id = ? " , id)

	if err := row.Scan(&alb.ID , &alb.Title , &alb.Artist , &alb.Price); err != nil {
		if err == sql.ErrNoRows{
			return alb , fmt.Errorf("albumsById %d no such album " , id )
		}
		return alb , fmt.Errorf("albumsById %d %v " , id ,err )
	}

	return alb , nil

}



var db *sql.DB



func main() {
	
	cfg := mysql.Config{
		User: "",
		Passwd: "",
		Net: "",
		Addr: "",
		DBName: "",
	}

	var err error

	db , err = sql.Open("mysql" , cfg.FormatDSN())

	if err != nil{
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	// router := gin.Default()
	// router.GET("/albums" , getAlbums)
	// router.GET("/albums/:id" , getAlbumById)
	// router.POST("/albums" , postAlbums)
	
	// albums , err := getAlbmsByArtist("John Coltrane")

	album , err := getAlbumsByID(2)

	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(album)
	// fmt.Println(albums)

	// router.Run("localhost:8080")

}