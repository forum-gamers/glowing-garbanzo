package base

import (
	"context"
	"database/sql"
)

type BaseRepo interface {
	CreateData(ctx context.Context, table DBNAME, data any) (string, error)
}

type BaseRepoImpl struct{ Db *sql.DB }

type DBNAME = string
type ErrorMsg = string

const (
	COMMUNITY DBNAME = "Community"
)

const (
	DB_UNAVAILABLE ErrorMsg = "database is unavailable"
)
