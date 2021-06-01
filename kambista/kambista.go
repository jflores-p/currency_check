package kambista

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type kambistaResponse struct {
	TP tipoCambio `json:"tc"`
}
type tipoCambio struct {
	Compra float32 `json:"bid"`
	Venta  float32 `json:"ask"`
}

const url = "https://api.kambista.com/v1/exchange/calculates?originCurrency=PEN&destinationCurrency=USD&amount=1500&active=S"


func GetKambistaValues() *tipoCambio {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Bad Kambista request")
		return nil
	}

	bytesL, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Bad Kambista response body")
		return nil
	}

	var result kambistaResponse
	if err = json.Unmarshal(bytesL, &result); err != nil {
		fmt.Println("Error while unmarshalling Kambista response")
		return nil
	}

	return &result.TP
}