package service

import (
	"context"
	"errors"

	"github.com/isjyi/demo/pb"
)

type LoginSvr struct{}

func NewLoginSrv() *LoginSvr  {
	return &LoginSvr{}
}

func (s *LoginSvr) Login(ctx context.Context,req *pb.LoginRequest) (*pb.LoginResponse,error)  {
	if (req.GetUsername() == "") {
		return nil,errors.New("username is null")
	}
	res := &pb.LoginResponse{AccessToken: "1234"}
	return res,nil;
}