package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/josenaldo/fc-arquitetura-hexagobal-jom/adapters/db"
	"github.com/josenaldo/fc-arquitetura-hexagobal-jom/application"
	"github.com/stretchr/testify/require"
)

var (
	Db  *sql.DB
	id1 string = uuid.New().String()
)

func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	product1 := application.Product{
		ID:     id1,
		Name:   "Product 1",
		Price:  10.0,
		Status: "active",
	}
	createProduct(Db, product1)

}

func teardown() {
	Db.Close()
}

func createTable(db *sql.DB) {
	table := "CREATE TABLE products (id string, name string, price float, status string)"
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB, product application.Product) {
	insert := "INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec(product.ID, product.Name, product.Price, product.Status)
}

func TestItShouldGetProductFromDb(t *testing.T) {
	setup()
	defer teardown()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get(id1)

	require.Nil(t, err)
	require.Equal(t, id1, product.GetID())
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, "active", product.GetStatus())
}

func TestItShouldNotGetProductFromDb(t *testing.T) {
	setup()
	defer teardown()

	productDb := db.NewProductDb(Db)
	_, err := productDb.Get("invalid-id")

	require.NotNil(t, err)
	require.Equal(t, application.ErrProductNotFound, err)
}
