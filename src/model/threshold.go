package model

import (
	"log"

	"../utils"
)

type Threshold struct {
	Id             uint32
	UserId         uint32
	TableName      string
	PkName         string
	PkValue        string
	ThresholdName  string
	ThresholdValue string
	Operator       string
}

func SelectFromThresholdsWhere(where string) ([]Threshold, error) {
	var thresholds []Threshold
	sql := "SELECT id,user_id,table_name,pk_name,pk_value,threshold_name,threshold_value,operator FROM thresholds"
	if len(where) > 0 {
		sql = sql + " Where " + where
	}
	rows, err := utils.SqlDB().Query(sql)
	if err != nil {
		log.Println("[WARN]", "model.SelectFromThresholdsWhere()", "SqlDB().Query()", err)
		return thresholds, err
	}
	defer rows.Close()

	for rows.Next() {
		var threshold Threshold
		err := rows.Scan(
			&threshold.Id,
			&threshold.UserId,
			&threshold.TableName,
			&threshold.PkName,
			&threshold.PkValue,
			&threshold.ThresholdName,
			&threshold.ThresholdValue,
			&threshold.Operator,
		)
		if err != nil {
			log.Println("[WARN]", "model.SelectFromThresholdsWhere()", "rows.Scan()", err)
			break
		}
		thresholds = append(thresholds, threshold)
	}
	return thresholds, nil
}

func (v *Threshold) Check() {
	sql := "SELECT * FROM " + v.TableName + " WHERE " + v.PkName + "='" + v.PkValue + "' AND " + v.ThresholdName + v.Operator + v.ThresholdValue
	log.Println("[INFO]", sql)
}
