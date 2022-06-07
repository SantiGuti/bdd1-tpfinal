package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	createDatabase()
	fmt.Print("alo")

	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=guarani sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`create table alumne (legajo int, nombre text, apellido text)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`insert into alumne values (1, 'Cristina', 'Kirchner');
					insert into alumne values (2, 'Juan Domingo', 'Peron');`)

	if err != nil {
		log.Fatal(err)
	}
}

func createDatabase() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`create database guarani`)
	if err != nil {
		log.Fatal(err)
	}
}
