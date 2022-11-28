package goPg

import (
	"ecommerce/adapter/repository/goPg/entity"
	"ecommerce/domain/model"
	"github.com/go-pg/pg/v10"
)

type ProductCruds struct {
	Database *pg.DB
}

func (thisPC *ProductCruds) UpdateProducts(products []*model.ProductStock) error {
	productsEntity := entity.MapToProductEntityList(products)
	_, err := thisPC.Database.Model(&productsEntity).OnConflict("(sku) DO NOTHING").
		Insert()
	if err != nil {
		return err
	}
	return nil
}
func (thisPC *ProductCruds) GetProductBySKU(sku string) ([]*model.ProductStock, error) {
	var productEntity []*entity.ProductStock
	err := thisPC.Database.Model(&productEntity).
		ColumnExpr("product_stock.sku,product_stock.country,product_stock.stock_change,p.name").
		Join("left join products as p").
		JoinOn("product_stock.sku = p.sku").
		Where("product_stock.sku = ?", sku).
		Select()
	if err != nil {
		return nil, err
	}
	productStockList := entity.MapToModel(productEntity)

	return productStockList, nil
}
