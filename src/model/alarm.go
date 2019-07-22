package model

import (
	"log"

	"../utils"
)

type Alarm struct {
	Id      uint32
	UserId  uint32
	StockId string
	Conditions string
	Sql     string
	Period  string
	Status  uint32
}

func SelectFromAlarmsWhere(where string) ([]Alarm, error) {
	var alarms []Alarm
	sql := "SELECT id,user_id,stock_id,conditions,`sql`,period,status FROM alarms"
	if len(where) > 0 {
		sql = sql + " Where " + where
	}
	rows, err := utils.SqlDB().Query(sql)
	if err != nil {
		log.Println("[WARN]", "model.SelectFromAlarmsWhere()", "SqlDB().Query()", err)
		return alarms, err
	}
	defer rows.Close()

	for rows.Next() {
		var alarm Alarm
		err := rows.Scan(
			&alarm.Id,
			&alarm.UserId,
			&alarm.StockId,
			&alarm.Conditions,
			&alarm.Sql,
			&alarm.Period,
			&alarm.Status,
		)
		if err != nil {
			log.Println("[WARN]", "model.SelectFromAlarmsWhere()", "rows.Scan()", err)
			break
		}
		alarms = append(alarms, alarm)
	}
	return alarms, nil
}

func (v *Alarm) Check() (int, error) {
	var cnt int

	rows, err := utils.SqlDB().Query(v.Sql)
	if err != nil {
		log.Println("[WARN]", "model.Check()", "SqlDB().Query()", err)
		return cnt, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&cnt)
		if err != nil {
			log.Println("[ERROR]", "model.Check()", "rows.Scan()", err)
			return cnt, err
		}
	}

	return cnt, nil
}

func (v *Alarm) Update() error {
	var err error

	sql := "UPDATE alarms SET status=? WHERE id=?"
	stmt, err := utils.SqlDB().Prepare(sql)
	if err != nil {
		log.Println("[ERROR]", "Alarm.Update()", "SqlDB().Prepare()", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		v.Status,
		v.Id,
	)
	if err != nil {
		log.Println("[WARN]", "Alarm.Update()", "Exec()", err)
		return err
	}
	return nil
}
