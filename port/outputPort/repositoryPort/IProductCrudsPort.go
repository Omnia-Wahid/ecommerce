package repositoryPort

import "ecommerce/domain/model"

type IProductCrudsPort interface {
	UpdateProducts(products []*model.ProductStock) error
	GetProductBySKU(sku string) ([]*model.ProductStock, error)
}
