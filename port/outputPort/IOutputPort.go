package outputPort

import "ecommerce/port/outputPort/repositoryPort"

type IOutputPort struct {
	Repository repositoryPort.IRepositoryPort
}
