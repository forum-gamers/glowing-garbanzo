package community

import (
	"context"
	"database/sql"

	"github.com/forum-gamers/glowing-garbanzo/pkg/base"
)

type CommunityRepo interface {
	FindByName(ctx context.Context, name string) (Community, error)
	Create(ctx context.Context, data *Community) error
	FindById(ctx context.Context, id string) (Community, error)
	DeleteById(ctx context.Context, id string) error
	UpdateImage(ctx context.Context, id, imageUrl, imageId string) error
	UpdateBackground(ctx context.Context, id, backgroundUrl, backgroundId string) error
	UpdateDesc(ctx context.Context, id, text string) error
	StartTransaction(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	ChangeOwnerShip(ctx context.Context, communityId, targetId string) error
}

type CommunityRepoImpl struct {
	base.BaseRepo
	Db *sql.DB
}
