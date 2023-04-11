package models

import (
	"time"
)

const (
	OneTimeVerifyPrimaryEmailCode   = iota
	OneTimeVerifySecondaryEmailCode = iota
	OneTimeVerifyPhoneNumberCode    = iota
	OneTimeDangerousPasscode
)

type OneTimePasscode struct {
	Model

	Type        int        `json:"type"`
	Passcode    string     `json:"passcode" gorm:"uniqueIndex"`
	RefreshedAt *time.Time `json:"refreshed_at"`
	ExpiredAt   *time.Time `json:"expired_at"`
	UserID      uint       `json:"user_id"`
}

type UserPersonalToken struct {
	Model

	Name        string     `json:"name"`
	Description string     `json:"description"`
	ExpiredAt   *time.Time `json:"expired_at"`
	TokenID     string     `json:"token"`
	UserID      uint       `json:"user_id"`
}
