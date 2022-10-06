package handler

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"one.now/backend/controller"
	pb "one.now/backend/gen/proto/auth/v1"
)

var errSession = status.Error(codes.Internal, "session error")

type AuthService struct {
	pb.UnimplementedAuthServiceServer

	ctrler controller.AuthCtrler
}

func (s AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	session, err := GetSession(ctx)
	if err != nil {
		log.Println(err)
		return nil, errSession
	}

	return &pb.LoginResponse{
		Ok: s.ctrler.Login(ctx, session, req.Email),
	}, nil
}

func (s AuthService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	session, err := GetSession(ctx)
	if err != nil {
		log.Println(err)
		return nil, errSession
	}

	s.ctrler.Logout(ctx, session)
	return &pb.LogoutResponse{}, nil
}

func NewAuthService(c controller.AuthCtrler) AuthService {
	return AuthService{
		ctrler: c,
	}
}
