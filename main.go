package main

import (
	"net/http"
)

func Main(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/main.html")
}

func main() {
	http.HandleFunc("/", Main)
	http.ListenAndServe(":8181", nil)
}
