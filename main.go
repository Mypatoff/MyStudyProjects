package main

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("----Main----"))
}
func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
