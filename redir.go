package main

import (
	"google.golang.org/appengine"
    "log"
    "net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {

    http.Redirect(w, r, "http://www.google.com", 301)
}

func main() {
	appengine.Main()
    http.HandleFunc("/", redirect)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}