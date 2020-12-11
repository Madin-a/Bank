package main

import (
	"BankHumo/db"
	"BankHumo/pkg/core/services"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Start(database *sql.DB) {
	for {
		login, password := services.Entry(database)
		user := services.Login(database, login, password)
		if user.Status == "admin" {
			fmt.Println(" You are admin")
			err := services.AdminAction(database)
			if err != nil {
				fmt.Println(" = ", err)
			}
		} else if user.Status == "user" {
			fmt.Println(" You are user")
			err := services.UserAction(database, user)
			fmt.Println("err = ", err)
		}

	}
}
func main() {
	database, err := sql.Open("sqlite3", "bank")
	if err != nil {
		log.Fatalf("Err is %e")
	}
	db.DbInit(database)
	Start(database)
}
