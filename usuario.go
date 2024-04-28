package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	// Verificar el tipo de contenido
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	var usuario UsuarioDB
	err := json.NewDecoder(r.Body).Decode(&usuario)
	// Cerrar el cuerpo de la solicitud
	defer r.Body.Close()

	if err != nil {
		// Registrar el error en el servidor
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//agrega debug para ver el contenido de usuario
	fmt.Println(usuario)
	err = PostUsuarioDB(usuario)
	if err != nil {
		// Registrar el error en el servidor
		log.Printf("Error posting to database: %v", err)
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
