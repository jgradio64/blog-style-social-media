package routes

import (
    "net/http"
	"strconv"
	"time"
	"fmt"

	models "blbr.com/main/models"
	extns "blbr.com/main/extensions"
)

type PostData struct {
	IsOwner		bool
	IsAuth		bool
	AuthUserID	int
	ThePost		models.Post
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {

	isAuthenticated := CheckSession(w, r) 
	if isAuthenticated {
		queryId := r.URL.Query().Get("id")
		c, _ := r.Cookie("session_token")
		userSession := sessions[c.Value]

		id, err := strconv.Atoi(queryId)
		extns.CheckError(err)

		post := models.Post{PostID: id}
		post = post.GetPost()

		if post.UserID == userSession.UserID {
			newPostData := PostData{IsOwner: true, ThePost: post, IsAuth: isAuthenticated, AuthUserID: userSession.UserID }
			page := &newPostData
			renderPostTemplate(w, "post", page)
		} else {
			newPostData := PostData{IsOwner: false, ThePost: post, IsAuth: isAuthenticated, AuthUserID: userSession.UserID }
			page := &newPostData
			renderPostTemplate(w, "post", page)
		}
	} else {
		Backtologin(w, r)
	}
}

func NewPostHandler(w http.ResponseWriter, r *http.Request) {

	isAuthenticated := CheckSession(w, r) 
	if isAuthenticated {
		c, _ := r.Cookie("session_token")
		userSession := sessions[c.Value]
		newPostData := PostData{IsAuth: isAuthenticated, AuthUserID: userSession.UserID }
		page := &newPostData
		renderPostTemplate(w, "post_create", page)
	} else {
		Backtologin(w, r)
	}
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated := CheckSession(w, r) 
	if isAuthenticated {
		// Get the post information
		c, _ := r.Cookie("session_token")
		userSession := sessions[c.Value]

		queryId := r.URL.Query().Get("id")
		id, err := strconv.Atoi(queryId)
		extns.CheckError(err)
		var post models.Post
		post.PostID = id

		switch r.Method {
		case "GET":		
			post = post.GetPost()
			newPostData := PostData{IsAuth: isAuthenticated, AuthUserID: userSession.UserID, ThePost: post }
			renderPostTemplate(w, "post_edit", &newPostData)
		case "POST":
			post.PostContent = r.FormValue("postContent")
			post.UpdatePost()

			fmt.Fprintf(w, "<h1>Post Successfully Edited</h1>")
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
				time.Sleep(1 * time.Second)
			}
			fmt.Fprintf(w, "<script>window.location.href=\"/post/?id=%v\";</script>", id) 
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
		
	} else {
		Backtologin(w, r)
	}
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated := CheckSession(w, r) 
	if isAuthenticated {
		c, _ := r.Cookie("session_token")
		userSession := sessions[c.Value]

		queryId := r.URL.Query().Get("id")
		id, err := strconv.Atoi(queryId)
		extns.CheckError(err)

		var post models.Post
		post.PostID = id
		newPostData := PostData{IsAuth: isAuthenticated, AuthUserID: userSession.UserID, ThePost: post }
		renderPostTemplate(w, "post_delete", &newPostData)
	} else {
		Backtologin(w, r)
	}
}

func DeletePostConfirmHandler(w http.ResponseWriter, r *http.Request){
	isAuthenticated := CheckSession(w, r) 
	if isAuthenticated {
		queryId := r.URL.Query().Get("id")
		id, err := strconv.Atoi(queryId)
		extns.CheckError(err)

		post := models.Post{PostID: id}
		post.DeletePost()

		fmt.Fprintf(w, "<h1>Post Deletion Successful!</h1>" + "<p>Were you too embarrassed by it?</p>")
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
			time.Sleep(1 * time.Second)
		}
		fmt.Fprintf(w, `<script>window.location.href="/";</script>`) 
	} else {
		Backtologin(w, r)
	}
}

func NewPostConfirmHandler(w http.ResponseWriter, r *http.Request){
	isAuthenticated := CheckSession(w, r) 
	if isAuthenticated {
		var newPost models.Post

		c, _ := r.Cookie("session_token")
		userSession := sessions[c.Value]

		newPost.UserID = userSession.UserID
		newPost.PostTitle = r.FormValue("postTitle")
		newPost.PostContent = r.FormValue("postContent")

		newPost.CreatePost()

		fmt.Fprintf(w, "<h1>New Post Created!</h1>" + "<p>Rending the new post now</p>")
		
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
			time.Sleep(1 * time.Second)
		}

		fmt.Fprintf(w, `<script>window.location.href="/";</script>`) 
	} else {
		Backtologin(w, r)
	}
}

func renderPostTemplate(w http.ResponseWriter, tmpl string, page *PostData) {
	err := templates.ExecuteTemplate(w, tmpl, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}