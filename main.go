package main

import (
	"fmt"
	"log"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func main() {
	fmt.Println("Server started at localhost:8000")
	root := goji.NewMux()
	usuario := goji.SubMux()
	actividades := goji.SubMux()
	tipoActividades := goji.SubMux()
	frecuenciaActividades := goji.SubMux()
	registrosActividades := goji.SubMux()
	subregistrosActividades := goji.SubMux()

	err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	//Generamos rutas para cada submódulo
	root.Handle(pat.New("/usuario/*"), usuario)
	root.Handle(pat.New("/actividades/*"), actividades)
	root.Handle(pat.New("/tipoActividades/*"), tipoActividades)
	root.Handle(pat.New("/frecuenciaActividades/*"), frecuenciaActividades)
	root.Handle(pat.New("/registrosActividades/*"), registrosActividades)
	root.Handle(pat.New("/subregistrosActividades/*"), subregistrosActividades)

	//Generamos request para cada submódulo
	//Usuarios
	usuario.HandleFunc(pat.Get("/"), UsuariosGetFunc)
	usuario.HandleFunc(pat.Post("/"), UsuarioPostFunc)
	usuario.HandleFunc(pat.Get("/:id"), UsuariosGetFunc)
	usuario.HandleFunc(pat.Put("/:id"), UsuarioPutFunc)
	usuario.HandleFunc(pat.Delete("/:id"), UsuarioDeleteFunc)

	//Actividades
	actividades.HandleFunc(pat.Get("/"), getFunc)
	actividades.HandleFunc(pat.Post("/"), postFunc)
	actividades.HandleFunc(pat.Get("/:id"), getFunc)
	actividades.HandleFunc(pat.Put("/:id"), putFunc)
	actividades.HandleFunc(pat.Delete("/:id"), deleteFunc)

	//TipoActividades
	tipoActividades.HandleFunc(pat.Get("/"), getFunc)
	tipoActividades.HandleFunc(pat.Post("/"), postFunc)
	tipoActividades.HandleFunc(pat.Get("/:id"), getFunc)
	tipoActividades.HandleFunc(pat.Put("/:id"), putFunc)
	tipoActividades.HandleFunc(pat.Delete("/:id"), deleteFunc)

	//FrecuenciaActividades
	frecuenciaActividades.HandleFunc(pat.Get("/"), getFunc)
	frecuenciaActividades.HandleFunc(pat.Post("/"), postFunc)
	frecuenciaActividades.HandleFunc(pat.Get("/:id"), getFunc)
	frecuenciaActividades.HandleFunc(pat.Put("/:id"), putFunc)
	frecuenciaActividades.HandleFunc(pat.Delete("/:id"), deleteFunc)

	http.ListenAndServe("localhost:8000", root)
}
