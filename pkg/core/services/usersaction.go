package services

import (
	"BankHumo/models"
	"database/sql"
	"fmt"
)

func UserAction(database *sql.DB, user models.User) (err error) {
	fmt.Println(UserFunction)
	var number int64
	fmt.Scan(&number)
	switch number {
	case 1:
		fmt.Println("Ваш баланс: ")
		CheckBalance(database, user.ID)
	case 2:
		fmt.Println("Перевести ")
		var gainerId, summa int64
		fmt.Println("введиту сумму")
		fmt.Scan(&summa)
		fmt.Println("введите ID получателя")
		fmt.Scan(&gainerId)
		err := Translation(database, user.ID, gainerId, summa)
		if err != nil {
			fmt.Println("ошибка при переводе")
		} else {
			fmt.Println("успешно переведено")
		}
		return err

	case 3:
		fmt.Println("Cписок банкоматов")
		ShowATMS(database)
	case 4:
		fmt.Println("История транзакций")
		Archive(database, user.ID)
	case 0:
		fmt.Println("Good Bye")
	}
	return err
}
