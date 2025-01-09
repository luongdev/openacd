package database

import (
	"context"
	"errors"
	"github.com/luongdev/openacd/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DbClient struct {
	C  *mongo.Client
	db *mongo.Database

	ctx      context.Context
	cancelFn context.CancelFunc
}

func ConnectCtx(dbConfig *config.DbConfig, ctx context.Context) (*DbClient, error) {
	client, err := mongo.Connect(
		options.
			Client().
			ApplyURI(dbConfig.GetDsn()).
			SetConnectTimeout(dbConfig.ConnectTimeout).
			SetTimeout(dbConfig.Timeout).
			SetMinPoolSize(5).
			SetMaxPoolSize(dbConfig.PoolSize),
	)
	if err != nil {
		return nil, err
	}

	c := &DbClient{C: client}
	c.ctx, c.cancelFn = context.WithCancel(ctx)

	pingCtx, pingCancel := context.WithTimeout(c.ctx, dbConfig.ConnectTimeout)
	defer pingCancel()
	if err := c.C.Ping(pingCtx, nil); err != nil {
		return nil, err
	}

	return c, nil
}

func Connect(dbConfig *config.DbConfig) (*DbClient, error) {
	return ConnectCtx(dbConfig, context.Background())
}

func (c *DbClient) Database(name string) *DbClient {
	c.db = c.C.Database(name)

	return c
}

func (c *DbClient) Collection(name string) (*mongo.Collection, error) {
	if c.db == nil {
		return nil, errors.New("database is not selected")
	}

	return c.db.Collection(name), nil

}

func (c *DbClient) Disconnect() error {
	return c.C.Disconnect(c.ctx)
}
