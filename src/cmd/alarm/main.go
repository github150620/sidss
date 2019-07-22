package main

import (
	"log"
	"time"

	"../../model"
	"../../utils"
)

func main() {
	log.Println("[INFO]", "Start...")
	for {
		CheckAndSendMail()
		t := time.Now().Format("15:04:05")
		if t < "09:00:00" {
			break
		} else if t > "12:00:00" && t < "13:00:00" {
			break
		} else if t > "15:00:00" {
			break
		}
		log.Println("[INFO]", "Wait 10 seconds...")
		time.Sleep(time.Duration(10) * time.Second)
	}
	log.Println("[INFO]", "Over.")
}

func CheckAndSendMail() {
	alarms, err := model.SelectFromAlarmsWhere("status=0")
	if err != nil {
		log.Println("[ERROR]", err)
	}
	
	for _, alarm := range alarms {
		log.Println("[INFO]", "Alarm sql: ", alarm.Sql)
		cnt, err := alarm.Check()
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		if cnt == 0 {
			continue
		}

		log.Println("[INFO]", "count(1)>=1")
		alarm.Status = 1
		err = alarm.Update()
		if err != nil {
			continue
		}

		log.Println("[INFO] Send alarm mail.")
		utils.SendMail("[SIDSS]" + alarm.StockId + " " + alarm.Conditions, alarm.Sql)
	}
}
