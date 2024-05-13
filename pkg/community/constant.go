package community

import (
	"context"
	"database/sql"

	"github.com/forum-gamers/glowing-garbanzo/pkg/base"
)

type CommunityRepo interface {
	FindByName(ctx context.Context, name string) (Community, error)
	Create(ctx context.Context, data *Community) error
}

type CommunityRepoImpl struct {
	base.BaseRepo
	Db *sql.DB
}
