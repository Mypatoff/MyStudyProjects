package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Davron!")
}
func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}

func main() {
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}
