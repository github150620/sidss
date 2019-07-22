package model

import (
	"log"

	"../utils"
)

type User struct {
	Id       uint32
	Username string
	Password string
	Mail     string
}

func SelectFromUsersWhere(where string) ([]User, error) {
	var users []User
	sql := "SELECT id,username,password,mail FROM users"
	if len(where) > 0 {
		sql = sql + " Where " + where
	}
	rows, err := utils.SqlDB().Query(sql)
	if err != nil {
		log.Println("[WARN]", "model.SelectFromUsersWhere()", "SqlDB().Query()", err)
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Mail)
		if err != nil {
			log.Println("[WARN]", "model.SelectFromUsersWhere()", "rows.Scan()", err)
			break
		}
		users = append(users, user)
	}
	return users, nil
}
