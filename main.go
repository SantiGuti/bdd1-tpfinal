package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	_ "github.com/lib/pq"
)
//
type cliente struct {
	nrocliente                  int
	nombre, apellido, domicilio string
	telefono                    string
}

func main() {
	//Abre la conexión a la base de datos
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetascredito sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//Despliega el menú visible al usuario
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
	var selec int
	fmt.Scanf("%s", &selec)
	//OPCIÓN 1
	if selec == 1 {
		fmt.Printf("\nUsted ha seleccionado la opción 1: Crear una base de datos.\n")
	}
	//Crea una nueva base de datos
	createDatabase()
	//OPCIÓN 2
	if selec == 2 {
		fmt.Printf("\nUsted ha seleccionado la opción 2: Crear las tablas.\n")
	}
	//Lee las tablas (Modelos de datos)
	_, err = db.Exec(leerArchivo("tablas.sql"))
	if err != nil {
		log.Fatal(err)
	}
	//OPCIÓN 3
	if selec == 3 {
		fmt.Printf("\nUsted ha seleccionado la opción 3: Completar las tablas.\n")	
	}
	//OPCIÓN 4
	if selec == 4{
		fmt.Printf("\nUsted ha seleccionado la opción 4: Asignar las PK y FK.\n")
	}
	
	//OPCIÓN 5
	if selec == 5{
		fmt.Printf("\nUsted ha seleccionado la opción 5: Borrar las PK y FK.\n")
	}

	//OPCIÓN 6
	if selec == 6{
		fmt.Printf("\nUsted ha seleccionado la opción 6: Crear las stored procedures y los triggers.\n")
	}

	//OPCIÓN 7
	if selec == 7{
		fmt.Printf("\nUsted ha seleccionado la opción 7: Autorizar las compras.\n")
	}

	//OPCIÓN 8
	if selec == 8{
		fmt.Printf("\nUsted ha seleccionado la opción 8: Generar el resumen de la compra.\n")
	}

	//OPCIÓN 9
	if selec == 9{
		fmt.Printf("\nUsted ha seleccionado la opción 9: Generar BoldDB.\n")
	}
	//Carga las tablas con los datos
	_, err = db.Exec(leerArchivo("datos.sql"))
	if err != nil {
		log.Fatal(err)
	}
	//ESCANEA E IMPRIME LOS DATOS DE LA TABLA CLIENTE	
	//Lectura de la tabla cliente
	rows, err := db.Query(`select * from cliente`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	//Scan de los datos contenidos en la tabla
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
	fmt.Printf("\nNueva base de datos creada.\n")
}

func leerArchivo(archivo string) string {
	datos, err := ioutil.ReadFile(archivo)
	if err != nil {
		log.Fatal(err)
	}
	ret := string(datos)
	return ret
}

func mostrarDatos(archivo string) string {
	tablas, err := ioutil.ReadFile(archivo)
	if err != nil {
		log.Fatal(err)
	}
	contenido := string(tablas)
	fmt.Printf("%s", contenido)
	ret := string(tablas)
	return ret
}