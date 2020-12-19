package db

const CreateUsersAccount = `Create table if not exists users(
	id integer primary key autoincrement,
	name text not null,
	surname text not null,
	age integer not null,
	gender text not null,
	login text not null unique,
	password text not null,
	status text not null,
	remove boolean not null default false
);`
const CreateATMsTAble = `Create table if not exists atms(
	id integer primary key autoincrement,
	name text not null,
	status boolean not null default true
);`
const CreateAccountTable = `Create table if not exists account (
	id integer primary key autoincrement,
	userId integer references users(id) not null,
	Number integer not null,
	Amount integer not null,
	Currency integer not null
);`
const CreateTransactionTable = `create table if not exists archive(
	id integer primary key autoincrement,
	date text not null,
	time text not null,
	amount integer not null,
	giverID integer not null,
	gainerID integer not null,
	availableLimit integer not null
)`

