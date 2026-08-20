package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("Pages_Frontend/home.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", homePage)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
