package db

import (
	"context"
	"errors"
	"fmt"
	"universe-auth/config"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
)

type Connection struct {
	db *pgx.Conn
}
type ExecWorker interface {
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

func New(dbConfig config.DbConfig) (*Connection, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database)

	conn, err := pgx.Connect(context.Background(), connectionString)

	if err != nil {
		return nil, err
	}

	return &Connection{
		db: conn,
	}, nil
}

func (con *Connection) GetDb() (interface{}, error) {
	if con.db != nil {
		return con.db, nil
	}

	return nil, errors.New("database instanse not exist")
}
