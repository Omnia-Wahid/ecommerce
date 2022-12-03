package repositoryMock

import "ecommerce/domain/model"

type ProductsCrudsMock struct {
	UpdateProductsMock  func(products []*model.ProductStock) error
	GetProductBySKUMock func(sku string) ([]*model.ProductStock, error)
}

func (thisPM *ProductsCrudsMock) UpdateProducts(products []*model.ProductStock) error {
	return thisPM.UpdateProductsMock(products)
}
func (thisPM *ProductsCrudsMock) GetProductBySKU(sku string) ([]*model.ProductStock, error) {
	return thisPM.GetProductBySKUMock(sku)
}
