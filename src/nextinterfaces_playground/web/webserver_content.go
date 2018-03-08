package main

import (
	"fmt"
	"log"
	"net/http"
)

type Content struct {
	Title string
	Body  string
}

func loadContent(title string) (*Content, error) {
	return &Content{Title: title, Body: "Le Body"}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadContent(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Print("sdfgsdfg")
}
