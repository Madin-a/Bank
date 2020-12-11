package models

import (
	"BankHumo/db"
	"database/sql"
	"fmt"
)

type ATMs struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

func AddATMs(database *sql.DB, Address string) (ok bool, err error) {
	_, err = database.Exec(db.AddNewATM, Address)
	if err != nil {
		fmt.Println(`Can't insert to ATMs table new address, err is`, err)
		return false, err
	}
	return true, nil
}
