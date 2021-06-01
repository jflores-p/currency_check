package main

import (
	"fmt"
	"github.com/jflores-p/currency_check/kambista"
	"github.com/jflores-p/currency_check/rextie"
)

func main() {

	rextieResult := rextie.GetRextieValues()
	fmt.Printf("Rextie Venta:  %.4f\nRextie Compra: %.4f \n", rextieResult.Venta, rextieResult.Compra)

	fmt.Println("---------------------")
	kambistaResult :=kambista.GetKambistaValues()
	fmt.Printf("Kambista Venta:  %.4f\nKambista Compra: %.4f ", kambistaResult.Venta, kambistaResult.Compra)
}
