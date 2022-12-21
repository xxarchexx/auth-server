package interfaces

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type RowWorker interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

// Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
type ExecWorker interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

type DbConnector interface {
	GetDb() (interface{}, error)
}
