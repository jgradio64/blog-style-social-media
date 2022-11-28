package models

import (
	"database/sql"
	"time"

	extensions "blbr.com/main/extensions"

    _ "github.com/lib/pq"
)

type Comment struct {
	CommentID	int 			`db:"comment_id`
	UserID 		int				`db:"user_id`
	PostID 		int				`db:"post_id`
	Content 	string			`db:"comment_content`
	DateCreated	time.Time 		`db:"date_created"`
	EditDate	sql.NullTime 	`db:"date_edited"`
	NumOfLikes	int				`db:"number_likes"`
	Author		string			`db:"user_name"`
}

func (c Comment) GetComment() Comment {
	connectionString := extensions.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	CheckError(err)

	defer db.Close()

	queryStatement := `SELECT users.user_name, comments.user_id, comments.comment_content, comments.date_created, comments.date_edited 
	FROM comments 
	INNER JOIN users ON comments.user_id=users.user_id 
	WHERE comment_id=$1`
	rows, err := db.Query(queryStatement, c.CommentID)
	CheckError(err)

	var comment Comment
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&comment.Author, &comment.UserID, &comment.Content, &comment.DateCreated, &comment.EditDate)
		CheckError(err)
	}

	// Gets the number of likes on a Comment
	numLikesQuery := `SELECT COUNT(user_id) FROM comment_likes WHERE comment_id=$1`
	row := db.QueryRow(numLikesQuery, c.CommentID)
	err = row.Scan(&comment.NumOfLikes)
	CheckError(err)

	return comment
}

func (c Comment) CreateComment() {
	connectionString := extensions.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	CheckError(err)

	defer db.Close()

	insertStatement := `INSERT INTO "comments"("user_id", "post_id", "comment_content", "date_created") values($1, $2, $3, $4)`
	_, err = db.Query(insertStatement, c.UserID, c.PostID, c.Content, time.Now())
	CheckError(err)
}

func (c Comment) UpdateComment() {
	connectionString := extensions.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	CheckError(err)

	defer db.Close()

	updateStatement := `UPDATE comments SET comment_content=$1, date_edited=$2 WHERE comment_id=$3`
	_, err = db.Exec(updateStatement, c.Content, time.Now(), c.CommentID)
	CheckError(err)
}

func (c Comment) DeleteComment() {
	connectionString := extensions.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	CheckError(err)

	defer db.Close()

	deleteStatement := `DELETE FROM comments WHERE comment_id=$1`
	_, err = db.Exec(deleteStatement, c.CommentID)
	CheckError(err)
}