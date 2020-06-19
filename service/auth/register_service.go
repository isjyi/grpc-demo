package auth

import (
	"context"
	"errors"

	"github.com/isjyi/grpc-demo/pb"
)

func (s *AuthSrv) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	username, repeatPasswrod, password := req.GetUsername(), req.GetRepeatPassword(), req.GetPassword()

	if len(username) < 4 || len(username) > 20 || repeatPasswrod != password || len(password) < 6 || len(password) > 35 {
		return nil, errors.New("Validation failed")
	}

	return &pb.RegisterResponse{Code: 200, Msg: "success"}, nil
}
