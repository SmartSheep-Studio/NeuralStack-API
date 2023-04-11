package models

import (
	"gorm.io/datatypes"
	"time"
)

type User struct {
	Model

	AvatarID    string              `json:"avatar"`
	Username    string              `json:"username" gorm:"uniqueIndex"`
	Nickname    string              `json:"nickname"`
	Password    string              `json:"-"`
	Description string              `json:"description"`
	Details     UserDetails         `json:"details" gorm:"embedded;embeddedPrefix:details_"`
	Friends     []*User             `json:"friends" gorm:"many2many:user_friends"`
	Projects    []Project           `json:"projects" gorm:"many2many:project_developers"`
	Punishes    []Punish            `json:"punishes"`
	Passcodes   []OneTimePasscode   `json:"passcodes"`
	Sessions    []UserSession       `json:"sessions"`
	Tokens      []UserPersonalToken `json:"tokens"`
	GroupID     *uint               `json:"group_id"`
	Permissions datatypes.JSON      `json:"permissions"`
	LockedAt    *time.Time          `json:"locked_at"`
}

type UserDetails struct {
	Firstname                string     `json:"firstname"`
	Lastname                 string     `json:"lastname"`
	PrimaryEmail             string     `json:"primary_email"`
	PrimaryEmailVerifiedAt   *time.Time `json:"primary_email_verified_at"`
	SecondaryEmail           string     `json:"secondary_email"`
	SecondaryEmailVerifiedAt *time.Time `json:"secondary_email_verified_at"`
	PhoneNumber              string     `json:"phone_number"`
	PhoneNumberVerifiedAt    *time.Time `json:"phone_number_verified_at"`
}

type UserGroup struct {
	Model

	Name        string         `json:"name"`
	Description string         `json:"description"`
	Users       []User         `json:"users" gorm:"foreignKey:GroupID"`
	Permissions datatypes.JSON `json:"permissions"`
}
