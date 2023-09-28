package main

import (
	"log"

	"github.com/fbriansyah/micro-biller-service/internal/adapter/chi"
	db "github.com/fbriansyah/micro-biller-service/internal/adapter/database"
	"github.com/fbriansyah/micro-biller-service/internal/application"
	"github.com/fbriansyah/micro-biller-service/util"
	_ "github.com/lib/pq"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(&logWriter{})

	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	sqlDB := connectToDB(config.DBDriver, config.DBSource)
	if sqlDB == nil {
		log.Fatal("cannot connect to db:", err)
	}

	runDBMigration(config.MigrationURL, config.DBSource)

	databaseAdapter := db.NewDatabaseAdapter(sqlDB)

	bs := application.NewBillerService(databaseAdapter)

	httpAdapter := chi.NewChiAdapter(bs)
	httpAdapter.Run(chi.ChiAdapterConfig{
		ServerAddress: config.HTTPServerAddress,
	})
}
