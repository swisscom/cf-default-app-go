package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

// page contains the contents of the index page.
type page struct {
	EnvVars    map[string]string
	ReqHeaders http.Header
	ReqParams  url.Values
}

// createVarMap creates a string map of a set of environment variables.
func createVarMap(data []string) map[string]string {
	vars := make(map[string]string)

	for _, v := range data {
		splits := strings.Split(v, "=")
		key := splits[0]
		val := splits[1]
		vars[key] = val
	}

	return vars
}

// indexHandler renders the index page.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := page{
		createVarMap(os.Environ()),
		r.Header,
		r.URL.Query(),
	}

	ip := path.Join("templates", "index.html.tmpl")

	t, err := template.ParseFiles(ip)
	if err != nil {
		log.Println("error parsing template files:", err)
		return
	}

	err = t.ExecuteTemplate(w, "index", p)
	if err != nil {
		log.Println("error executing template:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", indexHandler)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
