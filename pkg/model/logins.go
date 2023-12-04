package model

import "time"

type Login struct {
	Id                     string    `json:"id"`
	UserId                 string    `json:"userId"`
	Provider               string    `json:"provider"`
	LoginName              string    `json:"loginName"`
	RefreshToken           string    `json:"refreshToken"`
	RefreshTokenExpiryDate time.Time `json:"refreshTokenExpiryDate"`
	RegisteredDate         time.Time `json:"registeredDate"`
	LastLoginDateUtc       time.Time `json:"lastLoginDateUtc"`
}
