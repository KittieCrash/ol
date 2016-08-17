package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *DBHandler) BusinessesIndexHandler(w http.ResponseWriter, r *http.Request) {
	var businesses Businesses
	perPage := 50
	page := 1
	offset := (perPage * page) + 1
	h.db.Limit(perPage).Offset(offset).Find(&businesses)

	if err := json.NewEncoder(w).Encode(businesses); err != nil {
		panic(err)
	}
}

func (h *DBHandler) BusinessesShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	businessId := vars["businessId"]

	var business Business
	h.db.First(&business, businessId);

	if err := json.NewEncoder(w).Encode(business); err != nil {
		panic(err)
	}
}