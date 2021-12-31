package users

import (
	"fmt"

	"github.com/ceeghee/bookstore_users-api/datasources/mysql/users_db"
	"github.com/ceeghee/bookstore_users-api/utils/date_utils"
	"github.com/ceeghee/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError("error when tying to get user")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return errors.NewInternalServerError("error when tying to get user")
	}
	// result := usersDB[user.Id]
	// if result == nil {
	// 	return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	// }
	// user.Id = result.Id
	// user.FirstName = result.FirstName
	// user.LastName = result.LastName
	// user.Email = result.Email
	// user.DateCreated = result.DateCreated
	return nil
}

func (user *User) SaveToMap() *errors.RestError {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("user %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	user.DateCreated = date_utils.GetNowString()
	usersDB[user.Id] = user
	return nil
}

func (user *User) Save() *errors.RestError {
	if users_db.Client == nil {
		return errors.NewInternalServerError("unknown error occured please try again")
	}
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	// execute the prepared statement
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if saveErr != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", saveErr.Error()))
	}
	userId, err := insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", saveErr.Error()))
	}
	user.Id = userId

	return nil
}
