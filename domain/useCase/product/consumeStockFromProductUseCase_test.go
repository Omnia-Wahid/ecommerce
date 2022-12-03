package product

import (
	"ecommerce/domain/model"
	"ecommerce/test/mocks/adapter/repositoryMock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConsumeStockFromProductSuccess(t *testing.T) {
	GetStockChangeMock := func(sku string, country string) (*model.ProductStock, error) {
		change := new(int)
		*change = 30
		productStock := &model.ProductStock{
			Country:     "dz",
			Sku:         "e920c573f128",
			Name:        "Ramirez-Molina Granite Pizza",
			StockChange: change,
		}
		return productStock, nil
	}
	UpdateProductsStocksMock := func(products []*model.ProductStock) error {
		return nil
	}
	productsStockMock := repositoryMock.ProductsStockCrudsMock{
		UpdateProductsStocksMock: UpdateProductsStocksMock,
		GetStockChangeMock:       GetStockChangeMock,
	}
	repositoryMock := &repositoryMock.RepositoryMock{
		ProductsCrudsMock:      repositoryMock.ProductsCrudsMock{},
		ProductsStockCrudsMock: productsStockMock,
	}
	productUseCaseMock := productUseCase{
		repositoryMock,
	}
	consumedAmount := new(int)
	*consumedAmount = 10
	err := productUseCaseMock.ConsumeStockFromProduct("e920c573f128", consumedAmount, "dz")
	assert.Nil(t, err)
}

func TestConsumeStockFromProductErrorProductNotAvailable(t *testing.T) {
	GetStockChangeMock := func(sku string, country string) (*model.ProductStock, error) {
		return nil, nil
	}
	productsStockMock := repositoryMock.ProductsStockCrudsMock{
		UpdateProductsStocksMock: nil,
		GetStockChangeMock:       GetStockChangeMock,
	}
	repositoryMock := &repositoryMock.RepositoryMock{
		ProductsCrudsMock:      repositoryMock.ProductsCrudsMock{},
		ProductsStockCrudsMock: productsStockMock,
	}
	productUseCaseMock := productUseCase{
		repositoryMock,
	}
	consumedAmount := new(int)
	*consumedAmount = 10
	err := productUseCaseMock.ConsumeStockFromProduct("e920c573f128", consumedAmount, "dz")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "product stock not available")
}

func TestConsumeStockFromProductErrorProductStockNotEnough(t *testing.T) {
	GetStockChangeMock := func(sku string, country string) (*model.ProductStock, error) {
		change := new(int)
		*change = 5
		productStock := &model.ProductStock{
			Country:     "dz",
			Sku:         "e920c573f128",
			Name:        "Ramirez-Molina Granite Pizza",
			StockChange: change,
		}
		return productStock, nil
	}
	productsStockMock := repositoryMock.ProductsStockCrudsMock{
		UpdateProductsStocksMock: nil,
		GetStockChangeMock:       GetStockChangeMock,
	}
	repositoryMock := &repositoryMock.RepositoryMock{
		ProductsCrudsMock:      repositoryMock.ProductsCrudsMock{},
		ProductsStockCrudsMock: productsStockMock,
	}
	productUseCaseMock := productUseCase{
		repositoryMock,
	}
	consumedAmount := new(int)
	*consumedAmount = 10
	err := productUseCaseMock.ConsumeStockFromProduct("e920c573f128", consumedAmount, "dz")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "product stock not enough for consumed amount")
}
