package repositoryPort

import "ecommerce/domain/model"

type IProductStockCrudsPort interface {
	UpdateProductsStocks(products []*model.ProductStock) error
	GetStockChange(sku string, country string) (*model.ProductStock, error)
}
