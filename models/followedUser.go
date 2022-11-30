package models

import (
	"database/sql"
	
	extns "blbr.com/main/extensions"

    _ "github.com/lib/pq"
)

type FollowedUser struct {
	UserID 			int	`db:"user_id"`
	FollowedUserID 	int	`db:"followed_user_id"`
}

func (flw FollowedUser) GetFollowers() []int {
	connectionString := extns.GetEnvVariable("CONNECTIONSTRING")
	db, err := sql.Open("postgres", connectionString)
	extns.CheckError(err)

	defer db.Close()

	var followerIDs []int

	queryStatement := `SELECT user_id FROM followed_users WHERE followed_user_id=$1`
	rows, err := db.Query(queryStatement, flw.FollowedUserID)
	extns.CheckError(err)
	defer rows.Close()

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		extns.CheckError(err)

		followerIDs = append(followerIDs, id)
	}

	return followerIDs
}

func (flw FollowedUser) CreateFollowedUser() {
	connectionString := extns.GetEnvVariable("CONNECTIONSTRING")
	db, err := sql.Open("postgres", connectionString)
	extns.CheckError(err)

	defer db.Close()

	insertStatement := `INSERT INTO "followed_users"("user_id","followed_user_id") values($1,$2)`
	_, err = db.Exec(insertStatement, flw.UserID, flw.FollowedUserID)
	extns.CheckError(err)
}

func (flw FollowedUser) DeleteFollowedUser() {
	connectionString := extns.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	extns.CheckError(err)

	defer db.Close()

	insertStatement := `DELETE FROM followed_users WHERE user_id=$1 AND followed_user_id=$2`
	_, err = db.Exec(insertStatement, flw.UserID, flw.FollowedUserID)
	extns.CheckError(err)
}