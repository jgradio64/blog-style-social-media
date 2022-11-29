package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
	"net/http"
	"html/template"
	"time"
)

func postviewHandler(w http.ResponseWriter, r *http.Request) {
	if !checkSession(w, r) {
		backtologin(w, r)
		return
	}
    issueList, err := template.ParseFiles("pages/post_view.html")
    checkError(err)
    
    err = issueList.Execute(w, nil)
    checkError(err)
    
}

func postviewconfirmHandler(w http.ResponseWriter, r *http.Request) {

	if !checkSession(w, r) {
		backtologin(w, r)
		return
	}
	host     := HOST
    port     := PORT
    dbname   := DBNAME
	
	c, _ := r.Cookie("session_token")
	userSession := sessions[c.Value]
	user_name := userSession.username
    user_password := userSession.password
	
	
	target_user_name:= r.FormValue("Username")
	
	start := r.FormValue("Start")
	end:= r.FormValue("End")
	
	if target_user_name == "" {
		target_user_name = user_name
	} else if target_user_name == "all" || target_user_name == "All" || target_user_name == "ALL" {
		target_user_name = "all"
	}
	
	if start == "" {
		y,m,d := time.Now().AddDate(0, 0, -60).Date()
		start = fmt.Sprintf("%04d-%02d-%02d", y,m,d)
	}
	
	if end == "" {
		y,m,d := time.Now().AddDate(0, 0, 1).Date()
		end = fmt.Sprintf("%04d-%02d-%02d", y,m,d)
	}
	
	
	//fmt.Printf("start is %T, and is %s\n", start, start)
	
	// connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user_name, user_password, dbname)
         
    // open database
    db, err := sql.Open("postgres", psqlconn)
    checkError(err)
     
    // close database
    defer db.Close()
	
	// query
	var userconn string
	if target_user_name != "all" {
	userconn = fmt.Sprintf("select \"UserName\",\"title\",  \"post\", \"timestamp\" FROM \"Users\" INNER JOIN \"Posts\" ON \"UserID\" = \"authorID\" where \"UserName\" = '%s' AND \"timestamp\" > timestamp '%s' AND \"timestamp\" <= timestamp '%s'",
		target_user_name, start, end)
		
	} else {
	userconn = fmt.Sprintf("select \"UserName\", \"title\", \"post\", \"timestamp\" FROM \"Users\" INNER JOIN \"Posts\" ON \"UserID\" = \"authorID\" where \"timestamp\" > timestamp '%s' AND \"timestamp\" <= timestamp '%s'",
		start, end)
	}
	//fmt.Println(userconn)
	
	rows, err := db.Query(userconn)
	defer rows.Close()
	
	var new_post post
	index := 0
	
	for rows.Next() {
		
		err = rows.Scan(&new_post.username, &new_post.posttitle, &new_post.userpost, &new_post.posttime)
		checkError(err)
		posts = append(posts, new_post)
		index++
		//fmt.Printf("This is %s %s %s\n", new_post.username, new_post.userpost, new_post.posttime.String())

	}
	
	htmlFile := fmt.Sprintf("<h1>Total %d Posts</h1>", index)
	htmlFile += fmt.Sprintf("<table><tr><th>User Name</th><th>Post Title</th><th>Post Time</th><th>Post Content</th></tr>")
	
	for _, resultpost:= range posts {
		htmlFile += fmt.Sprintf("<tr><th>%s</th><th>%s</th><th>%s</th><th>%s</th></tr>", resultpost.username,resultpost.posttitle, resultpost.posttime.Format(time.RFC3339),resultpost.userpost)
	}
	
	htmlFile += fmt.Sprintf("</table><a href=\"http://localhost:8000/homepage\">Back to Homepage</a>")
	
	refreshSession(w, r)
	fmt.Fprintf(w, htmlFile)
	posts = posts[:0]
	
}
