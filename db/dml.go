package db

const AddNewATM = `insert into atms(name) values ($1) `
const AddNewUser = `insert into users(name, surname, age, gender, login, password, status, remove) VALUES (($1),($2),($3),($4),($5), ($6), ($7), ($8))`
const AddTransaction = `insert into archive(date, time, amount, giverID, gainerID, availableLimit) values(($1),($2),($3),($4),($5),($6))`

const CheckBalance = `SELECT * FROM account WHERE userId = ($1)`

const GetAmountByID = `select * from account where UserID = ($1)`
const UpdateAmountByID = `update account set amount = ($1) where UserID = ($2)`

const SelectATMs = `select * from atms`
const SelectUsers = `select * from users`
const SelectAccounts = `select * from account`

const Archive = `select * from archive where giverID = ($1)`
