package main

import (
	"log"
	"time"

	"../../model"
)

func DownloadAll() {
	//stocks, err := model.SelectFromStocksWhere("id>=3500 AND id<4000")
	//stocks, err := model.SelectFromStocksWhere("code in (SELECT `code` FROM favourites WHERE user_id=1)")
	//stocks, err := model.SelectFromStocksWhere("code not in (SELECT `code` FROM reports WHERE `date`='2019-03-31')")
	stocks, err := model.SelectFromStocksWhere("")
	if err != nil {
		log.Println(err)
		return
	}

	for _, stock := range stocks {
		log.Println("Downloading", stock, "...")
		reports, err := model.DownloadReportsFromWebsite(stock.Id)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println("Inserting", stock, "...")
		for _, statement := range reports {
			err := statement.Insert()
			if err != nil {
				log.Println(err)
			}
		}
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	log.Println("Start...")
	DownloadAll()
	log.Println("Over.")
}
