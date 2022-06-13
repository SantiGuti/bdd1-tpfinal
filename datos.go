package main

import(
	"encoding/json"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"log"
	"strconv"
)

func main() {
    db, err := bolt.Open("datos.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    clientes := cliente{1, "Juan", "Rosas", "Serano 701", "011-68943567"}
    data, err := json.Marshal(clientes)
    if err != nil {
        log.Fatal(err)
    }

    CreateUpdate(db, "cliente", []byte(strconv.Itoa(clientes.nrocliente)), data)

    resultado, err := ReadUnique(db, "cliente", []byte(strconv.Itoa(clientes.nrocliente)))

    fmt.Printf("%s\n", resultado)
}

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
