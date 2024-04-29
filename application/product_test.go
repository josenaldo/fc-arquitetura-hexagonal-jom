package application_test

import (
	"testing"

	"github.com/josenaldo/fc-arquitetura-hexagobal-jom/application"
	"github.com/stretchr/testify/require"
)

func TestItShouldCreateAValidProduct(t *testing.T) {
	p, err := application.NewProduct("Product 1", 100)
	require.Nil(t, err)

	valid, errValid := p.IsValid()

	require.True(t, valid)
	require.Nil(t, errValid)
	require.Equal(t, "Product 1", p.Name)
	require.Equal(t, 100.0, p.Price)
	require.Equal(t, application.DISABLED, p.Status)

}

func TestItShouldEnableProductWithValidPrice(t *testing.T) {
	p, err := application.NewProduct("Product 1", 10)
	require.Nil(t, err)

	err = p.Enable()

	require.Nil(t, err)
	require.Equal(t, application.ENABLED, p.GetStatus())
}

func TestItShouldNotEnableAProductWithPriceEqualsToZero(t *testing.T) {
	p, err := application.NewProduct("Product 1", 0)
	require.Nil(t, err)

	err = p.Enable()

	require.NotNil(t, err)
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestItShouldDisableProductOnlyIfPriceIsequalZero(t *testing.T) {
	p, err := application.NewProduct("Product 1", 10)
	require.Nil(t, err)

	p.Enable()
	p.Price = 0

	err = p.Disable()

	require.Nil(t, err)
	require.Equal(t, application.DISABLED, p.GetStatus())
}

func TestItShouldNotDisableProductIfPriceIsGreaterThanZero(t *testing.T) {
	p, err := application.NewProduct("Product 1", 10)
	require.Nil(t, err)

	p.Enable()

	err = p.Disable()

	require.NotNil(t, err)
	require.Equal(t, "the price must be equal to zero to disable the product", err.Error())
}

func TestItShouldReturnValidWhenProductIsValid(t *testing.T) {
	p, err := application.NewProduct("Product 1", 10)
	require.Nil(t, err)

	p.Enable()

	valid, error := p.IsValid()

	require.True(t, valid)
	require.Nil(t, error)
}

func TestItShouldReturnValidWhenStatusIsEmpty(t *testing.T) {
	p, err := application.NewProduct("Product 1", 10)
	require.Nil(t, err)

	p.Status = ""

	valid, error := p.IsValid()

	require.True(t, valid)
	require.Nil(t, error)
	require.Equal(t, application.DISABLED, p.GetStatus())

}

func TestItShouldReturnInvalidWhenPriceIsLessThanZero(t *testing.T) {
	p, err := application.NewProduct("Product 1", 10)
	require.Nil(t, err)

	p.Price = -1

	valid, error := p.IsValid()

	require.False(t, valid)
	require.Equal(t, "the price must be greater than zero", error.Error())
}

func TestItShouldReturnInvalidWhenNameIsEmpty(t *testing.T) {
	p, err := application.NewProduct("", 10)
	require.Equal(t, "the product is invalid: the name is required", err.Error())
	require.Nil(t, p)
}

func TestItShouldReturnInvalidWhenNameIsNotSet(t *testing.T) {
	var name string
	p, err := application.NewProduct(name, 10)
	require.Equal(t, "the product is invalid: the name is required", err.Error())
	require.Nil(t, p)
}

func TestItShouldReturnInvalidWhenStatusIsInvalid(t *testing.T) {
	p, err := application.NewProduct("Product 1", 10)
	require.Nil(t, err)

	p.Status = "Other Status"

	valid, error := p.IsValid()

	require.False(t, valid)
	require.Equal(t, "the status must be ENABLED or DISABLED", error.Error())
}
