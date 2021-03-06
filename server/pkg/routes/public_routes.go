package routes

import (
	"net/http"

	"github.com/febzey/port/pkg/controllers"
	"github.com/gorilla/mux"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
	{
		Method:      http.MethodGet,
		Pattern:     "/test",
		HandlerFunc: controllers.Test,
	},
}

func Public(router *mux.Router) {
	for _, route := range routes {
		router.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method)
	}
}
