package repositoryPort

type IRepositoryPort interface {
	InitAdapter()
	IProductCrudsPort
	IProductStockCrudsPort
}
