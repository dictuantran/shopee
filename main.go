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
	prodController := controllers.NexProductController()
	tc := controllers.NewTaxCodeController()
	cc := controllers.NewCartController()
	oc := controllers.NewOrderController()
	web := controllers.NewWebController()

	// REST API
	router.GET("/tax_code", tc.GetTaxCode)

	// Web Frontend
	router.GET("/", web.Index)
	router.GET("/order_view/:store_id", web.ViewBill)

	// API
	router.GET("/order/GetOrder", oc.GetOrder)
	router.GET("/order/GetOrderDetail/:ID", oc.GetOrderDetail)
	router.GET("/product/GetProduct", prodController.GetProductByCategory)

	// POST object tax into cart or order detail
	router.POST("/cart", cc.CreateCart)

	// add CORS support (Cross Origin Resource Sharing)
	// all origins accepted with simple methods (GET, POST). See
	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":3001", handler))
}
