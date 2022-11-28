package models

import (
	"database/sql"

	extns "blbr.com/main/extensions"

    _ "github.com/lib/pq"
)

type BlockedUser struct {
	UserID 			int	`db:"user_id"`
	BlockedUserID 	int `db:"blocked_user_id"`
}

func (blu BlockedUser) CreateBlockedUser() {
	connectionString := extns.GetEnvVariable("CONNECTIONSTRING")
	db, err := sql.Open("postgres", connectionString)
	extns.CheckError(err)

	defer db.Close()

	insertStatement := `INSERT INTO "blocked_users"("user_id","blocked_user_id") values($1,$2)`
	_, err = db.Exec(insertStatement, blu.UserID, blu.BlockedUserID)
	extns.CheckError(err)
}

func (blu BlockedUser) DeleteBlockedUser() {
	connectionString := extns.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	extns.CheckError(err)

	defer db.Close()

	insertStatement := `DELETE FROM blocked_users WHERE user_id=$1 AND blocked_user_id=$2`
	_, err = db.Exec(insertStatement, blu.UserID, blu.BlockedUserID)
	extns.CheckError(err)
}