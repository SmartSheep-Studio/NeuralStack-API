package datasource

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"

	_ "repo.smartsheep.studio/smartsheep/neuralstack-api/configs"
)

var C *gorm.DB

func init() {
	var level logger.LogLevel
	if viper.GetBool("debug") {
		level = logger.Info
	} else {
		level = logger.Silent
	}

	cfg := &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  level,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		),
	}

	start := time.Now()
	if conn, err := gorm.Open(postgres.Open(viper.GetString("datasource.dsn")), cfg); err != nil {
		log.Fatalf("Error when connecting to database: %w", err)
	} else {
		C = conn
	}

	log.Infof("Successfully created connection to database took %dms", time.Since(start).Milliseconds())
}
