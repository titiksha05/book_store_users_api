package services

import (
	"github.com/titiksha05/book_store_users_api/domain/users"
	"github.com/titiksha05/book_store_users_api/utils/errors"
	//"strings"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	//if userId <= 0 {
		//return nil, errors.NewBadRequestError("invalid user id")
	//}
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr){
	//user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	//if user.Email == "" {
		//return nil, errors.NewBadRequestError("invalid email address")
	//}
	//return &user, nil
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

