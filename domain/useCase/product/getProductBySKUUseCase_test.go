package product

import (
	"ecommerce/domain/model"
	"ecommerce/test/mocks/adapter/repositoryMock"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProductBySKUSuccess(t *testing.T) {
	change := new(int)
	*change = 30
	getProductBySKUMock := func(sku string) ([]*model.ProductStock, error) {
		productInfo := []*model.ProductStock{{
			Country:     "dz",
			Sku:         "e920c573f128",
			Name:        "Ramirez-Molina Granite Pizza",
			StockChange: change,
		}}
		return productInfo, nil
	}
	productMock := repositoryMock.ProductsCrudsMock{
		UpdateProductsMock:  nil,
		GetProductBySKUMock: getProductBySKUMock,
	}
	repositoryMock := &repositoryMock.RepositoryMock{
		ProductsCrudsMock:      productMock,
		ProductsStockCrudsMock: repositoryMock.ProductsStockCrudsMock{},
	}
	product, err := repositoryMock.GetProductBySKU("e920c573f128")
	assert.Nil(t, err)
	assert.Equal(t, []*model.ProductStock{{
		Country:     "dz",
		Sku:         "e920c573f128",
		Name:        "Ramirez-Molina Granite Pizza",
		StockChange: change,
	}}, product)
}

func TestGetProductBySKUError(t *testing.T) {
	change := new(int)
	*change = 30
	getProductBySKUMock := func(sku string) ([]*model.ProductStock, error) {
		return nil, errors.New("connection refused")
	}
	productMock := repositoryMock.ProductsCrudsMock{
		UpdateProductsMock:  nil,
		GetProductBySKUMock: getProductBySKUMock,
	}
	repositoryMock := &repositoryMock.RepositoryMock{
		ProductsCrudsMock:      productMock,
		ProductsStockCrudsMock: repositoryMock.ProductsStockCrudsMock{},
	}
	product, err := repositoryMock.GetProductBySKU("e920c573f128")
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "connection refused")
}
