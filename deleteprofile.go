package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
	"net/http"
	"html/template"
	"time"
)



func userdeleteHandler(w http.ResponseWriter, r *http.Request) {

	if !checkSession(w, r) {
		backtologin(w, r)
		return
	}
    issueList, err := template.ParseFiles("pages/user_delete.html")
    checkError(err)
	
    err = issueList.Execute(w, nil)
    checkError(err)
}



func userdeleteconfirmHandler (w http.ResponseWriter, r *http.Request) {
	if !checkSession(w, r) {
		backtologin(w, r)
		return
	}
	host     := HOST
    port     := PORT
    admin    := "postgres"
    admin_password := "postgres"
    dbname   := DBNAME
	
	c, _ := r.Cookie("session_token")
	userSession := sessions[c.Value]
	user_id := userSession.userid
	user_name := userSession.username
	
    // connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, admin, admin_password, dbname)
         
    // open database
    db, err := sql.Open("postgres", psqlconn)
    checkError(err)
     
    // close database
    defer db.Close()
	
	// delete	
	deleteStmt := `delete from "Posts" where "authorID"=$1`
	_, e := db.Exec(deleteStmt, user_id)
	checkError(e)
	
	deleteStmt = `delete from "Users" where "UserID"=$1`
	_, e = db.Exec(deleteStmt, user_id)
	checkError(e)
	
	deleteCmd := fmt.Sprintf("DROP ROLE %s", user_name)
	_, e = db.Exec(deleteCmd)
	checkError(e)
	
	delete(sessions, c.Value)
	
	fmt.Fprintf(w, "<h1> Profile Delete Successful!</h1>" + "<p>See you in the future...</p>")
	
	if f, ok := w.(http.Flusher); ok {
			f.Flush()
			time.Sleep(2 * time.Second)
		}
		

	fmt.Fprintf(w, "<script>window.location.href=\"/\";</script>") 
	
}

