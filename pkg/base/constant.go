package base

import (
	"context"
	"database/sql"
)

type BaseRepo interface {
	CreateData(ctx context.Context, table DBNAME, data any) (string, error)
	DeleteById(ctx context.Context, table DBNAME, id string) error
	StartTransaction(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type BaseRepoImpl struct{ Db *sql.DB }

type DBNAME = string
type ErrorMsg = string

const (
	COMMUNITY DBNAME = "Community"
	MEMBER    DBNAME = "Member"
)

const (
	DB_UNAVAILABLE ErrorMsg = "database is unavailable"
)
