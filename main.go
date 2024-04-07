package main

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func rootFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Root")
}

func main() {
	fmt.Println("Server started at localhost:8000")
	root := goji.NewMux()
	usuario := goji.SubMux()
	root.Handle(pat.New("/usuario/*"), usuario)
	root.HandleFunc(pat.Get("/"), rootFunc)
	usuario.HandleFunc(pat.Post("/:id"), UsuarioPostFunc)
	usuario.HandleFunc(pat.Get("/"), UsuarioFunc)
	usuario.HandleFunc(pat.Put("/:id"), UsuarioPutFunc)
	usuario.HandleFunc(pat.Delete("/:id"), UsuarioDeleteFunc)

	//mux := goji.NewMux()
	//mux.HandleFunc(pat.Get("/hello/:name"), hello)

	http.ListenAndServe("localhost:8000", root)
}
