package models

import (
	"database/sql"

	extns "blbr.com/main/extensions"

    _ "github.com/lib/pq"
)

type PostLike struct {
	UserID 	int		`db:user_id`
	PostID 	int		`db:post_id`
}

func (pl PostLike) CreatePostLike() {
	connectionString := extns.GetEnvVariable("CONNECTIONSTRING")
	db, err := sql.Open("postgres", connectionString)
	extns.CheckError(err)

	defer db.Close()

	insertStatement := `INSERT INTO "post_likes"("user_id","post_id") values($1,$2)`
	_, err = db.Exec(insertStatement, pl.UserID, pl.PostID)
	extns.CheckError(err)
}

func (pl PostLike) DeletePostLike() {
	connectionString := extns.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	extns.CheckError(err)

	defer db.Close()

	insertStatement := `DELETE FROM post_likes WHERE user_id=$1 AND post_id=$2`
	_, err = db.Exec(insertStatement, pl.UserID, pl.PostID)
	extns.CheckError(err)
}