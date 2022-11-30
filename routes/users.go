package routes
 
import (
    "fmt"
	"net/http"
	"time"
	"html/template"
	"strconv"
	
	models "blbr.com/main/models"
	extns "blbr.com/main/extensions"
)

type UserData struct {
	IsOwner		bool
	IsAuth		bool
	AuthUserID	int
	TheUser		models.User
	IsFollower	bool
}

var templates = template.Must(template.ParseGlob("pages/*"))


func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Base get route without any parameters will return the user's own profile
	// Get the user's id from the url parameters
	isAuthenticated := CheckSession(w, r) 
	if isAuthenticated {
		queryId := r.URL.Query().Get("id")
		c, _ := r.Cookie("session_token")
		sessionToken := c.Value
		userSession := sessions[sessionToken]

		id, err := strconv.Atoi(queryId)
		extns.CheckError(err)

		user  := models.User{UserID: id}
		user = user.GetUser()

		var flwUser models.FollowedUser
		flwUser.FollowedUserID = id
		followerIDs := flwUser.GetFollowers()

		isFollower := contains(followerIDs, userSession.UserID)

		if id == userSession.UserID {
			newUserData := UserData{IsOwner: true, TheUser: user, IsAuth: isAuthenticated, AuthUserID: userSession.UserID, IsFollower: isFollower }
			page := &newUserData
			renderUserTemplate(w, "userprofile", page)
		} else {
			newUserData := UserData{IsOwner: false, TheUser: user, IsAuth: isAuthenticated, AuthUserID: userSession.UserID, IsFollower: isFollower}
			page := &newUserData
			renderUserTemplate(w, "userprofile", page)
		}
	} else {
		Backtologin(w, r)
	}
}

func EditUserHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated := CheckSession(w, r) 
	if isAuthenticated {
		c, _ := r.Cookie("session_token")
		sessionToken := c.Value
		userSession := sessions[sessionToken]

		user := models.User{UserID: userSession.UserID, Username: userSession.Username}
		user = user.GetUser()

		newUserData := UserData{IsOwner: true, TheUser: user, IsAuth: isAuthenticated, AuthUserID: userSession.UserID}
		page := &newUserData

		renderUserTemplate(w, "edituser", page)
	} else {
		Backtologin(w, r)
	}
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated := CheckSession(w, r) 
	if isAuthenticated {
		issueList, err := template.ParseFiles("pages/user_delete.html")
		extns.CheckError(err)
		
		err = issueList.Execute(w, nil)
		extns.CheckError(err)
	} else {
		Backtologin(w, r)
	}
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	newUser.Username = r.FormValue("Username")
    newUser.Password = r.FormValue("Password")	
	newUser.AboutUser = r.FormValue("About")

	// Runs a query to check if the username is taken
	userAlreadyExists := newUser.CheckUserNameExists()
	
	if userAlreadyExists {
		fmt.Fprintf(w, "<h1> That username already exits. Try a new one. </h1>")
	} else {
		
		fmt.Fprintf(w, "<h1> User %s Created Successfully!</h1>" + "<p>Now go to the log in page...</p>", newUser.Username)
		newUser.CreateUser()
		
		
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
			time.Sleep(1 * time.Second)
		}
		

		fmt.Fprintf(w, `<script>window.location.href="/";</script>`) 
	} 
}

func EditUserConfirmHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated := CheckSession(w, r) 
	if isAuthenticated {
		about := r.FormValue("About")
		c, _ := r.Cookie("session_token")
		sessionToken := c.Value
		userSession := sessions[sessionToken]

		user := models.User{UserID: userSession.UserID, AboutUser: about}
		user.UpdateUser()

		RefreshSession(w, r)

		fmt.Fprintf(w, "<h1> Profile Update Successful!</h1>" + "<p>Now go to the user homepage...</p>")

		if f, ok := w.(http.Flusher); ok {
			f.Flush()
			time.Sleep(1 * time.Second)
		}

		fmt.Fprintf(w, `<script>window.location.href="/homepage";</script>`) 
	} else {
		Backtologin(w, r)
	}
}

func DeleteUserConfirmHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated := CheckSession(w, r) 
	if isAuthenticated {
		c, _ := r.Cookie("session_token")
		sessionToken := c.Value
		userSession := sessions[sessionToken]

		user := models.User{UserID: userSession.UserID}
		user.DeleteUser()

		delete(sessions, c.Value)

		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   "",
			Expires: time.Now(),
		})

		fmt.Fprintf(w, "<h1>Profile Deletion Successful!</h1>" + "<p>See you in the future. Maybe?</p>")
	
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
			time.Sleep(1 * time.Second)
		}
			
		fmt.Fprintf(w, `<script>window.location.href="/";</script>`) 
	} else {
		Backtologin(w, r)
	}
}

func FollowUserHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated := CheckSession(w, r) 
	if isAuthenticated {
		// Get user session for the user id
		c, _ := r.Cookie("session_token")
		sessionToken := c.Value
		userSession := sessions[sessionToken]

		// Get the user to follow's id from the quest param
		queryId := r.URL.Query().Get("id")
		id, err := strconv.Atoi(queryId)
		extns.CheckError(err)

		user := models.FollowedUser{UserID: userSession.UserID, FollowedUserID: id}
		user.CreateFollowedUser()

		fmt.Fprintf(w, "<h1>User followed! :)</h1>")
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
			time.Sleep(1 * time.Second)
		}
		fmt.Fprintf(w, "<script>window.location.href=\"/user/?id=%v\";</script>", id) 
	} else {
		Backtologin(w, r)
	}
}

func UnfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated := CheckSession(w, r) 
	if isAuthenticated {
		// Get user session for the user id
		c, _ := r.Cookie("session_token")
		sessionToken := c.Value
		userSession := sessions[sessionToken]

		// Get the user to follow's id from the quest param
		queryId := r.URL.Query().Get("id")
		id, err := strconv.Atoi(queryId)
		extns.CheckError(err)

		user := models.FollowedUser{UserID: userSession.UserID, FollowedUserID: id}
		user.DeleteFollowedUser()

		fmt.Fprintf(w, "<h1>User unfollowed! :(</h1>")
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
			time.Sleep(1 * time.Second)
		}
		fmt.Fprintf(w, "<script>window.location.href=\"/user/?id=%v\";</script>", id) 
	} else {
		Backtologin(w, r)
	}
}


func renderUserTemplate(w http.ResponseWriter, tmpl string, page *UserData) {
	err := templates.ExecuteTemplate(w, tmpl, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}