package goPg

import (
	"context"
	"ecommerce/port/outputPort/repositoryPort"
	"ecommerce/utils"
	"fmt"
	"github.com/go-pg/pg/v10"
)

var config = utils.GetEnvConfig()

type DatabaseGoPg struct {
	pgConnection *pg.DB
	repositoryPort.IProductCrudsPort
	repositoryPort.IProductStockCrudsPort
}

func (thisDB *DatabaseGoPg) InitAdapter() {
	thisDB.pgConnection = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%v", config.RepositoryHost, config.RepositoryPort),
		User:     config.RepositoryUsername,
		Password: config.RepositoryPassword,
		Database: config.RepositoryDatabaseName,
	})
	thisDB.pgConnection.AddQueryHook(dbLogger{})
	thisDB.IProductStockCrudsPort = &ProductStockCruds{Database: thisDB.pgConnection}
	thisDB.IProductCrudsPort = &ProductCruds{Database: thisDB.pgConnection}
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	fmt.Println(q.FormattedQuery())
	return ctx, nil
}
func (d dbLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	//fmt.Println(q.FormattedQuery())
	return nil
}
