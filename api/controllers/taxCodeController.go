package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dictuantran/shopee/api/models"
	"github.com/julienschmidt/httprouter"
)

type (
	TaxCodeController struct{}
)

type TaxCodeResp struct {
	RespCode string           `json:"response_code"`
	RespDesc string           `json:"response_description"`
	Data     []models.TaxCode `json:"data"`
}

func NewTaxCodeController() *TaxCodeController {
	return &TaxCodeController{}
}

// GetUser retrieves an individual user resource
func (tc TaxCodeController) GetTaxCode(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Fetch data from model
	taxCode := models.FetchTaxCode()

	//define response
	d := TaxCodeResp{"1", "success", taxCode}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(d)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}
