package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/josenaldo/fc-arquitetura-hexagobal-jom/adapters/db"
	"github.com/josenaldo/fc-arquitetura-hexagobal-jom/application"
	"github.com/stretchr/testify/require"
)

var (
	Db  *sql.DB
	id1 string
)

func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:")

	product1, err := application.NewProduct("Product 1", 10.0)
	if err != nil {
		log.Fatal(err.Error())
	}
	id1 = product1.GetID()

	createTable(Db)
	createProduct(Db, *product1)
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
	require.Equal(t, application.DISABLED, product.GetStatus())
}

func TestItShouldNotGetProductFromDb(t *testing.T) {
	setup()
	defer teardown()

	productDb := db.NewProductDb(Db)
	_, err := productDb.Get("invalid-id")

	require.NotNil(t, err)
	require.Equal(t, application.ErrProductNotFound, err)
}

func TestItShouldCreateANewProduct(t *testing.T) {
	setup()
	defer teardown()

	productDb := db.NewProductDb(Db)
	product, err := application.NewProduct("Product 2", 20.0)
	require.Nil(t, err)

	productSaved, errSave := productDb.Save(product)
	require.Nil(t, errSave)
	require.NotNil(t, productSaved)

	productFromDb, err := productDb.Get(product.GetID())

	require.Nil(t, err)
	require.Equal(t, product.GetID(), productFromDb.GetID())
	require.Equal(t, product.GetName(), productFromDb.GetName())
	require.Equal(t, product.GetPrice(), productFromDb.GetPrice())
	require.Equal(t, product.GetStatus(), productFromDb.GetStatus())
}

func TestItShouldUpdateAProduct(t *testing.T) {
	setup()
	defer teardown()

	productDb := db.NewProductDb(Db)

	productFromDb, err := productDb.Get(id1)
	require.Nil(t, err)
	require.NotNil(t, productFromDb)
	require.Equal(t, id1, productFromDb.GetID())
	require.Equal(t, "Product 1", productFromDb.GetName())
	require.Equal(t, 10.0, productFromDb.GetPrice())
	require.Equal(t, application.DISABLED, productFromDb.GetStatus())

	product := productFromDb.(*application.Product)
	product.Name = "Product 2 Updated"
	product.Price = 30.0
	product.Enable()

	updated, errUpdate := productDb.Save(product)
	require.Nil(t, errUpdate)
	require.NotNil(t, updated)

	productFromDb, err = productDb.Get(product.GetID())

	require.Nil(t, err)
	require.Equal(t, id1, productFromDb.GetID())
	require.Equal(t, "Product 2 Updated", productFromDb.GetName())
	require.Equal(t, 30.0, productFromDb.GetPrice())
	require.Equal(t, application.ENABLED, productFromDb.GetStatus())
}
