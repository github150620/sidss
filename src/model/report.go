package model

import (
	"log"

	"../utils"
)

type Report struct {
	Id uint32
	StockId string
	Date string
	Eps float64
	Roe float64
	Gpr float64
	Asset float64
	Debt float64
}

func (v *Report) Insert() error {
	var err error

	stmt, err := utils.SqlDB().Prepare("INSERT INTO reports(report_id,`date`,eps,roe,gpr,asset,debt) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		log.Println("[WARN]", "Report.Insert():", "SqlDB().Prepare()", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		v.StockId,
		v.Date,
		v.Eps,
		v.Roe,
		v.Gpr,
		v.Asset,
		v.Debt,
	)
	if err != nil {
		log.Println("[WARN]", "Report.Insert()", "Exec()", err)
		return err
	}

	return nil
}

func SelectFromReportsWhere(where string) ([]Report, error) {
	var reports []Report
	sql := "SELECT id,stock_id,`date`,eps,roe,gpr,asset,debt FROM reports"
	if len(where) > 0 {
		sql = sql + " Where " + where
	}
	rows, err := utils.SqlDB().Query(sql)
	if err != nil {
		log.Println("[WARN]", "model.SelectFromReportsWhere()", "SqlDB().Query()", err)
		return reports, err
	}
	defer rows.Close()

	for rows.Next() {
		var report Report
		err := rows.Scan(
			&report.Id,
			&report.StockId,
			&report.Date,
			&report.Eps,
			&report.Roe,
			&report.Gpr,
			&report.Asset,
			&report.Debt,
		)
		if err != nil {
			log.Println("[WARN]", "model.SelectFromReportsWhere()", "rows.Scan()", err)
			break
		}
		reports = append(reports, report)
	}
	return reports, nil

}
