package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DBHandler struct {
	db *gorm.DB
}

func main() {
	db, err := gorm.Open("sqlite3", "ol.db")
	if err != nil {
		panic("Failed to connect to the database!")
	}

	db.AutoMigrate(&Business{})
	h := DBHandler{db: db}
	h.printBusinessesHandler()
}

func (h *DBHandler) printBusinessesHandler() {
	var businesses Businesses
	h.db.Limit(5).Offset(1).Find(&businesses)
	fmt.Printf("%v\n", businesses)
}