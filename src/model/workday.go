// workday
package model

import (
	"log"

	"../utils"
)

type Workday struct {
	Id   uint32
	Date string
}

func SelectFromWorkdaysWhere(where string) ([]Workday, error) {
	var workdays []Workday
	sql := "SELECT id,date FROM workdays"
	if len(where) > 0 {
		sql = sql + " WHERE " + where
	}
	sql = sql + " ORDER BY date desc"
	rows, err := utils.SqlDB().Query(sql)
	if err != nil {
		log.Println("[WARN]", "model.SelectFromStocks()", "SqlDB().Query()", err)
		return workdays, err
	}
	defer rows.Close()

	for rows.Next() {
		var workday Workday
		err := rows.Scan(&workday.Id, &workday.Date)
		if err != nil {
			log.Println("[WARN]", "model.SelectFromStocks()", "rows.Scan()", err)
			break
		}
		workdays = append(workdays, workday)
	}
	return workdays, nil
}
