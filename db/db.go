package db

import (
	"context"
	"fmt"
	"strings"
)

type Database interface {
	Close() error

	ListTables(ctx context.Context, schema string) ([]string, error)
	ListColumns(ctx context.Context, schema, table string) ([]TableColumn, error)
}

type TableColumn struct {
	Name     string
	Type     string
	Nullable bool
	Default  string
}

func New(url string) (Database, error) {
	protocol := strings.Split(url, "://")[0]
	switch protocol {
	case "postgres", "postgresql", "pgx":
		return NewPostgres(url)
	default:
		return nil, fmt.Errorf("unsupported protocol: %s", protocol)
	}
}
