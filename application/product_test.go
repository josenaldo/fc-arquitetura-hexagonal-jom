package application_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/josenaldo/fc-arquitetura-hexagobal-jom/application"
	"github.com/stretchr/testify/require"
)

func TestItShouldEnableProductWithValidPrice(t *testing.T) {
	p := &application.Product{
		ID:     uuid.NewString(),
		Name:   "Product 1",
		Price:  10,
		Status: application.DISABLED,
	}

	err := p.Enable()

	require.Nil(t, err)
	require.Equal(t, application.ENABLED, p.GetStatus())
}

func TestItShouldNotEnableAProductWithPriceEqualsToZero(t *testing.T) {
	p := &application.Product{
		ID:     uuid.NewString(),
		Name:   "Product 1",
		Price:  0,
		Status: application.DISABLED,
	}
	err := p.Enable()

	require.NotNil(t, err)
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestItShouldDisableProductOnlyIfPriceIsequalZero(t *testing.T) {
	p := &application.Product{
		Name:   "Product 1",
		Price:  0,
		Status: application.ENABLED,
	}

	err := p.Disable()

	require.Nil(t, err)
	require.Equal(t, application.DISABLED, p.GetStatus())
}

func TestItShouldNotDisableProductIfPriceIsGreaterThanZero(t *testing.T) {
	p := &application.Product{
		Name:   "Product 1",
		Price:  10,
		Status: application.ENABLED,
	}

	err := p.Disable()

	require.NotNil(t, err)
	require.Equal(t, "the price must be equal to zero to disable the product", err.Error())
}

func TestItShouldReturnValidWhenProductIsValid(t *testing.T) {
	p := &application.Product{
		ID:     uuid.NewString(),
		Name:   "Product 1",
		Price:  10,
		Status: application.ENABLED,
	}

	valid, error := p.IsValid()

	require.True(t, valid)
	require.Nil(t, error)
}

func TestItShouldReturnValidWhenStatusIsEmpty(t *testing.T) {
	p := &application.Product{
		ID:     uuid.NewString(),
		Name:   "Product 1",
		Price:  10,
		Status: "",
	}

	valid, error := p.IsValid()

	require.True(t, valid)
	require.Nil(t, error)
	require.Equal(t, application.DISABLED, p.GetStatus())

}

func TestItShouldReturnInvalidWhenPriceIsLessThanZero(t *testing.T) {
	p := &application.Product{
		Name:   "Product 1",
		Price:  -10,
		Status: application.ENABLED,
	}

	valid, error := p.IsValid()

	require.False(t, valid)
	require.Equal(t, "the price must be greater than zero", error.Error())
}

func TestItShouldReturnInvalidWhenNameIsEmpty(t *testing.T) {
	p := &application.Product{
		Name:   "",
		Price:  10,
		Status: application.ENABLED,
	}

	valid, error := p.IsValid()

	require.False(t, valid)
	require.Equal(t, "the name is required", error.Error())
}

func TestItShouldReturnInvalidWhenNameIsNotSet(t *testing.T) {
	p := &application.Product{
		Price:  10,
		Status: application.ENABLED,
	}

	valid, error := p.IsValid()

	require.False(t, valid)
	require.Equal(t, "the name is required", error.Error())
}

func TestItShouldReturnInvalidWhenStatusIsInvalid(t *testing.T) {
	p := &application.Product{
		Name:   "Product 1",
		Price:  10,
		Status: "Other status",
	}

	valid, error := p.IsValid()

	require.False(t, valid)
	require.Equal(t, "the status must be ENABLED or DISABLED", error.Error())
}
