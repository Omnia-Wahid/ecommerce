package repositoryMock

import "ecommerce/domain/model"

type ProductsStockCrudsMock struct {
	UpdateProductsStocksMock func(products []*model.ProductStock) error
	GetStockChangeMock       func(sku string, country string) (*model.ProductStock, error)
}

func (thisPSM *ProductsStockCrudsMock) UpdateProductsStocks(products []*model.ProductStock) error {
	return thisPSM.UpdateProductsStocksMock(products)
}
func (thisPSM *ProductsStockCrudsMock) GetStockChange(sku string, country string) (*model.ProductStock, error) {
	return thisPSM.GetStockChangeMock(sku, country)
}
