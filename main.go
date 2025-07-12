package main

import (
	"arbitraj/internal/arbitrage"
	"arbitraj/internal/config"
	"fmt"
	"log"
)

func main() {
	exchanges, err := config.LoadExchanges("../../configs/exchanges.json")
	if err != nil {
		log.Fatalf("Borsa bilgileri yüklenemedi: %v", err)
	}
	fmt.Println("Yüklü Borsalar:")
	for _, ex := range exchanges {
		fmt.Printf("%s (%s)\n", ex.Name, ex.BaseURL)
	}

	// Mock fiyat verilerini çek
	prices := arbitrage.GetMockPrices()
	// Arbitraj fırsatlarını bul
	opps := arbitrage.FindArbitrage(prices, []arbitrage.Fee{
		{Exchange: "Binance", Rate: 0.001}, // %0.1 komisyon
		{Exchange: "Kraken", Rate: 0.001},  // %0.1 komisyon
	}, 0.1) // Minimum kar %0.1

	fmt.Println("Arbitraj Fırsatları:")
	for _, opp := range opps {
		fmt.Printf("%s borsasından al, %s borsasında sat: %s, Kar: %.2f USDT\n", opp.BuyExchange, opp.SellExchange, opp.Symbol, opp.Profit)
	}
}
