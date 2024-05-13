package controllers

import (
	"context"
	"errors"
	"time"

	protobuf "github.com/forum-gamers/glowing-garbanzo/generated/community"
	h "github.com/forum-gamers/glowing-garbanzo/helpers"
	"github.com/forum-gamers/glowing-garbanzo/pkg/community"
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

	payload := community.Community{
		Name:        req.Name,
		Owner:       s.GetUser(ctx).Id,
		Description: req.Desc,
		ImageUrl:    req.ImageUrl,
		ImageId:     req.ImageId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := s.CommunityRepo.Create(ctx, &payload); err != nil {
		return nil, err
	}

	return &protobuf.Community{
		Id:          payload.Id,
		Name:        payload.Name,
		ImageUrl:    payload.ImageUrl,
		ImageId:     payload.ImageId,
		Description: payload.Description,
		CreatedAt:   payload.CreatedAt.String(),
		UpdatedAt:   payload.UpdatedAt.String(),
		Owner:       payload.Owner,
	}, nil
}
