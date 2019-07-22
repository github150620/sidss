package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../utils"
)

type ReportOrigin struct {
	Id          uint32
	Code        string // 股票代码
	Date        string // 报告期
	Jbmgsy      string // 基本每股收益(元)
	Kfmgsy      string // 扣非每股收益(元)
	Xsmgsy      string // 稀释每股收益(元)
	Mgjzc       string // 每股净资产(元)
	Mggjj       string // 每股公积金(元)
	Mgwfply     string // 每股未分配利润(元)
	Mgjyxjl     string // 每股经营现金流(元)
	Yyzsr       string // 营业总收入(元)
	Mlr         string // 毛利润(元)
	Gsjlr       string // 归属净利润(元)
	Kfjlr       string // 扣非净利润(元)
	Yyzsrtbzz   string // 营业总收入同比增长(%)
	Gsjlrtbzz   string // 归属净利润同比增长(%)
	Kfjlrtbzz   string // 扣非净利润同比增长(%)
	Yyzsrgdhbzz string // 营业总收入滚动环比增长(%)
	Gsjlrgdhbzz string // 归属净利润滚动环比增长(%)
	Kfjlrgdhbzz string // 扣非净利润滚动环比增长(%)
	Jqjzcsyl    string // 加权净资产收益率(%)
	Tbjzcsyl    string // 摊薄净资产收益率(%)
	Tbzzcsyl    string // 摊薄总资产收益率(%)
	Mll         string // 毛利率(%)
	Jll         string // 净利率(%)
	Sjsl        string // 实际税率(%)
	Yskyysr     string // 预收款/营业收入
	Xsxjlyysr   string // 销售现金流/营业收入
	Jyxjlyysr   string // 经营现金流/营业收入
	Zzczzy      string // 总资产周转率(次)
	Yszkzzts    string // 应收账款周转天数(天)
	Chzzts      string // 存货周转天数(天)
	Zcfzl       string // 资产负债率(%)
	Ldzczfz     string // 流动负债/总负债(%)
	Ldbl        string // 流动比率
	Sdbl        string // 速动比率
}

func (v *ReportOrigin) Insert() error {
	var err error

	stmt, err := utils.SqlDB().Prepare("INSERT INTO reports_origin(code,date,jbmgsy,kfmgsy,xsmgsy,mgjzc,mggjj,mgwfply,mgjyxjl,yyzsr,mlr,gsjlr,kfjlr,yyzsrtbzz,gsjlrtbzz,kfjlrtbzz,yyzsrgdhbzz,gsjlrgdhbzz,kfjlrgdhbzz,jqjzcsyl,tbjzcsyl,tbzzcsyl,mll,jll,sjsl,yskyysr,xsxjlyysr,jyxjlyysr,zzczzy,yszkzzts,chzzts,zcfzl,ldzczfz,ldbl,sdbl) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Println("[WARN]", "ReportOrigin.Insert():", "SqlDB().Prepare()", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		v.Code,
		v.Date,
		v.Jbmgsy,
		v.Kfmgsy,
		v.Xsmgsy,
		v.Mgjzc,
		v.Mggjj,
		v.Mgwfply,
		v.Mgjyxjl,
		v.Yyzsr,
		v.Mlr,
		v.Gsjlr,
		v.Kfjlr,
		v.Yyzsrtbzz,
		v.Gsjlrtbzz,
		v.Kfjlrtbzz,
		v.Yyzsrgdhbzz,
		v.Gsjlrgdhbzz,
		v.Kfjlrgdhbzz,
		v.Jqjzcsyl,
		v.Tbjzcsyl,
		v.Tbzzcsyl,
		v.Mll,
		v.Jll,
		v.Sjsl,
		v.Yskyysr,
		v.Xsxjlyysr,
		v.Jyxjlyysr,
		v.Zzczzy,
		v.Yszkzzts,
		v.Chzzts,
		v.Zcfzl,
		v.Ldzczfz,
		v.Ldbl,
		v.Sdbl,
	)
	if err != nil {
		log.Println("[WARN]", "ReportOrigin.Insert()", "Exec()", err)
		return err
	}

	return nil
}

func (v *ReportOrigin) Delete() error {
	return nil
}

func (v *ReportOrigin) Update() error {
	return nil
}

