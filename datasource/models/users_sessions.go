package models

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	UserSessionTypeOauth = iota
	UserSessionTypeAuth
	UserSessionTypeToken
)

type UserSession struct {
	Model

	IpAddress  string     `json:"ip"`
	Location   string     `json:"location"`
	Available  bool       `json:"available"`
	Type       int        `json:"type"`
	Code       string     `json:"code" gorm:"type:varchar(512)"`
	Access     string     `json:"access" gorm:"type:varchar(512)"`
	Refresh    string     `json:"refresh" gorm:"type:varchar(512)"`
	Scope      string     `json:"scope" gorm:"type:varchar(512)"`
	ExpiredAt  *time.Time `json:"expired_at"`
	IdentityID *uint      `json:"identity_id"`
	ClientID   *uint      `json:"client_id"`
	UserID     uint       `json:"user_id"`
}

const (
	UserClaimsTypeAccess  = "access_token"
	UserClaimsTypeRefresh = "refresh_token"
)

type UserClaims struct {
	jwt.RegisteredClaims

	Type            string `json:"typ"`
	SessionID       *uint  `json:"sid"`
	ClientID        *uint  `json:"cid"`
	PersonalTokenID *uint  `json:"tid"`
}
