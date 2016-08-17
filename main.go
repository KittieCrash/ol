package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DBHandler struct {
	db *gorm.DB
}

func main() {
	h := InitDB()

/*
	router := mux.NewRouter()
	router.HandleFunc("/businesses", h.BusinessesIndexHandler)
	router.HandleFunc("/businesses/{businessId:[0-9]+}", h.BusinessesShowHandler)
	*/

	router := NewRouter(&h)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func InitDB() DBHandler {
	db, err := gorm.Open("sqlite3", "ol.db")
	if err != nil {
		panic("Failed to connect to the database!")
	}

	db.AutoMigrate(&Business{})
	h := DBHandler{db: db}
	return h
}

