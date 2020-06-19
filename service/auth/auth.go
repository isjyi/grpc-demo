package auth

import (
	"context"
	"log"
	"time"

	"github.com/isjyi/grpc-demo/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type AuthSrv struct {
	db *mongo.Collection
}

func NewAuthSrv() *AuthSrv {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(global.DB_URI))
	if err != nil {
		log.Fatal("Error connect to DB: ", err.Error())
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Error connect to DB: ", err.Error())
	}

	return &AuthSrv{
		db: client.Database(global.DB_NAME).Collection(global.DB_COLLECTION),
	}
}
