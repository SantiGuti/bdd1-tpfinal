package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

//DEFINIMOS LOS TIPOS DE DATOS
type cliente struct {
	nrocliente                  int
	nombre, apellido, domicilio string
	telefono                    string
}

type tarjeta struct {
	nrotarjeta, nrocliente, validadesde, validahasta, codseguridad int
	limitecompra                                                   float64
	estado                                                         string
}

type comercio struct {
	nrocomercio  int
	nombre       string
	domicilio    string
	codigopostal string
	telefono     string
}

type compra struct {
	nrooperacion int
	nrotarjeta   string
	nrocomercio  int
	fecha        string
	monto        float64
	pagado       bool
}

type rechazo struct {
	nrorechazo  int
	nrotarjeta  string
	nrocomercio int
	fecha       string
	monto       float64
	motivo      string
}

type detalle struct {
	nroresumen     int
	nrolinea       int
	fecha          string
	nombrecomercio string
	monto          float64
}
type cabecera struct {
	nroresumen int
	nombre     string
	apellido   string
	domicilio  string
	nrotarjeta string
	desde      string
	hasta      string
	vence      string
	total      float64
}
type alerta struct {
	nroalerta   int
	nrotarjeta  string
	fecha       string
	nrorechazo  int
	codalerta   int
	descripcion string
}
type alerta2 struct {
	nroalerta   int
	nrotarjeta  string
	fecha       string
	codalerta   int
	descripcion string
}

