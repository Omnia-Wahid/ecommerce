package entity

import "ecommerce/domain/model"

type ProductStock struct {
	tableName   struct{} `pg:"product_stocks,discard_unknown_columns"`
	Sku         string   `pg:"sku,type:char(50),unique:stock_unique"`
	Country     string   `pg:"country,notnull,type:char(50),unique:stock_unique"`
	StockChange int      `pg:"stock_change,notnull,use_zero,type:integer"`
	Name        string   `pg:"-"`
}

func MapToModel(ps []*ProductStock) []*model.ProductStock {
	var productsStocks []*model.ProductStock
	for _, stock := range ps {
		productsStocks = append(productsStocks, &model.ProductStock{
			Country:     stock.Country,
			Sku:         stock.Sku,
			Name:        stock.Name,
			StockChange: stock.StockChange,
		})
	}
	return productsStocks
}
func MapToProductStockEntity(productModel *model.ProductStock) *ProductStock {
	return &ProductStock{
		Sku:         productModel.Sku,
		Country:     productModel.Country,
		StockChange: productModel.StockChange,
	}
}

func MapToProductStockEntityList(productsModel []*model.ProductStock) []*ProductStock {
	var productStockList = make([]*ProductStock, 0)
	for _, product := range productsModel {
		productStockEntity := MapToProductStockEntity(product)
		productStockList = append(productStockList, productStockEntity)
	}
	return productStockList
}
