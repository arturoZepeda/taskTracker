package main

import (
	"fmt"
	"net/http"
)

func getFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test get"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
}
func postFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST path:"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
}
func putFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test PUT"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
}
func deleteFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test DELETE"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
}
