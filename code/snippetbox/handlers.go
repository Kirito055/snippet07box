package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter,r *http.Request)  {
	if r.URL.Path!="/"{
		http.NotFound(w,r)
		return
	}
	files:=[]string{
		"code/snippetbox/home.page.html",
		"code/snippetbox/base.layout.html",
		"code/snippetbox/footer.partial.html",
	}
	ts,err:=template.ParseFiles(files...)
	if err!=nil{
		log.Println(err.Error())
		http.Error(w,"Internal Server Error",500)
		return
	}
	err=ts.Execute(w,nil)
	if err!=nil {
		log.Println(err.Error())
		http.Error(w,"Internal Server Error",500)

	}
}
