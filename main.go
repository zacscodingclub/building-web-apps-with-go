package main

import (
	"fmt"
	"net/http"

	"github.com/russross/blackfriday"
)

func main() {
	fmt.Println("Now serving to localhost:8080")
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("templates")))
	http.ListenAndServe("localhost:8080", nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	md := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(md)
}
