package services

import (
	"BankHumo/db"
	"BankHumo/models"
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
)

func CheckBalance(database *sql.DB, userId int64) {
	var account models.Account
	database.QueryRow(db.CheckBalance, userId).Scan(&account.ID, &account.UserID, &account.Number, &account.Amount, &account.Currency)
	fmt.Println(account.Amount)
}

func Translation(database *sql.DB, giverID int64, gainerID int64, summa int64) (err error) {
	var giverAccount, gainerAccount models.Account
	database.QueryRow(db.GetAmountByID, giverID).Scan(&giverAccount.ID, &giverAccount.UserID, &giverAccount.Number, &giverAccount.Amount, &giverAccount.Currency)
	database.QueryRow(db.GetAmountByID, gainerID).Scan(&gainerAccount.ID, &gainerAccount.UserID, &gainerAccount.Number, &gainerAccount.Amount, &gainerAccount.Currency)
	if summa < 0 {
		err = errors.New("сумма должна быть >= 0")
		return err
	}
	if giverAccount.Amount >= summa {
		gainerAccount.Amount = gainerAccount.Amount + summa
		database.Exec(db.UpdateAmountByID, gainerAccount.Amount, gainerID)
		giverAccount.Amount = giverAccount.Amount - summa
		database.Exec(db.UpdateAmountByID, giverAccount.Amount, giverID)
		err = errors.New("недостаточно средств")
	}
	models.AddTransaction(database, giverID, gainerID, summa, giverAccount.Amount)
	return err
}

func ShowATMS(database *sql.DB) {
	rows, err := database.Query(db.SelectATMs)
	if err != nil {
		fmt.Println("ошибка", err)
	}
	atms := []models.ATMs{}
	for rows.Next() {
		atm := models.ATMs{}
		err := rows.Scan(
			&atm.ID, &atm.Name, &atm.Status)
		if err != nil {
			fmt.Println("ошибка", err)
			continue
		}
		atms = append(atms, atm)
	}
	for _, atm := range atms {
		fmt.Println("ID", atm.ID, "Name ", atm.Name, "Status", atm.Status)
	}
}

func ShowUsers(database *sql.DB) {
	rows, err := database.Query(db.SelectUsers)
	if err != nil {
		fmt.Println("ошибка", err)
	}
	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(
			&user.ID, &user.Name, &user.Surname, &user.Age, &user.Gender, &user.Login, &user.Password, &user.Status, &user.Remove)
		if err != nil {
			fmt.Println("ошибка", err)
			continue
		}
		users = append(users, user)
	}
	for _, user := range users {
		fmt.Println("ID ", user.ID, "Name ", user.Name, "Surname ", user.Surname, "Age ", user.Age, "Gender", user.Gender, "Status", user.Status, "Remove", user.Remove)
	}
}
func ShowAccount(database *sql.DB) {
	rows, err := database.Query(db.SelectAccounts)
	if err != nil {
		fmt.Println("ошибка", err)
	}
	accounts := []models.Account{}
	for rows.Next() {
		account := models.Account{}
		err := rows.Scan(
			&account.ID, &account.UserID, &account.Number, &account.Amount, &account.Currency)
		if err != nil {
			fmt.Println("ошибка", err)
			continue
		}
		accounts = append(accounts, account)
	}
	for _, account := range accounts {
		fmt.Println("ID ", account.ID, "UserID ", account.UserID, "Number", account.Number, "Amount", account.Amount, "Currency", account.Currency)
	}
}

func AddNewATM(database *sql.DB) (ok bool) {
	var s string
	fmt.Scan(&s)
	reader := bufio.NewReader(os.Stdin)
	Address, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf(`can't read command: %v`, err)
	}
	fmt.Println(s)
	sprintf := fmt.Sprintf(`%s %s`, s, Address)
	fmt.Println(sprintf)
	_, err = models.AddATMs(database, sprintf)
	if err != nil {
		fmt.Println(`vse ploxo`, err)
	} else {
		fmt.Println(`vse ok`)
	}
	return
}

func AddUser(database *sql.DB) (err error) {
	var name, surname, gender, login, password, status string
	var age int64
	var remove bool
	fmt.Println("Добавить нового пользователя")
	fmt.Println("name: ")
	fmt.Scan(&name)
	fmt.Println("surname: ")
	fmt.Scan(&surname)
	fmt.Println("age: ")
	fmt.Scan(&age)
	fmt.Println("gender: ")
	fmt.Scan(&gender)
	fmt.Println("login: ")
	fmt.Scan(&login)
	fmt.Println("password")
	fmt.Scan(&password)
	fmt.Println("status: ")
	fmt.Scan(&status)
	fmt.Println("remove: ")
	fmt.Scan(&remove)
	if err != nil {
		fmt.Println("owibka", err)
		return err
	}
	models.AddUserToDB(database, name, surname, age, gender, login, password, status, remove)
	return
}
func Archive(database *sql.DB, giver int64) {
	transaction := models.Transaction{}

	rows, err := database.Query(db.Archive, giver)
	if err != nil {
		log.Println(err, `users are not selected`)
		defer rows.Close()
	}

	for rows.Next() {
		err = rows.Scan(
			&transaction.ID,
			&transaction.Date,
			&transaction.Time,
			&transaction.Amount,
			&transaction.GiverID,
			&transaction.GainerID,
			&transaction.AvailableLimit,
		)
		if err != nil {
			log.Println(err, ` not selected archive`)
		}
		fmt.Println("ID ", transaction.ID, "Date", transaction.Date, "Time", transaction.Time, "Amount ", transaction.Amount, "GiverID", transaction.GiverID, "GainerID", transaction.GainerID, "AvailableLimit", transaction.AvailableLimit)
	}
}
