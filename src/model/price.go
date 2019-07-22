package model

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"../utils"
)

type Price struct {
	StockId string
	Price   float64
	Chg     float64
	High    float64
	Low     float64
	TtmPe	float64
}

func (v *Price) Update() error {
	var err error

	sql := "UPDATE prices SET price=?,chg=?,high=?,low=?,ttm_pe=? WHERE stock_id=?"
	stmt, err := utils.SqlDB().Prepare(sql)
	if err != nil {
		log.Println("[ERROR]", "Price.Update()", "SqlDB().Prepare()", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		v.Price,
		v.Chg,
		v.High,
		v.Low,
		v.TtmPe,
		v.StockId,
	)
	if err != nil {
		log.Println("[WARN]", "Price.Update()", "Exec()", err)
		return err
	}
	return nil
}

func SelectFromPricesWhere(where string) ([]Price, error) {
	var prices []Price
	sql := "SELECT stock_id,price,chg,high,low,ttm_pe FROM prices"
	if len(where) > 0 {
		sql = sql + " Where " + where
	}
	rows, err := utils.SqlDB().Query(sql)
	if err != nil {
		log.Println("[WARN]", "model.SelectFromPricesWhere()", "SqlDB().Query()", err)
		return prices, err
	}
	defer rows.Close()

	for rows.Next() {
		var price Price
		err := rows.Scan(
			&price.StockId,
			&price.Price,
			&price.Chg,
			&price.High,
			&price.Low,
			&price.TtmPe,
		)
		if err != nil {
			log.Println("[WARN]", "model.SelectFromPricesWhere()", "rows.Scan()", err)
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func DownloadPricesFromWebsite(stocks []string) ([]Price, error) {
	var prices []Price
	//"http://hq.sinajs.cn/list=sh601006"
	url := "http://qt.gtimg.cn/q="
	for _, code := range stocks {
		url += code + ","
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Println("[WARN]", "http.Get():", url, err)
		return prices, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		log.Println("http.Get():", "404 page not find.")
		return prices, errors.New("HTTP 404.")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll()", err)
		return prices, err
	}

	str := string(body)
	rows := strings.Split(str, ";")
	for _, row := range rows {
		start := strings.Index(row, "\"")
		end := strings.LastIndex(row, "\"")
		if start >= 0 && end >= 0 && start < end {
			var price Price

			fields := strings.Split(row[start+1:end], "~")
			price.StockId = AddPrefix(fields[2])
			price.Price, _ = strconv.ParseFloat(fields[3], 64)
			price.Chg, _ = strconv.ParseFloat(fields[32], 64)
			price.High, _ = strconv.ParseFloat(fields[33], 64)
			price.Low, _ = strconv.ParseFloat(fields[34], 64)
			price.TtmPe, _ = strconv.ParseFloat(fields[39], 64)
			prices = append(prices, price)
		}
	}
	return prices, nil
}

func AddPrefix(symbol string) string {
	if symbol[0:2] == "60" {
		return "sh" + symbol
	} else if symbol[0:2] == "00" {
		return "sz" + symbol
	} else if symbol[0:2] == "13" {
		return "sz" + symbol
	} else if symbol[0:2] == "20" {
		return "sh" + symbol
	} else if symbol[0:2] == "30" {
		return "sz" + symbol
	} else {
		log.Println("ao")
		return symbol
	}
}
