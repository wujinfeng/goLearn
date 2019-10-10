package main
import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)
func main()  {
	r:=httprouter.New()
	r.GET("/", HomeHandler)

	// posts collection
	r.GET("/posts", PostsIndexHandle)
	r.POST("/posts", PostsCreateHandle)

	// posts singular
	r.GET("/posts/:id", PostShowHandle)
	r.PUT("/posts/:id", PostUpdateHandle)
	r.DELETE("/posts/:id", PostDeleteHandle)
	r.GET("/posts/:id/edit", PostEditHandle)

	fmt.Println("starting server on :8080")
	http.ListenAndServe(":8080", r)

}

func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "home")
}

func PostsIndexHandle(rw http.ResponseWriter, r *http.Request, p httprouter.Params){
	fmt.Fprintln(rw, "posts index")
}

func PostsCreateHandle(rw http.ResponseWriter, r *http.Request, p httprouter.Params){
	fmt.Fprintln(rw, "posts create")
}

func PostShowHandle(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id:=p.ByName("id")
	fmt.Fprintln(rw, "showing post ", id)
}

func PostUpdateHandle(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts update")
}

func PostDeleteHandle(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts delete")
}

func PostEditHandle(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts edit")
}