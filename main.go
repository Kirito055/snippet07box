package  main

import "net/http"

func home(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("Snippet"))

}

func  main()  {
	mux:=http.NewServeMux()
	mux.HandleFunc("/4000",home)
}