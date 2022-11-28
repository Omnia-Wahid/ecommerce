package repository

import (
	"ecommerce/adapter/repository/goPg"
	"ecommerce/port/outputPort/repositoryPort"
)

type databaseAdapter struct {
	Adapter repositoryPort.IRepositoryPort
}

func NewDatabaseAdapter() repositoryPort.IRepositoryPort {
	instance := &databaseAdapter{}
	instance.Adapter = new(goPg.DatabaseGoPg)
	instance.Adapter.InitAdapter()
	return instance.Adapter
}
