package main

import (
	"log"
	"time"

	"../../model"
	"../../utils"
)

func DownloadByYear(year string) {
	stocks, err := model.SelectFromStocksWhere("")
	if err != nil {
		log.Println(err)
		return
	}

	for _, stock := range stocks {
		log.Println("Downloading", stock, "...")
		klines, err := model.DownloadKlinesFromWebsite(stock.Id, year)
		if err != nil {
			log.Println("[WARN]", err, stock)
			continue
		}
		log.Println("Inserting", stock, "...")
		for _, kline := range klines {
			err := kline.Insert()
			if err != nil {
				log.Println("[WARN]", err, stock)
			}
		}
	}
}

func DownloadByStock(code string) {
	//years := []string{"14", "15", "16", "17", "18"}
	years := []string{"19"}
	for _, y := range years {
		log.Println("Downloading", y, "...")
		klines, err := model.DownloadKlinesFromWebsite(code, y)
		if err != nil {
			log.Println("[WARN]", err)
			continue
		}
		log.Println("Inserting", y, "...")
		for _, kline := range klines {
			err := kline.Insert()
			if err != nil {
				log.Println("[WARN]", err)
				continue
			}
			log.Println("[INFO]", kline.StockId, kline.Date)
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	log.Println("Start...")
	utils.SendMail("kline start", "")
	DownloadByYear("19")
	utils.SendMail("kline over", "")
	log.Println("Over.")
}
