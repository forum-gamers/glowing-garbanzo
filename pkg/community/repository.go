package community

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/forum-gamers/glowing-garbanzo/database"
	h "github.com/forum-gamers/glowing-garbanzo/helpers"
	"github.com/forum-gamers/glowing-garbanzo/pkg/base"
	"google.golang.org/grpc/codes"
)

func NewCommunityRepo(b base.BaseRepo) CommunityRepo {
	return &CommunityRepoImpl{b, database.DB}
}

func (r *CommunityRepoImpl) FindByName(ctx context.Context, name string) (result Community, err error) {
	rows, err := r.Db.QueryContext(ctx, fmt.Sprintf("SELECT * FROM %s WHERE name = $1", base.COMMUNITY), name)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&result.Id,
			&result.Name,
			&result.ImageUrl,
			&result.ImageId,
			&result.Description,
			&result.Owner,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.BackgroundUrl,
			&result.BackgroundId,
		); err != nil {
			return
		}
	}

	if result.Id == "" {
		err = h.NewAppError(codes.NotFound, "data not found")
		return
	}
	return
}

func (r *CommunityRepoImpl) Create(ctx context.Context, data *Community) error {
	id, err := r.CreateData(ctx, base.COMMUNITY, data)
	if err != nil {
		return err
	}

	data.Id = id
	return nil
}

func (r *CommunityRepoImpl) FindById(ctx context.Context, id string) (result Community, err error) {
	rows, err := r.Db.QueryContext(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id = $1", base.COMMUNITY), id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&result.Id,
			&result.Name,
			&result.ImageUrl,
			&result.ImageId,
			&result.Description,
			&result.Owner,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.BackgroundUrl,
			&result.BackgroundId,
		); err != nil {
			return
		}
	}

	if result.Id == "" {
		err = h.NewAppError(codes.NotFound, "data not found")
		return
	}

	return
}

func (r *CommunityRepoImpl) DeleteById(ctx context.Context, id string) error {
	return r.BaseRepo.DeleteById(ctx, base.COMMUNITY, id)
}

func (r *CommunityRepoImpl) UpdateImage(ctx context.Context, id, imageUrl, imageId string) error {
	_, err := r.Db.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET imageUrl = $1, imageId = $2 WHERE id = $3", base.COMMUNITY), imageUrl, imageId, id)
	return err
}

func (r *CommunityRepoImpl) UpdateBackground(ctx context.Context, id, backgroundUrl, backgroundId string) error {
	_, err := r.Db.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET backgroundUrl = $1, backGroundId = $2 WHERE id = $3", base.COMMUNITY), backgroundUrl, backgroundId, id)
	return err
}

func (r *CommunityRepoImpl) UpdateDesc(ctx context.Context, id, text string) error {
	_, err := r.Db.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET description = $1 WHERE id = $2", base.COMMUNITY), text, id)
	return err
}

func (r *CommunityRepoImpl) StartTransaction(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return r.BaseRepo.StartTransaction(ctx, opts)
}
