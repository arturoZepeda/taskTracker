package main

import (
	"fmt"
	"net/http"

	"goji.io/pat"
)

func getFunc(w http.ResponseWriter, r *http.Request) {
	if pat.Param(r, "id") == "" {
		switch r.URL.Path {
		case "/":
			fmt.Fprintf(w, "Test GET path:"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
		case "/usuario":
			UsuariosGetFunc(w, r)
			fmt.Fprintf(w, "Test GET path:"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
		case "/actividades":
			fmt.Fprintf(w, "Test GET path:"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
		case "/tipoActividades":
			fmt.Fprintf(w, "Test GET path:"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
		case "/frecuenciaActividades":
			fmt.Fprintf(w, "Test GET path:"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
		case "/registrosActividades":
			fmt.Fprintf(w, "Test GET path:"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
		case "/subregistrosActividades":
			fmt.Fprintf(w, "Test GET path:"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
		}
	} else {
		switch r.URL.Path {
		case "/usuario/1":
			UsuarioGetFunc(w, r)
		case "/actividades/1":
			fmt.Fprintf(w, "Test GET path:"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
		case "/tipoActividades/1":
			fmt.Fprintf(w, "Test GET path:"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
		case "/frecuenciaActividades/1":
			fmt.Fprintf(w, "Test GET path:"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
		case "/registrosActividades/1":
			fmt.Fprintf(w, "Test GET path:"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
		case "/subregistrosActividades/1":
			fmt.Fprintf(w, "Test GET path:"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
		}
	}
}
func postFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/usuario" {
		UsuarioPostFunc(w, r)
	}
	fmt.Fprintf(w, "Test POST path:"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
}
func putFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test PUT"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
}
func deleteFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test DELETE"+r.URL.Path+" query:"+r.URL.Query().Encode()+" method:"+r.Method)
}
