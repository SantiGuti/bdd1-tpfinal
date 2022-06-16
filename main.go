package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
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

//ESTRUCTURAS DE DATOS PARA JSON (DATOS PÚBLICOS)
type clientes struct {
	Nrocliente                  int
	Nombre, Apellido, Domicilio string
	Telefono                    string
}

type tarjetas struct {
	Nrotarjeta, Nrocliente, Validadesde, Validahasta, Codseguridad int
	Limitecompra                                                   float64
	Estado                                                         string
}

type comercios struct {
	Nrocomercio  int
	Nombre       string
	Domicilio    string
	Codigopostal string
	Telefono     string
}

type compras struct {
	Nrooperacion int
	Nrotarjeta   int
	Nrocomercio  int
	Fecha        string
	Monto        float64
	Pagado       bool
}

func main() {
	//ABRE LA CONEXIÓN A LA BASE DE DATOS.
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetascredito sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Printf("¡Bienvenido!\n\n")

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

			//Continuar?
			fmt.Printf("\n¿Desea continuar operando? Presione y/n\n")
			var consulta string
			fmt.Scan(&consulta)
			if consulta == "y" {
				continue
			}
			break

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

			//Continuar?
			fmt.Printf("\n¿Desea continuar operando? Presione y/n\n")
			var consulta string
			fmt.Scan(&consulta)
			if consulta == "y" {
				continue
			}
			break
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

			//Continuar?
			fmt.Printf("\n¿Desea continuar operando? Presione y/n\n")
			var consulta string
			fmt.Scan(&consulta)
			if consulta == "y" {
				continue
			}
			break
		}

		//OPCIÓN 4: ASIGNAR LAS PRIMARY KEYS Y FOREIGN KEYS.
		if selec == 4 {
			fmt.Printf("\nUsted ha seleccionado la opción 4: Asignar las PK y FK.\n")
			_, err = db.Query(mostrarDatos("PK_FK.sql"))
			if err != nil {
				log.Fatal(err)
			}

			//Continuar?
			fmt.Printf("\n¿Desea continuar operando? Presione y/n\n")
			var consulta string
			fmt.Scan(&consulta)
			if consulta == "y" {
				continue
			}
			break
		}

		//OPCIÓN 5: BORRAR LAS PRIMARY KEYS Y FOREIGN KEYS.
		if selec == 5 {
			fmt.Printf("\nUsted ha seleccionado la opción 5: Borrar las PK y FK.\n")
			_, err = db.Exec(leerArchivo("drop_pk_fk.sql"))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\nPK Y FK eliminadas.\n")

			//Continuar?
			fmt.Printf("\n¿Desea continuar operando? Presione y/n\n")
			var consulta string
			fmt.Scan(&consulta)
			if consulta == "y" {
				continue
			}
			break
		}

		//OPCIÓN 6: CARGAR FUNCIONES.
		if selec == 6 {
			fmt.Printf("\nUsted ha seleccionado la opción 6: Cargar funciones.\n")
			_, err = db.Query(leerArchivo("SP&T.sql"))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\nFunciones cargadas.\n")

			//Continuar?
			fmt.Printf("\n¿Desea continuar operando? Presione y/n\n")
			var consulta string
			fmt.Scan(&consulta)
			if consulta == "y" {
				continue
			}
			break
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
			fmt.Printf("\n\n")

			//Continuar?
			fmt.Printf("\n¿Desea continuar operando? Presione y/n\n")
			var consulta string
			fmt.Scan(&consulta)
			if consulta == "y" {
				continue
			}
			break
		}

		//OPCIÓN 8: GENERAR EL RESUMEN DE LAS COMPRAS.
		if selec == 8 {
			fmt.Printf("\nUsted ha seleccionado la opción 8: Generar el resumen de las compras.\n")
			_, err = db.Exec(`select generar_resumen(01, '202205')`)
			if err != nil {
				log.Fatal(err)
			}
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
			//Continuar?
			fmt.Printf("\n¿Desea continuar operando? Presione y/n\n")
			var consulta string
			fmt.Scan(&consulta)
			if consulta == "y" {
				continue
			}
			break
		}

		//OPCIÓN 9: GENERAR DATOS EN BOLDDB.
		if selec == 9 {
			/*CONECTA CON LA BASE DE DATOS BOLT*/
			fmt.Printf("\nUsted ha seleccionado la opción 9: Crear una base de datos Bolt.\n")
			db, err := bolt.Open("tdbpi.boltdb", 0600, nil)
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()
			fmt.Printf("\nCreando nueva base de datos...\n")
			fmt.Printf("\nNueva base de datos Bolt creada.\n")

			/*TRES CLIENTES*/
			fmt.Printf("\nPresione 1 para cargar los datos de clientes.\n\n")
			var client int
			fmt.Scan(&client)
			fmt.Printf("\nCargando datos de tres clientes...\n\n")

			//Cliente 1
			cliente1 := clientes{Nrocliente: 1, Nombre: "Juan", Apellido: "Rosas", Domicilio: "Serano 701", Telefono: "011-68943567"}
			data, err := json.Marshal(cliente1)
			if err != nil {
				panic(err)
			}
			CreateUpdate(db, "cliente", []byte(strconv.Itoa(cliente1.Nrocliente)), data)
			resultado, err := ReadUnique(db, "cliente", []byte(strconv.Itoa(cliente1.Nrocliente)))
			fmt.Printf("%s\n", resultado)

			//Cliente 2
			cliente2 := clientes{Nrocliente: 2, Nombre: "Martin", Apellido: "Valdez", Domicilio: "Mendoza 293", Telefono: "011-78908583"}
			data2, err := json.Marshal(cliente2)
			if err != nil {
				panic(err)
			}
			CreateUpdate(db, "cliente", []byte(strconv.Itoa(cliente2.Nrocliente)), data2)
			resultado2, err := ReadUnique(db, "cliente", []byte(strconv.Itoa(cliente2.Nrocliente)))
			fmt.Printf("%s\n", resultado2)

			//Cliente 3
			cliente3 := clientes{Nrocliente: 3, Nombre: "Roberto", Apellido: "Gonzalez", Domicilio: "Las Heras 552", Telefono: "011-23587387"}
			cli3, err := json.Marshal(cliente3)
			if err != nil {
				panic(err)
			}
			CreateUpdate(db, "cliente", []byte(strconv.Itoa(cliente3.Nrocliente)), cli3)
			resulcli3, err := ReadUnique(db, "cliente", []byte(strconv.Itoa(cliente3.Nrocliente)))
			fmt.Printf("%s\n", resulcli3)

			/*TRES TARJETAS*/
			fmt.Printf("\nPresione 2 para cargar los datos de tarjetas.\n\n")
			var card int
			fmt.Scan(&card)
			fmt.Printf("\nCargando datos de tres tarjetas...\n\n")

			//Tarjeta 1
			tarjeta1 := tarjetas{Nrotarjeta: 4756326984155476, Nrocliente: 1, Validadesde: 201807, Validahasta: 202302, Codseguridad: 6713, Limitecompra: 500000.00, Estado: "vigente"}
			tar1, err := json.Marshal(tarjeta1)
			if err != nil {
				panic(err)
			}
			CreateUpdate(db, "tarjeta", []byte(strconv.Itoa(tarjeta1.Nrotarjeta)), tar1)
			resultar1, err := ReadUnique(db, "tarjeta", []byte(strconv.Itoa(tarjeta1.Nrotarjeta)))
			fmt.Printf("%s\n", resultar1)

			//Tarjeta 2
			tarjeta2 := tarjetas{Nrotarjeta: 4532969538877007, Nrocliente: 2, Validadesde: 202003, Validahasta: 202504, Codseguridad: 6646, Limitecompra: 200000.00, Estado: "vigente"}
			tar2, err := json.Marshal(tarjeta2)
			if err != nil {
				panic(err)
			}
			CreateUpdate(db, "tarjeta", []byte(strconv.Itoa(tarjeta2.Nrotarjeta)), tar2)
			resultar2, err := ReadUnique(db, "tarjeta", []byte(strconv.Itoa(tarjeta2.Nrotarjeta)))
			fmt.Printf("%s\n", resultar2)

			//Tarjeta 3
			tarjeta3 := tarjetas{Nrotarjeta: 4929941716451245, Nrocliente: 3, Validadesde: 202204, Validahasta: 202702, Codseguridad: 2312, Limitecompra: 100000.00, Estado: "vigente"}
			tar3, err := json.Marshal(tarjeta3)
			if err != nil {
				panic(err)
			}
			CreateUpdate(db, "tarjeta", []byte(strconv.Itoa(tarjeta3.Nrotarjeta)), tar3)
			resultar3, err := ReadUnique(db, "tarjeta", []byte(strconv.Itoa(tarjeta3.Nrotarjeta)))
			fmt.Printf("%s\n", resultar3)

			/*TRES COMERCIOS*/
			fmt.Printf("\nPresione 3 para cargar los datos de comercios.\n\n")
			var commerce int
			fmt.Scan(&commerce)
			fmt.Printf("\nCargando datos de tres comercios...\n\n")

			//Comercio 1
			comercio1 := comercios{Nrocomercio: 01, Nombre: "Libreria El patito feo", Domicilio: "Av. San Luis 1687", Codigopostal: "B1663HGK", Telefono: "011-93155601"}
			com1, err := json.Marshal(comercio1)
			if err != nil {
				panic(err)
			}
			CreateUpdate(db, "comercio", []byte(strconv.Itoa(comercio1.Nrocomercio)), com1)
			resulcom1, err := ReadUnique(db, "comercio", []byte(strconv.Itoa(comercio1.Nrocomercio)))
			fmt.Printf("%s\n", resulcom1)

			//Comercio 2
			comercio2 := comercios{Nrocomercio: 02, Nombre: "Heladeria Gustavo", Domicilio: "Serrano 1523", Codigopostal: "B1722NHC", Telefono: "011-97684470"}
			com2, err := json.Marshal(comercio2)
			if err != nil {
				panic(err)
			}
			CreateUpdate(db, "comercio", []byte(strconv.Itoa(comercio2.Nrocomercio)), com2)
			resulcom2, err := ReadUnique(db, "comercio", []byte(strconv.Itoa(comercio2.Nrocomercio)))
			fmt.Printf("%s\n", resulcom2)

			//Comercio 3
			comercio3 := comercios{Nrocomercio: 03, Nombre: "Carniceria El cordero feliz", Domicilio: "Ituzaingo 4896", Codigopostal: "B1669FUE", Telefono: "011-40346435"}
			com3, err := json.Marshal(comercio3)
			if err != nil {
				panic(err)
			}
			CreateUpdate(db, "comercio", []byte(strconv.Itoa(comercio3.Nrocomercio)), com3)
			resulcom3, err := ReadUnique(db, "comercio", []byte(strconv.Itoa(comercio3.Nrocomercio)))
			fmt.Printf("%s\n", resulcom3)

			/*TRES COMPRAS*/
			fmt.Printf("\nPresione 4 para cargar los datos de compras.\n\n")
			var purchase int
			fmt.Scan(&purchase)
			fmt.Printf("\nCargando datos de tres compras...\n\n")

			//Compra 1
			compra1 := compras{Nrooperacion: 26281872, Nrotarjeta: 4756326984155476, Nrocomercio: 01, Fecha: "05/06/2022", Monto: 5000.00, Pagado: false}
			comp1, err := json.Marshal(compra1)
			if err != nil {
				panic(err)
			}
			CreateUpdate(db, "compra", []byte(strconv.Itoa(compra1.Nrooperacion)), comp1)
			resulcomp1, err := ReadUnique(db, "compra", []byte(strconv.Itoa(compra1.Nrooperacion)))
			fmt.Printf("%s\n", resulcomp1)

			//Compra 2
			compra2 := compras{Nrooperacion: 26281872, Nrotarjeta: 4532969538877007, Nrocomercio: 02, Fecha: "06/06/2022", Monto: 7000.00, Pagado: false}
			comp2, err := json.Marshal(compra2)
			if err != nil {
				panic(err)
			}
			CreateUpdate(db, "compra", []byte(strconv.Itoa(compra2.Nrooperacion)), comp2)
			resulcomp2, err := ReadUnique(db, "compra", []byte(strconv.Itoa(compra2.Nrooperacion)))
			fmt.Printf("%s\n", resulcomp2)

			//Compra 3
			compra3 := compras{Nrooperacion: 26283535, Nrotarjeta: 4929941716451245, Nrocomercio: 03, Fecha: "07/06/2022", Monto: 4000.00, Pagado: false}
			comp3, err := json.Marshal(compra3)
			if err != nil {
				panic(err)
			}
			CreateUpdate(db, "compra", []byte(strconv.Itoa(compra3.Nrooperacion)), comp3)
			resulcomp3, err := ReadUnique(db, "compra", []byte(strconv.Itoa(compra3.Nrooperacion)))
			fmt.Printf("%s\n", resulcomp3)

			//Continuar?
			fmt.Printf("\n¿Desea continuar operando? Presione y/n\n")
			var consulta string
			fmt.Scan(&consulta)
			if consulta == "y" {
				continue
			}
			break
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
	fmt.Printf("Escriba 0 para salir.\n\n")
}

/*TRANSACCIÓN DE ESCRITURA*/
func CreateUpdate(db *bolt.DB, bucketName string, key []byte, val []byte) error {

	//Abre la transacción de escritura
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))
	err = b.Put(key, val)
	if err != nil {
		return err
	}

	//Cierra la transacción
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

/*TRANSACCIÓN DE LECTURA*/
func ReadUnique(db *bolt.DB, bucketName string, key []byte) ([]byte, error) {

	var buf []byte
	//Abre una transacción de lectura
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		buf = b.Get(key)
		return nil
	})
	return buf, err
}
