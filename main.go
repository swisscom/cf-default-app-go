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

type content struct {
	EnvVars    map[string]string
	ReqHeaders http.Header
	ReqParams  url.Values
}

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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	c := content{
		createVarMap(os.Environ()),
		r.Header,
		r.URL.Query(),
	}

	lp := path.Join("templates", "layout.html")
	ip := path.Join("templates", "index.html")

	t, err := template.ParseFiles(lp, ip)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "layout", c)
	if err != nil {
		panic(err)
	}
}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "3000"
	}

	http.HandleFunc("/", indexHandler)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
