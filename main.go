package main
 
import (
	"log"
	"net/http"

	routes "blbr.com/main/routes"
)

// Just handles the route functions
func main () {
	http.HandleFunc("/", routes.IndexHandler)
	http.HandleFunc("/home/", routes.HomepageHandler)
	http.HandleFunc("/all/", routes.AllPageHandler)

	// User routes
	http.HandleFunc("/user/", routes.GetUserHandler)
	http.HandleFunc("/user/new", routes.SignupHandler)
	http.HandleFunc("/user/edit", routes.EditUserHandler)
	http.HandleFunc("/user/delete", routes.DeleteUserHandler)
	http.HandleFunc("/user/follow/", routes.FollowUserHandler)
	http.HandleFunc("/user/unfollow/", routes.UnfollowUserHandler)
	http.HandleFunc("/usereditconfirm", routes.EditUserConfirmHandler)
	http.HandleFunc("/userdeleteconfirm", routes.DeleteUserConfirmHandler)	
	http.HandleFunc("/signupconfirm", routes.CreateUserHandler)
	
	// auth Routes
	http.HandleFunc("/login", routes.LoginHandler)
	http.HandleFunc("/loginconfirm", routes.LoginConfirmHandler)
	http.HandleFunc("/logout", routes.LogoutHandler)
	
	// user Post routes
	http.HandleFunc("/post/", routes.GetPostHandler)
	http.HandleFunc("/post/new", routes.NewPostHandler)
	http.HandleFunc("/post/delete/", routes.DeletePostHandler)
	http.HandleFunc("/post/edit/", routes.EditPostHandler)
	http.HandleFunc("/deletepostconfirm/", routes.DeletePostConfirmHandler)
	http.HandleFunc("/postconfirm/", routes.NewPostConfirmHandler)

	// Comment Routes
	http.HandleFunc("/comment/new/", routes.CreateCommentHandler)
	http.HandleFunc("/comment/edit/", routes.EditCommentHandler)
	http.HandleFunc("/comment/delete/", routes.DeleteCommentHandler)

	log.Println("server running on port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}