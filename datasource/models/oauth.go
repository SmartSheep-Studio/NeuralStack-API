package models

type OauthClient struct {
	Model

	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Secret       string        `json:"secret" gorm:"type:varchar(512)"`
	Domain       string        `json:"domain" gorm:"type:varchar(512)"`
	Sessions     []UserSession `json:"sessions" gorm:"foreignKey:ClientID"`
	IsDanger     bool          `json:"is_danger"`
	IsOfficial   bool          `json:"is_official"`
	IsVerified   bool          `json:"is_verified"`
	IsDeveloping bool          `json:"is_developing"`
	ProjectID    uint          `json:"project_id"`
}
