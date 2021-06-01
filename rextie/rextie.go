package rextie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type rextieRequest struct {
	SourceCurr   string `json:"source_currency"`
	SourceAmount int    `json:"source_amount"`
	TargetCurr   string `json:"target_currency"`
}

type rextieResponse struct {
	RateBuy  string `json:"fx_rate_buy"`
	RateSell string `json:"fx_rate_sell"`
	Compra   float32
	Venta    float32
}

func (r *rextieResponse) setFloats() {
	compra, err := strconv.ParseFloat(r.RateBuy, 32)
	if err != nil {
		fmt.Println("Error while converting string to float")
		return
	}
	venta, err:= strconv.ParseFloat(r.RateSell, 32)
	if err != nil {
		fmt.Println("Error while converting string to float")
		return
	}

	r.Compra = float32(compra)
	r.Venta = float32(venta)
}


func GetRextieValues() *rextieResponse {
	url := "https://app.rextie.com/api/v1/fxrates/rate/"

	requestStruct := rextieRequest{
		SourceCurr:   "PEN",
		SourceAmount: 1,
		TargetCurr:   "USD",
	}

	jsonBytes, err := json.Marshal(requestStruct)
	if err != nil {
		fmt.Println("Bad Rextie body")
		return nil
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		fmt.Println("Bad Rextie request")
		return nil
	}
	request.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Bad Rextie response")
		return nil
	}

	bytesL, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Bad Rextie response body")
		return nil
	}

	var result rextieResponse
	if err = json.Unmarshal(bytesL, &result); err != nil {
		fmt.Println("Error while unmarshalling Rextie response")
		return nil
	}

	result.setFloats()
	return &result
}
