package domainPort

import "ecommerce/domain/model"

type IProductDomainPort interface {
	UpdateProductsStocks(productStockList []*model.ProductStock) (err error)
	GetProductBySKU(sku string) ([]*model.ProductStock, error)
	ConsumeStockFromProduct(sku string, consumedAmount *int, country string) error
}
