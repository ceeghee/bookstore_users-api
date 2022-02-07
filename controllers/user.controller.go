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
	result, saveErr := services.UsersService.CreateUser(user)

	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}
	// fmt.Println("err", err)
	// fmt.Println("bytes to string", string(bytes))
	ctx.JSON(http.StatusCreated, result)
}

func Get(ctx *gin.Context) {
	userId, userErr := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		ctx.JSON(err.Status, err)
		return
	}
	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		ctx.JSON(getErr.Status, getErr)
		return
	}
	ctx.JSON(http.StatusOK, user.Marshall(false))
}

func SearchUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "implment me!")
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, err := services.UsersService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(false))
}
