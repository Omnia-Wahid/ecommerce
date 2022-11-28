package entity

import (
	"ecommerce/domain/model"
)

type Product struct {
	tableName struct{} `pg:"products,discard_unknown_columns"`
	Sku       string   `pg:"sku,pk,type:char(50)"`
	Name      string   `pg:"name,notnull,type:char(250)"`
}

func MapToProductEntity(productModel *model.ProductStock) *Product {
	return &Product{
		Sku:  productModel.Sku,
		Name: productModel.Name,
	}
}

func MapToProductEntityList(productsModel []*model.ProductStock) []*Product {
	var products = make([]*Product, 0)
	for _, product := range productsModel {
		productEntity := MapToProductEntity(product)
		products = append(products, productEntity)
	}
	return products
}
