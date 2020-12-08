/*
 * ACME Reviews - PSIDI
 *
 * Swagger proposed server for the Review Product infrastructure at ACME .Inc
 *
 * API version: 0.1
 * Contact: 1171071@isep.ipp.pt
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"GetCustomer",
		strings.ToUpper("Get"),
		"/v1/customers/{customerId}",
		GetCustomer,
	},

	Route{
		"RegisterCustomer",
		strings.ToUpper("Post"),
		"/v1/customers",
		RegisterCustomer,
	},
}
