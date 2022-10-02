package handler

import (
	"context"

	"one.now/backend/controller"
	pb "one.now/backend/gen/proto/auth/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer

	ctrler controller.AuthCtrler
}

func (s AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{
		Ok: s.ctrler.Login(ctx, getSession(ctx), req.Email),
	}, nil
}

func (s AuthService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	s.ctrler.Logout(ctx, getSession(ctx))
	return &pb.LogoutResponse{}, nil
}

func NewAuthService(c controller.AuthCtrler) AuthService {
	return AuthService{
		ctrler: c,
	}
}
