package web

import (
	"repo.smartsheep.studio/smartsheep/neuralstack-api/datasource/models"
	"repo.smartsheep.studio/smartsheep/neuralstack-api/services"
)

func RenderLaunchpad(user *models.User) []map[string]any {
	var rendered []map[string]any
	for _, app := range AppliedApps {
		m := make(map[string]any)
		if !app.DisplayOnLaunchpad {
			continue
		} else if app.Authorized && user == nil {
			continue
		} else if !services.HasPermissions(user, app.Perms...) {
			continue
		} else {
			m["id"] = app.ID
			m["icon"] = app.Icon
			m["name"] = app.Name
			m["descriptions"] = app.Descriptions
		}

		if len(app.Descriptions) == 0 {
			m["descriptions"] = nil
		} else {
			m["descriptions"] = app.Descriptions
		}

		rendered = append(rendered, m)
	}
	return rendered
}
