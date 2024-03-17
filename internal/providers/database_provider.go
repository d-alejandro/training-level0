package providers

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type DatabaseProvider struct {
}

func NewDatabaseProvider() *DatabaseProvider {
	return &DatabaseProvider{}
}

func (databaseProvider *DatabaseProvider) InitGorm() *gorm.DB {
	const sslMode = "disable"

	timezone := viper.GetString("DB_TIME_ZONE")

	dataSourceName := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_DATABASE"),
		viper.GetString("DB_PORT"),
		sslMode,
		timezone,
	)

	dialector := postgres.Open(dataSourceName)

	gormDB, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
		NowFunc: func() time.Time {
			location, _ := time.LoadLocation(timezone)
			return time.Now().In(location)
		},
	})

	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}

	return gormDB
}