func SelectFromReportsOriginWhere(where string) ([]ReportOrigin, error) {
	var reports []ReportOrigin
	sql := "SELECT id,code,date,jbmgsy,mgjzc,jqjzcsyl,mll,zcfzl FROM reports_origin"
	if len(where) > 0 {
		sql = sql + " Where " + where
	}
	sql = sql + " order by code,date"
	rows, err := utils.SqlDB().Query(sql)
	if err != nil {
		log.Println("[ERROR]", "model.SelectFromReports()", "SqlDB().Query()", err)
		return reports, err
	}
	defer rows.Close()

	for rows.Next() {
		var report ReportOrigin
		err := rows.Scan(
			&report.Id,
			&report.Code,
			&report.Date,
			&report.Jbmgsy,
			&report.Mgjzc,
			&report.Jqjzcsyl,
			&report.Mll,
			&report.Zcfzl,
		)
		if err != nil {
			log.Println("[WARN]", "model.SelectFromReports()", "rows.Scan()", err)
			break
		}
		reports = append(reports, report)
	}
	return reports, err
}

func DownloadReportsFromWebsite(code string) ([]ReportOrigin, error) {
	// Url parameter:
	//   type=0 -- quarter
	//   type=1 -- year
	var reports []ReportOrigin

	url := fmt.Sprintf("http://emweb.securities.eastmoney.com/NewFinanceAnalysis/MainTargetAjax?ctype=4&type=0&code=%s", code)
	resp, err := http.Get(url)
	if err != nil {
		return reports, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return reports, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return reports, err
	}

	var rows []map[string]interface{}
	err = json.Unmarshal(body, &rows)
	if err != nil {
		return reports, err
	}
	for _, row := range rows {
		var v ReportOrigin
		v.Code = code
		v.Date, _ = row["date"].(string)
		v.Jbmgsy, _ = row["jbmgsy"].(string)
		v.Kfmgsy, _ = row["kfmgsy"].(string)
		v.Xsmgsy, _ = row["xsmgsy"].(string)
		v.Mgjzc, _ = row["mgjzc"].(string)
		v.Mggjj, _ = row["mggjj"].(string)
		v.Mgwfply, _ = row["mgwfply"].(string)
		v.Mgjyxjl, _ = row["mgjyxjl"].(string)
		v.Yyzsr, _ = row["yyzsr"].(string)
		v.Mlr, _ = row["mlr"].(string)
		v.Gsjlr, _ = row["gsjlr"].(string)
		v.Kfjlr, _ = row["kfjlr"].(string)
		v.Yyzsrtbzz, _ = row["yyzsrtbzz"].(string)
		v.Gsjlrtbzz, _ = row["gsjlrtbzz"].(string)
		v.Kfjlrtbzz, _ = row["kfjlrtbzz"].(string)
		v.Yyzsrgdhbzz, _ = row["yyzsrgdhbzz"].(string)
		v.Gsjlrgdhbzz, _ = row["gsjlrgdhbzz"].(string)
		v.Kfjlrgdhbzz, _ = row["kfjlrgdhbzz"].(string)
		v.Jqjzcsyl, _ = row["jqjzcsyl"].(string)
		v.Tbjzcsyl, _ = row["tbjzcsyl"].(string)
		v.Tbzzcsyl, _ = row["tbzzcsyl"].(string)
		v.Mll, _ = row["mll"].(string)
		v.Jll, _ = row["jll"].(string)
		v.Sjsl, _ = row["sjsl"].(string)
		v.Yskyysr, _ = row["yskyysr"].(string)
		v.Xsxjlyysr, _ = row["xsxjlyysr"].(string)
		v.Jyxjlyysr, _ = row["jyxjlyysr"].(string)
		v.Zzczzy, _ = row["zzczzy"].(string)
		v.Yszkzzts, _ = row["yszkzzts"].(string)
		v.Chzzts, _ = row["chzzts"].(string)
		v.Zcfzl, _ = row["zcfzl"].(string)
		v.Ldzczfz, _ = row["ldzczfz"].(string)
		v.Ldbl, _ = row["ldbl"].(string)
		v.Sdbl, _ = row["sdbl"].(string)
		reports = append(reports, v)
	}
	return reports, nil
}
