package models

import (
	"BankHumo/db"
	"database/sql"
	"fmt"
	"time"
)

type Transaction struct {
	ID             int64
	Date           string
	Time           string
	Amount         int64
	GiverID        int64
	GainerID       int64
	AvailableLimit int64
}

func AddTransaction(tx *sql.Tx, giverID, gainerID, operationAmount, newAmount int64) (err error) {
	err = nil
	var check Transaction
	data := time.Now()
	check.Date = data.Format("02-Jan-2006")
	check.Time = data.Format("15:04")
	check.Amount = operationAmount
	check.GiverID = giverID
	check.AvailableLimit = newAmount
	check.GainerID = gainerID
	_, err = tx.Exec(db.AddTransaction, check.Date, check.Time, check.Amount, check.GiverID, check.GainerID, check.AvailableLimit)
	if err != nil {
		//panic(err)
		fmt.Println("ошибка при добавлении в архив")
		return err
	}
	return err
}
