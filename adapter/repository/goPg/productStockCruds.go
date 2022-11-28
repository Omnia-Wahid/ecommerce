package goPg

import (
	"ecommerce/adapter/repository/goPg/entity"
	"ecommerce/domain/model"
	"github.com/go-pg/pg/v10"
)

type ProductStockCruds struct {
	Database *pg.DB
}

func (thisPS *ProductStockCruds) UpdateProductsStocks(products []*model.ProductStock) error {
	productStockEntities := entity.MapToProductStockEntityList(products)
	_, err := thisPS.Database.Model(&productStockEntities).OnConflict("(sku, country) DO UPDATE").Insert()
	if err != nil {
		return err
	}
	return nil
}

func (thisPS *ProductStockCruds) GetStockChange(sku string, country string) (*model.ProductStock, error) {
	var productStockEntity entity.ProductStock
	err := thisPS.Database.Model(&productStockEntity).Where("sku= ?  and country = ?", sku, country).Select()
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	productStockModel := &model.ProductStock{
		Country:     productStockEntity.Country,
		Sku:         productStockEntity.Sku,
		StockChange: productStockEntity.StockChange,
	}
	return productStockModel, nil
}
