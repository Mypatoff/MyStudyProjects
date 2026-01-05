package main

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("----Main----")); err != nil {
		panic(err)
	}
}
func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
