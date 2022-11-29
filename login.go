package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
	"net/http"
	"html/template"
	"github.com/google/uuid"
	"time"
)


func loginHandler(w http.ResponseWriter, r *http.Request) {
    issueList, err := template.ParseFiles("pages/user_login.html")
    checkError(err)
    
    err = issueList.Execute(w, nil)
    checkError(err)
    
}


func loginconfirmHandler (w http.ResponseWriter, r *http.Request) {

	host     := HOST
    port     := PORT
    dbname   := DBNAME

    user_name     := r.FormValue("Username")
    user_password := r.FormValue("Password")	
	
	//fmt.Println(user_name)
	//fmt.Println(user_password)
	
    // connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user_name, user_password, dbname)
         
        // open database
    db, err := sql.Open("postgres", psqlconn)
    checkError(err)
     
    // close database
    defer db.Close()
    err = db.Ping()
    //checkError(err)
	if err != nil {
		fmt.Fprintf(w, "<h1> Login Failed. Try again. </h1>")
		return
	} 
	
	useridconn := fmt.Sprintf("select \"UserID\" FROM \"Users\" where \"UserName\" = '%s'", user_name)
	
	userid, err := db.Query(useridconn)
	defer userid.Close()
	
	var id int
	for userid.Next() {
		err = userid.Scan(&id)
		checkError(err)
	}
	
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)
	
	sessions[sessionToken] = session{
		username: user_name,
		userid: id,
		password: user_password,
		expiry: expiresAt,
	}
	
	http.SetCookie(w, &http.Cookie{
		Name:	"session_token",
		Value:	sessionToken,
		Expires:expiresAt,
	})
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1> User %s Login Successful!</h1>" + "<p>Now go to the user homepage...</p>",user_name)
	
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
		time.Sleep(1 * time.Second)
	}
		

	fmt.Fprintf(w, "<script>window.location.href=\"/homepage\";</script>") 
		
}

