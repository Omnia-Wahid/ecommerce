package product

import (
	"ecommerce/domain/model"
	"errors"
)

func (thisPU *productUseCase) ConsumeStockFromProduct(sku string, consumedAmount int, country string) error {
	productStock, err := thisPU.repository.GetStockChange(sku, country)
	if err != nil {
		return err
	}
	if productStock == nil {
		return errors.New("product stock not available")
	}
	if productStock.StockChange >= consumedAmount {
		productStock.StockChange = productStock.StockChange - consumedAmount
		var updatedStock []*model.ProductStock
		updatedStock = append(updatedStock, productStock)
		err = thisPU.repository.UpdateProductsStocks(updatedStock)
		if err != nil {
			return err
		}
	} else {
		return errors.New("product stock not enough for consumed amount")
	}
	return nil
}
