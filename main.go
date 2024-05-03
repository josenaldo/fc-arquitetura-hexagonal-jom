package main

import (
	"database/sql"
	"log"

	dbAdapters "github.com/josenaldo/fc-arquitetura-hexagobal-jom/adapters/db"
	"github.com/josenaldo/fc-arquitetura-hexagobal-jom/application"
)

func main() {
	log.Println("Starting application")

	db, _ := sql.Open("sqlite3", "sqlite.db")
	log.Println("Database connected")

	productDbAdapter := dbAdapters.NewProductDb(db)
	log.Println("Product database adapter created")

	productService := application.NewProductService(productDbAdapter)
	log.Println("Product service created")

	created, err := productService.Create("Product 1", 10.0)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Product created: %v\n", created)

	product, err := productService.Get(created.GetID())
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Product returned: %v\n", product)

	productService.Enable(product)
	log.Printf("Product enabled: %v\n", product)

	productService.Disable(product)
	log.Printf("Product disabled: %v\n", product)

}
