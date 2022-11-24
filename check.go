package main
 
import (
	"net/http"
)


func checkSession(w http.ResponseWriter, r *http.Request) bool {

	c, err := r.Cookie("session_token")
	checkError(err)
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
