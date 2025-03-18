package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//usernamer:password@
const url = "root:123456@tcp(localhost:3306)/goweb_db"

// Guarda la conexion
var db *sql.DB

// Abre la conexio
func Connect() {
	connection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion exitosa")
	db = connection
}

// Cierra la conexion
func Close() {
	db.Close()
}

// Verifica la conexion a la bd
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

// Verifica si existe una tabala
func ExistTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(sql)
	if err != nil {
		fmt.Println("eeror: ", err)
	}

	// recorre la tabla con el nombre indicado y devuelve un boolean
	return rows.Next()
}

// Crea una tabla en la base
func CrateTable(schema string, name string) {
	if !ExistTable(name) {
		// si la tabla no existe ejecuta:
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

// Polimorfismos de Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return result, err
}

// Polimorifismo de Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return rows, err
}

// Reinicar los registro de una tabal
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}
