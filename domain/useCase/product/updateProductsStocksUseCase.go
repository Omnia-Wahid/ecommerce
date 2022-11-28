package product

import "ecommerce/domain/model"

func (thisPU *productUseCase) UpdateProductsStocks(productStockList []*model.ProductStock) (err error) {
	err = thisPU.repository.UpdateProducts(productStockList)
	if err != nil {
		return err
	}
	err = thisPU.repository.UpdateProductsStocks(productStockList)
	if err != nil {
		return err
	}
	return nil
}
