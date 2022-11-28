package productController

import (
	"ecommerce/port/domainPort"
	"ecommerce/port/inputPort/routerPort"
	"github.com/labstack/echo"
)

type ProductController struct {
	router     *echo.Echo
	domainPort domainPort.IProductDomainPort
}

func NewProductController(router *echo.Echo, domainPort domainPort.IProductDomainPort) routerPort.IController {
	return &ProductController{
		router:     router,
		domainPort: domainPort,
	}
}

func (thisPC *ProductController) Control() {
	UpdateProductsStocksByCSV(thisPC.router, thisPC.domainPort)
	getProductBySKU(thisPC.router, thisPC.domainPort)
	consumeStockFromProduct(thisPC.router, thisPC.domainPort)
}
