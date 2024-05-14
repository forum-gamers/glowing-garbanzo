package base

import (
	"context"
	"fmt"
	"reflect"

	"github.com/forum-gamers/glowing-garbanzo/database"
	h "github.com/forum-gamers/glowing-garbanzo/helpers"
	"google.golang.org/grpc/codes"
)

func NewBaseRepo() BaseRepo {
	return &BaseRepoImpl{database.DB}
}

func (r *BaseRepoImpl) CreateData(ctx context.Context, table DBNAME, data any) (id string, err error) {
	if reflect.ValueOf(data).Kind() != reflect.Ptr {
		err = h.NewAppError(codes.Internal, "data must be pointer")
		return
	}

	query, values := generateInsertQueryAndValue(table, data)
	row := r.Db.QueryRowContext(ctx, query, values...)

	if err = row.Scan(&id); err != nil {
		err = h.NewAppError(codes.Unavailable, DB_UNAVAILABLE)
		return
	}
	return
}

func (r *BaseRepoImpl) DeleteById(ctx context.Context, table DBNAME, id string) error {
	_, err := r.Db.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE id = $1", table), id)
	return err
}
