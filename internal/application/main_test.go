package application

import (
	"os"
	"testing"

	"github.com/fbriansyah/micro-biller-service/internal/adapter/memorydb"
	"github.com/fbriansyah/micro-biller-service/internal/port"
)

var (
	testServiceWithMemoryDB *BillerService
	testDBMemoryAdapter     port.DatabasePort
)

func TestMain(m *testing.M) {
	testDBMemoryAdapter = memorydb.NewMemoryDatabase()

	testServiceWithMemoryDB = NewBillerService(testDBMemoryAdapter)

	os.Exit(m.Run())
}
