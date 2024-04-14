package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"goji.io/pat"
)

func UsuariosGetFunc(w http.ResponseWriter, r *http.Request) {
	var usuariosResponse []UsuarioDB
	usuarios, err := GetUsuariosDB()
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	usuariosResponse = append(usuariosResponse, usuarios...)
	response, err := json.Marshal(usuariosResponse)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	fmt.Fprintf(w, "%s", response)
}

func UsuarioGetFunc(w http.ResponseWriter, r *http.Request) {
	var usuarioResponse []UsuarioDB
	id, err := strconv.Atoi(pat.Param(r, "id"))
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	usuario, err := GetUsuarioDB(id)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	usuarioResponse = append(usuarioResponse, usuario)
	response, err := json.Marshal(usuarioResponse)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	fmt.Fprintf(w, "%s", response)

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
