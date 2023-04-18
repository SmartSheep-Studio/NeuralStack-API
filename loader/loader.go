package loader

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
	"plugin"
	plugins "repo.smartsheep.studio/smartsheep/neuralstack-api"
	"repo.smartsheep.studio/smartsheep/neuralstack-api/datasource"
	"repo.smartsheep.studio/smartsheep/neuralstack-api/web-injection"
)

var Plugins []*plugins.Plugin

func LoadPlugin(dst string, router *gin.Engine) error {
	plug, err := plugin.Open(dst)
	if err != nil {
		return err
	}

	// Lookup instance
	var instance *plugins.Plugin
	if lookup, err := plug.Lookup("P"); err != nil {
		return fmt.Errorf("error when loading plugin %s: missing instance %w", dst, err)
	} else {
		instance = lookup.(*plugins.Plugin)
	}

	// Run init function
	if instance.Init != nil {
		instance.Init(instance)
	}

	// Run migrate function
	if instance.Migrate != nil {
		instance.Migrate(datasource.C)
	}

	// Run setup function
	if instance.Setup != nil {
		instance.Setup(instance, router)
	} else {
		log.Warnf("Plugin %s haven't `Setup` function, loading it probably won't bring anything.", instance.Manifest.Name)
	}

	// Index routes
	if instance.Assets != nil && instance.Assets.Apps != nil {
		for _, app := range instance.Assets.Apps {
			app.ID = fmt.Sprintf("plugins.%s", app.ID)
			web.AppliedApps = append(web.AppliedApps, app)
		}
	}

	// Index locale
	if instance.Assets != nil && instance.Assets.Locale != nil {
		maps.Copy(web.AppliedLocale, instance.Assets.Locale)
	}

	Plugins = append(Plugins, instance)
	log.Infof("Successfully plugged in plugin %s v%s!", (*instance).Manifest.Name, (*instance).Manifest.Version)

	return nil
}
