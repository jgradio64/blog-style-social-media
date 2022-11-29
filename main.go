package main
 
import (
    "fmt"
	"log"
	"net/http"
	"time"
)

type session struct {
	username string
	userid	 int
	password string
	expiry	 time.Time
}

var sessions = map[string]session{}

type post struct {
	username string
	posttitle string
	userpost string
	posttime time.Time
}

var posts []post

func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func checkError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}


func main () {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/homepage", homepageHandler)
	
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/confirm", confirmHandler)
	
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/loginconfirm", loginconfirmHandler)
	http.HandleFunc("/useredit", usereditHandler)
	http.HandleFunc("/usereditconfirm", usereditconfirmHandler)
	
	http.HandleFunc("/userdelete", userdeleteHandler)
	http.HandleFunc("/userdeleteconfirm", userdeleteconfirmHandler)	

	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/postconfirm", postconfirmHandler)
	
	http.HandleFunc("/userview", userviewHandler)
	http.HandleFunc("/userviewconfirm", userviewconfirmHandler)
	
	http.HandleFunc("/postview", postviewHandler)
	http.HandleFunc("/postviewconfirm", postviewconfirmHandler)
	
	
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, INDEXPAGE)
	
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, HOMEPAGE)
	
}