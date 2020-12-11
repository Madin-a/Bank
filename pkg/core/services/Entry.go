package services

import (
	"BankHumo/models"
	"database/sql"
	"fmt"
)

const UserFunction = `
1.Проверка баланса
2.Перевод денег
3.Список банкоматов
4.История транзакций
0.Выход
`
const AdminsFunction = `
1.Добавить банкомат
2.Список банкоматов
3.Добавить нового пользователя
4.Список пользователей
5.Список аккаунтов
0.Выйти
`
const LoginOperation = `Введите логин и пароль: `

func Entry(database *sql.DB) (login, password string) {
	fmt.Println(LoginOperation)
	fmt.Println("login: ")
	fmt.Scan(&login)
	fmt.Println("password")
	fmt.Scan(&password)
	return login, password
}

func Login(database *sql.DB, login, password string) models.User {
	var User models.User
	_ = database.QueryRow(`Select *from users where login = ($1) and password = ($2) `, login, password).Scan(
		&User.ID,
		&User.Name,
		&User.Surname,
		&User.Age,
		&User.Gender,
		&User.Login,
		&User.Password,
		&User.Status,
		&User.Remove)

	fmt.Println(User)
	//fmt.Println(User.ID)
	return User
}
