package users

import (
	"encoding/json"
	"time"
)

type PublicUser struct {
	Id int64 `json:"id"`
	//FirstName   string    `json:"first_name"`
	//LastName    string    `json:"last_name"`
	//Email       string    `json:"email"`
	DateCreated time.Time `json:"date_created"`
	Status      string    `json:"status"`
	//Password    string    `json:"password"`
}

type PrivateUser struct {
	Id          int64     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"date_created"`
	Status      string    `json:"status"`
	//Password    string    `json:"password"`
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}

	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)

	return privateUser
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}

	return result
}
