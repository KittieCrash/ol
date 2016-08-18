package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/unrolled/render"
)

type DBHandler struct {
	db 	*gorm.DB
	r 	*render.Render
}

func main() {
	h := InitDB()
	router := NewRouter(&h)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func InitDB() DBHandler {
	db, err := gorm.Open("sqlite3", "ol.db")
	if err != nil {
		panic("Failed to connect to the database!")
	}
	db.AutoMigrate(&Business{})

	r := render.New(render.Options{})
	h := DBHandler{db: db, r: r}

	return h
}

