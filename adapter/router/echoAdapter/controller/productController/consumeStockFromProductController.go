package productController

import (
	"ecommerce/domain/model"
	"ecommerce/port/domainPort"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func consumeStockFromProduct(router *echo.Echo, domainPort domainPort.IProductDomainPort) {
	router.PUT("/product/stock/:sku", func(c echo.Context) (err error) {
		sku := c.Param("sku")
		var consumeStockModel model.ConsumeStock

		if err := c.Bind(&consumeStockModel); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "invalid data",
			})
		}
		validateErr := consumeStockModel.ValidateConsumeStock()
		if validateErr != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": validateErr.Error(),
			})
		}
		err = domainPort.ConsumeStockFromProduct(sku, consumeStockModel.ConsumedAmount, consumeStockModel.Country)
		if err != nil {
			log.Println("consume stock from product error====> ", err)
			if err.Error() == "product stock not available" || err.Error() == "product stock not enough for consumed amount" {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"message": err.Error(),
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "internal server error",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "consumed successfully",
		})
	})
}
