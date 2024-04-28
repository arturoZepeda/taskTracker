package main

import (
	"fmt"
	"log"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test2312312312")
}
func rootFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Root")
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Query())
	fmt.Println(r.Method)
}

func generaRuta(root *goji.Mux, ruta string, handler http.HandlerFunc) {
	root.Handle(pat.New(ruta), handler)
	ruta = ruta + "/:id"
	root.HandleFunc(pat.Get(ruta), getFunc)
	root.HandleFunc(pat.Post(ruta), postFunc)
	root.HandleFunc(pat.Post(ruta), postFunc)
	root.HandleFunc(pat.Put(ruta), putFunc)
	root.HandleFunc(pat.Delete(ruta), deleteFunc)

}

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
