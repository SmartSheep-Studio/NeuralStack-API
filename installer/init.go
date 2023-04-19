package installer

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"

	_ "repo.smartsheep.studio/smartsheep/neuralstack-api/configs"
	"repo.smartsheep.studio/smartsheep/neuralstack-api/filesystem"
)

func AutomaticSetup() {
	if viper.GetBool("safe-mode") {
		log.Warn("Safe mode is activated. Plugins installer was disabled!")
		return
	}

	start := time.Now()
	if packs, err := ScanFolder(filesystem.GetPlugins()); err != nil {
		log.Fatalf("Error when scanning plugins: %w", err)
	} else {
		for _, pack := range packs {
			if err := InstallPlugin(pack, filesystem.GetPlugins()); err != nil {
				log.Fatalf("Error when install plugin %s: %w\n", pack, err)
			}
		}
		log.Infof("Total installed %d plugin(s) took %dms", len(packs), time.Since(start).Milliseconds())
	}

	CleanupFolder(filesystem.GetPlugins())
}
