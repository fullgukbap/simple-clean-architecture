package database

import (
	"context"
	"fmt"
	"simple-clean-architecture/internal/infra/config"
	"simple-clean-architecture/internal/repository/ent"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	*ent.Client
}

func NewDatabase(config *config.Config) (*Database, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database,
	)

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return &Database{
		Client: client,
	}, nil
}

func (d Database) AutoMigration(ctx context.Context) error {
	if err := d.Schema.Create(ctx); err != nil {
		return err
	}

	return nil
}

func (d Database) Close() error {
	return d.Client.Close()
}
