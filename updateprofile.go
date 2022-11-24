package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
	"net/http"
	"html/template"
	"time"
)



func usereditHandler(w http.ResponseWriter, r *http.Request) {

	if !checkSession(w, r) {
		fmt.Fprintf(w, "<h1> User Invalid. Please Log In Again. </h1>")
		return
	}
    issueList, err := template.ParseFiles("user_edit.html")
    checkError(err)
	
    err = issueList.Execute(w, nil)
    checkError(err)
}



func usereditconfirmHandler (w http.ResponseWriter, r *http.Request) {
	host     := HOST
    port     := PORT
    dbname   := DBNAME

    about := r.FormValue("About")	
	
	c, _ := r.Cookie("session_token")

	userSession := sessions[c.Value]
	
	user_name := userSession.username
	user_id := userSession.userid
    user_password := userSession.password
	
    // connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user_name, user_password, dbname)
         
        // open database
    db, err := sql.Open("postgres", psqlconn)
    checkError(err)
     
    // close database
    defer db.Close()
	
	// update
	updateStmt := `update "Users" set "About"=$1 where "UserID"=$2`
	_, e := db.Exec(updateStmt, about, user_id)
	checkError(e)
	
	
	
	refreshSession(w, r)
	
	fmt.Fprintf(w, "<h1> Profile Update Successful!</h1>" + "<p>Now go to the user homepage...</p>")
	
	if f, ok := w.(http.Flusher); ok {
			f.Flush()
			time.Sleep(5 * time.Second)
		}
		

	fmt.Fprintf(w, "<script>window.location.href=\"/homepage\";</script>") 
	
}

