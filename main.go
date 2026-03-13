package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// define a page structure
type Page struct {
	Title string
	Body  []byte
}

// TODO : check if this has any usecase now
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// the handler when accessing /
func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{}
	t, _ := template.ParseFiles("./templates/base.html", "./templates/main.html")
	t.Execute(w, p)
}

// TODO: create a contact page
func contactHandler(w http.ResponseWriter, r *http.Request) {
	/*title := r.URL.Path[len("/contact/"):]
	p, err := loadPage("./templates/base")
	if err != nil {
		p = &Page{Title: title}
	}*/
	p := &Page{Title: "random"}
	t, _ := template.ParseFiles("./templates/base.html", "./templates/main.html")
	t.Execute(w, p)
}

func main() {
	// load static path
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	// define a handler for /
	http.HandleFunc("/", homeHandler)

	// TODO : check contact handler
	//http.HandleFunc("/contact", contactHandler)

	// serve
	log.Fatal(http.ListenAndServe(":8080", nil))
}
