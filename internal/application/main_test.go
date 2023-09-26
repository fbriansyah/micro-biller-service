package application

import (
	"database/sql"
	"log"
	"os"
	"testing"

	db "github.com/fbriansyah/micro-biller-service/internal/adapter/database"
	"github.com/fbriansyah/micro-biller-service/internal/adapter/memorydb"
	"github.com/fbriansyah/micro-biller-service/internal/port"
	"github.com/fbriansyah/micro-biller-service/util"
	_ "github.com/lib/pq"
)

var (
	testServiceWithMemoryDB   *BillerService
	testServiceWithPostgresDB *BillerService
	testDBMemoryAdapter       port.DatabasePort
	testDBAdapter             port.DatabasePort
)

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testDBMemoryAdapter = memorydb.NewMemoryDatabase()
	testDBAdapter = db.NewDatabaseAdapter(testDB)

	testServiceWithMemoryDB = NewBillerService(testDBMemoryAdapter)
	testServiceWithPostgresDB = NewBillerService(testDBAdapter)

	os.Exit(m.Run())
}