func main() {
	//ABRE LA CONEXIÓN A LA BASE DE DATOS.
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetascredito sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var nombre string
	fmt.Printf("Escriba su nombre: ")
	fmt.Scanf("%s", &nombre)
	fmt.Printf("Hola, %s\n", nombre)

	var leave = false

	//MUESTRA UN MENÚ VISIBLE PARA EL USUARIO.
	for !leave {
		fmt.Printf("Seleccione un número del siguiente menú:\n")
		menu()
		var selec int
		fmt.Scanln(&selec)

		//OPCIÓN 1: CREAR UNA BASE DE DATOS.
		if selec == 1 {
			fmt.Printf("Usted ha seleccionado la opción 1: Crear una base de datos.\n")
			fmt.Printf("Por favor espere...")
			//Crea una nueva base de datos
			createDatabase()
		}

		//OPCIÓN 2: CREAR LAS TABLAS.
		if selec == 2 {
			fmt.Printf("\nUsted ha seleccionado la opción 2: Crear las tablas.\n")
			fmt.Printf("\nPor favor espere...\n\n")
			fmt.Printf("\nlisto\n\n")
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

			//IMPRIME POR PANTALLA LA TABLA COMERCIO
			fmt.Printf("\nDatos de la tabla comercio:\n\n")
			row, err := db.Query(`select * from comercio`)
			if err != nil {
				log.Fatal(err)
			}
			defer row.Close()
			//Scan de los datos contenidos en la tabla
			var co comercio
			for row.Next() {
				if err := row.Scan(&co.nrocomercio, &co.nombre, &co.domicilio, &co.codigopostal, &co.telefono); err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%v %v %v %v %v\n", co.nrocomercio, co.nombre, co.domicilio, co.codigopostal, co.telefono)
			}
			if err = row.Err(); err != nil {
				log.Fatal(err)
			}
		}

		//OPCIÓN 4: ASIGNAR LAS PRIMARY KEYS Y FOREIGN KEYS.
		if selec == 4 {
			fmt.Printf("\nUsted ha seleccionado la opción 4: Asignar las PK y FK.\n")
			_, err = db.Query(mostrarDatos("PK_FK.sql"))
			if err != nil {
				log.Fatal(err)
			}
			//Imprime los datos pero no funciona bien. Error: there ir no unique constraint matching given keys for referenced table "comercio"
			//fmt.Printf("\nSe asignará la primary key a la tabla cliente:\n")
			//_, err = db.Exec(`alter table cliente add constraint cliente_pk primary key (nrocliente)`)
		}

		//OPCIÓN 5: BORRAR LAS PRIMARY KEYS Y FOREIGN KEYS.
		if selec == 5 {
			fmt.Printf("\nUsted ha seleccionado la opción 5: Borrar las PK y FK.\n")
			_, err = db.Exec(leerArchivo("drop_pk_fk.sql"))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\nPK Y FK eliminadas.\n")
		}

		//OPCIÓN 6: CARGAR FUNCIONES.
		if selec == 6 {
			fmt.Printf("\nUsted ha seleccionado la opción 6: Cargar funciones.\n")
			_, err = db.Query(leerArchivo("SP&T.sql"))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\nFunciones cargadas.\n")
		}

		//OPCIÓN 7: AUTORIZAR LAS COMPRAS.
		if selec == 7 {
			fmt.Printf("\nUsted ha seleccionado la opción 7: Autorizar las compras.\n")
			_, err = db.Exec(`select autorizar_compras('4929597785365045', '6235', 011, 500.00)`)
			if err != nil {
				log.Fatal(err)
			}

			rows, err := db.Query(`select * from rechazo`)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			// Scan de los datos contenidos en la tabla
			var r rechazo
			for rows.Next() {
				if err := rows.Scan(&r.nrorechazo, &r.nrotarjeta, &r.nrocomercio, &r.fecha, &r.monto, &r.motivo); err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%v %v %v %v %v %v\n", r.nrorechazo, r.nrotarjeta, r.nrocomercio, r.fecha, r.monto, r.motivo)
			}
			if err = rows.Err(); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("\n------------------------------\n")
			fmt.Printf("\nTodas las alertas por rechazo:\n")

			rows, err = db.Query(`select * from alerta`)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			// Scan de los datos contenidos en la tabla
			var a alerta
			for rows.Next() {
				if err := rows.Scan(&a.nroalerta, &a.nrotarjeta, &a.fecha, &a.nrorechazo, &a.codalerta, &a.descripcion); err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%v %v %v %v %v %v\n", a.nroalerta, a.nrotarjeta, a.fecha, a.nrorechazo, a.codalerta, a.descripcion)
			}
			if err = rows.Err(); err != nil {
				log.Fatal(err)
			}
		}

		//OPCIÓN 8: GENERAR EL RESUMEN DE LAS COMPRAS.
		if selec == 8 {
			fmt.Printf("\nUsted ha seleccionado la opción 8: Generar el resumen de las compras.\n")
			_, err = db.Exec(`select generar_resumen(01, '202205')`)
			if err != nil {
				log.Fatal(err)
			}
			// fmt.Printf("\nPor favor, ingrese el número de cliente: ")
			// var nrocli int
			// fmt.Scanf("%s", &nrocli)
			// fmt.Printf("\nIngrese el periodo del año que desea generar el resumen:")
			// var fecha string
			// fmt.Scanf("%s", &fecha)
			fmt.Print("TABLA DETALLE VALORES ACTUALES: ")
			rows, err := db.Query(`select * from detalle`)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			// Scan de los datos contenidos en la tabla
			var d detalle
			for rows.Next() {
				if err := rows.Scan(&d.nroresumen, &d.nrolinea, &d.fecha, &d.nombrecomercio, &d.monto); err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%v %v %v %v %v \n", d.nroresumen, d.nrolinea, d.fecha, d.nombrecomercio, d.monto)
			}
			if err = rows.Err(); err != nil {
				log.Fatal(err)
			}

			fmt.Print("\nTABLA CABECERA VALORES ACTUALES: ")
			rows, err = db.Query(`select * from cabecera`)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			// Scan de los datos contenidos en la tabla
			var c cabecera
			for rows.Next() {
				if err := rows.Scan(&c.nroresumen, &c.nombre, &c.apellido, &c.domicilio, &c.nrotarjeta, &c.desde, &c.hasta, &c.vence, &c.total); err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%v %v %v %v %v %v %v %v %v\n", c.nroresumen, c.nombre, c.apellido, c.domicilio, c.nrotarjeta, c.desde, c.hasta, c.vence, c.total)
			}
			if err = rows.Err(); err != nil {
				log.Fatal(err)
			}
		}

		//OPCIÓN 9: GENERAR DATOS EN BOLDDB.
		if selec == 9 {
			fmt.Printf("\nUsted ha seleccionado la opción 9: Generar datos en BoldDB.\n")
		}

		//OPCIÓN 10: TESTEO 2 COMPRAS EN MENOS DE 1 MIN.
		if selec == 10 {
			fmt.Printf("\nTesteo 2 compras en menos de 1 min.\n")
			_, err = db.Query(`select autorizar_compras('5543040397793513', '4172', 017, 100.00)`)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\n1ra compra autorizada.\n")
			_, err = db.Query(`select autorizar_compras('5543040397793513', '4172', 016, 200.00)`)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\n2da compra autorizada.\n")
			rows, err := db.Query(`select nroalerta, nrotarjeta, fecha, codalerta, descripcion from alerta`)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			// Scan de los datos contenidos en la tabla
			var a2 alerta2
			for rows.Next() {
				if err := rows.Scan(&a2.nroalerta, &a2.nrotarjeta, &a2.fecha, &a2.codalerta, &a2.descripcion); err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%v %v %v %v %v \n", a2.nroalerta, a2.nrotarjeta, a2.fecha, a2.codalerta, a2.descripcion)
			}
			if err = rows.Err(); err != nil {
				log.Fatal(err)
			}
		}

		//OPCIÓN 11: TESTEO 2 COMPRAS EN MENOS DE 5 MIN.
		if selec == 11 {
			fmt.Printf("\nTesteo 2 compras en menos de 5 min.\n")
			_, err = db.Query(`select autorizar_compras('4823836840552412', '8748', 03, 900.00)`)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\n1ra compra autorizada.\n")
			_, err = db.Query(`select autorizar_compras('4823836840552412', '8748', 020, 500.00)`)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\n2da compra autorizada.\n")
			rows, err := db.Query(`select nroalerta, nrotarjeta, fecha, codalerta, descripcion from alerta`)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			// Scan de los datos contenidos en la tabla
			var a2 alerta2
			for rows.Next() {
				if err := rows.Scan(&a2.nroalerta, &a2.nrotarjeta, &a2.fecha, &a2.codalerta, &a2.descripcion); err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%v %v %v %v %v \n", a2.nroalerta, a2.nrotarjeta, a2.fecha, a2.codalerta, a2.descripcion)
			}
			if err = rows.Err(); err != nil {
				log.Fatal(err)
			}
		}

		//OPCIÓN 12: TESTEO ALERTA POR 2 LIMITES DE COMPRA MISMO DIA.

		if selec == 0 {
			leave = true
		}
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
	db, err = sql.Open("postgres", "user=postgres host=localhost dbname=tarjetascredito sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Printf("\nNueva base de datos creada.\n")
}

//LECTURA DE ARCHIVOS
func leerArchivo(archivo string) string {
	datos, err := ioutil.ReadFile(archivo)
	if err != nil {
		log.Fatal(err)
	}
	ret := string(datos)
	return ret
}

//IMPRIME POR PANTALLA EL CONTENIDO DE UN ARCHIVO
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

//MENÚ VISIBLE AL USUARIO
func menu() {
	fmt.Printf("1. Crear una nueva base de datos.\n")
	fmt.Printf("2. Crear las tablas.\n")
	fmt.Printf("3. Completar las tablas.\n")
	fmt.Printf("4. Asignar las PK y FK.\n")
	fmt.Printf("5. Borrar las PK y FK.\n")
	fmt.Printf("6. Cargar funciones.\n")
	fmt.Printf("7. Autorizar las compras.\n")
	fmt.Printf("8. Generar el resumen de las compra.\n")
	fmt.Printf("9. Generar datos en BoldDB.\n")
	fmt.Printf("10. TESTEO 2 COMPRAS EN MENOS DE 1 MIN.\n")
	fmt.Printf("11. TESTEO 2 COMPRAS EN MENOS DE 5 MIN.\n")
	fmt.Printf("12. TESTEO ALERTA POR 2 LIMITES DE COMPRA.\n")
	fmt.Printf("Escriba 0 para salir.\n")
}
