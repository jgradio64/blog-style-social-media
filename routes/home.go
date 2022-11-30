package routes
 
import (
	"net/http"

	models "blbr.com/main/models"
)

type HomepageData struct {
	IsAuth		bool
	AuthUserID	int
	HomePosts	[]models.Post
}

func HomepageHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated := CheckSession(w, r) 
	if isAuthenticated {
		c, _ := r.Cookie("session_token")
		sessionToken := c.Value
		userSession := sessions[sessionToken]

		user := models.User{UserID: userSession.UserID}
		var homePosts []models.Post
		homePosts = user.GetHomePagePosts()
		
		newHomepageData := HomepageData{HomePosts: homePosts, IsAuth: isAuthenticated, AuthUserID: userSession.UserID}
		page := &newHomepageData
		renderHomePage(w, "homepage", page)
	} else {
		Backtologin(w, r)
	}
}

func renderHomePage(w http.ResponseWriter, tmpl string, page *HomepageData) {
    err := templates.ExecuteTemplate(w, tmpl, page)
    if err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
         return
    }
}