package main
 
import (
    "fmt"
	"log"
	"net/http"
	"time"
)

var sessions = map[string]session{}

func checkError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

type session struct {
	username string
	userid	 int
	password string
	expiry	 time.Time
}

func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func main () {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/confirm", confirmHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/loginconfirm", loginconfirmHandler)
	http.HandleFunc("/useredit", usereditHandler)
	http.HandleFunc("/usereditconfirm", usereditconfirmHandler)
	
	
	http.HandleFunc("/homepage", homepageHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/postconfirm", postconfirmHandler)
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