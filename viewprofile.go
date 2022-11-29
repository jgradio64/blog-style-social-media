package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
	"net/http"
	"html/template"
	"time"
)



func userviewHandler(w http.ResponseWriter, r *http.Request) {
	if !checkSession(w, r) {
		backtologin(w, r)
		return
	}
    issueList, err := template.ParseFiles("pages/user_view.html")
    checkError(err)
    
    err = issueList.Execute(w, nil)
    checkError(err)
    
}

func userviewconfirmHandler(w http.ResponseWriter, r *http.Request) {

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
	
	if target_user_name == "" {
		target_user_name = user_name
	}
	
	// connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user_name, user_password, dbname)
         
    // open database
    db, err := sql.Open("postgres", psqlconn)
    checkError(err)
     
    // close database
    defer db.Close()
	
	// query	
	userconn := fmt.Sprintf("select \"UserID\", \"About\", \"JoinTime\" FROM \"Users\" where \"UserName\" = '%s'", target_user_name)
	rows, err := db.Query(userconn)
	
	var user_id int
	var about string
	var jointime time.Time
	
	for rows.Next() {
		err = rows.Scan(&user_id, &about, &jointime)
		checkError(err)
	}
	
	//fmt.Println(user_id)
	var profile string
	if user_id != 0 {
		y,m,d := jointime.Date()
		rows.Close()
	
		// query count
		userconn = fmt.Sprintf("select COUNT(*) FROM \"Posts\" where \"authorID\" = '%d'", user_id)
		rows, err = db.Query(userconn)
	
		var count string
	
		for rows.Next() {
			err = rows.Scan(&count)
			checkError(err)
		}
		rows.Close()
		
		profile = fmt.Sprintf("<h1> User Profile <h1><div>User name: %s</div><div>User since: %04d-%02d-%02d</div><div>About user: %s</div><div>User has %s post(s)</div><a href=\"http://localhost:8000/homepage\">Back to Homepage</a>",
			target_user_name, y,m,d,about, count)
	} else {
		rows.Close()
		profile = fmt.Sprintf("<h1> User %s Not Exists <h1><a href=\"http://localhost:8000/homepage\">Back to Homepage</a>",
			target_user_name)
	}
	refreshSession(w, r)
	fmt.Fprintf(w, profile)

	
}
