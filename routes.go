package main

import(
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name		string
	HttpMethod	string
	Path		string
	Handler		http.HandlerFunc
}

type Routes []Route

func NewRouter(h *DBHandler) *mux.Router {
	router := mux.NewRouter()
	routes := CreateRoutes(h)
	for _, route := range routes {
		router.
			Methods(route.HttpMethod).
			Path(route.Path).
			Name(route.Name).
			Handler(route.Handler)
	}

	return router
}

func CreateRoutes(h *DBHandler) Routes{
	var routes = Routes {
		Route {
			"BusinessesIndex",
			"GET",
			"/businesses",
			h.BusinessesIndexHandler,
		},
		Route {
			"BusinessesCreate",
			"POST",
			"/businesses",
			h.BusinessesCreateHandler,
		},
		Route {
			"BusinessesShow",
			"GET",
			"/businesses/{businessId:[0-9]+}",
			h.BusinessesShowHandler,
		},		
		Route {
			"BusinessesUpdate",
			"PUT",
			"/businesses/{businessId:[0-9]+}",
			h.BusinessesUpdateHandler,
		},
		Route {
			"BusinessesDelete",
			"DELETE",
			"/businesses/{businessId:[0-9]+}",
			h.BusinessesDeleteHandler,
		},
	}

	return routes
}
