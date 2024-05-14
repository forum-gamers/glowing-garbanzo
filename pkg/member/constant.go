package member

import (
	"context"
	"database/sql"

	"github.com/forum-gamers/glowing-garbanzo/pkg/base"
)

type MemberRepo interface {
	FindByCommunityId(ctx context.Context, id string) ([]Member, error)
	FindByCommunityIdAndUserId(ctx context.Context, communityId, userId string) (Member, error)
	Create(ctx context.Context, data Member) (string, error)
}

type MemberRepoImpl struct {
	base.BaseRepo
	Db *sql.DB
}
