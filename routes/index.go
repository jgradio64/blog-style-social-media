package routes

import (
	"net/http"
)

type Data struct {
	Body 		string
	IsAuth 	    bool
    AuthUserID	int
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated := CheckSession(w, r) 
    if isAuthenticated {
        c, _ := r.Cookie("session_token")
        sessionToken := c.Value
        userSession := sessions[sessionToken]
        page := &Data{Body:"Welcome to our brand new landing page.", IsAuth: isAuthenticated, AuthUserID: userSession.UserID}
        renderTemplate(w, "index", page)
    } else {
        page := &Data{Body:"Welcome to our brand new landing page.", IsAuth: isAuthenticated}
        renderTemplate(w, "index", page)
    }
}

func renderTemplate(w http.ResponseWriter, tmpl string, page *Data) {
    err := templates.ExecuteTemplate(w, tmpl, page)
    if err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
         return
    }
}