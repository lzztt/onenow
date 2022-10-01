package handler

import (
	"context"

	pb "one.now/backend/gen/proto/auth/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer

	allowedEmail string
}

func (s AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	sess, err := GetSession(ctx)
	if err != nil {
		return &pb.LoginResponse{
			Ok: false,
		}, nil
	}

	ok := sess.Loggedin
	if !ok {
		ok = req.Email == s.allowedEmail
		if ok {
			sess.Loggedin = ok
		}
	}

	return &pb.LoginResponse{
		Ok: ok,
	}, nil
}

func (s AuthService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	sess, err := GetSession(ctx)
	if err == nil {
		sess.Loggedin = false
	}

	return &pb.LogoutResponse{}, nil
}

func NewAuthService(email string) AuthService {
	return AuthService{
		allowedEmail: email,
	}
}
