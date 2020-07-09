package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/dictuantran/shopee/api/models"
	"github.com/dictuantran/shopee/api/repositories"
	"github.com/julienschmidt/httprouter"
)

// OrderController struct
type (
	OrderController struct {
		orderRepository *repositories.OrderRepository
	}
)

// NewOrderController controller
func NewOrderController() *OrderController {
	return &OrderController{}
}

// GetOrder get order
func (oc OrderController) GetOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Fetch data from model
	orders, err := oc.orderRepository.FetchOrder()
	if err != nil {
		log.Printf("Error in decoding the request body: %+v\n", err)
		return
	}
	// Define response
	d := models.OrderResponse{"1", "success", orders}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(d)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// GetOrderDetail get order detail
func (oc OrderController) GetOrderDetail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Fetch data from model
	orderID, _ := strconv.Atoi(p.ByName("ID"))
	orderDetail, err := oc.orderRepository.FetchOrderDetailByOrderID(orderID)
	if err != nil {
		log.Printf("Error in decoding the request body: %+v\n", err)
		return
	}

	fmt.Print(orderID)
	// Define response
	d := models.OrderDetailResponse{"1", "success", orderDetail}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(d)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// ImportOrder method
func (oc OrderController) ImportOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	OpenFileCSV(w, r)
}

// OpenFileCSV method
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
	ResponseFile(w, r, csvFile)
}

// ReadFileCSV method
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

// ResponseFile method
func ResponseFile(w http.ResponseWriter, r *http.Request, csvFile multipart.File) {
	w.Header().Set("Content-Type", "text/csv")
	w.WriteHeader(200)
	wr := csv.NewWriter(w)
	records := []string{"India", "Canada", "Japan"}

	err := wr.Write(records)
	if err != nil {
		http.Error(w, "Error sending csv: "+err.Error(), http.StatusInternalServerError)
		return
	}

}
