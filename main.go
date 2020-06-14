package main

import (
	"log"
	"net/http"

	"github.com/dictuantran/shopee/api/controllers"
	"github.com/julienschmidt/httprouter"
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

	log.Fatal(http.ListenAndServe(":3000", router))
}
