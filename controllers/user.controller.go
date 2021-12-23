package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ceeghee/bookstore_users-api/domain/users"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		// TODO: Handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		// TODO handle json.Unmarshal error
		return
	}
	fmt.Println("user", user)
	fmt.Println("err", err)
	fmt.Println("bytes to string", string(bytes))
	ctx.String(http.StatusNotImplemented, "implment me!")
}

func GetUsers(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func SearchUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "implment me!")
}
