package main

import (
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {
	// middleware stack
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(MyMiddleware),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
	)
	n.Run(":8080")
}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Logging on the way there...")
	if r.URL.Query().Get("password") == "123" {
		next(rw, r)
	} else {
		http.Error(rw, "Not Authorized", 401)
	}
	log.Println("Logging on the way back...")
}
