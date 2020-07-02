package auth

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/isjyi/grpc-demo/model"
	"github.com/isjyi/grpc-demo/pb"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
)

func (s *AuthSrv) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	err := s.validate.Struct(req)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return nil, status.Errorf(http.StatusUnprocessableEntity, err.Translate(s.trans))
		}
	}

	var user model.User

	s.db.FindOne(ctx, bson.M{"username": req.GetUsername()}).Decode(&user)

	if user == model.NilUser {
		return nil, status.Error(http.StatusUnprocessableEntity, "用户名不存在")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword())); err != nil {
		return nil, status.Error(http.StatusUnprocessableEntity, "密码错误")
	}

	token, _ := s.jwt.Generate(&user)

	return &pb.LoginResponse{AccessToken: token}, nil
}
