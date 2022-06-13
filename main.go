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

type tarjeta struct {
	nrotarjeta   int
	nrocliente   int
	validadesde  int
	validahasta  int
	codseguridad int
	limitecompra float64
	estado       string
}

type comercio struct {
	nrocomercio  int
	nombre       string
	domicilio    string
	codigopostal int
	telefono     int
}

func main() {
	//ABRE LA CONEXIÓN A LA BASE DE DATOS.
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//MUESTRA UN MENÚ VISIBLE PARA EL USUARIO.
	var nombre string
	fmt.Printf("Escriba su nombre: ")
	fmt.Scanf("%s", &nombre)
	fmt.Printf("Hola, %s\n", nombre)
	fmt.Printf("Seleccione un número del siguiente menú:\n")
	menu()
	var selec int
	fmt.Scanln(&selec)

	//OPCIÓN 1: CREAR UNA BASE DE DATOS.
	if selec == 1 {
		fmt.Printf("Usted ha seleccionado la opción 1: Crear una base de datos.\n")
		fmt.Printf("Por favor espere...")
		//Crea una nueva base de datos
		createDatabase(db, err)
	}

	//OPCIÓN 2: CREAR LAS TABLAS.
	if selec == 2 {
		fmt.Printf("\nUsted ha seleccionado la opción 2: Crear las tablas.\n")
		fmt.Printf("\nPor favor espere...\n\n")
		_, err = db.Query(mostrarDatos("tablas.sql"))
		if err != nil {
			log.Fatal(err)
		}
	}

	//OPCIÓN 3: CARGAR LOS DATOS.
	if selec == 3 {
		fmt.Printf("\nUsted ha seleccionado la opción 3: Completar las tablas.\n")
		fmt.Printf("\nPor favor espere...\n")
		//LEE LOS DATOS DE LA TABLA CLIENTE
		fmt.Printf("\nLeyendo la tabla de datos...\n")
		_, err = db.Query(leerArchivo("datos.sql"))
		//Presenta error a partir de este punto. Ver solución. Error: integer out of range
		if err != nil {
			log.Fatal(err)
		}
		//IMPRIME POR PANTALLA LA TABLA CLIENTE
		fmt.Printf("\nDatos de la tabla cliente:\n\n")
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

		//IMPRIME POR PANTALLA LA TABLA TARJETA
		fmt.Printf("\nDatos de la tabla tarjeta:\n\n")
		row, err := db.Query(`select * from tarjeta`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		//Scan de los datos contenidos en la tabla
		var t tarjeta
		for row.Next() {
			if err := rows.Scan(&t.nrotarjeta, &t.nrocliente, &t.validadesde, &t.validahasta, &t.codseguridad, &t.limitecompra, &t.estado); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%v %v %v %v %v %v %v\n", t.nrotarjeta, t.nrocliente, t.validadesde, t.validahasta, t.codseguridad, t.limitecompra, t.estado)
		}
		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}
	}

	//OPCIÓN 4: ASIGNAR LAS PRIMARY KEYS Y FOREIGN KEYS.
	if selec == 4 {
		fmt.Printf("\nUsted ha seleccionado la opción 4: Asignar las PK y FK.\n")
		/*_, err = db.Query(mostrarDatos("PK_FK.sql"))
		if err != nil {
			log.Fatal(err)
		}*/
		//Imprime los datos pero no funciona bien. Error: there ir no unique constraint matching given keys for referenced table "comercio"
		fmt.Printf("\nSe asignará la primary key a la tabla cliente:\n")
		_, err = db.Exec(`alter table cliente add constraint cliente_pk primary key (nrocliente)`)
	}

	//OPCIÓN 5: BORRAR LAS PRIMARY KEYS Y FOREIGN KEYS.
	if selec == 5 {
		fmt.Printf("\nUsted ha seleccionado la opción 5: Borrar las PK y FK.\n")
		fmt.Printf("Si desea eliminar las PK, presione 1. Si desea eliminar las FK, presione 2.\n")
		var selec1 int
		fmt.Scanln(&selec1)
		if selec1 == 1 {
			_, err = db.Exec(`alter table cliente drop primary key`)
		}
		_, err = db.Query(mostrarDatos("PK_FK.sql"))
		if err != nil {
			log.Fatal(err)
		}
	}

	//OPCIÓN 6: AUTORIZAR LAS COMPRAS.
	if selec == 6 {
		fmt.Printf("\nUsted ha seleccionado la opción 6: Autorizar las compras.\n")
		//Tomo 7 casos de consumo para abarcar todas las posibilidades.

	}

	//OPCIÓN 7: GENERAR EL RESUMEN DE LAS COMPRAS.
	if selec == 7 {
		fmt.Printf("\nUsted ha seleccionado la opción 7: Generar el resumen de las compras.\n")
		fmt.Printf("\nPor favor, ingrese el número de cliente:")
		var nrocli int
		fmt.Scanf("%s", &nrocli)
		fmt.Printf("\nIngrese el periodo del año que desea generar el resumen:")
		//var periodo string
		//fmt.Scanf("%s", &periodo)
		_, err = db.Query(`select nrocliente from cliente where nrocliente == &nrocli`)
		_, err = db.Query(`select `)
		_, err = db.Exec(`insert into cabecera values()`)
		_, err = db.Exec(`insert into detalle values()`)
	}

	//OPCIÓN 8: GENERAR ALERTAS A LOS CLIENTES.
	if selec == 8 {
		fmt.Printf("\nUsted ha seleccionado la opción 8: Generar alertas a los clientes.\n")
	}

	//OPCIÓN 9: GENERAR DATOS EN BOLDDB.
	if selec == 9 {
		fmt.Printf("\nUsted ha seleccionado la opción 9: Generar datos en BoldDB.\n")
	}
}

func createDatabase(db *sql.DB, err error) {
	_, err = db.Exec(`drop database if exists tarjetascredito`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`create database tarjetascredito`)
	if err != nil {
		log.Fatal(err)
	}
	db, err = sql.Open("postgres", "user=postgres host=localhost dbname=tarjetascredito sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
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

func menu() {
	fmt.Printf("1. Crear una nueva base de datos.\n")
	fmt.Printf("2. Crear las tablas.\n")
	fmt.Printf("3. Completar las tablas.\n")
	fmt.Printf("4. Asignar las PK y FK.\n")
	fmt.Printf("5. Borrar las PK y FK.\n")
	fmt.Printf("6. Autorizar las compras.\n")
	fmt.Printf("7. Generar el resumen de las compra.\n")
	fmt.Printf("8. Generar alertas a clientes.\n")
	fmt.Printf("9. Generar datos en BoldDB.\n")
}
