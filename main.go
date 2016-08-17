package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gorilla/mux"
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

	router := mux.NewRouter()
	router.HandleFunc("/businesses", h.businessesIndexHandler)
	router.HandleFunc("/businesses/{businessId:[0-9]+}", h.businessesShowHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
	//h.printBusinessesHandler()
}

func (h *DBHandler) printBusinessesHandler() {
	var businesses Businesses
	h.db.Limit(5).Offset(1).Find(&businesses)
	fmt.Printf("%v\n", businesses)
}

func (h *DBHandler) businessesIndexHandler(w http.ResponseWriter, r *http.Request) {
	var businesses Businesses
	perPage := 50
	page := 1
	offset := (perPage * page) + 1
	h.db.Limit(perPage).Offset(offset).Find(&businesses)

	if err := json.NewEncoder(w).Encode(businesses); err != nil {
		panic(err)
	}
}

func (h *DBHandler) businessesShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	businessId := vars["businessId"]

	var business Business
	h.db.First(&business, businessId);

	if err := json.NewEncoder(w).Encode(business); err != nil {
		panic(err)
	}
}