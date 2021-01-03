package main

import "net/http"

func createSnippet(w http.ResponseWriter,r *http.Request)  {
	if r.Method!=http.MethodPost{
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w,"Method NOT Allowed",405)
		return
	}
	w.Write([]byte("Create a new snippet"))

}