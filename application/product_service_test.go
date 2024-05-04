package application_test

import (
	"testing"

	"github.com/josenaldo/fc-arquitetura-hexagonal-jom/application"
	mock_application "github.com/josenaldo/fc-arquitetura-hexagonal-jom/application/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestItShouldGetSomeProduct(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Get(gomock.Any()).Return(product, nil)

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("1")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestItShouldReturnErrorWhenGetInexistentProduct(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Get(gomock.Any()).Return(nil, application.ErrProductNotFound)

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("1")
	require.Nil(t, result)
	require.Equal(t, application.ErrProductNotFound, err)
}

func TestItShouldCreateAProduct(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("Product 1", 10)
	require.Nil(t, err)
	require.NotNil(t, result)
	require.Equal(t, product, result)
}

func TestItShouldReturnErrorWhenCreateProductWithInvalidPrice(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(nil, application.ErrInvalidProductPrice).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("Product 1", -10)
	require.NotNil(t, err)
	require.Nil(t, result)
	require.Equal(t, application.ErrInvalidProductPrice, err)
}

func TestItShouldReturnErrorWhenCreateProductWithEmptyName(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(nil, application.ErrRequiredProductName).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("", 10)
	require.NotNil(t, err)
	require.Nil(t, result)
	require.Equal(t, application.ErrRequiredProductName, err)
}

func TestItShouldEnableAProduct(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product, err := application.NewProduct("Product 1", 10)
	require.Nil(t, err)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil)

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.NotNil(t, result)
	require.Equal(t, application.ENABLED, result.GetStatus())
}

func TestItShouldDisableAProduct(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product, err := application.NewProduct("Product 1", 0)
	require.Nil(t, err)
	product.Enable()

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil)

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Disable(product)
	require.Nil(t, err)
	require.NotNil(t, result)
	require.Equal(t, application.DISABLED, result.GetStatus())
}
