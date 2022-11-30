package routes

import (
    "fmt"
	"net/http"
	"time"
	"html/template"
	"github.com/google/uuid"

	models "blbr.com/main/models"
	extns "blbr.com/main/extensions"
)

// Use a map to store the sessions for now, but could turn into a db structure later
var sessions = map[string]Session{}

type Session struct {
	Username 	string
	UserID		int
	Expiry   	time.Time
}

func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    issueList, err := template.ParseFiles("pages/user_login.html")
    extns.CheckError(err)
    
    err = issueList.Execute(w, nil)
    extns.CheckError(err)
}

func LoginConfirmHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User
    user.Username = r.FormValue("Username")
    user.Password = r.FormValue("Password")	
	
    var validUser models.User
	validUser = user.ValidateUser()

	if (validUser.UserID == 0) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "<h1> Login Failed. Try again. </h1>")
		return
	}
	
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)
	
	sessions[sessionToken] = Session{
		Username: user.Username,
		UserID: validUser.UserID,
		Expiry: expiresAt,
	}
	
	http.SetCookie(w, &http.Cookie{
		Name:		"session_token",
		Value:		sessionToken,
		Expires:	expiresAt,
	})
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Login Successful!</h1>" + "<p>Redirecting to the user homepage</p>")
	
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
		time.Sleep(1 * time.Second)
	}

	fmt.Fprintf(w, `<script>window.location.href="/homepage";</script>`) 
		
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated := CheckSession(w, r) 

	if isAuthenticated {
		c, _ := r.Cookie("session_token")

		sessionToken := c.Value

		delete(sessions, sessionToken)

		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   "",
			Expires: time.Now(),
		})

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<h1>User is logged out</h1>" + "<p>Redirecting to the index page</p>")
		
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
			time.Sleep(1 * time.Second)
		}

		fmt.Fprintf(w, `<script>window.location.href="/";</script>`) 
	} else {
		Backtologin(w, r)
	}

}

func RefreshSession(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r) 
	
	c, _ := r.Cookie("session_token")

	userSession := sessions[c.Value]
	
	newSessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	sessions[newSessionToken] = Session{
		Username: userSession.Username,
		UserID:	  userSession.UserID,
		Expiry:   expiresAt,
	}
	
	delete(sessions, c.Value)
	
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: expiresAt,
	})
}

func CheckSession(w http.ResponseWriter, r *http.Request) bool {

	c, err := r.Cookie("session_token")
	if err != nil {
		return false
	}
	
	var userSession Session
	sessionToken := c.Value	
	userSession, exists := sessions[sessionToken]
	
	if !exists {
		return false
	}
	
	if userSession.IsExpired() {
		delete(sessions, sessionToken)
		return false
	}
	return true

}

func Backtologin(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1> Not Authenticated. Please try logging in again. </h1>")
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
		time.Sleep(2 * time.Second)
	}
	fmt.Fprintf(w, `<script>window.location.href="/login";</script>`)
}