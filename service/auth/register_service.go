package auth

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/isjyi/grpc-demo/model"
	"github.com/isjyi/grpc-demo/pb"
)

func (s *AuthSrv) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	username, repeatPasswrod, password := req.GetUsername(), req.GetRepeatPassword(), req.GetPassword()

	if len(username) < 4 || len(username) > 20 || repeatPasswrod != password || len(password) < 6 || len(password) > 35 {
		return nil, errors.New("Validation failed")
	}

	newUser, err := model.NewUser(username, password)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = s.db.InsertOne(ctx, newUser)
	if err != nil {
		log.Println("Error inserting newUser: ", err.Error())
		return nil, errors.New("Something went wrong")
	}

	return &pb.RegisterResponse{Code: 200, Msg: "success"}, nil
}
