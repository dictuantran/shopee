package main

import (
	"log"
	"net/http"

	"github.com/dictuantran/shopee/api/controllers"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	router := httprouter.New()

	// get Controller instance
	tc := controllers.NewTaxCodeController()
	web := controllers.NewWebController()

	// REST API
	router.GET("/tax_code", tc.GetTaxCode)

	// Web Frontend
	router.GET("/", web.Index)

	// add CORS support (Cross Origin Resource Sharing)
	// all origins accepted with simple methods (GET, POST). See
	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":3001", handler))
}
