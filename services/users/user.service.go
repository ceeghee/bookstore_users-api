package services

import (
	"github.com/ceeghee/bookstore_users-api/domain/users"
	"github.com/ceeghee/bookstore_users-api/utils/crypto_utils"
	"github.com/ceeghee/bookstore_users-api/utils/date_utils"
	"github.com/ceeghee/bookstore_users-api/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestError)
	CreateUser(users.User) (*users.User, *errors.RestError)
	LoginUser(users.LoginRequest) (*users.User, *errors.RestError)
	// UpdateUser(bool, users.User) (*users.User, errors.RestError)
	// DeleteUser(int64) errors.RestError
	// SearchUser(string) (users.Users, errors.RestError)
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestError) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, *errors.RestError) {
	result := &users.User{
		Email:    request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := result.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return result, nil
}
