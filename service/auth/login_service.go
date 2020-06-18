package auth

import (
	"context"
	"errors"
	"time"

	"github.com/isjyi/demo/global"
	"github.com/isjyi/demo/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection mongo.Collection

func (s *AuthSrv) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.GetUsername() == "" {
		return nil, errors.New("username is null")
	}
	res := &pb.LoginResponse{AccessToken: "1234"}
	return res, nil

	login, password := req.GetUsername(), req.GetPassword()
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	var user global.User
	userCollection.FindOne(ctx, bson.M{"$or": []bson.M{bson.M{"username": login}, bson.M{"email": login}}}).Decode(&user)
	if user == global.NilUser {
		return &proto.AuthResponse{}, errors.New("Wrong Login Credentials provided")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return &proto.AuthResponse{}, errors.New("Wrong Login Credentials provided")
	}
	return &proto.AuthResponse{Token: user.GetToken()}, nil
}
