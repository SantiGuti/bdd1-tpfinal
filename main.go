package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

type cliente struct {
	nrocliente                  int
	nombre, apellido, domicilio string
	telefono                    string
}

func main() {
	createDatabase() //Crea una nueva base de datos
	menu()
	//Conecta con la base de datos creada
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetascredito sslmode=disable")
	if err != nil {
		log.Fatal(err)

	}
	defer db.Close()

	//Lee las tablas (Modelos de datos)
	_, err = db.Exec(leerArchivo("tablas.sql"))
	if err != nil {
		log.Fatal(err)
	}

	//Carga las tablas con los datos
	_, err = db.Exec(leerArchivo("datos.sql"))
	if err != nil {
		log.Fatal(err)
	}

	////////////////////////////////////////////////////////////////
	// print de prueba lectura tabla cliente //

	rows, err := db.Query(`select * from cliente`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var c cliente
	for rows.Next() {
		if err := rows.Scan(&c.nrocliente, &c.nombre, &c.apellido, &c.domicilio, &c.telefono); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %v %v %v %v\n", c.nrocliente, c.nombre, c.apellido, c.domicilio, c.telefono)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

}

/*_, err = db.Exec(`create table cliente(nrocliente int, nombre text, apellido text, domicilio text, telefono char(12))`)
if err != nil {
	log.Fatal(err)
}*/

/*	_, err = db.Exec(`insert into cliente values (1, 'Juan', 'Lopez', 'Urquiza 2629', '1157845695');
	                  insert int cliente values (2, 'Maria', 'Ramirez', 'Oribe 1576', '1159674130')`) //ejemplo

	if err != nil {
		log.Fatal(err)
	}*/

/*
	rows, err := db.Query(`select * from cliente`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var c cliente

	for rows.Next() {
		if err := rows.Scan(&c.nrocliente, &c.nombre, &c.apellido, &c.domicilio, &c.telefono); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %v %v\n", c.nrocliente, c.nombre, c.apellido, c.domicilio, c.telefono)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}*/

func createDatabase() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`drop database if exists tarjetascredito`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`create database tarjetascredito`)
	if err != nil {
		log.Fatal(err)
	}
}

func leerArchivo(archivo string) string {
	datos, err := ioutil.ReadFile(archivo)
	if err != nil {
		log.Fatal(err)
	}
	ret := string(datos)
	return ret
}

func menu(){

	var nombre string
	fmt.Printf("Escriba su nombre: ")
	fmt.Scanf("%s", &nombre)
	fmt.Printf("Hola, %s\n",nombre)
	fmt.Printf("Seleccione un número del siguiente menú:\n")
	fmt.Printf("1. Crear una nueva base de datos.\n")
	fmt.Printf("2. Crear las tablas.\n")
	fmt.Printf("3. Completar las tablas.\n")
	fmt.Printf("4. Asignar las PK y FK.\n")
	fmt.Printf("5. Borrar las PK y FK.\n")
	fmt.Printf("6. Crear las stored procedures y los triggers.\n")
	fmt.Printf("7. Autorizar las compras.\n")
	fmt.Printf("8. Generar el resumen de la compra.\n")
	fmt.Printf("9. Generar BoldDB.\n")

}