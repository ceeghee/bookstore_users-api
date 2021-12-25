package controllers

import (
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	"net/http"
	"strconv"

	"github.com/ceeghee/bookstore_users-api/domain/users"
	services "github.com/ceeghee/bookstore_users-api/services/users"
	"github.com/ceeghee/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var user users.User
	// bytes, err := ioutil.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	// TODO: Handle error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	// TODO handle json.Unmarshal error
	// 	return
	// }

	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		ctx.JSON(restErr.Status, restErr)
		// TODO handle json error
		return
	}
	result, saveErr := services.CreateUser(user)

	if saveErr != nil {
		// TODO handle user creation err
		ctx.JSON(saveErr.Status, saveErr)
		return
	}
	// fmt.Println("err", err)
	// fmt.Println("bytes to string", string(bytes))
	ctx.JSON(http.StatusCreated, result)
}

func GetUsers(ctx *gin.Context) {
	userId, userErr := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		ctx.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		ctx.JSON(getErr.Status, getErr)
	}
	ctx.JSON(http.StatusOK, user)
}

func SearchUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "implment me!")
}
