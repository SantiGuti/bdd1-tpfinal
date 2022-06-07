package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type cliente struct {
	nroCliente                            int
	nombre, apellido, domicilio, telefono string
	//telefono string??? en el enunciado es char(12)
}

func createDatabase() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`create database tarjetascredito`)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database creada correctamente.")
	}
}

func main() {
	// acá va el código de mi "aplicacion-de-bases-de-datos"

	createDatabase()

	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetascredito sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`create table cliente(nrocliente int, nombre text, apellido text, domicilio text, telefono char(12))`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`insert into cliente values (1, 'Juan', 'Lopez', 'Urquiza 2629', '1157845695');
	                		`)

	if err != nil {
		log.Fatal(err)
	}
	/*
		rows, err := db.Query(`select * from cliente`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var c cliente

		for rows.Next() {
			if err := rows.Scan(&c.nroCliente, &c.nombre, &c.apellido, &c.domicilio, &c.telefono); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%v %v %v\n", c.nroCliente, c.nombre, c.apellido, c.domicilio, c.telefono)
		}
		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}*/
}
