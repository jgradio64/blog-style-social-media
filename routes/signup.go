package routes
 
import (
	"net/http"
    "fmt"
    "time"
	"html/template"
    
	extns "blbr.com/main/extensions"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
    isAuthenticated := CheckSession(w, r) 
    if isAuthenticated {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<h1>You are already logged into an account ... </h1>" + "<p>Redirecting to the index page</p>")
		
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
			time.Sleep(1 * time.Second)
		}

		fmt.Fprintf(w, `<script>window.location.href="/";</script>`) 
    } else {
        issueList, err := template.ParseFiles("pages/user_signup.html")
        extns.CheckError(err)
        
        err = issueList.Execute(w, nil)
        extns.CheckError(err)
    }
    
}