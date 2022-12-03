package repositoryMock

import "log"

type RepositoryMock struct {
	ProductsCrudsMock
	ProductsStockCrudsMock
}

func (thisRM *RepositoryMock) InitAdapter() {
	log.Println("initAdapter in mock")
}
