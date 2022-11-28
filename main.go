package main

import (
	"ecommerce/adapter/repository"
	"ecommerce/adapter/router"
	"ecommerce/domain"
	"ecommerce/port/outputPort"
)

var outputPortInstance outputPort.IOutputPort

func init() {
	repository := repository.NewDatabaseAdapter()
	outputPortInstance = outputPort.IOutputPort{
		Repository: repository,
	}
}
func main() {
	domain := domain.NewDomain(outputPortInstance)
	router.NewRouterAdapter(domain)
}
