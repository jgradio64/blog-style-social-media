package models

import (
	"database/sql"
	extns "blbr.com/main/extensions"

    _ "github.com/lib/pq"
)

var connectionString = extns.GetEnvVariable("CONNECTIONSTRING")

type User struct {
	UserID 		int		`db:"user_id"`
	Username 	string	`db:"user_name"`
	Password 	string	`db:"password"`
	AboutUser 	string	`db:"about_user"`
}

func (u User) GetUser() User {
	connectionString := extns.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	extns.CheckError(err)

	defer db.Close()

	row := db.QueryRow("select user_id, user_name, about_user from users where user_id=$1", u.UserID)
	var user User
	err = row.Scan(&user.UserID, &user.Username, &user.AboutUser)
	extns.CheckError(err)

	return user
}

func (u User) CreateUser() {
	connectionString := extns.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	extns.CheckError(err)

	defer db.Close()

	insertStatement := `INSERT INTO "users"("user_name","password","about_user") values($1,$2,$3)`
	_, err = db.Exec(insertStatement, u.Username, u.Password, u.AboutUser)
	extns.CheckError(err)
}

func (u User) UpdateUser() {
	connectionString := extns.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	extns.CheckError(err)

	defer db.Close()

	updateStatement := `UPDATE users SET about_user=$1 WHERE user_id=$2`
	_, err = db.Exec(updateStatement, u.AboutUser, u.UserID)
	extns.CheckError(err)
}

func (u User) DeleteUser() {
	connectionString := extns.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	extns.CheckError(err)
	defer db.Close()

	deleteStatement := `DELETE FROM users WHERE user_id=$1`
	_, err = db.Exec(deleteStatement, u.UserID)
	extns.CheckError(err)
}