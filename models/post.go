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
	Author		string			`db:"user_name"`
	NumOfLikes	int				`db:"number_likes"`
	NumComments	int				`db:"number_comments"`
	Comments	[]Comment		`db:"comments"`
}

func (p Post) GetPost() Post {
	connectionString := extensions.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	CheckError(err)

	defer db.Close()

	queryStatement := `SELECT posts.user_id, posts.post_content, posts.post_title, posts.post_date, posts.edit_date, users.user_name 
	FROM posts 
	INNER JOIN users ON posts.user_id=users.user_id 
	WHERE post_id=$1`

	rows, err := db.Query(queryStatement, p.PostID)
	CheckError(err)

	var post Post
	
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&post.UserID, &post.PostContent, &post.PostTitle, &post.PostDate, &post.EditDate, &post.Author)
		CheckError(err)
	}

	// Gets the number of likes on a post
	numLikesQuery := `SELECT COUNT(user_id) FROM post_likes WHERE post_id=$1`
	row := db.QueryRow(numLikesQuery, p.PostID)
	err = row.Scan(&post.NumOfLikes)
	CheckError(err)

	// Gets the number of comments on a post
	numPostsQuery := `SELECT COUNT(comment_id) FROM comments WHERE post_id=$1`
	row = db.QueryRow(numPostsQuery, p.PostID)
	err = row.Scan(&post.NumComments)
	CheckError(err)

	// Gets the information of the comments on a post, man this one was complex
	commentsStatement := `SELECT users.user_name, comments.comment_id, comments.user_id, comments.comment_content, comments.date_created, comments.date_edited 
	FROM comments 
	INNER JOIN users ON comments.user_id=users.user_id 
	WHERE comments.post_id=$1`
	rows, err = db.Query(commentsStatement, p.PostID)
	for rows.Next() {
		var com Comment
		err = rows.Scan(&com.Author, &com.CommentID, &com.UserID, &com.Content, &com.DateCreated, &com.EditDate)
		CheckError(err)
		post.Comments = append(post.Comments, com)
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