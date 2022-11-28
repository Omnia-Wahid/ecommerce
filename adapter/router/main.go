package router

import (
	"ecommerce/adapter/router/echoAdapter"
	"ecommerce/port/domainPort"
	"ecommerce/port/inputPort/routerPort"
)

type routerAdapter struct {
	Adapter routerPort.IRouterPort
}

func NewRouterAdapter(domainPort domainPort.IDomainPort) routerPort.IRouterPort {
	instance := routerAdapter{}
	instance.Adapter = &echoAdapter.RestEcho{}
	instance.Adapter.InitAdapter(domainPort)
	return instance.Adapter
}
