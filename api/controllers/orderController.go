package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/dictuantran/shopee/api/models"
	"github.com/julienschmidt/httprouter"
)

type (
	OrderController struct{}
)

type OrderResponse struct {
	RespCode string         `json:"response_code"`
	RespDesc string         `json:"response_description"`
	Data     []models.Order `json:"data"`
}

type OrderDetailResponse struct {
	RespCode string               `json:"response_code"`
	RespDesc string               `json:"response_description"`
	Data     []models.OrderDetail `json:"data"`
}

func NewOrderController() *OrderController {
	return &OrderController{}
}

func (oc OrderController) GetOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Fetch data from model
	orders := models.FetchOrder()

	// Define response
	d := OrderResponse{"1", "success", orders}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(d)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func (oc OrderController) GetOrderDetail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Fetch data from model
	orderID, _ := strconv.Atoi(p.ByName("ID"))
	orderDetail := models.FetchOrderDetailByOrderID(orderID)

	fmt.Print(orderID)
	// Define response
	d := OrderDetailResponse{"1", "success", orderDetail}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(d)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func (oc OrderController) ImportOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	OpenFileCSV(w, r)
}

func OpenFileCSV(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	csvFile, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer csvFile.Close()
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	ReadFileCSV(csvFile, handler)
}

func ReadFileCSV(csvFile multipart.File, handler *multipart.FileHeader) {
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, line := range csvLines {
		fmt.Println(line[0] + " " + line[1])
	}
}
