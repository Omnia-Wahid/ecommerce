package product

import (
	"ecommerce/domain/model"
	"ecommerce/test/mocks/adapter/repositoryMock"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateProductsStocksSuccess(t *testing.T) {
	UpdateProductsMock := func(products []*model.ProductStock) error {
		return nil
	}
	UpdateProductsStocksMock := func(products []*model.ProductStock) error {
		return nil
	}
	productMock := repositoryMock.ProductsCrudsMock{
		UpdateProductsMock:  UpdateProductsMock,
		GetProductBySKUMock: nil,
	}
	productsStockMock := repositoryMock.ProductsStockCrudsMock{
		UpdateProductsStocksMock: UpdateProductsStocksMock,
		GetStockChangeMock:       nil,
	}
	repositoryMock := &repositoryMock.RepositoryMock{
		ProductsCrudsMock:      productMock,
		ProductsStockCrudsMock: productsStockMock,
	}
	productUseCaseMock := productUseCase{
		repositoryMock,
	}
	change := new(int)
	*change = 30
	productStockList := []*model.ProductStock{{
		Country:     "dz",
		Sku:         "e920c573f128",
		Name:        "Ramirez-Molina Granite Pizza",
		StockChange: change,
	}}
	err := productUseCaseMock.UpdateProductsStocks(productStockList)
	assert.Nil(t, err)
}

func TestUpdateProductsStocksError(t *testing.T) {
	UpdateProductsMock := func(products []*model.ProductStock) error {
		return nil
	}
	UpdateProductsStocksMock := func(products []*model.ProductStock) error {
		return errors.New("duplicate value")
	}
	productMock := repositoryMock.ProductsCrudsMock{
		UpdateProductsMock:  UpdateProductsMock,
		GetProductBySKUMock: nil,
	}
	productsStockMock := repositoryMock.ProductsStockCrudsMock{
		UpdateProductsStocksMock: UpdateProductsStocksMock,
		GetStockChangeMock:       nil,
	}
	repositoryMock := &repositoryMock.RepositoryMock{
		ProductsCrudsMock:      productMock,
		ProductsStockCrudsMock: productsStockMock,
	}
	productUseCaseMock := productUseCase{
		repositoryMock,
	}

	err := productUseCaseMock.UpdateProductsStocks([]*model.ProductStock{{
		Country:     "",
		Sku:         "",
		Name:        "",
		StockChange: nil,
	}})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "duplicate value")
}
