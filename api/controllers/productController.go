package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dictuantran/shopee/api/models"
	"github.com/julienschmidt/httprouter"
)

type (
	ProductController struct{}
)

type ProductResponse struct {
	RespCode string           `json:"response_code"`
	RespName string           `json:"response_name"`
	Data     []models.Product `json:"data"`
}

func NexProductController() *ProductController {
	return &ProductController{}
}

// Get Product by Category ID
func (product ProductController) GetProductByCategory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Fetch data from model
	products := models.FetchProduct()

	// Define response
	d := ProductResponse{"1", "success", products}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(d)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}
