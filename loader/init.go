package loader

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "repo.smartsheep.studio/smartsheep/neuralstack-api/configs"
	"repo.smartsheep.studio/smartsheep/neuralstack-api/filesystem"
)

func AutomaticSetup(router *gin.Engine) {
	if viper.GetBool("safe-mode") {
		log.Warn("Safe mode is activated. Plugins scanner was disabled!")
		return
	}

	start := time.Now()
	if plugins, err := ScanFolder(filesystem.GetPlugins()); err != nil {
		log.Fatalf("Error when scanning plugins: %w", err)
	} else {
		for _, plug := range plugins {
			if err := LoadPlugin(plug, router); err != nil {
				log.Fatalf("Error when loading plugin %s: %w\n", plug, err)
			}
		}
		log.Infof("Total plugged in %d plugin(s) took %dms", len(plugins), time.Since(start).Milliseconds())
	}
}
