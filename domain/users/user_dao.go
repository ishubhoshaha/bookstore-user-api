package users

import (
	"fmt"
	"github.com/ishubhoshaha/bookstore-user-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr{
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	fmt.Println(user)
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequest(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequest(fmt.Sprintf("user %d already exist", user.Id))
	}
	userDB[user.Id] = user
	return nil
}

