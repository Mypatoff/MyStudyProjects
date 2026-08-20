package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi, I'm Davron</h1><p>This is my website, and it's running on Go.</p>")
}

func main() {
	http.HandleFunc("/", homePage)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
