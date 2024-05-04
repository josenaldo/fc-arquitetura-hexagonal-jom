package cli_test

import (
	"fmt"
	testing "testing"

	"github.com/google/uuid"
	"github.com/josenaldo/fc-arquitetura-hexagonal-jom/adapters/cli"
	"github.com/josenaldo/fc-arquitetura-hexagonal-jom/application"
	mock_application "github.com/josenaldo/fc-arquitetura-hexagonal-jom/application/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

var (
	productId     string  = uuid.New().String()
	productName   string  = "Product Test"
	productPrice  float64 = 10.0
	productStatus string  = application.ENABLED
	product       *mock_application.MockProductInterface
	service       *mock_application.MockProductServiceInterface
	ctrl          *gomock.Controller
	products      []application.ProductInterface
)

func setup(t *testing.T) {
	ctrl = gomock.NewController(t)
	product = mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().GetID().Return(productId).AnyTimes()
	product.EXPECT().GetName().Return(productName).AnyTimes()
	product.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	product.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	product2 := mock_application.NewMockProductInterface(ctrl)
	product2.EXPECT().GetID().Return(uuid.New().String()).AnyTimes()
	product2.EXPECT().GetName().Return("Product 2").AnyTimes()
	product2.EXPECT().GetPrice().Return(20.0).AnyTimes()
	product2.EXPECT().GetStatus().Return(application.ENABLED).AnyTimes()

	products = []application.ProductInterface{product, product2}

	service = mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(product, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(product, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(product, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(product, nil).AnyTimes()
	service.EXPECT().GetAll().Return(products, nil).AnyTimes()
}

func teardown() {
	ctrl.Finish()
}

func TestItShouldCreateProduct(t *testing.T) {
	setup(t)
	defer teardown()

	resultExpected := fmt.Sprintf(
		"Product ID %s with the name %s has been created with price %f and status %s",
		productName,
		productId,
		productPrice,
		productStatus)

	result, err := cli.Run(service, "create", "", productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
	require.Equal(t, productId, product.GetID())
	require.Equal(t, productName, product.GetName())
	require.Equal(t, productPrice, product.GetPrice())
	require.Equal(t, productStatus, product.GetStatus())
}

func TestItShouldUpdateProductStatusToEnabled(t *testing.T) {
	setup(t)
	defer teardown()

	resultExpected := fmt.Sprintf("Product %s has been enabled", productName)

	result, err := cli.Run(service, "enable", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
	require.Equal(t, productId, product.GetID())
	require.Equal(t, productName, product.GetName())
	require.Equal(t, productPrice, product.GetPrice())
	require.Equal(t, application.ENABLED, product.GetStatus())
}

func TestItShouldDisableProduct(t *testing.T) {
	setup(t)
	defer teardown()

	product = mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().GetID().Return(productId).AnyTimes()
	product.EXPECT().GetName().Return(productName).AnyTimes()
	product.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	product.EXPECT().GetStatus().Return(application.DISABLED).AnyTimes()

	service = mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Get(productId).Return(product, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(product, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product %s has been disabled", productName)

	result, err := cli.Run(service, "disable", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
	require.Equal(t, productId, product.GetID())
	require.Equal(t, productName, product.GetName())
	require.Equal(t, productPrice, product.GetPrice())
	require.Equal(t, application.DISABLED, product.GetStatus())
}

func TestItShouldGetProduct(t *testing.T) {
	setup(t)
	defer teardown()

	resultExpected := fmt.Sprintf(
		"Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		productId,
		productName,
		productPrice,
		productStatus)

	result, err := cli.Run(service, "", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
	require.Equal(t, productId, product.GetID())
	require.Equal(t, productName, product.GetName())
	require.Equal(t, productPrice, product.GetPrice())
	require.Equal(t, productStatus, product.GetStatus())
}

func TestItShouldGetAllProducts(t *testing.T) {
	setup(t)
	defer teardown()

	resultExpected := ""
	for _, p := range products {
		resultExpected += fmt.Sprintf(
			"Product ID: %s\nName: %s\nPrice: %f\nStatus: %s\n\n",
			p.GetID(),
			p.GetName(),
			p.GetPrice(),
			p.GetStatus())
	}

	result, err := cli.Run(service, "list", "", "", 0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
