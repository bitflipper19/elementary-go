package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type student struct {
	name   string
	branch []byte
}

func (p *student) save() error {
	filename := p.name + ".txt"
	return os.WriteFile(filename, p.branch, 0600)
}

func load(title string) (*student, error) {
	filename := title + ".txt"
	branch, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &student{name: title, branch: branch}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %s", r.URL.Path[1:])
}

func view(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := load(title)
	fmt.Fprintf(w, "<h1>%s</h1><br><h2>%s</h2>", p.name, p.branch)
}

func main() {
	http.HandleFunc("/view/", view)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
