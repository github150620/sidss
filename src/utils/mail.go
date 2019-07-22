package utils

import (
	"log"
	"net/smtp"
	"strings"
)

var (
	addr = "smtp.163.com:25"
	auth = smtp.PlainAuth("", "roc1005", "123456", "smtp.163.com")
	from = "userA@163.com"
	to   = []string{"userB@139.com"}
)

func SendMail(subject, content string) {
	msg := "To: " + strings.Join(to, ",")
	msg += "\r\nFrom: " + from
	msg += "\r\nSubject: " + subject
	msg += "\r\n\r\n" + content
	err := smtp.SendMail(addr, auth, from, to, []byte(msg))
	if err != nil {
		log.Println("[ERROR]", err)
		return
	}
}
