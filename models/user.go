package models

import (
	"database/sql"
	"fmt"
	"os"
	extensions "blbr.com/main/extensions"

    _ "github.com/lib/pq"
)

type User struct {
	UserID 		int		`db:"user_id"`
	Username 	string	`db:"user_name"`
	Password 	string	`db:"password"`
	AboutUser 	string	`db:"about_user"`
}

func (u User) GetUser() User {
	connectionString := extensions.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	CheckError(err)

	defer db.Close()

	rows, err := db.Query("select * from users where user_id=$1", u.UserID)
	CheckError(err)

	var user User
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.UserID, &user.Username, &user.Password, &user.AboutUser)
		CheckError(err)
	}

	return user
}

func (u User) CreateUser() {
	connectionString := extensions.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	CheckError(err)

	defer db.Close()

	insertStatement := `INSERT INTO "users"("user_name","password","about_user") values($1,$2,$3)`
	_, err = db.Exec(insertStatement, u.Username, u.Password, u.AboutUser)
	CheckError(err)
}

func (u User) UpdateUser() {
	connectionString := extensions.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	CheckError(err)

	defer db.Close()

	updateStatement := `UPDATE users SET about_user=$1 WHERE user_id=$2`
	_, err = db.Exec(updateStatement, u.AboutUser, u.UserID)
	CheckError(err)
}

func (u User) DeleteUser() {
	connectionString := extensions.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	CheckError(err)
	defer db.Close()

	deleteStatement := `DELETE FROM users WHERE user_id=$1`
	_, err = db.Exec(deleteStatement, u.UserID)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}
