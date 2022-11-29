package models

import (
	"database/sql"

	extns "blbr.com/main/extensions"

    _ "github.com/lib/pq"
)

type CommentLike struct {
	CommentID 	int	`db:comment_id`
	UserID 		int	`db:user_id`
}


func (cl CommentLike) CreateCommentLike() {
	connectionString := extns.GetEnvVariable("CONNECTIONSTRING")
	db, err := sql.Open("postgres", connectionString)
	extns.CheckError(err)

	defer db.Close()

	insertStatement := `INSERT INTO "comment_likes"("user_id","comment_id") values($1,$2)`
	_, err = db.Exec(insertStatement, cl.UserID, cl.CommentID)
	extns.CheckError(err)
}

func (cl CommentLike) DeleteCommentLike() {
	connectionString := extns.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	extns.CheckError(err)

	defer db.Close()

	insertStatement := `DELETE FROM comment_likes WHERE user_id=$1 AND comment_id=$2`
	_, err = db.Exec(insertStatement, cl.UserID, cl.CommentID)
	extns.CheckError(err)
}