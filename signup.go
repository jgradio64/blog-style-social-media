package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
	"net/http"
	"html/template"
	"time"
)


func signupHandler(w http.ResponseWriter, r *http.Request) {
    issueList, err := template.ParseFiles("pages/user_signup.html")
    checkError(err)
    
    err = issueList.Execute(w, nil)
    checkError(err)
    
}

func confirmHandler (w http.ResponseWriter, r *http.Request) {

	host     := HOST
    port     := PORT
    admin    := "postgres"
    admin_password := "postgres"
    dbname   := DBNAME

    user_name     := r.FormValue("Username")
    user_password := r.FormValue("Password")	
	user_about := r.FormValue("About")
	
    // connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, admin, admin_password, dbname)
         
        // open database
    db, err := sql.Open("postgres", psqlconn)
    checkError(err)
     
    // close database
    defer db.Close()
	userconn := fmt.Sprintf("select \"UserName\" FROM \"Users\" where \"UserName\" = '%s'", user_name)
	//fmt.Println(userconn)
	existUser, err := db.Query(userconn)
	//existUser, err := db.Query(`SELECT "UserName" FROM "Users" WHERE "UserName" = 'david'`)
	defer existUser.Close()
	
	var name string
	for existUser.Next() {
		err = existUser.Scan(&name)
		checkError(err)
	}
	//fmt.Println(name)
	if name != "" {
		fmt.Fprintf(w, "<h1> User Name %s is used. Try a new one. </h1>",user_name)
	} else {
		
		fmt.Fprintf(w, "<h1> User %s Created Successfully!</h1>" + "<p>Now go to the log in page...</p>",user_name)
		
		createStmt := fmt.Sprintf("CREATE USER %s WITH PASSWORD '%s'", user_name ,user_password)
	    _, e := db.Exec(createStmt)
		checkError(e)
		
		attributeStmt := fmt.Sprintf("GRANT PG_READ_ALL_DATA to %s", user_name)
		_, e = db.Exec(attributeStmt)
		checkError(e)
		
		attributeStmt = fmt.Sprintf("GRANT PG_WRITE_ALL_DATA to %s", user_name)
		_, e = db.Exec(attributeStmt)
		checkError(e)
		
		insertDynStmt := `insert into "Users"("UserName", "About","JoinTime") values($1, $2, $3)`
		_, e = db.Exec(insertDynStmt, user_name, user_about, time.Now())
		checkError(e)
		
		
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
			time.Sleep(1 * time.Second)
		}
		

		fmt.Fprintf(w, "<script>window.location.href=\"/\";</script>") 
	}
			 
}
