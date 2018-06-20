package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Now serving to localhost:", port)
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("templates")))
	http.ListenAndServe("localhost:"+port, nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	md := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(md)
}
