package product

import "ecommerce/domain/model"

func (thisPU *productUseCase) GetProductBySKU(sku string) ([]*model.ProductStock, error) {
	productStocks, err := thisPU.repository.GetProductBySKU(sku)
	if err != nil {
		return nil, err
	}
	return productStocks, nil
}
