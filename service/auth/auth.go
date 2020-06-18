package auth

import "go.mongodb.org/mongo-driver/mongo"

type AuthSrv struct {
	Model mongo.Collection
}

func NewAuthSrv() *AuthSrv {
	return &AuthSrv{}
}
