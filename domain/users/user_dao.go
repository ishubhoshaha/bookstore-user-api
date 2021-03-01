package users

import (
	"fmt"
	"github.com/ishubhoshaha/bookstore-user-api/mysql/users_db"
	"github.com/ishubhoshaha/bookstore-user-api/utils/date_utils"
	"github.com/ishubhoshaha/bookstore-user-api/utils/errors"
	"strings"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr{

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

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
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), "email_UNIQUE") {
			return errors.NewBadRequest(fmt.Sprintf("email %s already exists",user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to saving user %s", err.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to saving user %s", err.Error()))
	}
	//current := userDB[user.Id]
	//if current != nil {
	//	if current.Email == user.Email {
	//		return errors.NewBadRequest(fmt.Sprintf("email %s already registered", user.Email))
	//	}
	//	return errors.NewBadRequest(fmt.Sprintf("user %d already exist", user.Id))
	//}
	//user.DateCreated = date_utils.GetNowString()
	//userDB[user.Id] = user
	user.Id = userId
	return nil
}

