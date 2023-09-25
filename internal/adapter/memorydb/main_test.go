package memorydb

import (
	"os"
	"testing"
)

var testAdapter *MemoryDatabase

func TestMain(m *testing.M) {
	testAdapter = NewMemoryDatabase()

	os.Exit(m.Run())
}
