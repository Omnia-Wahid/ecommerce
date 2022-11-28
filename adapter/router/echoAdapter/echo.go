package echoAdapter

import (
	"ecommerce/adapter/router/echoAdapter/controller/productController"
	"ecommerce/port/domainPort"
	"ecommerce/utils"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type RestEcho struct {
	router *echo.Echo
	domain domainPort.IDomainPort
}

func (thisRE *RestEcho) InitAdapter(domainPort domainPort.IDomainPort) {
	thisRE.domain = domainPort
	thisRE.router = echo.New()
	thisRE.router.Use(middleware.Logger())
	thisRE.router.Use(middleware.Recover())

	productController := productController.NewProductController(thisRE.router, thisRE.domain)
	productController.Control()

	thisRE.router.Start(fmt.Sprintf(":%d", utils.GetEnvConfig().ServerPort))
}
