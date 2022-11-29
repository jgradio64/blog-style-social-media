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
	UserPosts	[]Post	`db:"user_posts"`
}

func (u User) GetUser() User {
	connectionString := extns.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	extns.CheckError(err)

	defer db.Close()

	// Get the user's info
	row := db.QueryRow("select user_id, user_name, about_user from users where user_id=$1", u.UserID)
	var user User
	err = row.Scan(&user.UserID, &user.Username, &user.AboutUser)
	extns.CheckError(err)

	// Get 10 of a users most recent posts to display on their profile.
	queryString := `SELECT post_id, user_id, post_content, post_title, post_date, edit_date
	 FROM posts WHERE posts.user_id=$1 
	 ORDER BY post_date DESC LIMIT 10`
	rows, err := db.Query(queryString, u.UserID)
	defer rows.Close()
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.PostID, &post.UserID, &post.PostContent, &post.PostTitle, &post.PostDate, &post.EditDate)
		extns.CheckError(err)

		post.Author = user.Username
		user.UserPosts = append(user.UserPosts, post)
	}

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

func (u User) GetHomePagePosts() []Post {
	connectionString := extns.GetEnvVariable("CONNECTIONSTRING")
    db, err := sql.Open("postgres", connectionString)
	extns.CheckError(err)
	defer db.Close()
	var followedUserPosts []Post

	queryStatement := `SELECT posts.post_id, posts.user_id, posts.post_content, posts.post_title, posts.post_date, posts.edit_date, users.user_name 
	FROM posts INNER JOIN users ON posts.user_id=users.user_id 
	WHERE posts.user_id IN (SELECT followed_user_id FROM followed_users WHERE user_id=$1)
	ORDER BY post_date DESC LIMIT 10`
	rows, err := db.Query(queryStatement, u.UserID)
	extns.CheckError(err)
	defer rows.Close()

	for rows.Next() {
		var post Post
		err = rows.Scan(&post.PostID, &post.UserID, &post.PostContent, &post.PostTitle, &post.PostDate, &post.EditDate, &post.Author)
		extns.CheckError(err)

		// Gets the number of likes on a post
		numLikesQuery := `SELECT COUNT(user_id) FROM post_likes WHERE post_id=$1`
		row := db.QueryRow(numLikesQuery, post.PostID)
		err = row.Scan(&post.NumOfLikes)
		extns.CheckError(err)

		// Gets the number of comments on a post
		numPostsQuery := `SELECT COUNT(comment_id) FROM comments WHERE post_id=$1`
		row = db.QueryRow(numPostsQuery, &post.PostID)
		err = row.Scan(&post.NumComments)
		extns.CheckError(err)

		followedUserPosts = append(followedUserPosts, post)
	}

	return followedUserPosts
}