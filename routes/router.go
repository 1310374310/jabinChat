package routes

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {

	// create mux.Router instance
	router := mux.NewRouter().StrictSlash(true)

	// travers the webRoutes
	for _, route := range webRoutes {
		// aplly webRoutes to mux.Router
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}

	return router
}
