package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginCredentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type User struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"-" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

func NewUser(username string, password string) User {
	user := User{CreatedAt: time.Now(), Id: primitive.NewObjectID(), Username: username}
	pass, err := user.HashPassword(password)

	if err != nil {
		panic(err)
	}

	user.Password = pass

	return user
}

func (u *User) HashPassword(plainText string) (string, error) {
	if len(plainText) == 0 {
		return "", errors.New("Password cannot be empty.")
	}

	h, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)

	return string(h), err
}

func (u *User) CheckPassword(plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))

	return err == nil
}
