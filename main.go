package main
 
import (
	"log"
	"net/http"
	"html/template"
	routes "blbr.com/main/routes"

)

type Data struct {
	Body string
}

var templates = template.Must(template.ParseGlob("pages/*"))

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
	http.HandleFunc("/usereditconfirm", routes.EditUserConfirmHandler)
	http.HandleFunc("/userdeleteconfirm", routes.DeleteUserConfirmHandler)	
	http.HandleFunc("/signupconfirm", routes.CreateUserHandler)
	
	// auth Routes
	http.HandleFunc("/login", routes.LoginHandler)
	http.HandleFunc("/loginconfirm", routes.LoginConfirmHandler)
	http.HandleFunc("/logout", routes.LogoutHandler)
	
	http.HandleFunc("/post/", routes.GetPostHandler)
	http.HandleFunc("/post/new", routes.NewPostHandler)
	http.HandleFunc("/post/delete/", routes.DeletePostHandler)
	http.HandleFunc("/post/edit/", routes.EditPostHandler)
	http.HandleFunc("/deletepostconfirm/", routes.DeletePostConfirmHandler)
	http.HandleFunc("/postconfirm/", routes.NewPostConfirmHandler)
	
	// http.HandleFunc("/userview", userviewHandler)	
	// http.HandleFunc("/postviewconfirm", postviewconfirmHandler)
	
	log.Println("server running on port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	page := &Data{Body:"Welcome to our brand new landing page."}
    renderTemplate(w, "index", page)
}

// func homepageHandler(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	fmt.Fprintf(w, HOMEPAGE)
	
// }

func renderHomePage(w http.ResponseWriter, tmpl string, page *Data) {
    err := templates.ExecuteTemplate(w, tmpl, page)
    if err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
         return
    }
}

func renderTemplate(w http.ResponseWriter, tmpl string, page *Data) {
    err := templates.ExecuteTemplate(w, tmpl, page)
    if err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
         return
    }
}