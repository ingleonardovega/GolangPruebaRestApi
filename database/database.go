package database

import (
	"database/sql"

	"github.com/GolangPruebaRestApi/helper"
)

func InitDB() *sql.DB {
	connectionString := "root:1q2w3e4r@tcp(localhost:3306)/northwind"
	databaseConnection, err := sql.Open("mysql", connectionString)

	helper.Catch(err)
	return databaseConnection

}
