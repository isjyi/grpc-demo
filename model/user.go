package model

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"golang.org/x/crypto/bcrypt"
)

//NilUser is the Nil value for User
var NilUser User

// User contains user's information
type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

// NewUser returns a new user
func NewUser(username string, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	user := &User{
		ID:       primitive.NewObjectID(),
		Username: username,
		Password: string(hashedPassword),
	}

	return user, nil
}

// IsCorrectPassword checks if the provided password is correct or not
func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// Clone returns a clone of this user
func (user *User) Clone() *User {
	return &User{
		Username: user.Username,
		Password: user.Password,
	}
}
