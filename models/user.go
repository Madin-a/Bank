package models

import (
	"BankHumo/db"
	"database/sql"
	"fmt"
)

type User struct {
	ID       int64
	Name     string
	Surname  string
	Age      int64
	Gender   string
	Login    string
	Password string
	Status   string
	Remove   bool
}

func AddUserToDB(database *sql.DB, name string, surname string, age int64, gender string, login string, password string, status string, remove bool) (ok bool, err error) {
	_, err = database.Exec(db.AddNewUser, name, surname, age, gender, login, password, status, remove)
	if err != nil {
		fmt.Println(`Can't add new user`, err)
		return false, err
	}
	return true, nil
}
