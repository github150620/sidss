package model

import (
	"log"

	"../utils"
)

type Stock struct {
	Id   string
	Name string
}

func SelectFromStocksWhere(where string) ([]Stock, error) {
	var stocks []Stock
	sql := "SELECT id,name FROM stocks"
	if len(where) > 0 {
		sql = sql + " Where " + where
	}
	rows, err := utils.SqlDB().Query(sql)
	if err != nil {
		log.Println("[WARN]", "model.SelectFromStocks()", "SqlDB().Query()", err)
		return stocks, err
	}
	defer rows.Close()

	for rows.Next() {
		var stock Stock
		err := rows.Scan(&stock.Id, &stock.Name)
		if err != nil {
			log.Println("[WARN]", "model.SelectFromStocks()", "rows.Scan()", err)
			break
		}
		stocks = append(stocks, stock)
	}
	return stocks, nil
}

func SelectFromFavouritesGroupByStockId() ([]Stock, error) {
	var stocks []Stock
	sql := "SELECT stock_id FROM favourites GROUP BY stock_id"
	rows, err := utils.SqlDB().Query(sql)
	if err != nil {
		log.Println("[WARN]", "model.SelectFromFavouritesGroupByStockId()", "SqlDB().Query()", err)
		return stocks, err
	}
	defer rows.Close()

	for rows.Next() {
		var stock Stock
		err := rows.Scan(&stock.Id)
		if err != nil {
			log.Println("[WARN]", "model.SelectFromFavouritesGroupByStockId()", "rows.Scan()", err)
			break
		}
		stocks = append(stocks, stock)
	}
	return stocks, nil
}
