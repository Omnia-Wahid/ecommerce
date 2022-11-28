package routerPort

import "ecommerce/port/domainPort"

type IRouterPort interface {
	InitAdapter(domainPort domainPort.IDomainPort)
}
