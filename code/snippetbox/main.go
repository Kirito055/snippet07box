package  main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request)  {

	w.Write([]byte("Hello from snippetbox"))

}
func showSnippet(w http.ResponseWriter,r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
	http.NotFound(w,r)
		return
	}

fmt.Fprint(w,"Display a scecific snippet with ID %d...",id)
}
func createSnippet(w http.ResponseWriter,r *http.Request)  {
	if r.Method!=http.MethodPost{
		w.Header().Add("Create-control","public")
		w.Header().Add("Create-control"," max-age=645454")
		w.Header().Set("Allow",http.MethodPost)
	http.Error(w,"Method Not Aloowd",405)

		w.Header().Get("Create-control")
		w.Header().Del("Create-control")
		return
	}
	w.Write([]byte("Create new snippet"))
}
func main()  {
	mux:=http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet",showSnippet)
	mux.HandleFunc("/snippet/create",createSnippet)
	log.Println("Starting server on :4001")
	err:=http.ListenAndServe(":4001",mux)
	log.Fatal(err)
	}