package application_test

import (
	"testing"

	"github.com/josenaldo/fc-arquitetura-hexagobal-jom/application"
	"github.com/stretchr/testify/require"
)

func TestShouldEnableProductWithValidPrice(t *testing.T) {
	p := &application.Product{}
	p.Name = "Product 1"
	p.Price = 10
	p.Status = application.DISABLED

	err := p.Enable()

	require.Nil(t, err)
	require.Equal(t, application.ENABLED, p.GetStatus())
}

func TestShouldNotEnableAProductWithPriceEqualsToZero(t *testing.T) {
	p := &application.Product{Price: 0}
	err := p.Enable()

	require.NotNil(t, err)
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestShouldDisableProductOnlyIfPriceIsequalZero(t *testing.T) {
	p := &application.Product{
		Name:   "Product 1",
		Price:  0,
		Status: application.ENABLED,
	}

	err := p.Disable()

	require.Nil(t, err)
	require.Equal(t, application.DISABLED, p.GetStatus())
}

func TestShouldNotDisableProductIfPriceIsGreaterThanZero(t *testing.T) {
	p := &application.Product{
		Name:   "Product 1",
		Price:  10,
		Status: application.ENABLED,
	}

	err := p.Disable()

	require.NotNil(t, err)
	require.Equal(t, "the price must be equal to zero to disable the product", err.Error())
}
