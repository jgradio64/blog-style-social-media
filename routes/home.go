package routes

import (
    "fmt"
    // "log"
    "net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}


	fmt.Fprintf(w, "This will be the home page")
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/users/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}


	fmt.Fprintf(w, "This will be the home page")
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/posts/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}


	fmt.Fprintf(w, "This will be the home page")
}