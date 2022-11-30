package routes
 
import (
	"net/http"

	models "blbr.com/main/models"
)

type AllPostsData struct {
	IsAuth		bool
	AuthUserID	int
	PagePosts	[]models.Post
}

func AllPageHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated := CheckSession(w, r) 
	if isAuthenticated{
		c, _ := r.Cookie("session_token")
		sessionToken := c.Value
		userSession := sessions[sessionToken]

		var posts models.Post
		allPosts := posts.GetAllPosts()

		allPostData := AllPostsData{PagePosts: allPosts, IsAuth: isAuthenticated, AuthUserID: userSession.UserID }
		page := &allPostData
		renderAllPage(w, "all", page)
	} else {
		var posts models.Post
		allPosts := posts.GetAllPosts()

		allPostData := AllPostsData{PagePosts: allPosts, IsAuth: isAuthenticated }
		page := &allPostData
		renderAllPage(w, "all", page)
	}
}

func renderAllPage(w http.ResponseWriter, tmpl string, page *AllPostsData) {
    err := templates.ExecuteTemplate(w, tmpl, page)
    if err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
         return
    }
}