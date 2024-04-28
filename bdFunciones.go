package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var detalleQuery = []string{
	`CREATE TABLE Usuarios (
		ID INTEGER PRIMARY KEY,
		nombre TEXT,
		correo TEXT
	);`,
	`CREATE TABLE Actividades (
		ID INTEGER PRIMARY KEY,
		nombre TEXT,
		descripcion TEXT,
		nombre_subregistro TEXT,
		id_tipo INTEGER,
		FOREIGN KEY(id_tipo) REFERENCES TipoActividad(ID)
	);`,
	`CREATE TABLE TipoActividad (
		ID INTEGER PRIMARY KEY,
		nombre TEXT
	);`,
	`CREATE TABLE Frecuencia (
		ID INTEGER PRIMARY KEY,
		frecuencia TEXT
	);`,
	`CREATE TABLE Registros (
		ID INTEGER PRIMARY KEY,
		ID_usuario INTEGER,
		ID_actividad INTEGER,
		ID_frecuencia INTEGER,
		fecha TEXT,
		estado INTEGER,
		FOREIGN KEY(ID_usuario) REFERENCES Usuarios(ID),
		FOREIGN KEY(ID_actividad) REFERENCES Actividades(ID),
		FOREIGN KEY(ID_frecuencia) REFERENCES Frecuencia(ID)
	);`,
	`CREATE TABLE Subregistros (
		ID INTEGER PRIMARY KEY,
		nombre TEXT,
		descripcion TEXT,
		ID_registro INTEGER,
		ID_actividad INTEGER,
		FOREIGN KEY(ID_registro) REFERENCES Registros(ID),
		FOREIGN KEY(ID_actividad) REFERENCES Actividades(ID)
	);`,
}

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

func PostUsuarioDB(u UsuarioDB) error {
	fmt.Println(u)
	var nombre string = u.Nombre
	var correo string = u.Correo
	db, err := sql.Open("sqlite3", "./usuarios.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO usuarios (nombre, correo) VALUES (?, ?)", nombre, correo)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUsuarioDB(id int) error {
	db, err := sql.Open("sqlite3", "./usuarios.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM usuarios WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func initDB() error {
	db, err := sql.Open("sqlite3", "./usuarios.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	for _, query := range detalleQuery {
		_, err = db.Exec(query)
		if err != nil {
			println("Error al ejecutar la sentencia SQL: %s", err)
			continue
		}
	}
	mensaje := "Tabla usuarios creada"
	println(mensaje)
	return nil
}
