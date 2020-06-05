package main

import (
	"net/http"

	"github.com/GolangPruebaRestApi/customer"
	"github.com/GolangPruebaRestApi/database"
	"github.com/GolangPruebaRestApi/employee"
	"github.com/GolangPruebaRestApi/product"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	databaseConnection := database.InitDB()
	defer databaseConnection.Close()

	var (
		productRepository  = product.NewRepository(databaseConnection)
		employeeRepository = employee.NewRepository(databaseConnection)
		customerRepository = customer.NewRepository(databaseConnection)
	)

	var (
		productService  product.Service
		employeeService employee.Service
		customerService customer.Service
	)

	productService = product.NewService(productRepository)
	employeeService = employee.NerService(employeeRepository)
	customerService = customer.NewService(customerRepository)

	r := chi.NewRouter()
	r.Mount("/products", product.MakeHttpHandler(productService))
	r.Mount("/employees", employee.MakeHttpHandler(employeeService))
	r.Mount("/customers", customer.MakeHTTPHandler(customerService))
	http.ListenAndServe(":3000", r)

}
