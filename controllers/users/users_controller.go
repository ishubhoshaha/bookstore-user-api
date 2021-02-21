package users

import (
	//"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ishubhoshaha/bookstore-user-api/domain/users"
	"github.com/ishubhoshaha/bookstore-user-api/services"
	"github.com/ishubhoshaha/bookstore-user-api/utils/errors"
	"net/http"
	"strconv"
)

func GetUser(context *gin.Context )  {
	userId, userErr := strconv.ParseInt(context.Param("user_id"),10,64)
	if userErr != nil {
		if userErr != nil {
			err := errors.NewBadRequest("user id should be a valid number")
			context.JSON(err.Status, err)
		}

	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		context.JSON(getErr.Status, getErr)
		return
	}
	context.JSON(http.StatusOK, user)
}

func CreateUser(context *gin.Context)  {
	var user users.User

	if err := context.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("Bad Request")
		fmt.Println(err.Error())
		context.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		context.JSON(saveErr.Status, saveErr)
		return
	}
	context.JSON(http.StatusCreated, result )
}

func FindUser(context *gin.Context)  {
	context.String(http.StatusNotImplemented, "implemented me!")
}