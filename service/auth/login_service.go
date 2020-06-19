package auth

import (
	"context"
	"errors"

	"github.com/isjyi/grpc-demo/pb"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection mongo.Collection

func (s *AuthSrv) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.GetUsername() == "" {
		return nil, errors.New("username is null")
	}
	res := &pb.LoginResponse{AccessToken: "1234"}
	return res, nil
}
