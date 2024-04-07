package main

import (
	"fmt"
	"net/http"

	"goji.io/pat"
)

func UsuarioFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Usuario")
}
func UsuarioPostFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Usuario Post ID %s", pat.Param(r, "id"))
}
func UsuarioPutFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Usuario Put ID %s", pat.Param(r, "id"))
}
func UsuarioDeleteFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Usuario Delete ID %s", pat.Param(r, "id"))
}
