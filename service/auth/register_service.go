package auth

import (
	"context"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc/status"

	"github.com/go-playground/validator/v10"
	"github.com/isjyi/grpc-demo/model"
	"github.com/isjyi/grpc-demo/pb"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *AuthSrv) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	err := s.validate.Struct(req)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return nil, status.Errorf(http.StatusUnprocessableEntity, err.Translate(s.trans))
		}
	}

	newUser, err := model.NewUser(req.GetUsername(), req.GetPassword())

	if err != nil {
		return nil, err
	}

	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if s.UsernameUsed(c, req.Username) {
		return nil, status.Error(http.StatusUnprocessableEntity, "用户已存在")
	}
	_, err = s.db.InsertOne(c, newUser)

	if err != nil {
		log.Println("Error inserting newUser: ", err.Error())
		return nil, status.Error(http.StatusInternalServerError, "请稍后重试")
	}

	return &pb.RegisterResponse{Code: 200, Msg: "success"}, nil
}

func (s *AuthSrv) UsernameUsed(ctx context.Context, username string) (r bool) {
	var res model.User
	s.db.FindOne(ctx, bson.M{"username": username}).Decode(&res)
	r = res != model.NilUser
	return
}
