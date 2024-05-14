package main

import (
	"log"
	"net"
	"os"

	c "github.com/forum-gamers/glowing-garbanzo/controllers"
	db "github.com/forum-gamers/glowing-garbanzo/database"
	communityProto "github.com/forum-gamers/glowing-garbanzo/generated/community"
	memberProto "github.com/forum-gamers/glowing-garbanzo/generated/member"
	h "github.com/forum-gamers/glowing-garbanzo/helpers"
	"github.com/forum-gamers/glowing-garbanzo/interceptor"
	b "github.com/forum-gamers/glowing-garbanzo/pkg/base"
	"github.com/forum-gamers/glowing-garbanzo/pkg/community"
	"github.com/forum-gamers/glowing-garbanzo/pkg/member"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	h.PanicIfError(godotenv.Load())
	db.Conn()

	addr := os.Getenv("PORT")
	if addr == "" {
		addr = "50054"
	}

	lis, err := net.Listen("tcp", ":"+addr)
	if err != nil {
		log.Fatalf("Failed to listen : %s", err.Error())
	}

	baseRepo := b.NewBaseRepo()
	communityRepo := community.NewCommunityRepo(baseRepo)
	memberRepo := member.NewMemberRepo(baseRepo)

	interceptor := interceptor.NewInterCeptor()
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptor.Logging, interceptor.UnaryAuthentication),
	)

	communityProto.RegisterCommunityServiceServer(grpcServer, &c.CommunityService{
		GetUser:       interceptor.GetUserFromCtx,
		CommunityRepo: communityRepo,
	})

	memberProto.RegisterMemberServiceServer(grpcServer, &c.MemberService{
		GetUser:       interceptor.GetUserFromCtx,
		MemberRepo:    memberRepo,
		CommunityRepo: communityRepo,
	})

	log.Printf("Starting to serve in port : %s", addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %s", err.Error())
	}
}
