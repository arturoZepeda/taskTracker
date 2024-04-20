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
	w.Header().Set("Content-Type", "application/json")
	var usuario UsuarioDB
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = PostUsuarioDB(usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(usuario)

}
func UsuarioPutFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Usuario Put ID %s", pat.Param(r, "id"))
}
func UsuarioDeleteFunc(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(pat.Param(r, "id"))
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	fmt.Fprintf(w, "Usuario Delete ID %s", pat.Param(r, "id"))
	error := DeleteUsuarioDB(id)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", error)
		return
	}

}
