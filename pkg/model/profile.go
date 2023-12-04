package model

import "time"

type Profile struct {
	Id                       string     `json:"id"`
	PersonCode               string     `json:"personCode"`
	FirstName                string     `json:"firstName"`
	LastName                 string     `json:"lastName"`
	FullName                 string     `json:"fullName"`
	Birthday                 *time.Time `json:"birthday"`
	Gender                   string     `json:"gender"`
	UserPictureUrl           string     `json:"userPictureUrl"`
	UserId                   string     `json:"userId"`
	Address                  string     `json:"address"`
	City                     string     `json:"city"`
	Postal                   string     `json:"postal"`
	BirthPlace               string     `json:"birthPlace"`
	BirthCountry             string     `json:"birthCountry"`
	BirthCountryDescription  string     `json:"birthCountryDescription"`
	MaritalStatus            string     `json:"maritalStatus"`
	MaritalStatusDescription string     `json:"maritalStatusDescription"`
	Religion                 string     `json:"religion"`
	IsDeleted                bool       `json:"isDeleted"`
}
