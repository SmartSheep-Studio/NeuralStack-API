package datasource

import (
	"repo.smartsheep.studio/smartsheep/neuralstack-api/datasource/models"
	"repo.smartsheep.studio/smartsheep/neuralstack-api/utils"
)

func init() {
	utils.HandleFatalError(
		"setting up join tables",
		C.SetupJoinTable(&models.User{}, "Projects", &models.ProjectDeveloper{}),
	)
	utils.HandleFatalError(
		"auto migrating",
		C.AutoMigrate(
			&models.UserGroup{},
			&models.User{},
			&models.UserGroup{},
			&models.UserSession{},
			&models.UserPersonalToken{},
			&models.Punish{},
			&models.OneTimePasscode{},
			&models.OauthClient{},
			&models.OauthIdentity{},
		),
	)
}
