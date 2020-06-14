package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dictuantran/shopee/api/config"
	"github.com/julienschmidt/httprouter"
)

type (
	WebController struct{}
)

type BaseUrl string

const (
	development BaseUrl = "http://localhost:3000/"
)

func NewWebController() *WebController {
	return &WebController{}
}

func (web WebController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// get Tax code
	var buffer bytes.Buffer
	buffer.WriteString(string(development))
	buffer.WriteString("tax_code")
	url := buffer.String()
	body := RequestGet(url)

	tax := TaxCodeResp{}
	jsonErr := json.Unmarshal(body, &tax)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	taxData := tax.Data
	config.TPL.ExecuteTemplate(w, "create.gohtml", taxData)
}

func RequestGet(url string) []byte {
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "tax-calculator")
	res, getErr := spaceClient.Do(req)

	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}
