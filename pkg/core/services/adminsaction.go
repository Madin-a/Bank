package services

import (
	"database/sql"
	"fmt"
)

func AdminAction(database *sql.DB) (err error) {
	fmt.Println(AdminsFunction)
	var number int64
	fmt.Scan(&number)
	switch number {
	case 1:
		fmt.Println("Enter ATMs address")
		AddNewATM(database)
	case 2:
		fmt.Println("Список банкоматов")
		ShowATMS(database)
	case 3:
		err = AddUser(database)
		return err
	case 4:
		fmt.Println("Список пользователей")
		ShowUsers(database)
	case 5:
		fmt.Println("Список аккаунтов")
		ShowAccount(database)
	case 0:
		fmt.Println("Good Bye")
		return

	}
	return nil
}
