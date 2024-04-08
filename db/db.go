package db

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/cod3rboy/practice-cqrs/config"
	"github.com/cod3rboy/practice-cqrs/ent"
	"go.uber.org/fx"
)

type DBClient struct {
	*ent.Client
}

type NewDatabaseClientParams struct {
	fx.In
	Config config.Config
}

func NewDatabaseClient(lc fx.Lifecycle, params NewDatabaseClientParams) *DBClient {
	dbClient := &DBClient{}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			client, err := ent.Open(dialect.Postgres,
				fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
					params.Config.DBHost,
					params.Config.DBPort,
					params.Config.DBUser,
					params.Config.DBPassword,
					params.Config.Database,
				),
			)
			if err != nil {
				return err
			}

			dbClient.Client = client
			return nil
		},
		OnStop: func(context.Context) error {
			if dbClient.Client == nil {
				return nil
			}
			return dbClient.Client.Close()
		},
	})
	return dbClient
}
