package model

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"../utils"
)

type Kline struct {
	Id      uint64
	StockId string
	Date    string
	O       string
	C       string
	High    string
	Low     string
	Vol     string
}

func (v *Kline) Insert() error {
	var err error

	stmt, err := utils.SqlDB().Prepare("INSERT INTO klines(stock_id,date,o,c,h,l,v) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		log.Println("[WARN]", "Kline.Insert()", "SqlDB.Prepare()", err);
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		v.StockId,
		v.Date,
		v.O,
		v.C,
		v.High,
		v.Low,
		v.Vol,
	)
	if err != nil {
		log.Println("[WARN]", "Kline.Insert()", "Exec()", err)
		return err
	}

	return nil
}

func DownloadKlinesFromWebsite(stockId string, year string) ([]Kline, error) {
	var klines []Kline
	url := fmt.Sprintf("http://data.gtimg.cn/flashdata/hushen/daily/%s/%s.js", year, stockId)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("[WARN]", "http.Get():", url, err)
		return klines, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		fmt.Println("http.Get():", "404 page not find.")
		return klines, errors.New("HTTP 404.")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll()", err)
		return klines, err
	}

	str := string(body)
	str = strings.Replace(str, "daily_data_"+year+"=\"\\n\\\n", "", -1)
	str = strings.Replace(str, "\";", "", -1)
	rows := strings.Split(str, "\\n\\\n")
	for _, row := range rows {
		fields := strings.Split(row, " ")
		if len(fields) != 6 {
			break
		}
		var kline Kline
		kline.StockId = stockId
		kline.Date = fields[0]
		kline.O = fields[1]
		kline.C = fields[2]
		kline.High = fields[3]
		kline.Low = fields[4]
		kline.Vol = fields[5]
		klines = append(klines, kline)
	}

	return klines, nil
}