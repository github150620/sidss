package main

import (
	"log"
	"sync"
	"time"

	"../../model"
	"../../utils"
)

var wg sync.WaitGroup

func main() {
	log.Println("Start...")
	utils.SendMail("ma start", "")

	d := time.Now().AddDate(0, 0, 0).Format("060102")

	stocks, err := model.SelectFromStocksWhere("")
	if err != nil {
		log.Println(err)
		return
	}
	for _, s := range stocks {
		var ma model.MA
		ma.StockId = s.Id
		ma.Date = d
		ma.RefreshMA5()
		ma.RefreshMA10()
		ma.RefreshMA20()
		ma.RefreshMA30()
		ma.RefreshMA60()
		log.Println(ma)
		if ma.MA60 != 0 {
			log.Println("Insert()")
			ma.Insert()
			//ma.Update()
		}
	}

	utils.SendMail("ma over", "")
	log.Println("Over.")
}
