package configs

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	api "repo.smartsheep.studio/smartsheep/neuralstack-api"
	"time"
)

func init() {
	LoadConfig()
}

func SafeLoadConfig() error {
	viper.NewWithOptions(viper.KeyDelimiter("."))
	viper.SetConfigName("settings")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	return viper.ReadInConfig()
}

func LoadConfig() {
	start := time.Now()
	if err := SafeLoadConfig(); err != nil {
		log.Fatalf("Error when reading configuration file: %w", err)
	} else {
		log.Infof("Successfully loaded configuration took %dms", time.Since(start).Milliseconds())
	}
}

func SaveConfig() error {
	return viper.SafeWriteConfig()
}

func AllConfig(p api.Plugin) any {
	return viper.Get(p.PackageID)
}

func GetConfig[T any](p api.Plugin, key string) T {
	return viper.Get(fmt.Sprintf("plugins.%s.%s", p.PackageID, key)).(T)
}
