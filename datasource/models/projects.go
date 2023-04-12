package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type ProjectDeveloper struct {
	IsOwner   bool           `json:"is_owned"`
	UserID    uint           `json:"user_id" gorm:"primaryKey"`
	ProjectID uint           `json:"project_id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Project struct {
	Model

	Name          string             `json:"name"`
	Url           string             `json:"url"`
	Descriptions  string             `json:"descriptions"`
	Configuration datatypes.JSON     `json:"configuration"`
	Developers    []ProjectDeveloper `json:"developers"`
	OauthClients  []OauthClient      `json:"oauth_clients"`
	IsLocked      bool               `json:"is_locked"`
	IsPublic      bool               `json:"is_public"`
}

func (record *Project) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Delete(&record.Developers)
	tx.Delete(&record.OauthClients)
	return
}
