package base

import (
	"context"
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
