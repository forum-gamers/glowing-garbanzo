package controllers

import (
	"context"
	"errors"
	"fmt"
	"time"

	protobuf "github.com/forum-gamers/glowing-garbanzo/generated/community"
	h "github.com/forum-gamers/glowing-garbanzo/helpers"
	"github.com/forum-gamers/glowing-garbanzo/pkg/base"
	"github.com/forum-gamers/glowing-garbanzo/pkg/community"
	"github.com/forum-gamers/glowing-garbanzo/pkg/member"
	"github.com/forum-gamers/glowing-garbanzo/pkg/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CommunityService struct {
	protobuf.UnimplementedCommunityServiceServer
	GetUser       func(context.Context) user.User
	CommunityRepo community.CommunityRepo
}

func (s *CommunityService) CreateCommunity(ctx context.Context, req *protobuf.CreateCommunityInput) (*protobuf.Community, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}

	if len(req.Name) < 3 || len(req.Name) > 30 {
		return nil, status.Error(codes.InvalidArgument, "name characters must be between 3 and 30")
	}

	if exists, err := s.CommunityRepo.FindByName(ctx, req.Name); err != nil && !errors.Is(err, h.NewAppError(codes.NotFound, "data not found")) {
		return nil, err
	} else if exists.Id != "" {
		return nil, status.Error(codes.AlreadyExists, "data is already exists")
	}

	tx, err := s.CommunityRepo.StartTransaction(ctx, nil)
	if err != nil {
		tx.Rollback()
		return nil, status.Error(codes.Unavailable, "failed to open transaction")
	}

	userId := s.GetUser(ctx).Id
	payload := community.Community{
		Name:          req.Name,
		Owner:         userId,
		Description:   req.Desc,
		ImageUrl:      req.ImageUrl,
		ImageId:       req.ImageId,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		BackgroundUrl: req.BackgroundUrl,
		BackgroundId:  req.BackgroundId,
	}

	query, values := base.GenerateInsertQueryAndValue(base.COMMUNITY, payload)
	if err := tx.QueryRowContext(ctx, query, values...).Scan(&payload.Id); err != nil {
		tx.Rollback()
		return nil, status.Error(codes.Unavailable, base.DB_UNAVAILABLE)
	}

	member := member.Member{
		UserId:      userId,
		CommunityId: payload.Id,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	query, values = base.GenerateInsertQueryAndValue(base.MEMBER, member)
	if err := tx.QueryRowContext(ctx, query, values...).Scan(&member.Id); err != nil {
		tx.Rollback()
		return nil, status.Error(codes.Unavailable, base.DB_UNAVAILABLE)
	}

	tx.Commit()
	return &protobuf.Community{
		Id:            payload.Id,
		Name:          payload.Name,
		ImageUrl:      payload.ImageUrl,
		ImageId:       payload.ImageId,
		Description:   payload.Description,
		CreatedAt:     payload.CreatedAt.String(),
		UpdatedAt:     payload.UpdatedAt.String(),
		Owner:         payload.Owner,
		BackgroundUrl: payload.BackgroundUrl,
		BackgroundId:  payload.BackgroundId,
	}, nil
}

func (s *CommunityService) DeleteCommunity(ctx context.Context, in *protobuf.CommunityIdInput) (*protobuf.ImageIdResp, error) {
	if in.CommunityId == "" {
		return nil, status.Error(codes.InvalidArgument, "communityId is required")
	}

	exists, err := s.CommunityRepo.FindById(ctx, in.CommunityId)
	if err != nil {
		return nil, err
	}

	me := s.GetUser(ctx)
	if exists.Owner != me.Id && me.AccountType != user.ADMIN {
		return nil, status.Error(codes.PermissionDenied, "Forbidden")
	}

	tx, err := s.CommunityRepo.StartTransaction(ctx, nil)
	if err != nil {
		tx.Rollback()
		return nil, status.Error(codes.Unavailable, base.DB_UNAVAILABLE)
	}

	if _, err := tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE communityId = $1", base.MEMBER), in.CommunityId); err != nil {
		tx.Rollback()
		return nil, err
	}

	if _, err := tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE id = $1", base.COMMUNITY), in.CommunityId); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &protobuf.ImageIdResp{
		ImageId:      exists.ImageId,
		BackgroundId: exists.BackgroundId,
	}, nil
}

func (s *CommunityService) UpdateImage(ctx context.Context, in *protobuf.UpdateImgInput) (*protobuf.Community, error) {
	if in.Url == "" || in.Id == "" || in.CommunityId == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	data, err := s.CommunityRepo.FindById(ctx, in.CommunityId)
	if err != nil {
		return nil, err
	}

	id := s.GetUser(ctx).Id
	if data.Owner != id {
		return nil, status.Error(codes.PermissionDenied, "Forbidden")
	}

	if err := s.CommunityRepo.UpdateImage(ctx, data.Id, in.Url, in.Id); err != nil {
		return nil, err
	}

	return &protobuf.Community{
		Id:            data.Id,
		Name:          data.Name,
		ImageUrl:      in.Url,
		ImageId:       in.Id,
		Description:   data.Description,
		CreatedAt:     data.CreatedAt.String(),
		UpdatedAt:     data.UpdatedAt.String(),
		Owner:         data.Owner,
		BackgroundUrl: data.BackgroundUrl,
		BackgroundId:  data.BackgroundId,
	}, nil
}

func (s *CommunityService) UpdateBackground(ctx context.Context, in *protobuf.UpdateImgInput) (*protobuf.Community, error) {
	if in.Url == "" || in.Id == "" || in.CommunityId == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	data, err := s.CommunityRepo.FindById(ctx, in.CommunityId)
	if err != nil {
		return nil, err
	}

	id := s.GetUser(ctx).Id
	if data.Owner != id {
		return nil, status.Error(codes.PermissionDenied, "Forbidden")
	}

	if err := s.CommunityRepo.UpdateBackground(ctx, data.Id, in.Url, in.Id); err != nil {
		return nil, err
	}
	return &protobuf.Community{
		Id:            data.Id,
		Name:          data.Name,
		ImageUrl:      data.ImageUrl,
		ImageId:       data.ImageId,
		Description:   data.Description,
		CreatedAt:     data.CreatedAt.String(),
		UpdatedAt:     data.UpdatedAt.String(),
		Owner:         data.Owner,
		BackgroundUrl: in.Url,
		BackgroundId:  in.Id,
	}, nil
}

func (s *CommunityService) UpdateDesc(ctx context.Context, in *protobuf.TextInput) (*protobuf.Community, error) {
	data, err := s.CommunityRepo.FindById(ctx, in.CommunityId)
	if err != nil {
		return nil, err
	}

	id := s.GetUser(ctx).Id
	if data.Owner != id {
		return nil, status.Error(codes.PermissionDenied, "Forbidden")
	}

	if err := s.CommunityRepo.UpdateDesc(ctx, data.Id, in.Text); err != nil {
		return nil, err
	}
	return &protobuf.Community{
		Id:            data.Id,
		Name:          data.Name,
		ImageUrl:      data.ImageUrl,
		ImageId:       data.ImageId,
		Description:   in.Text,
		CreatedAt:     data.CreatedAt.String(),
		UpdatedAt:     data.UpdatedAt.String(),
		Owner:         data.Owner,
		BackgroundUrl: data.BackgroundUrl,
		BackgroundId:  data.BackgroundId,
	}, nil
}
