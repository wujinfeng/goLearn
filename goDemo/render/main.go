package main
import (
	"net/http"
	"github.com/unrolled/render"
)
func main(){
	r:=render.New(render.Options{})
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
		w.Write([]byte("Welcome, visit sub pages now."))
	})
	mux.HandleFunc("/data", func(w http.ResponseWriter, req *http.Request){
		r.Data(w, http.StatusOK, []byte("some binary data here."))
	})
	mux.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request){
		r.JSON(w, http.StatusOK, map[string]string{"hello":"json"})
	})
	mux.HandleFunc("/html", func(w http.ResponseWriter, req *http.Request){
		r.HTML(w, http.StatusOK, "example", "填充")
	})
	http.ListenAndServe(":8080", mux)
}
