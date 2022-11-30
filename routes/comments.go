package routes

import (
    "net/http"
	"strconv"
	"time"
	"fmt"
	_ "html/template"


	models "blbr.com/main/models"
	extns "blbr.com/main/extensions"
)

type CommentData struct {
	IsOwner		bool
	IsAuth		bool
	AuthUserID	int
	Com			models.Comment
}

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated := CheckSession(w, r) 
	if isAuthenticated {
		// First get the user's id from the session
		c, _ := r.Cookie("session_token")
		userSession := sessions[c.Value]

		var comment models.Comment
		// Get the post id from the form
		postID, err := strconv.Atoi(r.FormValue("post_id"))
		extns.CheckError(err)

		comment.PostID = postID
		comment.UserID = userSession.UserID
		comment.Content = r.FormValue("comment_text")
		comment.CreateComment()

		fmt.Fprintf(w, "<h1>Comment Successfully Added</h1>")
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
			time.Sleep(1 * time.Second)
		}
		redirectScript := fmt.Sprintf("<script>window.location.href=\"/post/?id=%v\";</script>", postID)
		fmt.Fprintf(w, redirectScript) 
		
	} else {
		Backtologin(w, r)
	}
}

func EditCommentHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated := CheckSession(w, r) 
	if isAuthenticated {
		// First get the user's id from the session
		c, _ := r.Cookie("session_token")
		userSession := sessions[c.Value]

		var comment models.Comment
		// Get the commentID from url search params
		queryId := r.URL.Query().Get("id")
		id, err := strconv.Atoi(queryId)
		extns.CheckError(err)
		comment.CommentID = id

		switch r.Method {
		case "GET":		
			comment = comment.GetComment()
			commentData := CommentData{IsAuth: isAuthenticated, AuthUserID: userSession.UserID, Com: comment}
			renderCommentTemplate(w, "comment_edit", &commentData)
			
		case "POST":
			comment.Content = r.FormValue("comment_text")
			comment.UpdateComment()
			comment = comment.GetComment()

			fmt.Fprintf(w, "<h1>Comment Successfully Added</h1>")
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
				time.Sleep(1 * time.Second)
			}
			redirectScript := fmt.Sprintf("<script>window.location.href=\"/post/?id=%v\";</script>", comment.PostID)
			fmt.Fprintf(w, redirectScript) 

		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	} else {
		Backtologin(w, r)
	}
}

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	isAuthenticated := CheckSession(w, r) 
	if isAuthenticated {
		// First get the user's id from the session
		c, _ := r.Cookie("session_token")
		userSession := sessions[c.Value]

		var comment models.Comment
		// Get the commentID from url search params
		queryId := r.URL.Query().Get("id")
		id, err := strconv.Atoi(queryId)
		extns.CheckError(err)
		comment.CommentID = id

		switch r.Method {
		case "GET":		
			commentData := CommentData{IsAuth: isAuthenticated, AuthUserID: userSession.UserID, Com: comment}
			renderCommentTemplate(w, "comment_delete", &commentData)
		case "POST":
			postID := comment.GetComment().PostID
			comment.DeleteComment()

			fmt.Fprintf(w, "<h1>Comment Successfully Deleted</h1>")
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
				time.Sleep(1 * time.Second)
			}
			redirectScript := fmt.Sprintf("<script>window.location.href=\"/post/?id=%v\";</script>", postID)
			fmt.Fprintf(w, redirectScript) 
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	} else {
		Backtologin(w, r)
	}
}

func renderCommentTemplate(w http.ResponseWriter, tmpl string, page *CommentData) {
	err := templates.ExecuteTemplate(w, tmpl, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}