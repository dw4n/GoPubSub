package model

import "time"

type User struct {
	Id                     string    `json:"id"`
	Fullname               string    `json:"fullname"`
	Username               string    `json:"username"`
	Email                  string    `json:"email"`
	IsEmailVerified        bool      `json:"isEmailVerified"`
	EmailLastLogin         time.Time `json:"emailLastLogin"`
	MobilePhone            string    `json:"mobilePhone"`
	MobileLastLogin        time.Time `json:"mobileLastLogin"`
	HashPassword           string    `json:"HashPassword"`
	IsDeleted              bool      `json:"isDeleted"`
	IsLocked               bool      `json:"isLocked"`
	LockLimitUtc           time.Time `json:"lockLimitUtc"`
	InvalidPasswordCounter int       `json:"invalidPasswordCounter"`
	ForgotPasswordCounter  int       `json:"forgotPasswordCounter"`
	SignUpDate             time.Time `json:"signUpDate"`
}
