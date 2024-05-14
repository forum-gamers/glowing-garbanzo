package controllers

import (
	"context"
	"errors"
	"time"

	protobuf "github.com/forum-gamers/glowing-garbanzo/generated/member"
	h "github.com/forum-gamers/glowing-garbanzo/helpers"
	"github.com/forum-gamers/glowing-garbanzo/pkg/community"
	"github.com/forum-gamers/glowing-garbanzo/pkg/member"
	"github.com/forum-gamers/glowing-garbanzo/pkg/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MemberService struct {
	protobuf.MemberServiceServer
	GetUser       func(context.Context) user.User
	CommunityRepo community.CommunityRepo
	MemberRepo    member.MemberRepo
}

func (s *MemberService) JoinCommunity(ctx context.Context, in *protobuf.CommunityIdInput) (*protobuf.Message, error) {
	if in.CommunityId == "" {
		return nil, status.Error(codes.InvalidArgument, "communityId is required")
	}

	// TODO refactor this implementation to be a single query
	if _, err := s.CommunityRepo.FindById(ctx, in.CommunityId); err != nil {
		return nil, err
	}

	userId := s.GetUser(ctx).Id
	if exists, err := s.MemberRepo.FindByCommunityIdAndUserId(
		ctx,
		in.CommunityId, userId,
	); err != nil && !errors.Is(err, h.NewAppError(codes.NotFound, "data not found")) {
		return nil, err
	} else if exists.Id != "" {
		return nil, status.Error(codes.AlreadyExists, "data is already exists")
	}

	if _, err := s.MemberRepo.Create(ctx, member.Member{
		UserId:      userId,
		CommunityId: in.CommunityId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}); err != nil {
		return nil, err
	}

	return &protobuf.Message{Message: "success"}, nil
}

func (s *MemberService) LeaveCommunity(ctx context.Context, in *protobuf.CommunityIdInput) (*protobuf.Message, error) {
	if in.CommunityId == "" {
		return nil, status.Error(codes.InvalidArgument, "communityId is required")
	}

	userId := s.GetUser(ctx).Id
	communityData, err := s.CommunityRepo.FindById(ctx, in.CommunityId)
	if err != nil {
		return nil, err
	}

	if communityData.Owner == userId {
		return nil, status.Error(codes.FailedPrecondition, "cannot leave community, please transfer the ownership first")
	}

	if err := s.MemberRepo.DeleteOne(ctx, communityData.Id, userId); err != nil {
		return nil, err
	}
	return &protobuf.Message{Message: "success"}, nil
}
