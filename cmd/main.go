package main

import (
	"database/sql"
	"log"
	"math"
	"time"

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

	databaseAdapter := db.NewDatabaseAdapter(sqlDB)

	bs := application.NewBillerService(databaseAdapter)

	httpAdapter := chi.NewChiAdapter(bs)
	httpAdapter.Run(chi.ChiAdapterConfig{
		ServerAddress: config.HTTPServerAddress,
	})
}

func openDB(dbDriver, dbSource string) (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB(dbDriver, dbSource string) *sql.DB {
	var connectionCounts int64
	var backOff = 1 * time.Second
	for {
		connection, err := openDB(dbDriver, dbSource)
		if err != nil {
			log.Println("Postgres not yet ready ...")
			connectionCounts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if connectionCounts > 10 {
			log.Println(err)
			return nil
		}

		backOff = time.Duration(math.Pow(float64(connectionCounts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backOff)
		continue
	}
}
