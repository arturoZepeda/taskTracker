package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type UsuarioDB struct {
	ID     int
	Nombre string
	Correo string
}

type DB struct {
	Database *sql.DB
}

func (db *DB) GetUsuario(id int) (UsuarioDB, error) {
	var usuario UsuarioDB
	err := db.Database.QueryRow("SELECT id, nombre, correo FROM usuarios WHERE id = ?", id).Scan(&usuario.ID, &usuario.Nombre, &usuario.Correo)
	return usuario, err
}

func GetUsuarioDB(id int) (UsuarioDB, error) {
	db, err := sql.Open("sqlite3", "./usuarios.db")
	if err != nil {
		return UsuarioDB{}, err
	}
	defer db.Close()
	usuario := UsuarioDB{}
	err = db.QueryRow("SELECT id, nombre, correo FROM usuarios WHERE id = ?", id).Scan(&usuario.ID, &usuario.Nombre, &usuario.Correo)
	if err != nil {
		return UsuarioDB{}, err
	}
	return usuario, nil
}

func GetUsuariosDB() ([]UsuarioDB, error) {
	db, err := sql.Open("sqlite3", "./usuarios.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, nombre, correo FROM usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usuarios := []UsuarioDB{}
	for rows.Next() {
		var u UsuarioDB
		err := rows.Scan(&u.ID, &u.Nombre, &u.Correo)
		if err != nil {
			return nil, err
		}
		usuarios = append(usuarios, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return usuarios, nil
}

func main() {
	db, err := sql.Open("sqlite3", "./usuarios.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS usuarios (id INTEGER PRIMARY KEY, nombre TEXT, correo TEXT)")
	if err != nil {
		panic(err)
	}

	db.Exec("INSERT INTO usuarios (nombre, correo) VALUES ('Juan', 'aaze92@gmai.com')")
	fmt.Println("Server started at localhost:8000")

	usuarios, err := GetUsuariosDB()
	if err != nil {
		log.Fatal(err)
	}

	for _, usuario := range usuarios {
		fmt.Println(usuario.ID, usuario.Nombre, usuario.Correo)
	}
	var idUsuario int
	idUsuario = 1
	usuario, err := GetUsuarioDB(idUsuario)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Usuario con ID 1:")
	fmt.Println(usuario.ID, usuario.Nombre, usuario.Correo)
}
