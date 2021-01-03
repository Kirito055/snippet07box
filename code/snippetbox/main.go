package  main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello from snippetbossssx"))

}
func showSnippet(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("Display a specific snippet..."))
}
func createSnippet(w http.ResponseWriter,r *http.Request)  {
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