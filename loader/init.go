package loader

import (
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"

	_ "repo.smartsheep.studio/smartsheep/neuralstack-api/configs"
)

func AutomaticSetup(router *gin.Engine) {
	if viper.GetBool("safe-mode") {
		log.Warn("Safe mode is activated. Plugins scanner was disabled!")
		return
	}

	start := time.Now()
	if plugins, err := ScanFolder("plugins"); err != nil {
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
