package users

type User struct {
	Id          int64  `json:"id"`
	firstName   string `json:"first_name"`
	lastName    string `json:"last_name"`
	email       string `json:"email"`
	dateCreated string `json:"date_created"`
}
