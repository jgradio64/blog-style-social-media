package main
 
import (
	"net/http"
	"github.com/google/uuid"
	"time"
)

func refreshSession(w http.ResponseWriter, r *http.Request) {
	checkSession(w, r) 
	
	c, _ := r.Cookie("session_token")

	userSession := sessions[c.Value]
	
	newSessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	sessions[newSessionToken] = session{
		username: userSession.username,
		userid:	  userSession.userid,
		password: userSession.password,
		expiry:   expiresAt,
	}
	
	delete(sessions, c.Value)
	
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: expiresAt,
	})

}