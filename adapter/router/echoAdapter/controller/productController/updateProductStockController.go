package productController

import (
	"ecommerce/domain/model"
	"ecommerce/port/domainPort"
	"encoding/csv"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)

func prepareProductStockList(data [][]string) []*model.ProductStock {
	var productStockList []*model.ProductStock
	for i, line := range data {
		if i > 0 { // omit header line
			change, _ := strconv.Atoi(line[3])
			rec := &model.ProductStock{
				Country:     line[0],
				Sku:         line[1],
				Name:        line[2],
				StockChange: change,
			}

			productStockList = append(productStockList, rec)
		}
	}
	return productStockList
}
func UpdateProductsStocksByCSV(router *echo.Echo, domainPort domainPort.IProductDomainPort) {
	router.PUT("/product/stock", func(c echo.Context) (err error) {

		file, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "invalid data",
			})
		}
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "invalid data",
			})
		}
		csvReader := csv.NewReader(src)
		data, err := csvReader.ReadAll()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "invalid data",
			})
		}
		productStockList := prepareProductStockList(data)
		err = domainPort.UpdateProductsStocks(productStockList)
		if err != nil {
			log.Println("update products stocks error====> ", err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "internal server error",
			})
		}
		defer src.Close()
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "updated successfully",
		})
	})
}
