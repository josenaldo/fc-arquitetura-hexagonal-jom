package integration_test

import (
	"database/sql"
	"log"
	"testing"

	dbAdapters "github.com/josenaldo/fc-arquitetura-hexagonal-jom/adapters/db"
	"github.com/josenaldo/fc-arquitetura-hexagonal-jom/application"
	"github.com/stretchr/testify/require"
)

var (
	Db             *sql.DB
	id1            string
	productService application.ProductServiceInterface
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

	productDbAdapter := dbAdapters.NewProductDb(Db)
	productService = application.NewProductService(productDbAdapter)
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

func TestItShouldCreateProduct(t *testing.T) {
	setup()
	defer teardown()

	created, err := productService.Create("Product 1", 10.0)
	require.Nil(t, err)
	require.NotNil(t, created)
	require.Equal(t, "Product 1", created.GetName())
	require.Equal(t, 10.0, created.GetPrice())
	require.Equal(t, application.DISABLED, created.GetStatus())

}

func TestItShouldGetProduct(t *testing.T) {
	setup()
	defer teardown()

	product, err := productService.Get(id1)
	require.Nil(t, err)
	require.NotNil(t, product)
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, application.DISABLED, product.GetStatus())
}

func TestItShouldEnableProduct(t *testing.T) {
	setup()
	defer teardown()

	product, err := productService.Get(id1)
	require.Nil(t, err)
	require.NotNil(t, product)

	saved, err := productService.Enable(product)
	require.Nil(t, err)
	require.NotNil(t, saved)
	require.Equal(t, "Product 1", saved.GetName())
	require.Equal(t, 10.0, saved.GetPrice())
	require.Equal(t, application.ENABLED, saved.GetStatus())
}

func TestItShouldDisableProduct(t *testing.T) {

	setup()
	defer teardown()

	product, err := productService.Get(id1)
	require.Nil(t, err)
	require.NotNil(t, product)

	saved, err := productService.Enable(product)
	require.Nil(t, err)
	require.NotNil(t, saved)
	require.Equal(t, "Product 1", saved.GetName())
	require.Equal(t, 10.0, saved.GetPrice())
	require.Equal(t, application.ENABLED, saved.GetStatus())

	productImpl := product.(*application.Product)
	productImpl.Price = 0

	saved, err = productService.Disable(product)
	require.Nil(t, err)
	require.NotNil(t, saved)
	require.Equal(t, "Product 1", saved.GetName())
	require.Equal(t, 0.0, saved.GetPrice())
	require.Equal(t, application.DISABLED, saved.GetStatus())

}

func TestItShouldNotDisableProductWithPriceGreaterThanZero(t *testing.T) {
	setup()
	defer teardown()

	product, err := productService.Get(id1)
	require.Nil(t, err)
	require.NotNil(t, product)

	saved, err := productService.Enable(product)
	require.Nil(t, err)
	require.NotNil(t, saved)
	require.Equal(t, "Product 1", saved.GetName())
	require.Equal(t, 10.0, saved.GetPrice())
	require.Equal(t, application.ENABLED, saved.GetStatus())

	_, err = productService.Disable(product)
	require.NotNil(t, err)
	require.Equal(t, "the price must be equal to zero to disable the product", err.Error())

}

func TestItShouldNotCreateProductWithInvalidPrice(t *testing.T) {
	setup()
	defer teardown()

	_, err := productService.Create("Product 1", -10.0)
	require.NotNil(t, err)
	require.Equal(t, "the price must be greater than zero", err.Error())
}

func TestItShouldNotCreateProductWithEmptyName(t *testing.T) {
	setup()
	defer teardown()

	_, err := productService.Create("", 10.0)
	require.NotNil(t, err)
	require.Equal(t, "the name is required", err.Error())
}
