package services

import (
	"github.com/hdomin/bookstore_users-api/domain/users"
	"github.com/hdomin/bookstore_users-api/utils/crypto_utils"
	"github.com/hdomin/bookstore_users-api/utils/errors"
)

var (
	UsersService userServiceInterface = &usersService{}
)

type usersService struct {
}

type userServiceInterface interface {
	Get(int64) (*users.User, *errors.RestErr)
	Create(users.User) (*users.User, *errors.RestErr)
	Update(bool, users.User) (*users.User, *errors.RestErr)
	Delete(int64) *errors.RestErr
	Search(string) (users.Users, *errors.RestErr)
}

func (s *usersService) Get(userId int64) (*users.User, *errors.RestErr) {

	result := &users.User{Id: userId}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *usersService) Create(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) Update(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current := &users.User{Id: user.Id}

	if err := current.Get(); err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}

	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email

	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *usersService) Delete(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}

	return user.Delete()
}

func (s *usersService) Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}

	return dao.FindByStatus(status)

}
