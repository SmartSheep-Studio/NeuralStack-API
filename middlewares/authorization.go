package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"net/http"
	"repo.smartsheep.studio/smartsheep/neuralstack-api/datasource"
	"repo.smartsheep.studio/smartsheep/neuralstack-api/datasource/models"
	"strconv"
	"strings"
	"time"
)

func UseAuthorization(c *gin.Context, abort bool) {
	// Extract
	var token string
	if header := c.Request.Header.Get("Authorization"); header != "" {
		token = strings.TrimPrefix(header, "Bearer ")
	} else if cookie, err := c.Cookie("authorization"); err == nil {
		token = cookie
	} else if query := c.Query("token"); query != "" {
		token = query
	} else {
		if abort {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		return
	}

	// Decode
	claims, err := jwt.ParseWithClaims(token, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("secret")), nil
	})
	if err != nil || !claims.Valid {
		if abort {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		return
	} else if claims.Claims.(*models.UserClaims).Type != models.UserClaimsTypeAccess {
		if abort {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		return
	}

	// Parse Subject
	var uid uint
	if subject, err := claims.Claims.GetSubject(); err != nil {
		if abort {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		return
	} else {
		id, _ := strconv.Atoi(subject)
		uid = uint(id)
	}

	// Load out data
	if claims.Claims.(*models.UserClaims).PersonalTokenID != nil {
		var personal models.UserPersonalToken
		if datasource.C.Where("id = ?", claims.Claims.(*models.UserClaims).PersonalTokenID).First(&personal).Error != nil {
			if abort {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			return
		} else if personal.ExpiredAt != nil && personal.ExpiredAt.Unix() < time.Now().Unix() {
			if abort {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			return
		}
		c.Set("user-personal-token", personal)
	} else {
		var session models.UserSession
		if datasource.C.Where("id = ?", claims.Claims.(*models.UserClaims).SessionID).First(&session).Error != nil || !session.Available {
			if abort {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			return
		}
		c.Set("user-session", session)

		if session.IdentityID != nil {
			var identity models.OauthIdentity
			if datasource.C.Where("id = ?", *session.IdentityID).First(&identity).Error != nil {
				if abort {
					c.AbortWithStatus(http.StatusUnauthorized)
				}
				return
			}
			c.Set("user-identity", identity)
		}
	}

	var user models.User
	if datasource.C.Where("id = ?", uid).Preload("Punishes").First(&user).Error != nil {
		if abort {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		return
	} else {
		for _, punish := range user.Punishes {
			if punish.ExpiredAt == nil || punish.ExpiredAt.Unix() < time.Now().Unix() {
				continue
			}
			if punish.Level >= models.PunishCannotVisit {
				if abort {
					c.AbortWithStatus(http.StatusForbidden)
				}
				return
			}
		}
	}

	c.Set("user-account", user)
	c.Set("user-claims", claims.Claims)
	c.Set("user-authorized", true)
}
