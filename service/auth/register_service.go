package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/isjyi/grpc-demo/model"
	"github.com/isjyi/grpc-demo/pb"
)

func (s *AuthSrv) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	username, _, password := req.GetUsername(), req.GetRepeatPassword(), req.GetPassword()

	if err := s.validate.Struct(req); err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil, err
		}
		fmt.Println(err.Error())
		for _, err := range err.(validator.ValidationErrors) {
			// fmt.Println(err.Translate(s.trans)) //年龄必须大于18
			fmt.Println(err.Namespace()) // can differ when a custom TagNameFunc is registered or
			fmt.Println(err.Field())     // by passing alt name to ReportError like below
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
		return nil, errors.New("Something went wrong")
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
