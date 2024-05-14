package member

import (
	"context"
	"fmt"

	"github.com/forum-gamers/glowing-garbanzo/database"
	h "github.com/forum-gamers/glowing-garbanzo/helpers"
	"github.com/forum-gamers/glowing-garbanzo/pkg/base"
	"google.golang.org/grpc/codes"
)

func NewMemberRepo(b base.BaseRepo) MemberRepo {
	return &MemberRepoImpl{b, database.DB}
}

func (r *MemberRepoImpl) FindByCommunityId(ctx context.Context, id string) (result []Member, err error) {
	rows, err := r.Db.QueryContext(ctx, fmt.Sprintf("SELECT * FROM %s WHERE communityId = $1", base.MEMBER), id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var data Member
		if err = rows.Scan(
			&data.Id,
			&data.CommunityId,
			&data.UserId,
			&data.CreatedAt,
			&data.UpdatedAt,
		); err != nil {
			return
		}
		result = append(result, data)
	}
	if len(result) < 1 {
		err = h.NewAppError(codes.NotFound, "data not found")
		return
	}
	return
}

func (r *MemberRepoImpl) FindByCommunityIdAndUserId(ctx context.Context, communityId, userId string) (result Member, err error) {
	rows, err := r.Db.QueryContext(ctx, fmt.Sprintf("SELECT * FROM %s WHERE communityId = $1 AND userId = $2", base.MEMBER), communityId, userId)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&result.Id,
			&result.CommunityId,
			&result.UserId,
			&result.CreatedAt,
			&result.UpdatedAt,
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

func (r *MemberRepoImpl) Create(ctx context.Context, data Member) (id string, err error) {
	err = r.Db.QueryRowContext(
		ctx,
		fmt.Sprintf("INSERT INTO %s (userId, communityId, createdAt, updatedAt) VALUES ($1, $2, $3, $4) RETURNING id", base.MEMBER),
		data.UserId, data.CommunityId, data.CreatedAt, data.UpdatedAt,
	).Scan(&id)
	return
}
