package main

import(
	"net/http"
	"log"
	"strconv"
)

type PaginationResponse struct {
	PaginationData map[string]int `json:"_metadata"`
	Businesses []Business `json:"businesses"`
}

var DefaultPerPage = 50

func ParseQueryValues(req *http.Request, value string) int {
	vals := req.URL.Query()
	val := vals[value]
	if val != nil {
		v, err := strconv.ParseInt(val[0], 10, 0)
		if err != nil {
			log.Println(err)
			return 0
		}
		return int(v)
	}
	return 0
}

func GetPage(req *http.Request) int {
	page := ParseQueryValues(req, "page")
	if page < 1 {
		return 1
	}
	return page
}

func GetPerPage(req *http.Request) int {
	perPage := ParseQueryValues(req, "per_page")
	if perPage == 0 {
		return DefaultPerPage
	}
	return perPage
}