package filesystem

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	_ "repo.smartsheep.studio/smartsheep/neuralstack-api/configs"
	"repo.smartsheep.studio/smartsheep/neuralstack-api/utils"
)

var root string
var folders = []func() string{GetUserAssets, GetBuckets}

func init() {
	root = viper.GetString("datasource.user-contents")
	stat, err := os.Stat(root)
	if err != nil || !stat.IsDir() {
		if err := os.MkdirAll(root, 0755); err != nil {
			log.Fatalf("fatal error create user contents: %w", err)
		}

		for _, folder := range folders {
			utils.HandleFatalError("preparing filesystem", os.MkdirAll(folder(), 0755))
		}
	}
}

func GetPlugins() string {
	return filepath.Join(filepath.Dir(root), "plugins")
}

func GetBuckets() string {
	return filepath.Join(root, "buckets")
}

func GetUserAssets() string {
	return filepath.Join(root, "avatars")
}
