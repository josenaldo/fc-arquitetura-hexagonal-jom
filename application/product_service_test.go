package application_test

import (
	"testing"

	"github.com/josenaldo/fc-arquitetura-hexagobal-jom/application"
	"github.com/stretchr/testify/require"
)

func TestItShouldGetSomeProduct(t *testing.T) {
	productPersistence := application.NewProductPersistence()
	productService := application.ProductService{Persistence: productPersistence}

	product, err := productService.Get("some-id")

	require.Nil(t, err)
	require.NotNil(t, product)
}
