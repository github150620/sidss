package main

import (
	"log"
	"time"

	"../../model"
)

func main() {
	log.Println("[INFO]", "Start...")
	for {
		RefreshFavouritesPrices()
		//RefreshAllPrices()
		t := time.Now().Format("15:04:05")
		if t < "09:00:00" {
			break
		} else if t > "12:00:00" && t < "13:00:00" {
			break
		} else if t > "15:00:00" {
			break
		}
		log.Println("[INFO]", "Wait 10 seconds...")
		time.Sleep(time.Duration(10) * time.Second)
	}
	log.Println("[INFO]", "Over.")
}

func RefreshFavouritesPrices() {
	log.Println("[INFO]", "Preparing data...")
	stocks, err := model.SelectFromFavouritesGroupByStockId()
	if err != nil {
		log.Println(err)
		return
	}
	var codes []string
	for _, stock := range stocks {
		codes = append(codes, stock.Id)
	}
	log.Println("[INFO]", codes)
	log.Println("[INFO]", "Downloading...")
	prices, _ := model.DownloadPricesFromWebsite(codes)
	log.Println("[INFO]", "Updating...")
	for _, price := range prices {
		price.Update()
		log.Printf("[INFO] %v.Update()\n", price)
	}
}

func RefreshAllPrices() {
	stocks, err := model.SelectFromStocksWhere("")
	if err != nil {
		log.Println(err)
		return
	}

	var codes []string
	for i, stock := range stocks {
		codes = append(codes, stock.Id)
		if i%50 == 49 {
			log.Println("Downloading...", i)
			prices, _ := model.DownloadPricesFromWebsite(codes)
			for _, price := range prices {
				price.Update()
			}
			codes = codes[0:0]
		}
	}
	if len(codes) != 0 {
		log.Println("Downloading...")
		prices, _ := model.DownloadPricesFromWebsite(codes)
		for _, price := range prices {
			price.Update()
		}
	}
}