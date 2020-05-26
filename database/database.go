package database

import "database/sql"

func InitDB() *sql.DB {
	connectionString := "root:1q2w3e4r@tcp(localhost:3306)/northwind"
	databaseConnection, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err.Error()) //manejo de errores
	}
	return databaseConnection

}
