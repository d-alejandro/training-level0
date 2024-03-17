package main

import (
	"github.com/d-alejandro/training-level0/internal/migrations"
	"github.com/d-alejandro/training-level0/internal/providers"
	"log"
)

func main() {
	envReaderProvider := providers.NewEnvReaderProvider()
	envReaderProvider.InitViper()

	databaseProvider := providers.NewDatabaseProvider()
	gorm := databaseProvider.InitGorm()

	orderModelsTableMigration := migrations.NewOrderModelsTableMigration(gorm)
	err := orderModelsTableMigration.Migrate()
	if err != nil {
		log.Fatal("Failed to complete migrations")
	}
}
