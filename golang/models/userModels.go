package models

import "time"

type User struct {
	Guid_user   string    `json:"guid_user"`
	Email       string    `json:"email"`
	Password    string    `json:"password,omitempty"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Date_joined time.Time `json:"date_joined"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UserRelational struct {
	Guid_user   string    `json:"guid_user"`
	Email       string    `json:"email"`
	Password    string    `json:"password,omitempty"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Date_joined time.Time `json:"date_joined"`
}
