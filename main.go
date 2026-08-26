package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}

func greetPage(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "stranger"
	}

	tmpl := template.Must(template.New("greet").Parse(`
		<html>
		<body>
			<h1>Hello, {{.}}!</h1>
		</body>
		</html>
	`))

	tmpl.Execute(w, name)
}

func main() {
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/greet", greetPage)
	http.ListenAndServe(":8080", nil)
}

// we have learnt how to
//I think I am not gonna learn anything today...
