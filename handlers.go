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
	
	if len(businesses) == 0 {
		h.r.Text(w, http.StatusBadRequest, "400 Bad Request")
		return
	}

	ppr := PaginationResponse{
		map[string]int{"page": page, "per_page": perPage}, 
		businesses}

	h.r.JSON(w, http.StatusOK, ppr)
}

func (h *DBHandler) BusinessesShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	businessId := vars["businessId"]

	var business Business
	h.db.First(&business, businessId)
	
	if business == EmptyBusiness() {
		h.r.Text(w, http.StatusBadRequest, "400 Bad Request")
		return
	}

	h.r.JSON(w, http.StatusOK, business)
}

func (h *DBHandler) BusinessesCreateHandler(w http.ResponseWriter, r *http.Request) {
	h.r.Text(w, http.StatusNotImplemented, "501 Not Implemented")
}

func (h *DBHandler) BusinessesUpdateHandler(w http.ResponseWriter, r *http.Request) {
	h.r.Text(w, http.StatusNotImplemented, "501 Not Implemented")
}

func (h *DBHandler) BusinessesDeleteHandler(w http.ResponseWriter, r *http.Request) {
	h.r.Text(w, http.StatusNotImplemented, "501 Not Implemented")
}