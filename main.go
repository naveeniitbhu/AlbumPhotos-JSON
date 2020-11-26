package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type App struct {
	R  *gin.Engine
	Db *sqlx.DB
}

func main() {
	db, err := sqlx.Connect("mysql", "ngnaven@gmail.com:ngnaven@gmail.com@(143.110.190.177:3306)/ngnaven@gmail.com")
	// db, err := sqlx.Connect("mysql", "root:root@(localhost:3306)/ngnaven@gmail.com")

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("connected to database")
	}

	app := App{
		R:  gin.Default(),
		Db: db,
	}

	// r := gin.Default()
	app.R.GET("/albums", app.AlbumListing)
	app.R.GET("/photos", app.PhotoListing)
	app.R.GET("/albums/:id", app.Search)
	app.R.Run()
}

type User struct {
	UserID int64  `json:"userId" db:"id"`
	Id     int64  `json:"id" db:"userId"`
	Title  string `json:"title" db:"title"`
}

type Photo struct {
	Id           int64  `json:"id" db:"userId"`
	AlbumId      int64  `json:"albumId" db:"albumId"`
	Title        string `json:"title" db:"title"`
	Url          string `json:"url" db:"url"`
	ThumbNailUrl string `json:"thumbnailUrl" db:"thumbnailUrl"`
}

var (
	user  = []User{}
	photo = []Photo{}
)

func (app *App) AlbumListing(c *gin.Context) {

	url := "https://jsonplaceholder.typicode.com/albums"
	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(body, &user)

	db := app.Db

	for i := 0; i < len(user); i++ {
		fmt.Println(user[i].Id, user[i].UserID, user[i].Title)
		_, err = db.Exec(`INSERT INTO album(id, userId, title) VALUES(?,?,?)`, user[i].Id, user[i].UserID, user[i].Title)

		if err != nil {
			log.Println(err, "Insertion failed")
		} else {
			log.Println("INFO: Data Inserted Succesfully")
		}
	}
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	fmt.Println("json binding success")
	// } else {
	// 	fmt.Println("json binding failed")
	// 	log.Println(err)
	// }
}

func (app *App) PhotoListing(c *gin.Context) {

	url := "https://jsonplaceholder.typicode.com/photos"
	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(body, &photo)

	db := app.Db

	for i := 0; i < len(photo); i++ {
		// fmt.Println(photo[i].Id, photo[i].AlbumId, photo[i].Title, photo[i].Url)
		_, err = db.Exec(`INSERT INTO photo(id, albumId, title, url, thumbnailUrl) VALUES(?,?,?,?,?)`, photo[i].Id, photo[i].AlbumId, photo[i].Title, photo[i].Url, photo[i].ThumbNailUrl)

		if err != nil {
			log.Println(err, "Insertion failed")
		} else {
			log.Println("INFO: Data Inserted Succesfully")
		}
	}
}

func (app *App) Search(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
	}
	fmt.Println(id)

	db := app.Db

	row, err := db.Queryx(`SELECT * FROM album WHERE id=?`, id)
	if err != nil {
		log.Println("select failed")
	}

	user := User{}
	for row.Next() {
		err := row.StructScan(&user)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("successful", user)
		}
	}
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	fmt.Println("json binding success")
	// } else {
	// 	fmt.Println("json binding failed")
	// 	log.Println(err)
	// }
}
