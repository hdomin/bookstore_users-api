package users

import (
	"strings"
	"time"

	"github.com/hdomin/bookstore_users-api/utils/errors"
)

type User struct {
	Id          int64     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"date_created"`
	Status      string    `json:"status"`
	Password    string    `json:"password"`
}

type Users []User

const (
	StatusActive = "active"
)

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace((strings.ToLower(user.Email)))
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}
