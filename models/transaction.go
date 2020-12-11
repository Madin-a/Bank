package models

import (
	"BankHumo/db"
	"database/sql"
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

func AddTransaction(Db *sql.DB, giverID, gainerID, operationAmount, newAmount int64) (err error) {
	var check Transaction
	data := time.Now()
	check.Date = data.Format("02-Jan-2006")
	check.Time = data.Format("15:04")
	check.Amount = operationAmount
	check.GiverID = giverID
	check.AvailableLimit = newAmount
	check.GainerID = gainerID
	_, err = Db.Exec(db.AddTransaction, check.Date, check.Time, check.Amount, check.GiverID, check.GainerID, check.AvailableLimit)
	if err != nil {
		panic(err)
	}
	return
}
