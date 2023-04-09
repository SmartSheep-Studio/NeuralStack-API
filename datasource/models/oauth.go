package models

import (
	"gorm.io/datatypes"
)

type OauthClient struct {
	Model

	Name                 string          `json:"name"`
	Description          string          `json:"description"`
	Secret               string          `json:"secret" gorm:"type:varchar(512)"`
	Domain               string          `json:"domain" gorm:"type:varchar(512)"`
	Sessions             []UserSession   `json:"sessions" gorm:"foreignKey:ClientID"`
	Identities           []OauthIdentity `json:"identities" gorm:"foreignKey:ClientID"`
	IsDanger             bool            `json:"is_danger"`
	IsOfficial           bool            `json:"is_official"`
	IsVerified           bool            `json:"is_verified"`
	IsDeveloping         bool            `json:"is_developing"`
	IsIdentitiesReadable bool            `json:"is_identities_readable"`
	IsIdentitiesEditable bool            `json:"is_identities_editable"`
	ProjectID            uint            `json:"project"`
}

type OauthIdentity struct {
	Model

	Nickname    string         `json:"nickname"`
	Data        datatypes.JSON `json:"data"`
	Permissions datatypes.JSON `json:"permissions"`
	Sessions    []UserSession  `json:"sessions" gorm:"foreignKey:IdentityID"`
	UserID      uint           `json:"user"`
	ClientID    uint           `json:"client"`
}
