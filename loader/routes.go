package loader

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	api "repo.smartsheep.studio/smartsheep/neuralstack-api"
	"repo.smartsheep.studio/smartsheep/neuralstack-api/web-injection"
)

func InitPluginRoutes(router gin.IRouter) {
	router.GET("/api/plugins", func(c *gin.Context) {
		var p *api.Plugin
		for _, plugin := range Plugins {
			if plugin.Manifest.Package == c.Query("pkg") {
				p = plugin
				break
			}
		}
		if p == nil {
			c.JSON(http.StatusOK, Plugins)
		} else {
			c.JSON(http.StatusOK, p)
		}
	})

	router.GET("/plugins/:pkg/*path", func(c *gin.Context) {
		var pkg string
		if dst, err := base64.RawURLEncoding.DecodeString(c.Param("pkg")); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		} else {
			pkg = string(dst)
		}
		for _, app := range web.AppliedApps {
			if app.ID == pkg {
				if len(c.Param("path")) <= 0 {
					c.FileFromFS(app.RootFile, app.Assets)
				} else {
					c.FileFromFS(c.Param("path"), app.Assets)
				}
				return
			}
		}
		c.AbortWithStatus(http.StatusNotFound)
	})
}
