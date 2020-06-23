package auth

import (
	"context"
	"log"
	"reflect"
	"time"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/isjyi/grpc-demo/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type AuthSrv struct {
	db       *mongo.Collection
	validate *validator.Validate
	trans    ut.Translator
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

	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	//验证器注册翻译器
	err = zh_translations.RegisterDefaultTranslations(validate, trans)

	if err != nil {
		log.Fatal(err)
	}

	return &AuthSrv{
		db:       client.Database(global.DB_NAME).Collection(global.DB_COLLECTION),
		validate: validate,
		trans:    trans,
	}
}
