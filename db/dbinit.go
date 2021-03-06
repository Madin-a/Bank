package db

import (
	"database/sql"
	"log"
)

func DbInit(database *sql.DB) {
	DDLs := []string{CreateUsersAccount, CreateATMsTAble, CreateAccountTable, CreateTransactionTable}
	for _, ddl := range DDLs {
		_, err := database.Exec(ddl)
		if err != nil {
			log.Fatalf("Can't init %s err is %e", ddl, err)

		}
	}
}
