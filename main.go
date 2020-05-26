package main

import (
	"GolangPruebaRestApi/database"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	databaseConnection := database.InitDB()
	defer databaseConnection.Close()
	fmt.Println(databaseConnection)
}
