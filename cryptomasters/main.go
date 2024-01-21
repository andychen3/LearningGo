package main

import (
	"fmt"

	"github.com/andychen3/cryptomasters/api"
)

func main() {
	currencies := []string{"BTC", "ETH", "DOGE"}

	for _, currency := range currencies {
		getCurrencyData(currency)
	}

}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}
}
