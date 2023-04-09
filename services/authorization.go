package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/exp/slices"
	"net/http"
	"repo.smartsheep.studio/smartsheep/neuralstack-api/datasource"
	"repo.smartsheep.studio/smartsheep/neuralstack-api/datasource/models"
	"repo.smartsheep.studio/smartsheep/neuralstack-api/middlewares"
)

var JWTSigningMethod = jwt.SigningMethodHS512

func GetAuthorization(c *gin.Context, abort bool, requires ...string) (*models.User, error) {
	middlewares.UseAuthorization(c, abort)
	user, authorized := c.Get("user-account")
	if !authorized {
		if abort {
			c.AbortWithStatus(http.StatusUnauthorized)
			return nil, fmt.Errorf("unauthorized")
		} else {
			return nil, nil
		}
	} else {
		u := user.(models.User)

		if !HasPermissions(&u, requires...) {
			c.AbortWithStatus(http.StatusForbidden)
			return nil, fmt.Errorf("forbidden")
		}

		return &u, nil
	}
}

func HasPermissions(user *models.User, requires ...string) bool {
	// Skip permission check if is possible
	if len(requires) <= 0 {
		return true
	}

	// Collecting permissions
	var perms []string
	json.Unmarshal(user.Permissions, &perms)

	var group models.UserGroup
	if err := datasource.C.Where("id = ?", user.GroupID).First(&group).Error; err == nil {
		var t []string
		json.Unmarshal(group.Permissions, &t)
		perms = append(perms, t...)
	}

	if !slices.Contains(perms, "*") {
		for _, require := range requires {
			if !slices.Contains(perms, require) {
				return false
			}
		}
	}

	return true
}

const (
	ProjectPermissionOwnedProject = iota
	ProjectPermissionDevelopProject
)

func HasPermissionOnProject(uid uint, pid uint, perms int) error {
	var role models.ProjectDeveloper
	if err := datasource.C.Where("user_id = ? AND project_id = ?", uid, pid).First(&role).Error; err != nil {
		return err
	}

	switch perms {
	case ProjectPermissionOwnedProject:
		if !role.IsOwner {
			return fmt.Errorf("forbidden")
		}
		break
	case ProjectPermissionDevelopProject:
		break
	}

	return nil
}
