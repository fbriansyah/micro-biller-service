package memorydb

import (
	"sync"

	db "github.com/fbriansyah/micro-biller-service/internal/adapter/database"
)

var data *map[string]db.Billing
var once sync.Once

type MemoryDatabase struct {
}

// NewMemoryDatabase instantiate data in memory using singleton pattern
func NewMemoryDatabase() *MemoryDatabase {
	once.Do(func() {
		dt := make(map[string]db.Billing)

		data = &dt
	})

	return &MemoryDatabase{}
}
