package model

import (
	"log"

	"../utils"
)

type Favourite struct {
	Id      uint32
	UserId  uint32
	StockId string
}

func SelectFromFavouritesWhere(where string) ([]Favourite, error) {
	var favourites []Favourite
	sql := "SELECT id,user_id,stock_id FROM favourites"
	if len(where) > 0 {
		sql = sql + " Where " + where
	}
	rows, err := utils.SqlDB().Query(sql)
	if err != nil {
		log.Println("[WARN]", "model.SelectFromFavouritesWhere()", "SqlDB().Query()", err)
		return favourites, err
	}
	defer rows.Close()

	for rows.Next() {
		var favourite Favourite
		err := rows.Scan(&favourite.Id, &favourite.UserId, &favourite.StockId)
		if err != nil {
			log.Println("[WARN]", "model.SelectFromFavouritesWhere()", "rows.Scan()", err)
			break
		}
		favourites = append(favourites, favourite)
	}
	return favourites, nil
}
