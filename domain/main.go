package domain

import (
	"ecommerce/domain/useCase/product"
	"ecommerce/port/domainPort"
	"ecommerce/port/outputPort"
)

type domainUseCase struct {
	domainPort.IProductDomainPort
}

func NewDomain(outputPort outputPort.IOutputPort) domainPort.IDomainPort {
	productUseCases := product.NewProductUseCases(outputPort.Repository)
	return &domainUseCase{
		productUseCases,
	}
}
