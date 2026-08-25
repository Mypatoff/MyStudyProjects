import (
	"html/template"
	"net/http"
)

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
