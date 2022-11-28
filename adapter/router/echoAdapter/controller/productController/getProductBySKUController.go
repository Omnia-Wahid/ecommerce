package productController

import (
	"ecommerce/port/domainPort"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func getProductBySKU(router *echo.Echo, domainPort domainPort.IProductDomainPort) {
	router.GET("/product/stock/:sku", func(c echo.Context) (err error) {
		sku := c.Param("sku")
		productStocks, err := domainPort.GetProductBySKU(sku)
		if err != nil {
			log.Println("get product by sku error====> ", err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "internal server error",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "found successfully",
			"result":  productStocks,
		})
	})
}
