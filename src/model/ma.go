package model

import (
	"log"

	"../utils"
)

type MA struct {
	Id      uint64
	StockId string
	Date    string
	MA5     float64
	MA10    float64
	MA20    float64
	MA30    float64
	MA60    float64
}

func (v *MA) Insert() error {
	var err error

	stmt, err := utils.SqlDB().Prepare("INSERT INTO ma(stock_id,date,ma5,ma10,ma20,ma30,ma60) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		log.Println("[WARN]", "MA.Insert():", "SqlDB().Prepare()", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		v.StockId,
		v.Date,
		v.MA5,
		v.MA10,
		v.MA20,
		v.MA30,
		v.MA60,
	)
	if err != nil {
		log.Println("[WARN]", "MA.Insert()", "Exec()", err)
		return err
	}

	return nil
}

func (v *MA) Update() error {
	var err error

	sql := "UPDATE ma SET ma5=?,ma10=?,ma20=?,ma30=?,ma60=? WHERE stock_id=? AND date=?"
	stmt, err := utils.SqlDB().Prepare(sql)
	if err != nil {
		log.Println("[ERROR]", "Price.Update()", "SqlDB().Prepare()", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		v.MA5,
		v.MA10,
		v.MA20,
		v.MA30,
		v.MA60,
		v.StockId,
		v.Date,
	)
	if err != nil {
		log.Println("[WARN]", "Price.Update()", "Exec()", err)
		return err
	}
	return nil
}

func (v *MA) RefreshMA5() error {
	sql := "select avg(c),count(1),max(date) from (select * from klines where stock_id='" + v.StockId + "' and date<='" + v.Date + "' order by `date` desc limit 5) t HAVING count(1)=5 AND max(date)='" + v.Date + "'"
	rows, err := utils.SqlDB().Query(sql)
	if err != nil {
		log.Println("[ERROR]", "model.RefreshMA5()", "SqlDB().Query()", err)
		return err
	}
	defer rows.Close()
	if rows.Next() {
		var count int
		var date string
		err := rows.Scan(
			&v.MA5,
			&count,
			&date,
		)
		if err != nil {
			log.Println("[WARN]", "MA.RefreshMA5()", "rows.Scan()", err)
			return err
		}
	}
	return nil
}

func (v *MA) RefreshMA10() error {
	sql := "select avg(c),count(1),max(date) from (select * from klines where stock_id='" + v.StockId + "' and date<='" + v.Date + "' order by `date` desc limit 10) t HAVING count(1)=10 AND max(date)='" + v.Date + "'"
	rows, err := utils.SqlDB().Query(sql)
	if err != nil {
		log.Println("[ERROR]", "model.RefreshMA10()", "SqlDB().Query()", err)
		return err
	}
	defer rows.Close()
	if rows.Next() {
		var count int
		var date string
		err := rows.Scan(
			&v.MA10,
			&count,
			&date,
		)
		if err != nil {
			log.Println("[WARN]", "MA.RefreshMA10()", "rows.Scan()", err)
			return err
		}
	}
	return nil
}

func (v *MA) RefreshMA20() error {
	sql := "select avg(c),count(1),max(date) from (select * from klines where stock_id='" + v.StockId + "' and date<='" + v.Date + "' order by `date` desc limit 20) t HAVING count(1)=20 AND max(date)='" + v.Date + "'"
	rows, err := utils.SqlDB().Query(sql)
	if err != nil {
		log.Println("[ERROR]", "model.RefreshMA20()", "SqlDB().Query()", err)
		return err
	}
	defer rows.Close()
	if rows.Next() {
		var count int
		var date string
		err := rows.Scan(
			&v.MA20,
			&count,
			&date,
		)
		if err != nil {
			log.Println("[WARN]", "MA.RefreshMA20()", "rows.Scan()", err)
			return err
		}
	}
	return nil
}

func (v *MA) RefreshMA30() error {
	sql := "select avg(c),count(1),max(date) from (select * from klines where stock_id='" + v.StockId + "' and date<='" + v.Date + "' order by `date` desc limit 30) t HAVING count(1)=30 AND max(date)='" + v.Date + "'"
	rows, err := utils.SqlDB().Query(sql)
	if err != nil {
		log.Println("[ERROR]", "model.RefreshMA30()", "SqlDB().Query()", err)
		return err
	}
	defer rows.Close()

	if rows.Next() {
		var count int
		var date string
		err := rows.Scan(
			&v.MA30,
			&count,
			&date,
		)
		if err != nil {
			log.Println("[WARN]", "MA.RefreshMA30()", "rows.Scan()", err)
			return err
		}
	}
	return nil
}

func (v *MA) RefreshMA60() error {
	sql := "select avg(c),count(1),max(date) from (select * from klines where stock_id='" + v.StockId + "' and date<='" + v.Date + "' order by `date` desc limit 60) t HAVING count(1)=60 AND max(date)='" + v.Date + "'"
	rows, err := utils.SqlDB().Query(sql)
	if err != nil {
		log.Println("[ERROR]", "model.RefreshMA30()", "SqlDB().Query()", err)
		return err
	}
	defer rows.Close()

	if rows.Next() {
		var count int
		var date string
		err := rows.Scan(
			&v.MA60,
			&count,
			&date,
		)
		if err != nil {
			log.Println("[WARN]", "MA.RefreshMA30()", "rows.Scan()", err)
			return err
		}
	}
	return nil
}
