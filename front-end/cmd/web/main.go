package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Sets up the server by registering an HTTP request handler function with http.HandleFunc().
// The handler function renders the specified template file when the root URL is requested.
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	fmt.Println("Starting front end service on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Panic(err)
	}
}

// Renders the specified template file and writes the output to an http.ResponseWriter instance.
// The function uses template.ParseFiles() to parse all the template files required to render the page.
func render(w http.ResponseWriter, t string) {

	partials := []string{
		"./cmd/web/templates/base.layout.gohtml",
		"./cmd/web/templates/header.partial.gohtml",
		"./cmd/web/templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./cmd/web/templates/%s", t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
