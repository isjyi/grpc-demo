package auth

import (
	"context"

	"github.com/isjyi/demo/pb"
)

func (s *AuthSrv) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{Code: 200, Msg: "success"}, nil
}
