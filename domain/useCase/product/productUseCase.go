package product

import (
	"ecommerce/port/domainPort"
	"ecommerce/port/outputPort/repositoryPort"
)

type productUseCase struct {
	repository repositoryPort.IRepositoryPort
}

func NewProductUseCases(repository repositoryPort.IRepositoryPort) domainPort.IProductDomainPort {
	return &productUseCase{repository: repository}
}
