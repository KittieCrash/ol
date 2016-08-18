package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *DBHandler) BusinessesIndexHandler(w http.ResponseWriter, r *http.Request) {
	var businesses []Business
	perPage := GetPerPage(r)
	page := GetPage(r)
	offset := (perPage * (page-1)) + 1

	h.db.Limit(perPage).Offset(offset).Find(&businesses)
	
	ppr := PrettyPaginationResponse{
		map[string]int{"page": page, "per_page": perPage}, 
		businesses}

	h.r.JSON(w, http.StatusOK, ppr)
}

func (h *DBHandler) BusinessesShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	businessId := vars["businessId"]

	var business Business
	h.db.First(&business, businessId)
	
	h.r.JSON(w, http.StatusOK, business)
}

func (h *DBHandler) BusinessesCreateHandler(w http.ResponseWriter, r *http.Request) {

}