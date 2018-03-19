package main

import (
	"net/http"
	"io"
)

type hotdog int

func (d hotdog)ServeHTTP(res http.ResponseWriter, req *http.Request)  {
	io.WriteString(res, "dog dog dogggy")
}

type hotcat int

func (c hotcat)ServeHTTP(res http.ResponseWriter, req *http.Request)  {
	io.WriteString(res, "cat cat cat")
}


func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat/", c)

	http.ListenAndServe(":8080", mux)
}
