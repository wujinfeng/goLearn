package main

import (
	"net/http"
	"os"
	"gopkg.in/russross/blackfriday.v2"
)

func main() {
	port:=os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":"+port, nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown:=blackfriday.Run([]byte(r.FormValue("body")))
	rw.Write(markdown)
}