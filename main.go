package main

import (
	"fmt"
	"time"

	"khwahishIsCooking.com/crypto/api"
)

func main() {
	go getCurrencyData("BTC")
	go getCurrencyData("ETH")
	go getCurrencyData("BCH")
	go getCurrencyData("LTC")
	go getCurrencyData("XRP")
	time.Sleep(time.Second * 5)
}
func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("Rate for %v is %.2f \n", rate.Currency, rate.Price)
	}
}
