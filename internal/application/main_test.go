package application

import (
	"os"
	"testing"

	"github.com/fbriansyah/micro-biller-service/internal/adapter/memorydb"
	"github.com/fbriansyah/micro-biller-service/internal/port"
)

var testService *BillerService
var testDatabaseAdapter port.DatabasePort

func TestMain(m *testing.M) {

	testDatabaseAdapter = memorydb.NewMemoryDatabase()

	testService = NewBillerService(testDatabaseAdapter)

	os.Exit(m.Run())
}
