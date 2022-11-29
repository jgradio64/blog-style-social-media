package main
 
import (
	"net/http"
	"time"
	"fmt"
)


func checkSession(w http.ResponseWriter, r *http.Request) bool {

	c, err := r.Cookie("session_token")
	//checkError(err)
	if err != nil {
		return false
	}
	
	
	sessionToken := c.Value	
	userSession, exists := sessions[sessionToken]
	
	if !exists {
		return false
	}
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		return false
	}
	
	return true

}

func backtologin(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1> User Invalid. Please Log In Again. </h1>")
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
		time.Sleep(2 * time.Second)
	}
	fmt.Fprintf(w, "<script>window.location.href=\"/login\";</script>")
}