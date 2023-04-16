package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	fmt.Println("Starting front end service on port 80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {

	wd, err := os.Getwd()

	partials := []string{
		wd + "/templates/base.layout.gohtml",
		wd + "/templates/header.partial.gohtml",
		wd + "/templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf(wd+"/templates/test.page.gohtml"))

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
