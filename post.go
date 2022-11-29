package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
	"net/http"
	"html/template"
	"time"
)



func postHandler(w http.ResponseWriter, r *http.Request) {

	if !checkSession(w, r) {
		backtologin(w, r)
		return
	}
    issueList, err := template.ParseFiles("pages/user_post.html")
    checkError(err)
	
	//c , _ := r.Cookie("session_token")
    //fmt.Println(sessions[c.Value].username)
	
    err = issueList.Execute(w, nil)
    checkError(err)
}



func postconfirmHandler (w http.ResponseWriter, r *http.Request) {
	if !checkSession(w, r) {
		backtologin(w, r)
		return
	}
	host     := HOST
    port     := PORT
    dbname   := DBNAME

    title     := r.FormValue("Title")
    content := r.FormValue("Content")	
	
	c, _ := r.Cookie("session_token")

	userSession := sessions[c.Value]
	
	user_name := userSession.username
	user_id := userSession.userid
    user_password := userSession.password
	
	//fmt.Println(user_name)
	//fmt.Println(user_password)
	
    // connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user_name, user_password, dbname)
         
        // open database
    db, err := sql.Open("postgres", psqlconn)
    checkError(err)
     
    // close database
    defer db.Close()
	
	insertStmt := `insert into "Posts"("authorID", "post","title","timestamp") values($1, $2, $3, $4)`
    _, e := db.Exec(insertStmt, user_id, content, title, time.Now())
    checkError(e)
	
	refreshSession(w, r)
	
	fmt.Fprintf(w, "<h1> Post Create Successful!</h1>" + "<p>Now go to the user homepage...</p>")
	
	if f, ok := w.(http.Flusher); ok {
			f.Flush()
			time.Sleep(1 * time.Second)
		}
		

	fmt.Fprintf(w, "<script>window.location.href=\"/homepage\";</script>") 
	
}

