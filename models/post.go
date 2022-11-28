package models

import (
	"database/sql"
	"time"

	extensions "blbr.com/main/extensions"

    _ "github.com/lib/pq"
)

type Post struct {
	UserID 		int				`db:"user_id"`
	PostID 		int				`db:"post_id"`
	PostContent string			`db:"post_content"`
	PostTitle 	string			`db:"post_title"`
	PostDate 	time.Time		`db:"post_date"`
	EditDate 	sql.NullTime	`db:"edit_date"`
}

func (p Post) GetPost() Post {
	connectionString := extensions.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	CheckError(err)

	defer db.Close()

	queryStatement := `select * from posts where post_id=$1`
	rows, err := db.Query(queryStatement, p.PostID)
	CheckError(err)

	var post Post
	
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&post.PostID, &post.PostContent, &post.PostDate, &post.EditDate, &post.UserID, &post.PostTitle)
		CheckError(err)
	}

	return post
}

func (p Post) CreatePost() {
	connectionString := extensions.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	CheckError(err)

	defer db.Close()

	insertStatement :=  `INSERT INTO "posts"("user_id", "post_content", "post_title", "post_date") values($1, $2, $3, $4)`
	_, err = db.Exec(insertStatement, p.UserID, p.PostContent, p.PostTitle, time.Now())
	CheckError(err)
}

func (p Post) UpdatePost() {
	connectionString := extensions.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	CheckError(err)

	defer db.Close()

	updateStatement := `UPDATE posts SET post_content=$1, edit_date=$2 WHERE post_id=$3`
	_, err = db.Exec(updateStatement, p.PostContent, time.Now(), p.PostID)
	CheckError(err)
}

func (p Post) DeletePost() {
	connectionString := extensions.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	CheckError(err)

	defer db.Close()

	deleteStatement := `DELETE FROM posts WHERE post_id=$1`
	_, err = db.Exec(deleteStatement, p.PostID)
	CheckError(err)
}