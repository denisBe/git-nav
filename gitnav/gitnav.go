package main

import (
	"fmt"
	"net/http"
	//"github.com/gorilla/mux"
)

func main() {
    fmt.Println("Hello gitnav world !")
    http.HandleFunc("/", indexHandler)
    http.ListenAndServe(":2340", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
